package promotion

import "github.com/wendylau87/bcg/entities"

func(d *domain) findPromotionByID(id int)(*entities.Promotion, error){
	result := entities.Promotion{}
	rows, err := d.SQLHandler.Query(ReadPromotionByID, id)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.MinimumItem,
			&result.MinimumItemDiscount,
			&result.CountOrigin,
			&result.CountDestination,
			&result.FreeItem); err != nil {
			return nil, err
		}

	}

	return &result, nil
}

func(d *domain) findPromotionItemByID(id int)(*entities.PromotionItem, error){
	result := entities.PromotionItem{}
	rows, err := d.SQLHandler.Query(ReadPromotionItemByID, id)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(
			&result.ID,
			&result.PromotionID,
			&result.ItemID,
			&result.FreeItemID,
			&result.IsActive,
			&result.Priority); err != nil {
			return nil, err
		}

	}

	return &result, nil
}

func(d *domain) findAllPromotionItemByItemID(id int)([]entities.PromotionItem, error){
	promos := []entities.PromotionItem{}
	rows, err := d.SQLHandler.Query(ReadAllPromotionItemByItemID, id)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var result entities.PromotionItem
		if err = rows.Scan(
			&result.ID,
			&result.PromotionID,
			&result.ItemID,
			&result.FreeItemID,
			&result.IsActive,
			&result.Priority); err != nil {
			return nil, err
		}
		promos = append(promos, result)
	}

	return promos, nil
}
