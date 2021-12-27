package mall_category

import "gorm.io/gorm"

func (s *MallCategory) Create(session *gorm.DB) error {
	result := session.Table(TableMallCategory).Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallCategoryList) SelectCategoryByShopId(session *gorm.DB, shopId uint) error {
	result := session.Table(TableMallCategory).Where("is_deleted = 0 AND shop_id = ?", shopId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallCategoryList) SelectCategoryByShopIds(session *gorm.DB, shopIds []uint) error {
	result := session.Table(TableMallCategory).Where("is_deleted = 0 AND shop_id IN ?", shopIds).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
