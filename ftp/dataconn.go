package ftp

import (
	"fmt"
	"net"
)

type dataPort struct {
	h1, h2, h3, h4 int
	p1, p2         int
}

func dataPortFromHostPort(hostPort string) (*dataPort, error) {
	var dp dataPort
	_, err := fmt.Sscanf(hostPort, "%d,%d,%d,%d,%d,%d",
		&dp.h1, &dp.h2, &dp.h3, &dp.h4, &dp.p1, &dp.p2)
	if err != nil {
		return nil, err
	}
	return &dp, nil
}

func (d *dataPort) toAddress() string {
	if d == nil {
		return ""
	}
	port := d.p1<<8 + d.p2
	return fmt.Sprintf("%d.%d.%d.%d:%d", d.h1, d.h2, d.h3, d.h4, port)
}

func (c *Conn) dataConnect() (net.Conn, error) {
	conn, err := net.Dial("tcp", c.dataPort.toAddress())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
