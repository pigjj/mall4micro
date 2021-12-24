package conf

import "testing"

func TestAuthConf_LoadConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "c1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := &MicroConf{}
			err := af.LoadConf()
			t.Logf("err: %+v, conf: %+v", err, af)
		})
	}
}
