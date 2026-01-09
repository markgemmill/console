package console

import (
	"fmt"

	"github.com/markgemmill/console/internal"
)

type LEVEL int

var (
	TRACE LEVEL = 0
	DEBUG LEVEL = 1
	WARN  LEVEL = 2
	INFO  LEVEL = 3
)

type Console struct {
	level LEVEL
}

func (c *Console) Level() LEVEL {
	return c.level
}

func (c *Console) SetLevel(level LEVEL) {
	c.level = level
}

func (c *Console) write(msg string, args ...any) {
	if internal.RequiresParsing(msg) {
		consoleText := internal.NewMessageParser(msg)
		consoleText.Parse()
		msg = consoleText.String()
	}

	fmt.Printf(msg, args...)
}

func (c *Console) Info(msg string, args ...any) {
	if c.level >= INFO {
		c.write(msg, args...)
	}
}

func (c *Console) Warn(msg string, args ...any) {
	if c.level >= WARN {
		c.write(msg, args...)
	}
}

func (c *Console) Debug(msg string, args ...any) {
	if c.level >= DEBUG {
		c.write(msg, args...)
	}
}

func (c *Console) Trace(msg string, args ...any) {
	if c.level >= TRACE {
		c.write(msg, args...)
	}
}

func (c *Console) Print(msg string, args ...any) {
	c.write(msg, args...)
}

func NewConsole(level LEVEL) *Console {
	return &Console{level: level}
}
