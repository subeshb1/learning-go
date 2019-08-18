// CH002 :: Pythagorean Triplet
package main

import "fmt"

func pythagoreanTriplet(number int) (int, int, int) {
	for first := 1; first <= number/2; first++ {
		for second := first + 1; second <= number/2; second++ {
			third := number - (first + second)
			if (first*first)+(second*second) == (third * third) {
				return first, second, third

			}
		}

	}

	return 0, 0, 0
}

func main() {
	fmt.Println(pythagoreanTriplet(1000)) // 200 375 425

}
