package main

import (
		"strings"
		"fmt"
)

func Spaces(n int) string {
  return strings.Repeat(" ", n)
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func MapKeys(m interface{}) ([]string, error) {
    switch t := m.(type) {
    case map[string]string:
        keys := make([]string, 0, len(t))
        for key := range t {
            keys = append(keys, key)
        }
        return keys, nil
    case map[string][]string:
        keys := make([]string, 0, len(t))
        for key := range t {
            keys = append(keys, key)
        }
        return keys, nil
    default:
        return nil, fmt.Errorf("unknown map type: %T", m)
    }
}
