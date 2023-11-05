package jsonld

import (
	"encoding/json"
	"fmt"
	"strings"
)

type NodeReference struct {
	ID string `json:"@id"`
}

type NodeReferences []NodeReference

func (ns *NodeReferences) UnmarshalJSON(bytes []byte) error {
	var id NodeReference
	if err := json.Unmarshal(bytes, &id); err == nil {
		*ns = []NodeReference{id}
		return nil
	}
	var ids []NodeReference
	if err := json.Unmarshal(bytes, &ids); err == nil {
		*ns = ids
		return nil
	}
	return nil
}

func (ns NodeReferences) MarshalJSON() ([]byte, error) {
	if len(ns) == 1 {
		return json.Marshal(ns[0])
	}
	return json.Marshal([]NodeReference(ns))
}

type Node struct {
	ID       string `json:"@id"`
	Type     Type   `json:"@type,omitempty"`
	Value    string `json:"@value,omitempty"`
	Language string `json:"@language,omitempty"`
	Values   map[string]any
}

func (node *Node) MarshalJSON() ([]byte, error) {
	d := make(map[string]any)
	if node.ID != "" {
		d["@id"] = node.ID
	}
	if len(node.Type) != 0 {
		// Only add the type if it is not empty.
		d["@type"] = node.Type
	}
	if node.Value != "" {
		d["@node"] = node.Value
	}
	if node.Language != "" {
		d["@language"] = node.Language
	}
	for k, v := range node.Values {
		d[k] = v
	}
	return json.Marshal(d)
}

func (node *Node) UnmarshalJSON(bytes []byte) error {
	return node.UnmarshalValueExtensible(bytes, node.UnmarshalValue)
}

func (node *Node) UnmarshalValueExtensible(bytes []byte, extensible func(key string, bytes []byte) error) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}
	if node.Values == nil {
		// Make sure that the values map is not nil.
		node.Values = make(map[string]any)
	}
	for k, v := range raw {
		if strings.HasPrefix(k, "@") {
			switch k {
			case _id:
				if err := json.Unmarshal(v, &node.ID); err != nil {
					return err
				}
			case _type:
				if err := json.Unmarshal(v, &node.Type); err != nil {
					return err
				}
			case _value:
				if err := json.Unmarshal(v, &node.Value); err != nil {
					return err
				}
			case _language:
				if err := json.Unmarshal(v, &node.Language); err != nil {
					return err
				}
			default:
				return fmt.Errorf("unknown keyword: %s", k)
			}
			continue
		}

		if err := extensible(k, v); err != nil {
			return err
		}
	}
	return nil
}

func (node *Node) UnmarshalValue(key string, bytes []byte) error {
	var other any
	if err := json.Unmarshal(bytes, &other); err != nil {
		return err
	}
	node.Values[key] = other
	return nil
}

type ValueExtensible interface {
	json.Unmarshaler
	UnmarshalValue(key string, bytes []byte) error
	UnmarshalValueExtensible(bytes []byte, extensible func(key string, bytes []byte) error) error
}
