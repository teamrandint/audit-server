package log

import (
	"io"
	"seng468/auditserver/commands"
)

// Log contains a list of user commands
type Log struct {
	Entries []commands.Command
}

// Write takes in a writer object and writes the log to a file
func Write(w io.Writer) {

}

// String returns an XML representation of the log
func String() string {
	return "FULL LOG HERE"
}

// Insert takes a command object and inserts it into the log
func Insert(c commands.Command) bool {
	return false
}
