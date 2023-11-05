package jsonld

import (
	"encoding/json"
)

type TermDefinition struct {
	ID   string `json:"@id"`
	Type string `json:"@type,omitempty"`
}

func (t *TermDefinition) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err == nil {
		t.ID = s
		return nil
	}

	var m map[string]string
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	if id, ok := m["@id"]; ok {
		t.ID = id
	}
	if typ, ok := m["@type"]; ok {
		t.Type = typ
	}
	return nil
}
