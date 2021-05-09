package checkout

import (
	"encoding/json"
	"github.com/wendylau87/bcg/entities"
	"github.com/wendylau87/bcg/infrastructure/sqlx"
	"time"
)

func (u *usecase) CreateCheckout(v entities.CreateCheckOut) (*entities.CheckOutResp, error) {
	//Convert Object
	obj := entities.CheckOut{
		ID:           0,
		CustomerName: v.CustomerName,
		TotalPrice:   0,
		CreatedAt:    sqlx.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	freeItems := make(map[int]int)
	itemNames := make(map[int]string)
	promotionNames := make(map[int]string)
	totalprice := 0.0

	for _, createCheckoutDetail := range v.Details{
		item, err := u.itemdom.GetItemBySku(createCheckoutDetail.SKU)
		if err != nil{
			u.logger.LogError("%s", err)
			return nil, err
		}
		itemNames[item.ID] = item.Name

		// Check promotion
		promos, err := u.promotiondom.GetAllPromotionItemByItemID(item.ID)
		if err != nil{
			u.logger.LogError("%s", err)
			return nil, err
		}
		promotionItemID := 0
		if len(promos) > 0{
			promotionItemID = promos[0].ID
		}

		// Calculate price
		subTotalPrice := 0.0
		if promotionItemID != 0 {

			promotion, err := u.promotiondom.GetPromotionByID(promos[0].PromotionID)
			if err != nil{
				u.logger.LogError("%s", err)
				return nil, err
			}
			promotionNames[promos[0].ID]= promotion.Name
			u.logger.LogAccess("Check Promotion for item %s is %s", item.Name, promotion.Name)

			if promotion.MinimumItem >= 1 && promotion.MinimumItem <= int64(createCheckoutDetail.TotalQuantity){
				//buy 3 discount 10%
				subTotalPrice = float64(createCheckoutDetail.TotalQuantity) * item.Price - (float64(createCheckoutDetail.TotalQuantity) * item.Price * (promotion.MinimumItemDiscount/100))
			}else if promotion.CountOrigin >= 1 && promotion.CountOrigin <= int64(createCheckoutDetail.TotalQuantity){
				//buy 3 pay 2, buy 6 pay 4
				multiplier := createCheckoutDetail.TotalQuantity/int(promotion.CountOrigin)
				subTotalPrice = float64(promotion.CountDestination) * item.Price * float64(multiplier) + (item.Price * float64(createCheckoutDetail.TotalQuantity % int(promotion.CountOrigin)))
			}else if promotion.FreeItem{
				// free item
				freeItems[promos[0].FreeItemID]++
				subTotalPrice = float64(createCheckoutDetail.TotalQuantity) * item.Price
			}else{
				subTotalPrice = float64(createCheckoutDetail.TotalQuantity) * item.Price
			}
		} else{
			u.logger.LogAccess("Item %s doesn't have promo", item.Name)
			subTotalPrice = float64(createCheckoutDetail.TotalQuantity) * item.Price
		}
		objDetail := entities.CheckOutDetail{
			ID:              0,
			CheckOutID:      0,
			ItemID:          item.ID,
			PromotionItemID: promotionItemID,
			TotalQuantity:   createCheckoutDetail.TotalQuantity,
			SubTotalPrice:   subTotalPrice,
		}
		totalprice+= subTotalPrice
		obj.Details = append(obj.Details, objDetail)

	}

	//recalculate free item
	for key, detail := range obj.Details{
		if _, ok := freeItems[detail.ItemID]; ok {
			itemPrice := detail.SubTotalPrice / float64(detail.TotalQuantity)
			obj.Details[key].SubTotalPrice = detail.SubTotalPrice - (float64(freeItems[detail.ItemID]) * itemPrice)
			totalprice-= float64(freeItems[detail.ItemID]) * itemPrice
		}
	}
	obj.TotalPrice = totalprice
	jsonObj,_ := json.Marshal(obj)
	u.logger.LogAccess("Obj: %s", string(jsonObj))

	checkout, err := u.checkoutdom.CreateCheckout(obj)
	if err != nil{
		u.logger.LogError("%s", err)
		return nil, err
	}
	result := entities.CheckOutResp{
		ID:           checkout.ID,
		CustomerName: checkout.CustomerName,
		TotalPrice:   checkout.TotalPrice,
		CreatedAt:    checkout.CreatedAt,
	}

	for _, detail := range checkout.Details{
		respDetail := entities.CheckOutDetailResp{
			ItemName:      itemNames[detail.ItemID],
			PromotionName: promotionNames[detail.PromotionItemID],
			TotalQuantity: detail.TotalQuantity,
			SubTotalPrice: detail.SubTotalPrice,
		}
		result.Details = append(result.Details, respDetail)
	}

	return &result, nil

}
