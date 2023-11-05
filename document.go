package jsonld

import (
	"encoding/json"
	"fmt"
	"strings"
)

var (
	// Used to define the shorthand names (terms) that are used throughout a JSON-LD document.
	_context = "@context"
	// Used to express a graph.
	_graph = "@graph"
	// Used to uniquely identify node objects that are being described in the document.
	// DOCS: https://www.w3.org/TR/json-ld11/#node-identifiers
	_id = "@id"
	// Used to set the type of node or the datatype of a typed value.
	_type = "@type"
	// Used to specify the data that is associated with a particular property in the graph.
	_value = "@value"
	// Used to specify the language for a particular string value or the default language of a JSON-LD document.
	_language = "@language"
)

type Document struct {
	Context Context `json:"@context"`
	Graph   Graph   `json:"@graph"`
	ID      string  `json:"@id"`
	Type    Type    `json:"@type"`
	Values  map[string]Node
}

func NewDocument() *Document {
	return &Document{
		Context: make(Context),
		Values:  make(map[string]Node),
	}
}

func (doc *Document) MarshalJSON() ([]byte, error) {
	d := make(map[string]any)
	if len(doc.Context) != 0 {
		// Only add the context if it is not empty.
		d["@context"] = doc.Context
	}
	if len(doc.Graph) != 0 {
		// Only add the graph if it is not empty.
		d["@graph"] = doc.Graph
	}
	for k, v := range doc.Values {
		d[k] = v
	}
	return json.Marshal(d)
}

func (doc *Document) UnmarshalJSON(bytes []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}
	if doc.Values == nil {
		// Make sure that the values map is not nil.
		doc.Values = make(map[string]Node)
	}
	for k, v := range raw {
		if strings.HasPrefix(k, "@") {
			switch k {
			case _context:
				if err := json.Unmarshal(v, &doc.Context); err != nil {
					return err
				}
			case _graph:
				if err := json.Unmarshal(v, &doc.Graph); err != nil {
					return err
				}
			case _id:
				if err := json.Unmarshal(v, &doc.ID); err != nil {
					return err
				}
			case _type:
				if err := json.Unmarshal(v, &doc.Type); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unknown keyword: %s", k)
			}
			continue
		}

		var value Node
		if err := json.Unmarshal(v, &value); err != nil {
			return err
		}
		doc.Values[k] = value
	}
	return nil
}
