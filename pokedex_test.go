package main

import "testing"

func TestCleanInput(t *testing.T) {
tests := []struct {
input    string
expected []string
}{
{" Hello World ", []string{"hello", "world"}},
{"GoLang is Great!", []string{"golang", "is", "great"}},
{"  Multiple   Spaces  ", []string{"multiple", "spaces"}},
}

for _, c := range tests {
actual := cleanInput(c.input)
if len(actual) != len(c.expected) {
t.Errorf("cleanInput(%q) length = %d; want %d; got %v", c.input, len(actual), len(c.expected), actual)
continue
}
for i := range actual {
if actual[i] != c.expected[i] {
t.Errorf("cleanInput(%q)[%d] = %q; want %q", c.input, i, actual[i], c.expected[i])
}
}
}
}
