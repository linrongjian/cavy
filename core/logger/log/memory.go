package log

import (
	"fmt"

	"github.com/linrongjian/cavy/core/util"
)

// memoryLog is default micro log
type memoryLog struct {
	*util.Buffer
}

// NewLog returns default Logger with
func MemoryLog(opts ...Option) Log {
	// get default options
	options := DefaultOptions()

	// apply requested options
	for _, o := range opts {
		o(&options)
	}

	return &memoryLog{
		Buffer: util.NewBuffer(options.Size),
	}
}

// Write writes logs into logger
func (l *memoryLog) Write(r Record) error {
	l.Buffer.Put(fmt.Sprint(r.Message))
	return nil
}

// Read reads logs and returns them
func (l *memoryLog) Read(opts ...ReadOption) ([]Record, error) {
	options := ReadOptions{}
	// initialize the read options
	for _, o := range opts {
		o(&options)
	}

	var entries []*util.Entry
	// if Since options ha sbeen specified we honor it
	if !options.Since.IsZero() {
		entries = l.Buffer.Since(options.Since)
	}

	// only if we specified valid count constraint
	// do we end up doing some serious if-else kung-fu
	// if since constraint has been provided
	// we return *count* number of logs since the given timestamp;
	// otherwise we return last count number of logs
	if options.Count > 0 {
		switch len(entries) > 0 {
		case true:
			// if we request fewer logs than what since constraint gives us
			if options.Count < len(entries) {
				entries = entries[0:options.Count]
			}
		default:
			entries = l.Buffer.Get(options.Count)
		}
	}

	records := make([]Record, 0, len(entries))
	for _, entry := range entries {
		record := Record{
			Timestamp: entry.Timestamp,
			Message:   entry.Value,
		}
		records = append(records, record)
	}

	return records, nil
}

// Stream returns channel for reading log records
// along with a stop channel, close it when done
func (l *memoryLog) Stream() (Stream, error) {
	// get stream channel from ring buffer
	stream, stop := l.Buffer.Stream()
	// make a buffered channel
	records := make(chan Record, 128)
	// get last 10 records
	last10 := l.Buffer.Get(10)

	// stream the log records
	go func() {
		// first send last 10 records
		for _, entry := range last10 {
			records <- Record{
				Timestamp: entry.Timestamp,
				Message:   entry.Value,
				Metadata:  make(map[string]string),
			}
		}
		// now stream continuously
		for entry := range stream {
			records <- Record{
				Timestamp: entry.Timestamp,
				Message:   entry.Value,
				Metadata:  make(map[string]string),
			}
		}
	}()

	return &logStream{
		stream: records,
		stop:   stop,
	}, nil
}
