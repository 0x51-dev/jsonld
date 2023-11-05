package jsonld

import "encoding/json"

type Graph []Node

func (graph *Graph) UnmarshalJSON(bytes []byte) error {
	var values []Node
	if err := json.Unmarshal(bytes, &values); err != nil {
		return err
	}
	*graph = values
	return nil
}
