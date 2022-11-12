package gobyexample

import (
	"fmt"
	st "strings"
)

var p = fmt.Println

func stringFunctions() {

	p("Contains:  ", st.Contains("test", "es"))
	p("Count:     ", st.Count("test", "t"))
	p("HasPrefix: ", st.HasPrefix("test", "te"))
	p("HasSuffix: ", st.HasSuffix("test", "st"))
	p("Index:     ", st.Index("test", "e"))
	p("Join:      ", st.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", st.Repeat("a", 5))
	p("Replace:   ", st.Replace("foo", "o", "0", -1))
	p("Replace:   ", st.Replace("foo", "o", "0", 1))
	p("Split:     ", st.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", st.ToLower("TEST"))
	p("ToUpper:   ", st.ToUpper("test"))
}
