package ftp

import (
	"fmt"
	"log"
)

const (
	STATUS150 = "150 File status okay; about to open data connection."
	STATUS200 = "200 Command okay."
	STATUS220 = "220 Service ready for new user."
	STATUS221 = "221 Service closing control connection."
	STATUS226 = "226 Closing data connection. Requested file action successful."
	STATUS230 = "230 User %s logged in, proceed."
	STATUS425 = "425 Can't open data connection."
	STATUS426 = "426 Connection closed; transfer aborted."
	STATUS501 = "501 Syntax error in parameters or argument."
	STATUS502 = "502 Command not implemented."
	STATUS504 = "504 Command not implemented for that parameter."
	STATUS550 = "550 Requested action not taken. File unavailable."
)

func (c *Conn) respond(s string) {
	log.Print(">> ", s)
	_, err := fmt.Fprint(c.conn, s, c.EOL())
	if err != nil {
		log.Print(err)
	}
}

func (c *Conn) EOL() string {
	switch c.dataType {
	case ASCII:
		return "\n\r"
	case BINARY:
		return "\n"
	default:
		return "\n"
	}
}
