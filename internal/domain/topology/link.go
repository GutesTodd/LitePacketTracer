// Package topology
package topology

import "errors"

type Link struct {
	a DeviceID
	b DeviceID
}

func NewLink(a, b DeviceID) (*Link, error) {
	if a == b {
		return nil, errors.New("links must be different")
	}
	if a == "" || b == "" {
		return nil, errors.New("links must not be blank")
	}
	return &Link{
		a: a,
		b: b,
	}, nil
}

func (l *Link) A() DeviceID {
	return l.a
}

func (l *Link) B() DeviceID {
	return l.b
}

func (l *Link) Connects(ida DeviceID, idb DeviceID) bool {
	return (l.a == ida && l.b == idb) || (l.a == idb && l.b == ida)
}

func (l *Link) Other(id DeviceID) (DeviceID, bool) {
	if l.a == id {
		return l.b, true
	}
	if l.b == id {
		return l.a, true
	}
	return "", false
}
