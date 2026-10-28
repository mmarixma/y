package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burgers := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}

	switch order {
	case "burger":
		return burgers.preptime
	case "chips":
		return chips.preptime
	case "nuggets":
		return nuggets.preptime
	default:
		return 404
	}
}

// burgers 15
// chips 10
// nuggets 12
