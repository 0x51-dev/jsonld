package schema

import (
	"encoding/json"
	"io"
	"net/http"
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
	if err := json.Unmarshal(raw, NewDocument()); err != nil {
		t.Fatal(err)
	}
}
