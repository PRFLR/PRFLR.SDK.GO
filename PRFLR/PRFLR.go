package PRFLR

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"time"
)

type Timer struct {
	Timer string
	start time.Time
}

var conn *net.UDPConn
var key, source string

func Init(dsn, src string) error {
	d, err := url.Parse(dsn)
	if err != nil || d.User == nil || len(d.Host) == 0 {
		return errors.New("Cannot parse PRFLR DSN")
	}
	serverIP, err2 := net.ResolveUDPAddr("udp", d.Host)
	if err2 != nil {
		return err2
	}
	key = d.User.Username()
	source = src
	conn, _ = net.DialUDP("udp", nil, serverIP)
	return nil
}

func New(timer string) *Timer {
	return &Timer{
		Timer: timer,
		start: time.Now(),
	}
}

func (p *Timer) End(info string) error {
	if conn == nil {
		return errors.New("PRFLR connection is nil")
	}
	dur := fmt.Sprintf("%.3f", millisecond(time.Since(p.start)))
	data := fmt.Sprintf("%.32s|%.32s|%.48s|%s|%.32s|%.32s\n", "-", source, p.Timer, dur, info, key)
	_, err := conn.Write([]byte(data))
	return err
}

func millisecond(d time.Duration) float64 {
	return float64(d/time.Millisecond) + float64(d%time.Millisecond)*1e-9
}
