package main

import (
	"sort"
	"strings"
	"testing"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) || (strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func Compare(a, b CaseInsensitive) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInsensitive_sort(t *testing.T) {
	cases := []struct {
		in   CaseInsensitive
		want CaseInsensitive
	}{
		{[]string{"iPhone", "iPad", "MacBook", "AppStore"}, []string{"AppStore iPad iPhone MacBook"}},
		{[]string{"b", "a", "d", "c"}, []string{"a b c d"}},
	}

	for _, c := range cases {
		got := CaseInsensitive(c.in)
		sort.Strings(got)
		if Compare(got, c.want) {
			t.Errorf("%s == %s, want %s", c.in, got, c.want)
		}
	}
}
