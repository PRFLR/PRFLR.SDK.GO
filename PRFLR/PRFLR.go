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

var host   string
var key    string
var source string

func New(timer string) *Timer {
	return &Timer{
		Timer:  timer,
		start:  time.Now(),
	}
}
func Init(dsn, src string) (err error) {
	key, host, err = parseDSN(dsn)
	source = src
	return
}

func (p *Timer) End(info string) error {
	dur := fmt.Sprintf("%.3f", millisecond(time.Since(p.start)))
	
	conn, err := getConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	data := fmt.Sprintf("%.32s|%.32s|%.48s|%s|%.32s|%.32s\n", "0", source, p.Timer, dur, info, key)

	_, err = conn.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

func millisecond(d time.Duration) float64 {
	msec := d / time.Millisecond
	nsec := d % time.Millisecond
	return float64(msec) + float64(nsec)*1e-9
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

func parseDSN(dsn string) (key, host string, err error) {
    d, err := url.Parse(dsn)
    if err != nil || d.User == nil || len(d.Host) == 0 {
        return "", "", errors.New("Cannot parse PRFLR DSN")
    }

    return d.User.Username(), d.Host, nil
}
