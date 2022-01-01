package discovery

import (
	"github.com/google/uuid"
	"github.com/pigjj/mall4micro/mall4micro-common/dto"
	"testing"
)

func TestNewAliveRegister(t *testing.T) {
	type args struct {
		d *dto.ConsulServiceDTO
	}
	c1 := args{
		&dto.ConsulServiceDTO{
			ID:      uuid.New().String(),
			Name:    "mall4micro-user",
			Tags:    []string{"auth"},
			Address: "192.168.7.5",
			Port:    8080,
			ServiceCheck: dto.Check{
				DeregisterCriticalServiceAfter: "10s",
				Args: []string{
					"curl",
					"http://192.168.7.5:8080/api/auth/ping",
				},
				Interval: "2s",
				Timeout:  "10s",
			},
			Proxy: dto.Proxy{
				DestinationServiceName: "UserService",
				DestinationServiceId:   "UserService",
				LocalServiceAddress:    "192.168.7.5",
				LocalServicePort:       8081,
				Mode:                   "transparent",
				TransparentProxy:       dto.TransparentProxy{OutboundListenerPort: 22500},
				Config:                 map[string]interface{}{"foo": "bar"},
				Upstreams: []dto.Upstream{
					{DestinationType: "service", DestinationName: "db", LocalBindPort: 9191},
				},
				Expose: dto.Expose{
					Check: true,
					Paths: []dto.Path{
						{
							Path:          "/healthz",
							LocalPathPort: 8080,
							ListenerPort:  21500,
							Protocol:      "http2",
						},
					},
				},
			},
		},
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "c1", args: c1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAliveRegister(tt.args.d)
			got.Register()
			t.Logf("%+v", got)
		})
	}
}

func TestAliveRegister_Register(t *testing.T) {
	type fields struct {
		ConsulServiceDTO *dto.ConsulServiceDTO
	}
	f := fields{
		ConsulServiceDTO: &dto.ConsulServiceDTO{
			ID:      uuid.New().String(),
			Name:    "mall4micro-user",
			Tags:    []string{"auth"},
			Address: "192.168.0.105",
			Port:    8080,
			ServiceCheck: dto.Check{
				DeregisterCriticalServiceAfter: "90m",
				Args:                           []string{"curl", "http://192.168.0.105:8080/api/auth/ping"},
				Interval:                       "1s",
				Timeout:                        "10s",
			},
			ServiceWeights: dto.Weights{
				Passing: 10,
				Warning: 1,
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "c1", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &AliveRegister{
				ConsulServiceDTO: tt.fields.ConsulServiceDTO,
			}
			err := ar.Register()
			t.Logf("%+v", err)
		})
	}
}
