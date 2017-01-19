package jsonutil

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

// A SyntaxError is a description of a JSON syntax error.
type SyntaxError struct {
	msg    string // description of error
	Offset int64  // error occurred after reading Offset bytes
	Line   int64  // the line number where error occurred
	Pos    int64  // the line pos where error occurred
}

func (e *SyntaxError) Error() string {
	return fmt.Sprintf("{offset %v, line %v:%v} - %v", e.Offset, e.Line, e.Pos, e.msg)
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, v)
	if err != nil {
		syntax, ok := err.(*json.SyntaxError)
		if ok {
			str := string(data[:syntax.Offset])
			line, pos := 1, 1
			for i, w := 0, 0; i < len(str); i += w {
				runeValue, width := utf8.DecodeRuneInString(str[i:])
				if runeValue == '\n' {
					line++
					pos = 1
				} else {
					pos++
				}
				w = width
			}
			err = &SyntaxError{syntax.Error(), syntax.Offset, int64(line), int64(pos)}
		}
	}
	return
}
