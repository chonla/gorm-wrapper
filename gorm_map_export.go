package gormwrapper

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
)

// ToGormMap convert Gorm Object Model to Map for Low level update purpose
// Exported Field Name will follow Gorm rules
// - if explicitly defined in format gorm:"column:xxx", xxx will be used as field name.
// - if there is not defined in gorm tag, field name will be converted in to snake case format and used as field name.
// - if gorm:"-" is defined, it will not be exported.
// - if mappable:"-" is defined, it will also not be exported.
func ToGormMap(ref interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(ref)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToGormMap only accepts structs; got %T", v)
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if fi := typ.Field(i); ast.IsExported(fi.Name) && mappable(fi) {
			fieldName := parseFieldName(fi)
			fieldValue := v.Field(i).Interface()
			out[fieldName] = fieldValue
		}
	}
	return out, nil
}

func mappable(fi reflect.StructField) bool {
	tagSettings := parseTagSetting(fi.Tag.Get("gorm"), ";")
	if tagSettings["-"] == "-" {
		return false
	}
	tagSettings = parseTagSetting(fi.Tag.Get("mappable"), ";")
	return tagSettings["-"] != "-"
}

func parseFieldName(fi reflect.StructField) string {
	tagSettings := parseTagSetting(fi.Tag.Get("gorm"), ";")
	fieldName := strcase.ToSnake(fi.Name)
	if value, ok := tagSettings["COLUMN"]; ok {
		fieldName = value
	}
	return fieldName
}

// parstTagSeting parses struct tag, copied from gorm package
func parseTagSetting(str string, sep string) map[string]string {
	settings := map[string]string{}
	names := strings.Split(str, sep)

	for i := 0; i < len(names); i++ {
		j := i
		if len(names[j]) > 0 {
			for {
				if names[j][len(names[j])-1] == '\\' {
					i++
					names[j] = names[j][0:len(names[j])-1] + sep + names[i]
					names[i] = ""
				} else {
					break
				}
			}
		}

		values := strings.Split(names[j], ":")
		k := strings.TrimSpace(strings.ToUpper(values[0]))

		if len(values) >= 2 {
			settings[k] = strings.Join(values[1:], ":")
		} else if k != "" {
			settings[k] = k
		}
	}

	return settings
}
