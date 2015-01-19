package sql_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/syst3mw0rm/sql-lexer-parser"
)

func TestParser_ParseStatement(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *sql.SelectStatement
		err  string
	}{
		// Single field statement
		{
			s: `SELEcT name from tbl`,
			stmt: &sql.SelectStatement{
				Fields:    []string{"name"},
				TableName: "tbl",
			},
		},

		// Multi field statement
		{
			s: `SELECT a, b from tbl`,
			stmt: &sql.SelectStatement{
				Fields:    []string{"a", "b"},
				TableName: "tbl",
			},
		},

		// Select all statement
		{
			s: `SELECT * FROM tbl`,
			stmt: &sql.SelectStatement{
				Fields:    []string{"*"},
				TableName: "tbl",
			},
		},

		// Errors
		{s: `foo`, err: `found "foo", expected SELECT`},
		{s: `SELECT !`, err: `found "!", expected field`},
		{s: `SELECT field xxx`, err: `found "xxx", expected FROM`},
		{s: `SELECT field FROM *`, err: `found "*", expected table name`},
	}

	for i, tt := range tests {
		stmt, err := sql.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n exp=%s\n got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}

	}

}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
