package highlight

import (
	"bytes"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

// Highlight the source using chroma in HTML format.
// It will return the original content in plain text when it fails.
func Highlight(source, lang, theme, linenr string) string {
	var buf bytes.Buffer
	lexer := lexers.Get(lang)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	// coalesce runs of identical token types into a single token
	lexer = chroma.Coalesce(lexer)

	style := styles.Get(theme)
	if style == nil {
		// default theme is dracula
		style = styles.Dracula
	}

	iterator, err := lexer.Tokenise(nil, source)
	if err != nil {
		return source
	}

	opts := []html.Option{html.Standalone(true)}

	if linenr != "" {
		opts = append(opts, html.WithLineNumbers(true))
	}

	formatter := html.New(opts...)
	err = formatter.Format(&buf, style, iterator)
	if err != nil {
		return source
	}

	return buf.String()
}
