package checkout

const(
	CreateCheckout=`INSERT INTO check_out(customer_name, total_price, created_at) VALUES(?,?,?)`
	CreateCheckoutDetail = `INSERT INTO check_out_detail(check_out_id,item_id,promotion_item_id, total_quantity, sub_total_price) VALUES(?,?,?,?,?)`
)
