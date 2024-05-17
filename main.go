package main

import (
	"math/rand"
	"time"
)

func main() {
	fds := NewFoodStore()
	fds.UpsertItems(ItemSeeds...)
	fds.UpsertDiscountPromotions(PromotionSeeds()...)
	fds.UpsertMembers(MemberSeeds...)
	runFoodSimulator(fds)
}

func runFoodSimulator(fds *FoodStore) {
	items := fds.GetItems()
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderItemsFunc := func(amount int) []OrderItem {
		randomItems := getRandomFoodItems(items, amount)
		out := make([]OrderItem, 0, len(randomItems))
		for _, v := range randomItems {
			out = append(out, OrderItem{
				Item: Item{
					Name:  v.Name,
					Price: v.Price,
				},
				Amount: uint(rd.Intn(100)),
			})
		}
		return out
	}
	members := fds.GetMembers()
	orders := []*Order{
		NewOrder(orderItemsFunc(5), nil, fds),
		NewOrder(orderItemsFunc(3), getRandomMember(members), fds),
		NewOrder(orderItemsFunc(3), getRandomMember(members), fds),
	}

	for _, v := range orders {
		v.Checkout()
	}
}

// Function to get a random key-value pair from a map items
func getRandomFoodItems(items map[string]*Item, amount int) []*Item {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Convert map keys to a slice
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}

	out := make([]*Item, 0, amount)
	for range amount {
		// Get a random index
		randomIndex := rd.Intn(len(keys))
		// Select a random key
		randomKey := keys[randomIndex]
		out = append(out, items[randomKey])
	}

	return out
}

// Function to get a random key-value pair from a map members
func getRandomMember(members map[MemberID]*Member) *Member {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Convert map keys to a slice
	keys := make([]MemberID, 0, len(members))
	for key := range members {
		keys = append(keys, key)
	}

	// Get a random index
	randomIndex := rd.Intn(len(keys))
	// Select a random key
	randomKey := keys[randomIndex]

	return members[randomKey]
}
