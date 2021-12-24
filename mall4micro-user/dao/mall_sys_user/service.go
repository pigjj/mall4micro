package mall_sys_user

import "gorm.io/gorm"

func (user *MallSysUser) Insert(tx *gorm.DB) error {
	result := tx.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *MallSysUser) SelectUserByUsername(tx *gorm.DB, username string) error {
	result := tx.Model(MallSysUser{}).Where("is_deleted = 0 AND username = ?", username).First(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
