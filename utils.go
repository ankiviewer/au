package main

import (
		"strings"
		"fmt"
)

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

func InAnkiRoot(fp string) bool {
  return strings.HasSuffix(fp, "anki_viewer_umbrella")
}

func InAnkiApp(fp string) bool {
  return strings.Contains(fp, "anki_viewer_umbrella")
}
