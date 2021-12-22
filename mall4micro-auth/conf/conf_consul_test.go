package conf

import "testing"

func TestAuthConf_Reload(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "c1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &AuthConf{}
			af.Reload()
		})
	}
}
