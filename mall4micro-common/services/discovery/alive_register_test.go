package discovery

import (
	"github.com/google/uuid"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/dto"
	"testing"
)

func TestNewAliveRegister(t *testing.T) {
	type args struct {
		d *dto.ConsulServiceDTO
	}
	c1 := args{
		&dto.ConsulServiceDTO{
			ID:      uuid.New().String(),
			Name:    "mall4micro-auth",
			Tags:    []string{"auth"},
			Address: "192.168.0.105",
			Port:    8080,
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
			Name:    "mall4micro-auth",
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
