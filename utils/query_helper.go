package utils

import (
	"fmt"
	"strings"
)

func RawFilter(filters map[string]interface{}, table string) string {
	if len(filters) == 0 {
		return "(1 = 1)"
	}
	return BuildRawFilter("", "and_", filters, table)
}

func BuildRawFilter(fieldName string, key string, value interface{}, table string) string {
	tablePrefix := ""
	if table != "" {
		tablePrefix = table + "."
	}

	if _, ok := value.(map[string]interface{}); ok && key == "and_" {
		var args []string
		for attr, v := range value.(map[string]interface{}) {
			args = append(args, BuildRawFilter("", attr, v, table))
		}
		return fmt.Sprintf("(%s)", strings.Join(args, " AND "))
	} else if _, ok := value.(map[string]interface{}); ok && key == "or_" {
		var args []string
		for attr, v := range value.(map[string]interface{}) {
			args = append(args, BuildRawFilter("", attr, v, table))
		}
		return fmt.Sprintf("(%s)", strings.Join(args, " OR "))
	} else if key == "gt_" {
		v := value
		if _, isString := value.(string); isString {
			v = fmt.Sprintf("'%s'", value)
		}
		return fmt.Sprintf("%s%s > %v", tablePrefix, fieldName, v)
	} else if key == "gte_" {
		v := value
		if _, isString := value.(string); isString {
			v = fmt.Sprintf("'%s'", value)
		}
		return fmt.Sprintf("%s%s >= %v", tablePrefix, fieldName, v)
	} else if key == "lt_" {
		v := value
		if _, isString := value.(string); isString {
			v = fmt.Sprintf("'%s'", value)
		}
		return fmt.Sprintf("%s%s < %v", tablePrefix, fieldName, v)
	} else if key == "lte_" {
		v := value
		if _, isString := value.(string); isString {
			v = fmt.Sprintf("'%s'", value)
		}
		return fmt.Sprintf("%s%s <= %v", tablePrefix, fieldName, v)
	} else if key == "not_" {
		if value == nil {
			return fmt.Sprintf("%s%s IS NOT NULL", tablePrefix, fieldName)
		} else if _, isList := value.([]interface{}); !isList {
			v := value
			if _, isString := value.(string); isString {
				v = fmt.Sprintf("'%s'", value)
			}
			return fmt.Sprintf("%s%s <> %v", tablePrefix, fieldName, v)
		} else {
			var vs []string
			for _, v := range value.([]interface{}) {
				if _, isString := v.(string); isString {
					vs = append(vs, fmt.Sprintf("'%s'", v.(string)))
				} else {
					vs = append(vs, fmt.Sprintf("%d", v))
				}
			}
			return fmt.Sprintf("%s%s NOT IN (%s)", tablePrefix, fieldName, strings.Join(vs, ","))
		}
	} else if _, ok := value.([]interface{}); ok && key == "between_" {
		v1 := value.([]interface{})[0]
		if _, isString := v1.(string); isString {
			v1 = fmt.Sprintf("'%s'", v1)
		}
		v2 := value.([]interface{})[1]
		if _, isString := v2.(string); isString {
			v2 = fmt.Sprintf("'%s'", v2)
		}
		return fmt.Sprintf("(%s%s > %v AND %s%s <= %v)", tablePrefix, fieldName, v1, tablePrefix, fieldName, v2)
	} else if _, ok := value.(map[string]interface{}); ok {
		var args []string
		for attr, v := range value.(map[string]interface{}) {
			args = append(args, BuildRawFilter(key, attr, v, table))
		}
		return fmt.Sprintf("(%s)", strings.Join(args, " AND "))
	} else if _, isList := value.([]interface{}); isList {
		var vs []string
		for _, v := range value.([]interface{}) {
			if _, isString := v.(string); isString {
				vs = append(vs, fmt.Sprintf("'%s'", v.(string)))
			} else {
				vs = append(vs, fmt.Sprintf("%v", v))
			}
		}
		return fmt.Sprintf("%s%s IN (%s)", tablePrefix, key, strings.Join(vs, ","))
	} else if value == nil {
		return fmt.Sprintf("%s%s IS NULL", tablePrefix, key)
	} else {
		v := value
		if _, isString := value.(string); isString {
			v = fmt.Sprintf("'%s'", value)
		}
		return fmt.Sprintf("%s%s = %v", tablePrefix, key, v)
	}
}
