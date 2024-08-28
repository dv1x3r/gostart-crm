package model

import (
	"encoding/json"
	"strings"
)

func getValue[T any](jsonMap map[string]json.RawMessage, jsonKey string, partialMap map[string]struct{}, partialKey string) T {
	var value T

	raw, ok := jsonMap[jsonKey]
	if !ok {
		return value
	}

	partialMap[partialKey] = struct{}{}

	if string(raw) == "null" || string(raw) == `""` {
		return value
	}

	if err := json.Unmarshal(raw, &value); err != nil {
		return value
	}

	if strValue, ok := any(value).(string); ok {
		value = any(strings.TrimSpace(strValue)).(T)
	}

	return value
}

func getValuePtr[T any](jsonMap map[string]json.RawMessage, jsonKey string, partialMap map[string]struct{}, partialKey string) *T {
	var value T

	raw, ok := jsonMap[jsonKey]
	if !ok {
		return nil
	}

	partialMap[partialKey] = struct{}{}

	if string(raw) == "null" || string(raw) == `""` {
		return nil
	}

	if err := json.Unmarshal(raw, &value); err != nil {
		return nil
	}

	if strValue, ok := any(value).(string); ok {
		value = any(strings.TrimSpace(strValue)).(T)
	}

	return &value
}
