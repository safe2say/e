package e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Kind is a subtype of string to scope error codes nicely
type Kind string

const (
	// ECONFLICT means action cannot be performed
	ECONFLICT Kind = "conflict"
	// EINTERNAL means internal error
	EINTERNAL Kind = "internal"
	// EINVALID means validation failed
	EINVALID Kind = "invalid"
	// ENOTFOUND means entity does not exist
	ENOTFOUND Kind = "not_found"
)

var toHTTPCode = map[Kind]int{
	ECONFLICT: http.StatusBadRequest,
	EINTERNAL: http.StatusInternalServerError,
	EINVALID:  http.StatusBadRequest,
	ENOTFOUND: http.StatusBadRequest,
}

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	Code Kind `json:"code"`

	// Human-readable message.
	Message string `json:"message"`

	// Logical operation and nested error.
	Op string `json:"op"`

	// Base error
	Err error `json:"err"`

	// Metadata from the app
	Meta map[string]string `json:"meta"`
}

// HTTPCode returns the http code of the root error, if available. Otherwise returns 500.
func HTTPCode(err error) int {
	if err == nil {
		return 0
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return toHTTPCode[e.Code]
	} else if ok && e.Err != nil {
		return HTTPCode(e.Err)
	}
	return http.StatusInternalServerError
}

// ErrorCode returns the code of the root error, if available. Otherwise returns EINTERNAL.
func ErrorCode(err error) Kind {
	if err == nil {
		return ""
	} else if e, ok := err.(*Error); ok && e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
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

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
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
