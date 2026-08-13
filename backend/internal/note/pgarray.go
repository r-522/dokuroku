package note

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type stringArray []string

func (a stringArray) Value() (driver.Value, error) {
	parts := make([]string, 0, len(a))
	for _, value := range a {
		escaped := strings.ReplaceAll(value, `"`, `\"`)
		parts = append(parts, `"`+escaped+`"`)
	}
	return "{" + strings.Join(parts, ",") + "}", nil
}

func (a *stringArray) Scan(src any) error {
	if src == nil {
		*a = []string{}
		return nil
	}
	var text string
	switch value := src.(type) {
	case string:
		text = value
	case []byte:
		text = string(value)
	default:
		return fmt.Errorf("unsupported urls type %T", src)
	}
	text = strings.TrimPrefix(strings.TrimSuffix(text, "}"), "{")
	if text == "" {
		*a = []string{}
		return nil
	}
	items := strings.Split(text, ",")
	result := make([]string, 0, len(items))
	for _, item := range items {
		result = append(result, strings.Trim(item, `"`))
	}
	*a = result
	return nil
}
