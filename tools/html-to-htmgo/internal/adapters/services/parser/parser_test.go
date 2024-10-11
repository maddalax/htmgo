package parser

import (
	"errors"
	"testing"
)

func FuzzFromBytes(f *testing.F) {
	serviceParser := New()
	f.Add([]byte("<html><body>Hello World</body></html>"))
	f.Add([]byte("<html><head><title>Test</title></head><body>Sample</body></html>"))
	f.Add([]byte("<div>Some random text</div>"))
	f.Add([]byte("Invalid HTML"))
	f.Add([]byte("<a/><p/>"))
	f.Add([]byte("</BodY><!0"))

	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) > 10000 { // (10KB)
			t.Skip()
		}
		_, err := serviceParser.FromBytes(data)
		if err != nil {
			return
		}
		if err != nil && !isExpectedError(err) {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func isExpectedError(err error) bool {
	return err != nil && errors.Is(err, ParseErr)
}
