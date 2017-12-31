package main

import (
  "testing"
  "reflect"
  "sort"
)

func TestStringInSlice(t *testing.T) {
  type testdata struct {
    s string
    sl []string
    res bool
  }

  var tests = []testdata{
    {"hello", []string{}, false},
    {"hello", []string{"hello"}, true},
    {"hello", []string{"h", "hello", "hello"}, true},
    {"hello", []string{"hllo", "ello"}, false},
    {"one", []string{"none", "ones"}, false},
  }

  for _, data := range tests {
    v := StringInSlice(data.s, data.sl)
    if v != data.res {
      t.Error(
        "For string:", data.s,
        "and slice:", data.sl,
        "expected:", data.res,
        "instead got:", v,
      )
    }
  }
}

func TestMapKeys1(t *testing.T) {
  input := map[string]string{
    "one": "first",
    "two": "second",
  }
  expected := []string{"one", "two"}
  actual, err := MapKeys(input)
  sort.Strings(actual)
  sort.Strings(expected)

  if !reflect.DeepEqual(actual, expected) || err != nil {
    t.Error(
      "For map:", input,
      "expected:", expected,
      "instead got:", actual,
    )
  }
}

func TestMapKeys2(t *testing.T) {
  input := map[string][]string{
    "one": []string{"first", "initial"},
    "two": []string{"second"},
  }
  expected := []string{"one", "two"}
  actual, err := MapKeys(input)
  sort.Strings(actual)
  sort.Strings(expected)

  if !reflect.DeepEqual(actual, expected) || err != nil {
    t.Error(
      "For map:", input,
      "expected:", expected,
      "instead got:", actual,
    )
  }
}

func TestMapKeys3(t *testing.T) {
  input := map[string]int{
    "one": 1,
    "two": 2,
  }
  _, err := MapKeys(input)

  if err == nil {
    t.Error(
      "For map:", input,
      "expected: err",
      "instead got: nil",
    )
  }
}
