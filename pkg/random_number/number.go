package random_number

import (
	"fmt"
	"math/rand"
)

func RandomNumber(maxNumber int) string {
	if maxNumber <= 0 {
		maxNumber = 6
	}
	number := ""
	for i := 0; i < maxNumber; i++ {
		number = fmt.Sprintf("%s%d", number, rand.Int31n(10))
	}
	return number
}
