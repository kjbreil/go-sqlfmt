package group

import (
	"bytes"
	"testing"

	"github.com/kjbreil/go-sqlfmt/sqlfmt/lexer"
)

func TestReindentSelectGroup(t *testing.T) {
	tests := []struct {
		name        string
		tokenSource []Reindenter
		want        string
	}{
		{
			name: "normal case",
			tokenSource: []Reindenter{
				lexer.Token{Type: lexer.SELECT, Value: "SELECT"},
				lexer.Token{Type: lexer.IDENT, Value: "name"},
				lexer.Token{Type: lexer.COMMA, Value: ","},
				lexer.Token{Type: lexer.IDENT, Value: "age"},
			},
			want: "\nSELECT name\n  , age",
		},
		{
			name: "normal case",
			tokenSource: []Reindenter{
				lexer.Token{Type: lexer.SELECT, Value: "SELECT"},
				lexer.Token{Type: lexer.IDENT, Value: "TEST"},
				lexer.Token{Type: lexer.COMMA, Value: ","},
				lexer.Token{Type: lexer.IDENT, Value: "TEST2\r"},
			},
			want: "\nSELECT\n  TEST\n  , TEST2\r",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			selectGroup := &Select{Element: tt.tokenSource}

			selectGroup.Reindent(buf)
			got := buf.String()
			if tt.want != got {
				t.Errorf("want%#v, got %#v", tt.want, got)
			}
		})
	}
}

func TestIncrementIndentLevel(t *testing.T) {
	s := &Select{}
	s.IncrementIndentLevel(1)
	got := s.IndentLevel
	want := 1
	if got != want {
		t.Errorf("want %#v got %#v", want, got)
	}
}
