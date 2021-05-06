package main

import (
	"encoding/json"
	"fmt"
	"os"

	logger "github.com/ElrondNetwork/elrond-go-logger"
)

// this is custom code to monkey patch
// logger to allow json logging
func init() {
	// logger.NewLogLineWrapperFormatter()
	logger.ClearLogObservers()
	subject := logger.GetLogOutputSubject()
	subject.AddObserver(os.Stdout, &JSONFormatter{})
}

// ConsoleFormatter implements formatter interface and is used to format log lines to be written on the console
// It uses ANSI-color for colorized console/terminal output.
type JSONFormatter struct {
}

// Output converts the provided LogLineHandler into a slice of bytes ready for output
func (cf *JSONFormatter) Output(line logger.LogLineHandler) []byte {
	if line == nil {
		return nil
	}

	type logline struct {
		Level     string            `json:"level"`
		Name      string            `json:"name"`
		Timestamp int64             `json:"timestamp"`
		Message   string            `json:"message"`
		Map       map[string]string `json:"arguments"`
	}

	l := logline{
		Level:     logger.LogLevel(line.GetLogLevel()).String(),
		Name:      line.GetLoggerName(),
		Timestamp: line.GetTimestamp(),
		Message:   line.GetMessage(),
		Map:       make(map[string]string),
	}

	args := line.GetArgs()
	formatArgs(l.Map, args)

	data, err := json.Marshal(l)
	if err != nil {
		fmt.Printf("stakecamp: failed ot serialize logger %s", err.Error())
	}

	data = append(data, '\n')
	return data
}

// formatArgs iterates through the provided arguments displaying the argument name and after that its value
// The arguments must be provided in the following format: "name1", "val1", "name2", "val2" ...
// It ignores odd number of arguments
func formatArgs(m map[string]string, args []string) {
	if len(args) == 0 {
		return
	}
	for index := 1; index < len(args); index += 2 {
		m[args[index-1]] = args[index]
	}
}

// IsInterfaceNil returns true if there is no value under the interface
func (cf *JSONFormatter) IsInterfaceNil() bool {
	return cf == nil
}
