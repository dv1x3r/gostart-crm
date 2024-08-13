package model

import (
	"encoding/json"
	"strings"
)

type FieldRaw = map[string]json.RawMessage
type FieldMap = map[string]struct{}

func getValue[T comparable](jsonKey string, jsonRaw *FieldRaw, fieldKey string, fieldMap *FieldMap) T {
	return *new(T)
}

func getValuePointer[T comparable](jsonKey string, jsonRaw *FieldRaw, fieldKey string, fieldMap *FieldMap) *T {
	return nil

}

func getPartialValue[T comparable](nullable bool, partials FieldRaw, partialKey string, validate *[]string, validateKey string) *T {
	value := new(T)

	raw, ok := partials[partialKey]
	if !ok {
		if nullable {
			return nil
		} else {
			return value
		}
	}

	*validate = append(*validate, validateKey)

	if string(raw) == "null" || string(raw) == `""` {
		if nullable {
			return nil
		} else {
			return value
		}
	}

	if err := json.Unmarshal(raw, value); err != nil {
		if nullable {
			return nil
		} else {
			return value
		}
	}

	if strValue, ok := any(*value).(string); ok {
		strValue = strings.TrimSpace(strValue)
		if tValue, ok := any(strValue).(T); ok {
			if nullable && strValue == "" {
				return nil
			} else {
				return &tValue
			}
		}
	}

	return value
}
