// Code generated by entc, DO NOT EDIT.

package issue

const (
	// Label holds the string label denoting the issue type in the database.
	Label = "issue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldIsPr holds the string denoting the is_pr field in the database.
	FieldIsPr = "is_pr"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeIssue holds the string denoting the issue edge name in mutations.
	EdgeIssue = "issue"
	// Table holds the table name of the issue in the database.
	Table = "tbl_issue"
	// IssueTable is the table that holds the issue relation/edge.
	IssueTable = "tbl_issue"
	// IssueInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	IssueInverseTable = "todos"
	// IssueColumn is the table column denoting the issue relation/edge.
	IssueColumn = "id"
)

// Columns holds all SQL columns for issue fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldIsPr,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultIsPr holds the default value on creation for the "is_pr" field.
	DefaultIsPr bool
)
