package main

type FoodStore struct {
	orders             map[OrderID]*Order
	items              map[string]*Item
	discountPromotions map[string]*DiscountPromotion
	members            map[MemberID]*Member
}

func NewFoodStore() *FoodStore {
	return &FoodStore{
		orders:             make(map[OrderID]*Order),
		items:              make(map[string]*Item),
		discountPromotions: make(map[string]*DiscountPromotion),
		members:            make(map[MemberID]*Member),
	}
}

func (fds *FoodStore) GetItems() map[string]*Item {
	return fds.items
}

func (fds *FoodStore) UpsertItems(items ...*Item) {
	for _, v := range items {
		fds.items[v.Name] = v
	}

}

func (fds *FoodStore) GetMembers() map[MemberID]*Member {
	return fds.members
}

func (fds *FoodStore) UpsertMembers(members ...*Member) {
	for _, v := range members {
		fds.members[v.ID] = v
	}
}

func (fds *FoodStore) UpsertDiscountPromotions(dc ...*DiscountPromotion) {
	for _, v := range dc {
		fds.discountPromotions[v.Name] = v
	}
}

func (fds *FoodStore) GetAvaialbleDiscountPromotions() map[string]*DiscountPromotion {
	return fds.discountPromotions
}
