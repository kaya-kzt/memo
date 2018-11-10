package main

import (
	"fmt"
	"time"
)

var tFormat = "Jan 02 2006"

func getContent(t time.Time, title, format string) string {
	var s string
	if format == "txt" {
		s += fmt.Sprintf("%s\nDate: %s\nAttendees: \nAgenda: \n----\n\n",
			title, t.Format(tFormat))
	} else if format == "md" {
		s += fmt.Sprintf("#%s\n*Date*: %s\n*Attendees*: \n*Agenda*: \n***\n\n",
			title, t.Format(tFormat))
	}
	return s
}
