package utils

import "unicode"

func ToClassName(s string) string{
  runes := []rune(s)
  runes[0] = unicode.ToUpper(runes[0])
  return string(runes)
}

func ToAttributeName(s string) string{
  runes := []rune(s)
  runes[0] = unicode.ToLower(runes[0])
  return string(runes)
}
