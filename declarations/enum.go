package main

import "fmt"

type Enum interface {
	name() string
	ordinal() int
	valueOf() *[]string
}

type LogLevel uint

func (log LogLevel) name() string {
	return logs[log]
}
func (log LogLevel) ordinal() int {
	return int(log)
}
func (log LogLevel) String() string {
	return logs[log]
}
func (log LogLevel) values() *[]string {
	return &logs
}

var logs = []string{"UNSPECIFIED", "TRACE", "INFO", "WARNING", "ERROR"}

const (
	UNSPECIFIED LogLevel = iota
	TRACE // 1
	INFO // 2
	WARNING // 3
	ERROR // 4
)

func main() {
	fmt.Println(UNSPECIFIED)
	fmt.Println(TRACE.name())
	fmt.Println(TRACE.String())
	fmt.Println(TRACE.ordinal())
	fmt.Println(*ERROR.values())
}
