// Code generated by ent, DO NOT EDIT.

package trades

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the trades type in the database.
	Label = "trades"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the trades in the database.
	Table = "trades"
)

// Columns holds all SQL columns for trades fields.
var Columns = []string{
	FieldID,
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

// OrderOption defines the ordering options for the Trades queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}
