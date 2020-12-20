package schemax

import (
	"fmt"
	"reflect"

	"github.com/goools/tools/strx"
)

type SchemaType string
type Children map[string]*SchemaNode
type Schema Children

const (
	SchemaTypeObject = "object"
	SchemaTypeInt32  = "int32"
	SchemaTypeUint32 = "uint32"
	SchemaTypeInt64  = "int64"
	SchemaTypeUint64 = "uint64"
	SchemaTypeString = "string"
	SchemaTypeList   = "list"
	SchemaTypeAny    = "any"
	SchemaTypeMap    = "map"
)

func (m SchemaType) haveChild() bool {
	return m == SchemaTypeObject || m == SchemaTypeList || m == SchemaTypeMap
}

type SchemaNode struct {
	Type     string   `json:"type"`
	Children Children `json:"children"`
}

func (node *SchemaNode) makeSchemaChildren(t reflect.Type) {
	switch t.Elem().Kind() {
	case reflect.Struct:
		fieldCount := t.Elem().NumField()
		for i := 0; i < fieldCount; i++ {
			field := t.Elem().Field(i)
			SchemaTag, ok := getSchemaTag(field)
			if !ok {
				continue
			}
			fieldType := field.Type
			fieldObj := reflect.New(fieldType).Interface()
			node.Children[SchemaTag] = NewSchemaNode(fieldObj)
		}
	case reflect.Slice:
		if haveChildren(t) {
			node.makeSchemaChildren(t.Elem())
		}
	case reflect.Map:
		if haveChildren(t) {
			node.makeSchemaChildren(t.Elem())
		}
	}
}

func (node *SchemaNode) initSchemaType(t reflect.Type) {
	node.Type = node.getSchemaType(t)
}

func (node *SchemaNode) getSchemaType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr:
		return node.getSchemaType(t.Elem())
	case reflect.Struct:
		return SchemaTypeObject
	case reflect.Int32:
		return SchemaTypeInt32
	case reflect.Uint32:
		return SchemaTypeUint32
	case reflect.Int64:
		return SchemaTypeInt64
	case reflect.Uint64:
		return SchemaTypeUint64
	case reflect.String:
		return SchemaTypeString
	case reflect.Interface:
		return SchemaTypeAny
	case reflect.Map:
		keyType := node.getSchemaType(t.Key())
		valueType := node.getSchemaType(t.Elem())
		return fmt.Sprintf("%s<%s,%s>", SchemaTypeMap, keyType, valueType)
	case reflect.Slice:
		valueType := node.getSchemaType(t.Elem())
		return fmt.Sprintf("%s<%s>", SchemaTypeList, valueType)
	default:
		panic(fmt.Errorf("can not find type %s", t.Elem().Kind()))
	}
}

func NewSchemaNode(obj interface{}) (res *SchemaNode) {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Ptr {
		panic(fmt.Errorf("MakeSchema obj parameter have to a pointer"))
	}
	res = &SchemaNode{}
	res.Children = make(Children)
	res.makeSchemaChildren(objType)
	res.initSchemaType(objType)
	return
}

func haveChildren(t reflect.Type) bool {
	return t.Elem().Kind() == reflect.Struct ||
		t.Elem().Kind() == reflect.Slice || t.Elem().Kind() == reflect.Map
}

func getSchemaTag(field reflect.StructField) (string, bool) {
	tag := field.Tag
	SchemaTag := tag.Get("_schema")
	name := field.Name
	if SchemaTag == "-" {
		return "", false
	} else if SchemaTag == "" {
		SchemaTag = strx.ToSnakeCase(name)
	}
	return SchemaTag, true
}
