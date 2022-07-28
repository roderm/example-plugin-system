package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"

	"github.com/rs/cors"
)

type svc struct {
	handler http.Handler
	timeout time.Duration
}
type GraphqlOption func(*svc)

func WithTimeout(t time.Duration) GraphqlOption {
	return func(s *svc) {
		s.timeout = t
	}
}
func Wrap(schema graphql.ExecutableSchema, opts ...GraphqlOption) http.Handler {
	hndl := handler.New(schema)
	hndl.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	hndl.AddTransport(transport.Options{})
	hndl.AddTransport(transport.GET{})
	hndl.AddTransport(transport.POST{})
	hndl.AddTransport(transport.MultipartForm{})

	hndl.SetQueryCache(lru.New(1000))

	hndl.Use(extension.Introspection{})
	hndl.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	corsHandler := cors.New(cors.Options{
		AllowOriginRequestFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		AllowCredentials:   true,
		OptionsPassthrough: true,
		Debug:              true,
	}).Handler(hndl)
	s := &svc{
		handler: corsHandler,
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

func (s *svc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ttlContext, cancel := context.WithTimeout(r.Context(), time.Second*30)
	defer cancel()

	// octx := metadata.NewOutgoingContext(ttlContext, metadata.New(map[string]string{
	// 	"authorization": r.Header.Get("Authorization"),
	// }))
	s.handler.ServeHTTP(w, r.WithContext(ttlContext))
}
