package simulation

import (
	"testing"

	"litepackettracer/internal/domain/topology"
)

func TestQueuePopFrontReturnsFalseWhenEmpty(t *testing.T) {
	queue := DeviceNewQueue()

	id, ok := queue.PopFront()
	if ok {
		t.Fatal("expected empty queue pop to return false")
	}
	if id != "" {
		t.Fatalf("expected empty id, got %q", id)
	}
}

func TestQueuePushBackRejectsBlankID(t *testing.T) {
	queue := DeviceNewQueue()

	if err := queue.PushBack(""); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestQueuePushBackAndPopFront(t *testing.T) {
	queue := DeviceNewQueue()

	if err := queue.PushBack("pc1"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	id, ok := queue.PopFront()
	if !ok {
		t.Fatal("expected pop to return true")
	}
	if id != topology.DeviceID("pc1") {
		t.Fatalf("expected pc1, got %q", id)
	}

	id, ok = queue.PopFront()
	if ok {
		t.Fatal("expected second pop to return false")
	}
	if id != "" {
		t.Fatalf("expected empty id, got %q", id)
	}
}

func TestQueueIsFIFO(t *testing.T) {
	queue := DeviceNewQueue()
	ids := []topology.DeviceID{"pc1", "router1", "server1"}

	for _, id := range ids {
		if err := queue.PushBack(id); err != nil {
			t.Fatalf("push %q: %v", id, err)
		}
	}

	for _, want := range ids {
		got, ok := queue.PopFront()
		if !ok {
			t.Fatalf("expected %q, got empty queue", want)
		}
		if got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	}
}
