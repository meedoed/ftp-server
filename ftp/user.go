package ftp

import (
	"fmt"
	"strings"
)

func (c *Conn) user(args []string) {
	c.respond(fmt.Sprintf(STATUS230, strings.Join(args, " ")))
}
