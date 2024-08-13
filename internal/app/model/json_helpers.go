package model

import (
	"encoding/json"
	"strings"
)

func getValue[T any](jsonKey string, fieldKey string, jsonRaw map[string]json.RawMessage, fieldMap map[string]struct{}) T {
	var value T

	raw, ok := jsonRaw[jsonKey]
	if !ok {
		return value
	}

	fieldMap[fieldKey] = struct{}{}

	if string(raw) == "null" || string(raw) == `""` {
		return value
	}

	if err := json.Unmarshal(raw, value); err != nil {
		return value
	}

	if strValue, ok := any(value).(string); ok {
		value = any(strings.TrimSpace(strValue)).(T)
	}

	return value
}

func getValuePtr[T any](jsonKey string, fieldKey string, jsonRaw map[string]json.RawMessage, fieldMap map[string]struct{}) *T {
	var value T

	raw, ok := jsonRaw[jsonKey]
	if !ok {
		return nil
	}

	fieldMap[fieldKey] = struct{}{}

	if string(raw) == "null" || string(raw) == `""` {
		return nil
	}

	if err := json.Unmarshal(raw, value); err != nil {
		return nil
	}

	if strValue, ok := any(value).(string); ok {
		trimmed := strings.TrimSpace(strValue)
		value = any(trimmed).(T)
	}

	return &value
}
