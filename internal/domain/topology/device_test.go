package topology

import "testing"

func TestNewDevice(t *testing.T) {
	tests := []struct {
		name    string
		id      DeviceID
		devName string
		kind    DeviceKind
		wantErr bool
	}{
		{
			name:    "valid pc",
			id:      "pc1",
			devName: "PC 1",
			kind:    DevicePC,
			wantErr: false,
		},
		{
			name:    "blank id",
			id:      "",
			devName: "PC 1",
			kind:    DevicePC,
			wantErr: true,
		},
		{
			name:    "blank name",
			id:      "pc1",
			devName: "",
			kind:    DevicePC,
			wantErr: true,
		},
		{
			name:    "invalid kind",
			id:      "pc1",
			devName: "PC 1",
			kind:    DeviceKind("switch"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			device, err := NewDevice(tt.id, tt.devName, tt.kind)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
			if device.id != tt.id {
				t.Fatalf("expected id %q, got %q", tt.id, device.id)
			}
			if device.name != tt.devName {
				t.Fatalf("expected name %q, got %q", tt.devName, device.name)
			}
			if device.kind != tt.kind {
				t.Fatalf("expected kind %q, got %q", tt.kind, device.kind)
			}
		})
	}
}

