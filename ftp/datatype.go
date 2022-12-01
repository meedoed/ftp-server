package ftp

type dataType int

const (
	ASCII dataType = iota
	BINARY
)

func (c *Conn) setDataType(args []string) {
	if len(args) == 0 {
		c.respond(STATUS501)
	}

	switch args[0] {
	case "A":
		c.dataType = ASCII
	case "I":
		c.dataType = BINARY
	default:
		c.respond(STATUS504)
		return
	}
	c.respond(STATUS200)
}
