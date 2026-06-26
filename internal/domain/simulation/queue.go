// Package simulation
package simulation

import (
	"errors"

	"litepackettracer/internal/domain/topology"
)

type LengthQueue int

type queueNode struct {
	value topology.DeviceID
	next  *queueNode
	prev  *queueNode
}

type queueDevice struct {
	head *queueNode
	tail *queueNode
}

func DeviceNewQueue() *queueDevice {
	return &queueDevice{
		head: nil,
		tail: nil,
	}
}

func (q *queueDevice) PushBack(id topology.DeviceID) error {
	if id == "" {
		return errors.New("id must not be blank")
	}
	node := &queueNode{
		value: id,
		next:  nil,
		prev:  nil,
	}
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		node.prev = q.tail
		q.tail = node
	}
	return nil
}

func (q *queueDevice) PopFront() (topology.DeviceID, bool) {
	var v topology.DeviceID
	if q.head == nil && q.tail == nil {
		return "", false
	}
	if q.head == q.tail {
		v = q.tail.value
		q.tail = nil
		q.head = nil
	} else {
		v = q.head.value
		q.head = q.head.next
		q.head.prev = nil
	}
	return v, true
}
