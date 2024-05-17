package main

type DiscountPromotion struct {
	Condition          func(order *Order) bool
	Name               string
	Description        string
	DiscountPercentage float64
}

func (dc *DiscountPromotion) isMatchCondition(order *Order) bool {
	return dc.Condition(order)
}

func PromotionSeeds() []*DiscountPromotion {
	isDoubleRedPinkGreen := func(order *Order) bool {
		greenCount, pinkCount, redCount := 0, 0, 0
		for _, v := range order.Items {
			switch v.Name {
			case "Green Coffee":
				greenCount++
			case "Pink Coffee":
				pinkCount++
			case "Red Coffee":
				redCount++
			}
		}
		isGreaterThanOne := func(counter int) bool {
			return counter > 1
		}
		return isGreaterThanOne(greenCount) || isGreaterThanOne(pinkCount) || isGreaterThanOne(redCount)
	}
	isMemberPriviledgeLevel1 := func(order *Order) bool {
		if order.Member == nil {
			return false
		}
		return order.Member.PriviledgeLevel == MemberPrivledgeLevel1
	}
	isMemberPriviledgeLevel2 := func(order *Order) bool {
		if order.Member == nil {
			return false
		}
		return order.Member.PriviledgeLevel == MemberPrivledgeLevel2
	}
	isMemberPriviledgeLevel3 := func(order *Order) bool {
		if order.Member == nil {
			return false
		}
		return order.Member.PriviledgeLevel == MemberPrivledgeLevel3
	}
	return []*DiscountPromotion{
		{
			Condition:          isDoubleRedPinkGreen,
			Name:               "DoubleRed|Pink|Green",
			Description:        "The order contains Red Coffee, Pink Coffee or Green Coffee more than 1",
			DiscountPercentage: 5,
		},
		{
			Condition:          isMemberPriviledgeLevel1,
			Name:               "MemberPriviledgeLevel-1",
			Description:        "The member priviledge level is 1, get the discount 10%",
			DiscountPercentage: 10,
		},
		{
			Condition:          isMemberPriviledgeLevel2,
			Name:               "MemberPriviledgeLevel-2",
			Description:        "The member priviledge level is 2, get the discount 12%",
			DiscountPercentage: 12,
		},
		{
			Condition:          isMemberPriviledgeLevel3,
			Name:               "MemberPriviledgeLevel-3",
			Description:        "The member priviledge level is 3, get the discount 15%",
			DiscountPercentage: 15,
		},
	}
}
