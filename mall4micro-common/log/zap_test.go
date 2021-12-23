package log

import "testing"

func TestCreateLogsDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "c1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createLogsDir()
			t.Log(got)
		})
	}
}
