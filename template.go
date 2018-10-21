package main

import (
	"fmt"
	"time"
)

var tFormat = "Jan 02 2006"

func getContent(t, f string, tt time.Time) string {
	if f == "txt" {
		return fmt.Sprintf("%s\n%s\n----\n\n",
			t, tt.Format(tFormat))
	}
	return ""
}
