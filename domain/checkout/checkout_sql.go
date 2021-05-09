package checkout

import "github.com/wendylau87/bcg/entities"

func(d *domain) createCheckout(v entities.CheckOut)(*entities.CheckOut, error){
	// NOTE: this is a transaction example.
	tx, err := d.SQLHandler.Begin()
	if err != nil {
		return nil, err
	}

	// Create Checkout
	result, err := tx.Exec(CreateCheckout, v.CustomerName, v.TotalPrice, v.CreatedAt.Time)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	v.ID = int(id)

	// Create Checkout Details
	for key, detail := range v.Details{
		result, err := tx.Exec(CreateCheckoutDetail, v.ID, detail.ItemID, detail.PromotionItemID, detail.TotalQuantity, detail.SubTotalPrice)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return nil, err
		}
		v.Details[key].ID = int(id);
	}


	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &v, nil
}
