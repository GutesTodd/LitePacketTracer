package simulation

import (
	"slices"
	"testing"

	"litepackettracer/internal/domain/topology"
)

func TestBFSPathFinderFindPathWhenFromEqualsTo(t *testing.T) {
	lab := newPathTestLab(t)
	addPathTestDevice(t, lab, "pc1", "PC 1", topology.DevicePC)

	path, err := new(BFSPathFinder).FindPath(lab, "pc1", "pc1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	want := []topology.DeviceID{"pc1"}
	if !slices.Equal(path, want) {
		t.Fatalf("expected path %v, got %v", want, path)
	}
}

func TestBFSPathFinderFindsDirectPath(t *testing.T) {
	lab := newPathTestLab(t)
	addPathTestDevice(t, lab, "pc1", "PC 1", topology.DevicePC)
	addPathTestDevice(t, lab, "router1", "Router 1", topology.DeviceRouter)
	connectPathTestDevices(t, lab, "pc1", "router1")

	path, err := new(BFSPathFinder).FindPath(lab, "pc1", "router1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	want := []topology.DeviceID{"pc1", "router1"}
	if !slices.Equal(path, want) {
		t.Fatalf("expected path %v, got %v", want, path)
	}
}

func TestBFSPathFinderFindsPathThroughRouter(t *testing.T) {
	lab := newPathTestLab(t)
	addPathTestDevice(t, lab, "pc1", "PC 1", topology.DevicePC)
	addPathTestDevice(t, lab, "router1", "Router 1", topology.DeviceRouter)
	addPathTestDevice(t, lab, "server1", "Server 1", topology.DeviceServer)
	connectPathTestDevices(t, lab, "pc1", "router1")
	connectPathTestDevices(t, lab, "router1", "server1")

	path, err := new(BFSPathFinder).FindPath(lab, "pc1", "server1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	want := []topology.DeviceID{"pc1", "router1", "server1"}
	if !slices.Equal(path, want) {
		t.Fatalf("expected path %v, got %v", want, path)
	}
}

func TestBFSPathFinderReturnsErrorWhenPathNotFound(t *testing.T) {
	lab := newPathTestLab(t)
	addPathTestDevice(t, lab, "pc1", "PC 1", topology.DevicePC)
	addPathTestDevice(t, lab, "server1", "Server 1", topology.DeviceServer)

	path, err := new(BFSPathFinder).FindPath(lab, "pc1", "server1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if path != nil {
		t.Fatalf("expected nil path, got %v", path)
	}
}

func TestBFSPathFinderReturnsErrorForUnknownFrom(t *testing.T) {
	lab := newPathTestLab(t)
	addPathTestDevice(t, lab, "server1", "Server 1", topology.DeviceServer)

	path, err := new(BFSPathFinder).FindPath(lab, "pc1", "server1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if path != nil {
		t.Fatalf("expected nil path, got %v", path)
	}
}

func TestBFSPathFinderReturnsErrorForUnknownTo(t *testing.T) {
	lab := newPathTestLab(t)
	addPathTestDevice(t, lab, "pc1", "PC 1", topology.DevicePC)

	path, err := new(BFSPathFinder).FindPath(lab, "pc1", "server1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if path != nil {
		t.Fatalf("expected nil path, got %v", path)
	}
}

func TestBFSPathFinderReturnsErrorForNilLab(t *testing.T) {
	path, err := new(BFSPathFinder).FindPath(nil, "pc1", "server1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if path != nil {
		t.Fatalf("expected nil path, got %v", path)
	}
}

func newPathTestLab(t *testing.T) *topology.Lab {
	t.Helper()

	lab, err := topology.NewLab("lab1", "Main lab")
	if err != nil {
		t.Fatalf("create lab: %v", err)
	}

	return lab
}

func addPathTestDevice(
	t *testing.T,
	lab *topology.Lab,
	id topology.DeviceID,
	name string,
	kind topology.DeviceKind,
) {
	t.Helper()

	device, err := topology.NewDevice(id, name, kind)
	if err != nil {
		t.Fatalf("create device %q: %v", id, err)
	}

	if err := lab.AddDevice(device); err != nil {
		t.Fatalf("add device %q: %v", id, err)
	}
}

func connectPathTestDevices(t *testing.T, lab *topology.Lab, a, b topology.DeviceID) {
	t.Helper()

	if err := lab.Connect(a, b); err != nil {
		t.Fatalf("connect %q and %q: %v", a, b, err)
	}
}

