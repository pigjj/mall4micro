package mall_sys_permission

import "gorm.io/gorm"

func (s *MallSysPermission) Create(tx *gorm.DB) error {
	result := tx.Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysPermissionList) SelectPermissionByIds(tx *gorm.DB, ids []uint) error {
	result := tx.Table(TableMallSysPermission).Where("is_deleted = 0 AND id IN ?", ids).Find(s)
	return result.Error
}
