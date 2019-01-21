package pricedata

import "math/rand"

// createRandomData creates randomwalking values and return it
// need to initialize the seed like as follows: rand.Seed(time.Now().UnixNano())
func CreateRandomData(initial, n int) []int {
	data := make([]int, n)
	value := initial
	for i := 0; i < n; i++ {
		data[i] = value
		// rand.Intn(3) - 1 => {-1, 0, 1}
		value += rand.Intn(3) - 1
	}
	return data
}
