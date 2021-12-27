package mall_shop

import "gorm.io/gorm"

//
// FindShopByUser
// @Description: 通过用户查找关联店铺
// @Document:
// @receiver s
// @param session
// @param userId
// @return error
//
func (s *MallShopList) FindShopByUser(session *gorm.DB, userId uint) error {
	result := session.Table(TableMallShop).Where("is_deleted = 0 AND user_id = ?", userId).Find(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//
// Create
// @Description: 创建商铺数据
// @Document:
// @receiver s
// @param session
// @return error
//
func (s *MallShop) Create(session *gorm.DB) error {
	result := session.Table(TableMallShop).Create(s)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
