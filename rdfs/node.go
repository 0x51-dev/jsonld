package rdfs

import (
	"encoding/json"
	"github.com/0x51-dev/jsonld"
)

type Node struct {
	jsonld.Node

	Label         jsonld.Literal `json:"rdfs:label"`
	Comment       jsonld.Literal `json:"rdfs:comment"`
	SubPropertyOf SubOf          `json:"rdfs:subPropertyOf"`
	SubClassOf    SubOf          `json:"subClassOf"`
}

type SubOf []jsonld.NodeReference

func (of *SubOf) UnmarshalJSON(bytes []byte) error {
	var id jsonld.NodeReference
	if err := json.Unmarshal(bytes, &id); err == nil {
		*of = []jsonld.NodeReference{id}
		return nil
	}
	var ids []jsonld.NodeReference
	if err := json.Unmarshal(bytes, &ids); err == nil {
		*of = ids
		return nil
	}
	return nil
}

func (of SubOf) MarshalJSON() ([]byte, error) {
	if len(of) == 1 {
		return json.Marshal(of[0])
	}
	return json.Marshal([]jsonld.NodeReference(of))
}

func (node *Node) UnmarshalJSON(bytes []byte) error {
	return node.UnmarshalValueExtensible(bytes, node.UnmarshalValue)
}

func (node *Node) UnmarshalValue(key string, bytes []byte) error {
	switch key {
	case "rdfs:label":
		return json.Unmarshal(bytes, &node.Label)
	case "rdfs:comment":
		return json.Unmarshal(bytes, &node.Comment)
	case "rdfs:subPropertyOf":
		return json.Unmarshal(bytes, &node.SubPropertyOf)
	case "rdfs:subClassOf":
		return json.Unmarshal(bytes, &node.SubClassOf)
	}
	return node.Node.UnmarshalValue(key, bytes)
}
