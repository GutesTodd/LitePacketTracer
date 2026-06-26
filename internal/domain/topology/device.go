// Package topology
package topology

import "errors"

type DeviceID string

type DeviceKind string

const (
	DevicePC     DeviceKind = "pc"
	DeviceRouter DeviceKind = "router"
	DeviceServer DeviceKind = "server"
)

type Device struct {
	id   DeviceID
	name string
	kind DeviceKind
}

func NewDevice(id DeviceID, name string, kind DeviceKind) (*Device, error) {
	if id == "" {
		return nil, errors.New("device id must not be blank")
	}

	if name == "" {
		return nil, errors.New("device name must not be blank")
	}

	if !kind.isValid() {
		return nil, errors.New("device kind is invalid")
	}
	return &Device{
		id:   id,
		name: name,
		kind: kind,
	}, nil
}

func (d *Device) ID() DeviceID {
	return d.id
}

func (d *Device) Name() string {
	return d.name
}

func (d *Device) Kind() DeviceKind {
	return d.kind
}

func (k DeviceKind) isValid() bool {
	if k == DevicePC || k == DeviceRouter || k == DeviceServer {
		return true
	}
	return false
}
