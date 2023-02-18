package log

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	DefaultSize   = 1024
	DefaultLog    = NewLog()
	DefaultFormat = TextFormat
)

type Log interface {
	Read(...ReadOption) ([]Record, error)
	Write(Record) error
	Stream() (Stream, error)
}

type Record struct {
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]string `json:"metadata"`
	Message   interface{}       `json:"proto"`
}

type Stream interface {
	Chan() <-chan Record
	Stop() error
}

type FormatFunc func(Record) string

func TextFormat(r Record) string {
	t := r.Timestamp.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s %v", t, r.Message)
}

func JSONFormat(r Record) string {
	b, _ := json.Marshal(r)
	return string(b)
}
