package schemax

import (
	"bytes"
	"encoding/json"
	"testing"
)

type SchemaDemoInner struct {
	InnerName  string            `json:"inner_name"`
	InnerValue map[string]*int64 `json:"inner_value"`
}

type SchemaDemo struct {
	Name        string                       `json:"name"`
	Tags        []string                     `json:"tags"`
	Have        map[string]int32             `json:"have"`
	Inner       SchemaDemoInner              `json:"inner"`
	Fields      map[string][]string          `json:"fields"`
	FieldInners map[string][]SchemaDemoInner `json:"field_inners"`
}

func TestMakeSchema(t *testing.T) {
	obj := &SchemaDemo{}
	res := NewSchemaNode(obj)
	byteBuffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(byteBuffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(res)
	if err != nil {
		t.Fatalf("json marshal res have an err: %v", err)
	}
	t.Logf("res: %s", byteBuffer.String())
}
