package mall_sys_user

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conf"
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	"github.com/pigjj/mall4micro/mall4micro-common/models"
	"gorm.io/gorm"
	"testing"
)

func init() {
	conf.ReloadConf("mall4micro-user")
}

func TestMallSysUser_SelectUserByUsername(t *testing.T) {
	type fields struct {
		MallBase models.MallBase
		Username string
		Password string
		Email    string
		Mobile   string
		Status   *int
		SaltStr  string
	}
	f1 := fields{}
	type args struct {
		tx       *gorm.DB
		username string
	}
	session, _ := conn.Conn()
	c1 := args{
		tx:       session,
		username: "user",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "c1", fields: f1, args: c1, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &MallSysUser{
				MallBase: tt.fields.MallBase,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Email:    tt.fields.Email,
				Mobile:   tt.fields.Mobile,
				Status:   tt.fields.Status,
				SaltStr:  tt.fields.SaltStr,
			}
			err := user.SelectUserByUsername(tt.args.tx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
