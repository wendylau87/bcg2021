package promotion
const(
	ReadPromotionByID = `SELECT id, name, minimum_item, minimum_item_discount, count_origin, count_destination, free_item FROM promotion WHERE id =?`
	ReadPromotionItemByID = `SELECT id, promotion_id, item_id, free_item_id, is_active, FROM promotion_item WHERE id =?`
	ReadAllPromotionItemByItemID = `SELECT id, promotion_id, item_id, free_item_id, is_active, priority FROM promotion_item WHERE is_active = 1 and item_id = ? order by priority desc`
)
