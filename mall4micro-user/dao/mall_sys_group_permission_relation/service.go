package mall_sys_group_permission_relation

import "gorm.io/gorm"

func (s *MallSysGroupPermissionRelation) Create(tx *gorm.DB) error {
	result := tx.Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysGroupPermissionRelationList) Create(tx *gorm.DB) error {
	result := tx.Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysGroupPermissionRelationList) SelectRelationByPermissionId(tx *gorm.DB, permissionId uint) error {
	result := tx.Model(MallSysGroupPermissionRelationList{}).Where("is_deleted = 0 AND permission_id = ?", permissionId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysGroupPermissionRelationList) SelectRelationByGroupIds(tx *gorm.DB, groupIds []uint) error {
	result := tx.Model(MallSysGroupPermissionRelationList{}).Where("is_deleted = 0 AND group_id IN ?", groupIds).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysGroupPermissionRelationList) SelectRelationByGroupId(tx *gorm.DB, groupId uint) error {
	result := tx.Model(MallSysGroupPermissionRelationList{}).Where("is_deleted = 0 AND group_id = ?", groupId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
