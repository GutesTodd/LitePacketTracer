// Package topology
package topology

import (
	"errors"
)

type LabID string

type Lab struct {
	id      LabID
	name    string
	devices map[DeviceID]*Device
	links   []*Link
}

func NewLab(id LabID, name string) (*Lab, error) {
	if id == "" {
		return nil, errors.New("id must not be blank")
	}
	if name == "" {
		return nil, errors.New("name must not be blank")
	}
	return &Lab{
		id:      id,
		name:    name,
		devices: make(map[DeviceID]*Device),
		links:   make([]*Link, 0),
	}, nil
}

func (l *Lab) ID() LabID {
	return l.id
}

func (l *Lab) Name() string {
	return l.name
}

func (l *Lab) Device(id DeviceID) (*Device, error) {
	if v, exists := l.devices[id]; exists {
		return v, nil
	}
	return nil, errors.New("Device does not exists")
}

func (l *Lab) AddDevice(d *Device) error {
	if d == nil {
		return errors.New("device is nil")
	}
	if _, exists := l.devices[d.id]; exists {
		return errors.New("device already exists")
	}
	l.devices[d.id] = d
	return nil
}

func (l *Lab) HasDevice(id DeviceID) bool {
	_, exists := l.devices[id]
	return exists
}

func (l *Lab) Connect(a, b DeviceID) error {
	if a == "" {
		return errors.New("a DeviceID must not be blank")
	}
	if b == "" {
		return errors.New("b DeviceID must not be blank")
	}
	if _, exists := l.devices[a]; !exists {
		return errors.New("device a not found")
	}
	if _, exists := l.devices[b]; !exists {
		return errors.New("device b not found")
	}
	for _, link := range l.links {
		if link.Connects(a, b) {
			return errors.New("link already exists")
		}
	}
	link, err := NewLink(a, b)
	if err != nil {
		return err
	}
	l.links = append(l.links, link)
	return nil
}

// Neighbors scans all links, so it is O(E) per call.
// For larger topologies, store an adjacency map in Lab.
func (l *Lab) Neighbors(id DeviceID) ([]*Device, error) {
	devices := []*Device{}
	if !l.HasDevice(id) {
		return nil, errors.New("device does not exist in lab")
	}
	for _, v := range l.links {
		if tid, exists := v.Other(id); exists {
			if l.HasDevice(tid) {
				devices = append(devices, l.devices[tid])
			}
		}
	}

	return devices, nil
}
