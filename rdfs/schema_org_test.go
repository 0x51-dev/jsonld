package rdfs

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

func Test_schemaOrg(t *testing.T) {
	// DOCS: https://schema.org/docs/developers.html
	resp, err := http.Get("https://schema.org/version/latest/schemaorg-current-http.jsonld")
	if err != nil {
		t.Fatal(err)
	}
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	doc := NewDocument()
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatal(err)
	}

	for _, n := range doc.Graph {
		for k, v := range n.Values {
			if strings.HasPrefix(k, "@") || strings.HasPrefix(k, "rdf:") || strings.HasPrefix(k, "rdfs:") {
				t.Error(n.ID, k, v)
			}
		}
	}
}
