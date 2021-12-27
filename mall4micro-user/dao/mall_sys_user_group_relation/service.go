package mall_sys_user_group_relation

import "gorm.io/gorm"

func (s *MallSysUserGroupRelation) Insert(tx *gorm.DB) error {
	result := tx.Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysUserGroupRelationList) SelectRelationByUserId(tx *gorm.DB, userId uint) error {
	result := tx.Model(MallSysUserGroupRelationList{}).Where("is_deleted = 0 AND user_id = ?", userId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallSysUserGroupRelationList) SelectRelationByGroupId(tx *gorm.DB, groupId uint) error {
	result := tx.Model(MallSysUserGroupRelationList{}).Where("is_deleted = 0 AND group_id = ?", groupId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
