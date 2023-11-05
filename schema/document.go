package schema

import (
	"encoding/json"
	"fmt"
	"github.com/0x51-dev/jsonld"
	"strings"
)

type Document struct {
	Context jsonld.Context `json:"@context"`
	Graph   Graph          `json:"@graph"`
}

func NewDocument() *Document {
	return &Document{
		Context: make(jsonld.Context),
	}
}

type DocumentError struct {
	Context string
	error
}

func (e DocumentError) Error() string {
	return fmt.Sprintf("document error: %s | %s", e.Context, e.error.Error())
}

func (doc *Document) UnmarshalJSON(bytes []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}
	for k, v := range raw {
		if strings.HasPrefix(k, "@") {
			switch k {
			case "@context":
				if err := json.Unmarshal(v, &doc.Context); err != nil {
					return DocumentError{Context: "@context", error: err}
				}
			case "@graph":
				if err := json.Unmarshal(v, &doc.Graph); err != nil {
					return DocumentError{Context: "@graph", error: err}
				}
			}
			continue
		}
	}
	return nil
}
