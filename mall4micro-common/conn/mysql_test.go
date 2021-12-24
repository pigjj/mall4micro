package conn

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"testing"
)

func TestConn(t *testing.T) {
	conf.ReloadConf(routers.MicroServiceName)
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "c1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Conn()
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("%+v", got)
		})
	}
}
