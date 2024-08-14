package group

import (
	"bytes"
	"testing"

	"github.com/kjbreil/go-sqlfmt/sqlfmt/lexer"
)

func TestReindentWhereGroup(t *testing.T) {
	tests := []struct {
		name        string
		tokenSource []Reindenter
		want        string
	}{
		{
			name: "normal case",
			tokenSource: []Reindenter{
				lexer.Token{Type: lexer.WHERE, Value: "WHERE"},
				lexer.Token{Type: lexer.IDENT, Value: "something1"},
				lexer.Token{Type: lexer.IDENT, Value: "="},
				lexer.Token{Type: lexer.IDENT, Value: "something2"},
				lexer.Token{Type: lexer.AND, Value: "AND"},
				lexer.Token{Type: lexer.IDENT, Value: "something1"},
				lexer.Token{Type: lexer.IDENT, Value: "="},
				lexer.Token{Type: lexer.IDENT, Value: "something2"},
			},
			want: "\nWHERE something1 = something2\n    AND something1 = something2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			buf := &bytes.Buffer{}
			whereGroup := &Where{Element: tt.tokenSource}

			whereGroup.Reindent(buf)
			got := buf.String()
			if tt.want != got {
				t.Errorf("want%#v, got %#v", tt.want, got)
			}
		})

	}
}
