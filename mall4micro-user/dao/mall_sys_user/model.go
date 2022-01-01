package mall_sys_user

import (
	"github.com/pigjj/mall4micro/mall4micro-common/conn"
	cm "github.com/pigjj/mall4micro/mall4micro-common/models"
)

//
// MallSysUser
// @Description: 商城用户表
//
type MallSysUser struct {
	cm.MallBase
	Username string `gorm:"column:username;size:50;comment:用户名;not null;index:idx_user_name,unique" json:"username"`
	Password string `gorm:"column:password;size:100;comment:密码;not null" json:"password"`
	Email    string `gorm:"column:email;size:50;comment:邮箱;default:null;index:idx_user_email,unique" json:"email"`
	Mobile   string `gorm:"column:mobile;size:50;comment:手机号;default:null;index:idx_user_mobile,unique" json:"mobile"`
	Status   *int   `gorm:"column:status;comment:状态,0：禁用,1：正常;default:null" json:"status"`
	SaltStr  string `gorm:"column:salt_str;size:100;comment:盐;not null" json:"-"`
}

const (
	TableMallSysUser = "mall_sys_user"
)

var (
	StatusDisable = 0
	StatusActive  = 1
)

func (MallSysUser) TableName() string {
	return TableMallSysUser
}

func (MallSysUser) Migrate() {
	session, err := conn.Conn()
	if err != nil {
		panic(err)
	}
	if session.Migrator().HasTable(&MallSysUser{}) {
		_ = session.AutoMigrate(&MallSysUser{})
	} else {
		_ = session.Migrator().CreateTable(&MallSysUser{})
	}
}
