package main

// CH006 :: Weird Number

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getDivisors(number int) []int {
	divisors := []int{}
	for i := 1; i <= number/2; i++ {
		if number%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func knapSack(weights []int, maxWeight int) int {
	table := make([][]int, len(weights)+1)
	for i := 0; i <= len(weights); i++ {
		table[i] = make([]int, maxWeight+1)
	}
	for i := 0; i <= len(weights); i++ {
		for w := 1; w <= maxWeight; w++ {
			if i == 0 || w == 0 {
				table[i][w] = 0
			} else if weights[i-1] <= w {

				table[i][w] = max(weights[i-1]+table[i-1][w-weights[i-1]], table[i-1][w])

			} else {
				table[i][w] = table[i-1][w]
			}
		}

	}
	return table[len(weights)][maxWeight]
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func isWeirdNumber(number int) bool {
	divisors := getDivisors(number)
	if sum(divisors...) <= number {
		return false
	}
	return knapSack(divisors, number) != number
}

func main() {
	for i := 1; i < 100000; i++ {
		if isWeirdNumber(i) {
			fmt.Println(i) //
		}
	}

}

// Result

// 70
// 836
// 4030
// 5830
// 7192
// 7912
// 9272
// 10430
// 10570
// 10792
// 10990
// 11410
// 11690
// 12110
// 12530
// 12670
// 13370
// 13510
// 13790
// 13930
// 14770
// 15610
// 15890
// 16030
// 16310
// 16730
// 16870
// 17272
// 17570
// 17990
// 18410
// 18830
// 18970
// 19390
// 19670
// 19810
// 20510
// 21490
// 21770
// 21910
// 22190
// 23170
// 23590
// 24290
// 24430
// 24710
// 25130
// 25690
// 26110
// 26530
// 26810
// 27230
// 27790
// 28070
// 28630
// 29330
// 29470
// 30170
// 30310
// 30730
// 31010
// 31430
// 31990
// 32270
// 32410
// 32690
// 33530
// 34090
// 34370
// 34930
// 35210
// 35630
// 36470
// 36610
// 37870
// 38290
// 38990
// 39410
// 39830
// 39970
// 40390
// 41090
// 41510
// 41930
// 42070
// 42490
// 42910
// 43190
// 43330
// 44170
// 44870
// 45010
// 45290
// 45356
// 45710
// 46130
// 46270
// 47110
// 47390
// 47810
// 48370
// 49070
// 49630
// 50330
// 50890
// 51310
// 51730
// 52010
// 52570
// 52990
// 53270
// 53830
// 54110
// 55090
// 55790
// 56630
// 56770
// 57470
// 57610
// 57890
// 58030
// 58730
// 59710
// 59990
// 60130
// 60410
// 61390
// 61670
// 61810
// 62090
// 63490
// 63770
// 64330
// 65030
// 65590
// 65870
// 66290
// 66710
// 67690
// 67970
// 68390
// 68810
// 69370
// 69790
// 70630
// 70910
// 71330
// 71470
// 72170
// 72310
// 72730
// 73430
// 73570
// 73616
// 74270
// 74410
// 74830
// 76090
// 76370
// 76510
// 76790
// 77210
// 77630
// 78190
// 78610
// 79030
// 80570
// 80710
// 81410
// 81970
// 82670
// 83090
// 83312
// 83510
// 84070
// 84910
// 85190
// 85610
// 86030
// 86170
// 86590
// 87430
// 88130
// 89390
// 89530
// 89810
// 90230
// 90370
// 90790
// 91070
// 91210
// 91388
// 91490
// 92330
// 92470
// 92890
// 95270
// 95690
// 96110
// 96670
// 97930
// 98630
// 99610
// 99890
