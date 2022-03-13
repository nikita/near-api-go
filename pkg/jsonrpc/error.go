package jsonrpc

import (
	"encoding/json"
	"fmt"
)

const (
	CodeParseError     = -32700
	CodeInvalidRequest = -32600
	CodeMethodNotFound = -32601
	CodeInvalidParams  = -32602
	CodeInternalError  = -32603

	CodeServerErrorRangeStart = -32099
	CodeServerErrorRangeEnd   = -32000
)

type Error struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    json.RawMessage  `json:"data"`
	Name    string           `json:"name"`
	Cause   ErrorCause       `json:"cause"`
	Info    *json.RawMessage `json:"info"`
}

type ErrorCause struct {
	Name string          `json:"name"`
	Info json.RawMessage `json:"info"`
}

func (err Error) Error() string {
	// TODO: use Name, Cause & Info
	//fmt.Printf("%#v\n", err)
	fmt.Printf("name: %s\n", err.Name)
	fmt.Printf("cause: name=%s, info=%s\n", err.Cause.Name, string(err.Cause.Info))
	info := "<nil>"
	if err.Info != nil {
		info = string(*err.Info)
	}
	fmt.Printf("info: %s\n", info)

	return fmt.Sprintf("JSON-RPC error '%s' (%d) %s", err.Message, err.Code, string(err.Data))
}
