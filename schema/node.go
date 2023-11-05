package schema

import (
	"encoding/json"
	"fmt"
	"github.com/0x51-dev/jsonld"
	"github.com/0x51-dev/jsonld/rdfs"
	"strings"
)

type Node struct {
	rdfs.Node

	SameAs         jsonld.NodeReference  `json:"schema:sameAs"`
	Source         jsonld.NodeReferences `json:"schema:source"`
	SupersededBy   jsonld.NodeReference  `json:"schema:supersededBy"`
	IsPartOf       jsonld.NodeReference  `json:"schema:isPartOf"`
	InverseOf      jsonld.NodeReference  `json:"schema:inverseOf"`
	Contributor    jsonld.NodeReferences `json:"schema:contributor"`
	DomainIncludes jsonld.NodeReferences `json:"domainIncludes"`
	RangeIncludes  jsonld.NodeReferences `json:"rangeIncludes"`
}

func (node *Node) UnmarshalJSON(bytes []byte) error {
	return node.UnmarshalValueExtensible(bytes, node.UnmarshalValue)
}

func (node *Node) UnmarshalValue(key string, bytes []byte) error {
	if strings.HasPrefix(key, "schema:") {
		switch key {
		case "schema:sameAs":
			return json.Unmarshal(bytes, &node.SameAs)
		case "schema:source":
			return json.Unmarshal(bytes, &node.Source)
		case "schema:supersededBy":
			return json.Unmarshal(bytes, &node.SupersededBy)
		case "schema:isPartOf":
			return json.Unmarshal(bytes, &node.IsPartOf)
		case "schema:inverseOf":
			return json.Unmarshal(bytes, &node.InverseOf)
		case "schema:contributor":
			return json.Unmarshal(bytes, &node.Contributor)
		case "schema:domainIncludes":
			return json.Unmarshal(bytes, &node.DomainIncludes)
		case "schema:rangeIncludes":
			return json.Unmarshal(bytes, &node.RangeIncludes)
		default:
			return fmt.Errorf("unknown keyword: %s", key)
		}
	}
	return node.Node.UnmarshalValue(key, bytes)
}
