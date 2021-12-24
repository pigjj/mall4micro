package conf

import "testing"

func Test_loadLocalConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "c1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadLocalConf("mall4micro-user")
		})
	}
}
