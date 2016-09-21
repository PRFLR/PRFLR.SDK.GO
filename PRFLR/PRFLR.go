package PRFLR

import (
	"net/url"
	"errors"
	"fmt"
	"net"
	"time"
)

type Timer struct {
	Timer  string
	start  time.Time
}

var conn   	*net.UDPConn 
var key 	string

func Init(dsn, src string) error {
	key, host, err = parseDSN(dsn)
	d, err := url.Parse(dsn)
    	if err != nil || d.User.Username() == nil || len(d.Host) == 0 {
        	return errors.New("Cannot parse PRFLR DSN")
    	}
	server, err2 := net.ResolveUDPAddr("udp", d.Host)
	if err2 != nil {
		return err2
	}
	key = d.User.Username()
	conn = net.DialUDP("udp", nil, server)
	return nil
}

func New(timer string) *Timer {
	return &Timer{
		Timer:  timer,
		start:  time.Now(),
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
	return float64( d/time.Millisecond ) + float64( d%time.Millisecond )*1e-9
}

func getConnection() (*net.UDPConn, error) {
	if len(host) == 0 || len(key) == 0 {
		return nil, errors.New("PRFLR Host/Key is not specified. Please call PRFLR.Init() BEFORE sending timers!")
	}
	serverAddr, err  := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return nil, err
	}
	return net.DialUDP("udp", nil, serverAddr)
}
