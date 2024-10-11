// Forked from https://github.com/PiotrKowalski/html-to-gomponents

package main

import (
	serviceformatter "html-to-htmgo/internal/adapters/services/formatter"
	serviceparser "html-to-htmgo/internal/adapters/services/parser"
)

func Parse(input []byte) []byte {
	parser := serviceparser.New()
	formatter := serviceformatter.New()
	parsed, err := parser.FromBytes(
		input,
	)

	if err != nil {
		return nil
	}

	return []byte(formatter.Format(parsed))
}
