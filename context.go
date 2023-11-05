package jsonld

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
)

// Context is a JSON-LD context.
type Context map[string]TermDefinition

func (ctx Context) Inverse() InverseContext {
	inverse := make(InverseContext)
	for id, term := range ctx {
		inverse[term.ID] = InverseContextEntry{
			ID:   id,
			Term: term,
		}
	}
	return inverse
}

func (ctx *Context) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err == nil {
		var bytes []byte
		if u, _ := url.Parse(s); u.IsAbs() {
			// Get the context from the web.
			resp, err := http.Get(s)
			if err != nil {
				return err
			}
			if bytes, err = io.ReadAll(resp.Body); err != nil {
				return err
			}
		} else {
			// Get the context from the file system.
			if bytes, err = os.ReadFile(s); err != nil {
				return err
			}
		}

		doc := NewDocument()
		if err := json.Unmarshal(bytes, &doc); err != nil {
			return err
		}
		*ctx = doc.Context
		return nil
	}

	var m map[string]TermDefinition
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	*ctx = m
	return nil
}

type InverseContext map[string]InverseContextEntry

type InverseContextEntry struct {
	ID   string
	Term TermDefinition
}
