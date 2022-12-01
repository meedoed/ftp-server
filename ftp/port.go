package ftp

import "log"

func (c *Conn) port(args []string) {
	if len(args) != 1 {
		c.respond(STATUS501)
		return
	}
	dataPort, err := dataPortFromHostPort(args[0])
	if err != nil {
		log.Print(err)
		c.respond(STATUS501)
		return
	}
	c.dataPort = dataPort
	c.respond(STATUS200)
}
