package main

import (
  "encoding/xml"
  "io"
  "strings"
)

func readXml(r io.Reader, fields []string) map[string][]string {
  result := make(map[string][]string)
  decoder := xml.NewDecoder(r)
  var scope []string
  var err error
  var t xml.Token
  for err == nil {
    t, err = decoder.Token()
    if start, ok := t.(xml.StartElement); ok {
      scope = append(scope, start.Name.Local)
    } else if _, ok := t.(xml.EndElement); ok {
      scope = scope[:len(scope)-1]
    } else if data, ok := t.(xml.CharData); ok {
      test := strings.Join(scope, ".")
      found := false
      for _, f := range fields {
        if f == test {
          found = true
          break
        }
      }
      if found {
        if val, ok := result[test]; ok {
          result[test] = append(val, string(data))
        } else {
          result[test] = []string{string(data)}
        }
      }
    }
  }
  return result
}
