package mall_product

import "gorm.io/gorm"

func (s *MallProduct) Create(session *gorm.DB) error {
	result := session.Table(TableMallProduct).Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallProductList) SelectProductByShopId(session *gorm.DB, shopId uint) error {
	result := session.Table(TableMallProduct).Where("is_deleted = 0 AND shop_id = ?", shopId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *MallProductList) SelectProductByShopIds(session *gorm.DB, shopIds []uint) error {
	result := session.Table(TableMallProduct).Where("is_deleted = 0 AND shop_id IN ?", shopIds).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
