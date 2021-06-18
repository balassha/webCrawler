package parser

import (
	"io"
)

//An Interface to abstract implementation
//This allows us to modify the application to use a different library in future
type Parser interface {
	ProcessHTML(data io.Reader) (map[string]interface{}, error)
}
