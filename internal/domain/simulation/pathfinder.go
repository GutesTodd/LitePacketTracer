// Package simulation
package simulation

import (
	"errors"
	"slices"

	"litepackettracer/internal/domain/topology"
)

type BFSPathFinder struct{}

func (f *BFSPathFinder) FindPath(lab *topology.Lab, from, to topology.DeviceID) ([]topology.DeviceID, error) {
	if lab == nil {
		return nil, errors.New("lab is nil")
	}
	if !lab.HasDevice(from) {
		return nil, errors.New("from does not exists")
	}
	if !lab.HasDevice(to) {
		return nil, errors.New("to does not exists")
	}
	if from == to {
		return []topology.DeviceID{from}, nil
	}
	queue := DeviceNewQueue()
	visited := map[topology.DeviceID]bool{}
	parent := map[topology.DeviceID]topology.DeviceID{}
	err := queue.PushBack(from)
	if err != nil {
		return nil, err
	}
	visited[from] = true
	for {
		current, ok := queue.PopFront()
		if !ok {
			break
		}
		neighbors, err := lab.Neighbors(current)
		if err != nil {
			return nil, err
		}
		for _, v := range neighbors {
			_, exists := visited[v.ID()]
			if !exists {
				id := v.ID()
				visited[id] = true
				parent[id] = current
				if id == to {
					return buildPath(from, to, parent), nil
				}
				err := queue.PushBack(id)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return nil, errors.New("path not found")
}

func buildPath(from, to topology.DeviceID, parent map[topology.DeviceID]topology.DeviceID) []topology.DeviceID {
	path := []topology.DeviceID{}
	current := to
	for current != from {
		path = append(path, current)
		current = parent[current]
	}
	path = append(path, from)
	slices.Reverse(path)
	return path
}
