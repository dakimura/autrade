package strategy

// Result of an analysis
type Result struct{
	// directions about what actions you need to take
	directions []Direction
}


// Direction to buy/sell stock
type Direction struct{
	stockId int
	action Action
	price float64
}

type Action int
const (
	Buy Action = iota
	Sell
)
