package net_address

import (
	"errors"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (a NetAddress) String() string {
	return "http://" + a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return ErrInvalidAddr
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil {
		return err
	}
	a.Host = hp[0]
	a.Port = port
	return nil
}

var ErrInvalidAddr = errors.New("need address in a form host:port")
