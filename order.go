package main

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type OrderID string

func NewOrderID() OrderID {
	return OrderID(ulid.Make().String())
}

type PromotionCenter interface {
	GetAvaialbleDiscountPromotions() map[string]*DiscountPromotion
}

type Order struct {
	ID                        OrderID
	Items                     []OrderItem
	AppliedDiscountPromotions map[string]*DiscountPromotion
	TotalPrice                float64
	promotionCenter           PromotionCenter
	Member                    *Member
}

func NewOrder(items []OrderItem, member *Member, pmc PromotionCenter) *Order {
	return &Order{
		ID:                        NewOrderID(),
		Items:                     items,
		AppliedDiscountPromotions: make(map[string]*DiscountPromotion),
		TotalPrice:                0,
		Member:                    member,
		promotionCenter:           pmc,
	}
}

type OrderItem struct {
	Item
	Amount uint
}

func (o *Order) Checkout() {
	discountPromotions := o.promotionCenter.GetAvaialbleDiscountPromotions()
	if o.AppliedDiscountPromotions == nil {
		o.AppliedDiscountPromotions = make(map[string]*DiscountPromotion)
	}
	totalPrice := 0.0
	for _, item := range o.Items {
		totalPrice += (item.Price * float64(item.Amount))
	}
	for _, promotion := range discountPromotions {
		if !promotion.isMatchCondition(o) {
			continue
		}
		_, ok := o.AppliedDiscountPromotions[promotion.Name]
		if !ok {
			o.AppliedDiscountPromotions[promotion.Name] = promotion
			continue
		}
	}

	totalDiscountPercentage := 0.0
	for _, promotion := range o.AppliedDiscountPromotions {
		totalDiscountPercentage += promotion.DiscountPercentage
	}
	discountAmount := totalPrice * (totalDiscountPercentage / 100)
	o.TotalPrice = totalPrice - discountAmount
	fmt.Printf("==================== OrderID: %s ====================\n", o.ID)
	fmt.Printf("Items: %+v\n", o.Items)
	fmt.Printf("Total: %v\n", totalPrice)
	fmt.Printf("AppliedDiscount: %+v\n", o.AppliedDiscountPromotions)
	fmt.Println("TotalDiscountPercentage: ", totalDiscountPercentage)
	fmt.Println("DiscountAmount: ", discountAmount)
	fmt.Printf("FinalTotal: %v\n", o.TotalPrice)
	fmt.Println("-----------------------------------------------------------------")
}
