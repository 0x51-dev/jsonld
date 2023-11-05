package jsonld

import "encoding/json"

type Literal struct {
	Language string `json:"@language"`
	Value    string `json:"@value"`
}

func (l *Literal) UnmarshalJSON(bytes []byte) error {
	var v string
	if err := json.Unmarshal(bytes, &v); err == nil {
		l.Value = v
		return nil
	}

	var m map[string]string
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}
	if v, ok := m["@value"]; ok {
		l.Value = v
	}
	if lang, ok := m["@language"]; ok {
		l.Language = lang
	}
	return nil
}

func (l *Literal) MarshalJSON() ([]byte, error) {
	if l.Language == "" {
		return json.Marshal(l.Value)
	}
	return json.Marshal(map[string]string{
		"@value":    l.Value,
		"@language": l.Language,
	})
}
