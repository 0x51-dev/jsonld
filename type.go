package jsonld

import (
	"encoding/json"
	"slices"
)

type Type []string

func (typ Type) IsType(t string) bool {
	return slices.Contains(typ, t)
}

func (typ *Type) UnmarshalJSON(bytes []byte) error {
	var types []string
	if err := json.Unmarshal(bytes, &types); err == nil {
		*typ = types
		return nil
	}

	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	*typ = []string{s}
	return nil
}
