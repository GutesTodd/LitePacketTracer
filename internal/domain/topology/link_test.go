package topology

import "testing"

func TestNewLink(t *testing.T) {
	tests := []struct {
		name    string
		a       DeviceID
		b       DeviceID
		wantErr bool
	}{
		{
			name:    "valid link",
			a:       "pc1",
			b:       "router1",
			wantErr: false,
		},
		{
			name:    "same endpoints",
			a:       "pc1",
			b:       "pc1",
			wantErr: true,
		},
		{
			name:    "blank a",
			a:       "",
			b:       "router1",
			wantErr: true,
		},
		{
			name:    "blank b",
			a:       "pc1",
			b:       "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link, err := NewLink(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if link.A() != tt.a {
				t.Fatalf("expected a %q, got %q", tt.a, link.A())
			}
			if link.B() != tt.b {
				t.Fatalf("expected b %q, got %q", tt.b, link.B())
			}
		})
	}
}

func TestLinkConnects(t *testing.T) {
	link, err := NewLink("pc1", "router1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if !link.Connects("pc1", "router1") {
		t.Fatal("expected link to connect pc1 and router1")
	}
	if !link.Connects("router1", "pc1") {
		t.Fatal("expected link to ignore endpoint order")
	}
	if link.Connects("pc1", "server1") {
		t.Fatal("expected link not to connect pc1 and server1")
	}
}

func TestLinkOther(t *testing.T) {
	link, err := NewLink("pc1", "router1")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	other, ok := link.Other("pc1")
	if !ok {
		t.Fatal("expected other endpoint for pc1")
	}
	if other != "router1" {
		t.Fatalf("expected router1, got %q", other)
	}

	other, ok = link.Other("server1")
	if ok {
		t.Fatal("expected no other endpoint for unrelated device")
	}
	if other != "" {
		t.Fatalf("expected empty endpoint, got %q", other)
	}
}

