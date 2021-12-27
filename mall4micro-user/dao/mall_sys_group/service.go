package mall_sys_group

import "gorm.io/gorm"

func (s *MallSysGroup) Create(tx *gorm.DB) error {
	result := tx.Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysGroupList) SelectGroupByUserId(tx *gorm.DB, userId uint) error {
	result := tx.Model(MallSysGroupList{}).Where("is_deleted = 0 AND user_id = ?", userId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
