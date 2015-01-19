package sql_test

import (
	"strings"
	"testing"

	"github.com/syst3mw0rm/sql-lexer-parser"
)

// Ensure the scanner can scan tokens correctly.
func TestSingleTokenScan(t *testing.T) {
	var testCases = []struct {
		s   string
		tok sql.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: sql.EOF, lit: ""},
		{s: `#`, tok: sql.ILLEGAL, lit: `#`},
		{s: ` `, tok: sql.WS, lit: " "},
		{s: "\t", tok: sql.WS, lit: "\t"},
		{s: "\n", tok: sql.WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: sql.ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, tok: sql.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: sql.IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: sql.FROM, lit: "FROM"},
		{s: `SELECT`, tok: sql.SELECT, lit: "SELECT"},
	}

	for i, tt := range testCases {
		s := sql.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()

		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got %q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}

func TestMultiTokenScan(t *testing.T) {

	var testCases = []struct {
		s   string
		ans []struct {
			tok sql.Token
			lit string
		}
	}{
		{s: `SELECT FROM`, ans: []struct {
			tok sql.Token
			lit string
		}{
			{tok: sql.SELECT, lit: "SELECT"},
			{tok: sql.WS, lit: " "},
			{tok: sql.FROM, lit: "FROM"},
		},
		},
	}

	for i, tt := range testCases {
		s := sql.NewScanner(strings.NewReader(tt.s))
		j := 0

		for {
			tok, lit := s.Scan()

			if tok == sql.EOF {
				break
			}

			if tok != tt.ans[j].tok {
				t.Errorf("%d. %q token mismatch: exp=%q got %q <%q>", i, tt.s, tt.ans[j].tok, tok, lit)
			} else if lit != tt.ans[j].lit {
				t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.ans[j].lit, lit)
			}

			j++
		}
	}
}
