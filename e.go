package e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// New E
func New(
	ID string,
	Message string,
	Err error,
	kvs ...string,
) *Error {

	meta := map[string]string{}
	if len(kvs)%2 == 0 {
		prev := ""
		for i, val := range kvs {
			if i%2 == 0 {
				meta[val] = ""
			} else {
				meta[prev] = val
			}
			prev = val
		}
	} else {
		fmt.Println("ERROR: Number of KVs not even")
	}
	err := &Error{
		ID:      ID,
		Message: Message,
		Err:     Err,
		Meta:    meta,
	}
	return err
}

// Error defines a standard application error.
type Error struct {
	// Error ID for grepping purposes
	ID string `json:"id"`

	// Human-readable message.
	Message string `json:"message"`

	// Base error
	Err error `json:"err"`

	// Metadata from the app
	Meta map[string]string `json:"meta"`
}

// ErrorMessage returns the human-readable message of the error, if available.
// Otherwise returns a generic error message.
func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return "An internal error has occurred. Please contact technical support."
}

// Error returns the string representation of the error message.
// Does not contain parent metadata
func (e *Error) Error() string {
	var buf bytes.Buffer

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if len(e.Meta) > 0 {
			metaArr := []string{}
			for k, v := range e.Meta {
				metaArr = append(metaArr, k+"="+v)
			}
			buf.WriteString("[")
			buf.WriteString(strings.Join(metaArr, " "))
			buf.WriteString("] ")
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}

// JSON representation of the Error
// Contains parent metadata
func JSON(err error) []byte {

	if err == nil {
		return []byte("")
	} else if e, ok := err.(*Error); ok {
		b, err := json.Marshal(e)
		if err != nil {
			panic("could not marshal JSON: " + err.Error())
		}
		return b
	}
	return []byte("{An internal error has occurred. Please contact technical support.}")

}
