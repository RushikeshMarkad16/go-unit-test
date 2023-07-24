package arithmatic

import "fmt"

func div(x, y float64) float64 {
	if y == 0 {
		return 0
	}
	ans := x / y
	fmt.Println("Answer is : ", ans)
	return ans
}
