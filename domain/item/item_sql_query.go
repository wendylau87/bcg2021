package item

const(
	ReadItemByID = `SELECT id, name, sku, price, quantity total FROM item WHERE id =?`
	ReadItemBySKU = `SELECT id, name, sku, price, quantity total FROM item WHERE sku =?`
)
