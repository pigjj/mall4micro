package services

import (
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/conf"
	"github.com/jianghaibo12138/mall4micro/mall4micro-common/response"
	"github.com/jianghaibo12138/mall4micro/mall4micro-user/http_dto"
	"reflect"
	"testing"
)

func init() {
	conf.ReloadConf("mall4micro-user")
}

func TestRegisterSrv(t *testing.T) {
	type args struct {
		r *http_dto.RegisterDTO
	}
	c1 := args{
		r: &http_dto.RegisterDTO{
			Username: "admin",
			Password: "admin",
			Platform: 0,
		},
	}
	tests := []struct {
		name    string
		args    args
		want    *response.Response
		wantErr bool
	}{
		{name: "c1", args: c1, want: response.SuccessResponse, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RegisterSrv(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterSrv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterSrv() got = %v, want %v", got, tt.want)
			}
		})
	}
}
