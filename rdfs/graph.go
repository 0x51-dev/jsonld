package rdfs

import (
	"encoding/json"
	"fmt"
)

type Graph []Node

type GraphError struct {
	Context string
	error
}

func (e GraphError) Error() string {
	return fmt.Sprintf("graph error: %s | %s", e.Context, e.error.Error())
}

func (graph *Graph) UnmarshalJSON(bytes []byte) error {
	var raw []json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return GraphError{Context: "array", error: err}
	}

	var values []Node
	for _, value := range raw {
		var v Node
		if err := json.Unmarshal(value, &v); err != nil {
			return GraphError{Context: "value", error: err}
		}
		values = append(values, v)
	}
	*graph = values
	return nil
}
