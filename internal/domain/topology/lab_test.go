package topology

import "testing"

func TestNewLab(t *testing.T) {
	lab, err := NewLab("lab1", "Main lab")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if lab.ID() != "lab1" {
		t.Fatalf("expected lab id %q, got %q", LabID("lab1"), lab.ID())
	}
	if lab.Name() != "Main lab" {
		t.Fatalf("expected lab name %q, got %q", "Main lab", lab.Name())
	}
	if lab.devices == nil {
		t.Fatal("expected devices map to be initialized")
	}
	if lab.links == nil {
		t.Fatal("expected links slice to be initialized")
	}
	if len(lab.devices) != 0 {
		t.Fatalf("expected no devices, got %d", len(lab.devices))
	}
	if len(lab.links) != 0 {
		t.Fatalf("expected no links, got %d", len(lab.links))
	}
}

func TestLabAddDevice(t *testing.T) {
	lab := newTestLab(t)
	device := newTestDevice(t, "pc1", "PC 1", DevicePC)

	if err := lab.AddDevice(device); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if !lab.HasDevice("pc1") {
		t.Fatal("expected lab to have pc1")
	}
}

func TestLabAddDeviceRejectsNil(t *testing.T) {
	lab := newTestLab(t)

	if err := lab.AddDevice(nil); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestLabAddDeviceRejectsDuplicate(t *testing.T) {
	lab := newTestLab(t)
	device := newTestDevice(t, "pc1", "PC 1", DevicePC)

	if err := lab.AddDevice(device); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err := lab.AddDevice(device); err == nil {
		t.Fatal("expected duplicate device error, got nil")
	}
}

func TestLabConnect(t *testing.T) {
	lab := newTestLab(t)
	addTestDevice(t, lab, "pc1", "PC 1", DevicePC)
	addTestDevice(t, lab, "router1", "Router 1", DeviceRouter)

	if err := lab.Connect("pc1", "router1"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(lab.links) != 1 {
		t.Fatalf("expected 1 link, got %d", len(lab.links))
	}
	if !lab.links[0].Connects("pc1", "router1") {
		t.Fatal("expected lab link to connect pc1 and router1")
	}
}

func TestLabConnectRejectsMissingDevices(t *testing.T) {
	lab := newTestLab(t)
	addTestDevice(t, lab, "pc1", "PC 1", DevicePC)

	if err := lab.Connect("pc1", "router1"); err == nil {
		t.Fatal("expected missing device error, got nil")
	}
}

func TestLabConnectRejectsDuplicateLink(t *testing.T) {
	lab := newTestLab(t)
	addTestDevice(t, lab, "pc1", "PC 1", DevicePC)
	addTestDevice(t, lab, "router1", "Router 1", DeviceRouter)

	if err := lab.Connect("pc1", "router1"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err := lab.Connect("router1", "pc1"); err == nil {
		t.Fatal("expected duplicate link error, got nil")
	}
}

func TestLabNeighbors(t *testing.T) {
	lab := newTestLab(t)
	addTestDevice(t, lab, "pc1", "PC 1", DevicePC)
	addTestDevice(t, lab, "router1", "Router 1", DeviceRouter)
	addTestDevice(t, lab, "server1", "Server 1", DeviceServer)

	if err := lab.Connect("pc1", "router1"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err := lab.Connect("router1", "server1"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	neighbors, err := lab.Neighbors("router1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if len(neighbors) != 2 {
		t.Fatalf("expected 2 neighbors, got %d", len(neighbors))
	}

	got := map[DeviceID]bool{
		neighbors[0].id: true,
		neighbors[1].id: true,
	}
	if !got["pc1"] {
		t.Fatal("expected pc1 to be router1 neighbor")
	}
	if !got["server1"] {
		t.Fatal("expected server1 to be router1 neighbor")
	}
}

func TestLabNeighborsRejectsUnknownDevice(t *testing.T) {
	lab := newTestLab(t)

	neighbors, err := lab.Neighbors("pc1")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if neighbors != nil {
		t.Fatalf("expected nil neighbors, got %v", neighbors)
	}
}

func newTestLab(t *testing.T) *Lab {
	t.Helper()

	lab, err := NewLab("lab1", "Main lab")
	if err != nil {
		t.Fatalf("create test lab: %v", err)
	}

	return lab
}

func newTestDevice(t *testing.T, id DeviceID, name string, kind DeviceKind) *Device {
	t.Helper()

	device, err := NewDevice(id, name, kind)
	if err != nil {
		t.Fatalf("create test device: %v", err)
	}

	return device
}

func addTestDevice(t *testing.T, lab *Lab, id DeviceID, name string, kind DeviceKind) {
	t.Helper()

	if err := lab.AddDevice(newTestDevice(t, id, name, kind)); err != nil {
		t.Fatalf("add test device: %v", err)
	}
}

