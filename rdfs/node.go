package rdfs

import (
	"encoding/json"
	"fmt"
	"github.com/0x51-dev/jsonld"
	"strings"
)

type Node struct {
	jsonld.Node

	Label         jsonld.Literal        `json:"rdfs:label"`
	Comment       jsonld.Literal        `json:"rdfs:comment"`
	SubPropertyOf jsonld.NodeReferences `json:"rdfs:subPropertyOf"`
	SubClassOf    jsonld.NodeReferences `json:"subClassOf"`
}

func (node *Node) UnmarshalJSON(bytes []byte) error {
	return node.UnmarshalValueExtensible(bytes, node.UnmarshalValue)
}

func (node *Node) UnmarshalValue(key string, bytes []byte) error {
	if strings.HasPrefix(key, "rdfs:") {
		switch key {
		case "rdfs:label":
			return json.Unmarshal(bytes, &node.Label)
		case "rdfs:comment":
			return json.Unmarshal(bytes, &node.Comment)
		case "rdfs:subPropertyOf":
			return json.Unmarshal(bytes, &node.SubPropertyOf)
		case "rdfs:subClassOf":
			return json.Unmarshal(bytes, &node.SubClassOf)
		default:
			return fmt.Errorf("unknown keyword: %s", key)
		}
	}
	return node.Node.UnmarshalValue(key, bytes)
}
