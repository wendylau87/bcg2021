package item

import "github.com/wendylau87/bcg/entities"

func(d *domain) findItemByID(id int)(*entities.Item, error){
	result := entities.Item{}
	rows, err := d.SQLHandler.Query(ReadItemByID, id)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.Name, &result.Sku, &result.Price, &result.Quantity); err != nil {
			return nil, err
		}

	}

	return &result, nil
}

func(d *domain) findItemBySKU(sku string)(*entities.Item, error){
	result := entities.Item{}
	rows, err := d.SQLHandler.Query(ReadItemBySKU, sku)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err = rows.Scan(&result.ID, &result.Name, &result.Sku, &result.Price, &result.Quantity); err != nil {
			return nil, err
		}

	}

	return &result, nil
}