package main

import "fmt"

func main() {
	num3 := []int{}
	var a1, a2, a3, count, max int

	num5 := digit5() //Creating the array of all 5 digit numbers

	for i, v := range num5 { //Main loop that determinator the starting point of tries

		count++

		if i == 0 { //Creating the first unuseable numbers
			a1 = f3(v)
			a2 = s3(v)
			a3 = t3(v)
			num3 = append(num3, digit3(a1, a2, a3)...)
		}

		for _, x := range num5 { //Second loop that compares the numbers are usuable or not
			if check(x, num3) == true { //Updating the unuseable numbers
				a1 = f3(x)
				a2 = s3(x)
				a3 = t3(x)
				num3 = append(num3, digit3(a1, a2, a3)...)
				count++

			}
		}

		if count > max {
			max = count
		}

		count, a1, a2, a3 = 0, 0, 0, 0
		num3 = num3[:0]
	}

	fmt.Println(max)

}

func digit5() []int {

	nums := []int{}

	var num, temp int

	for i := 1; i < 6; i++ {
		for j := 1; j < 6; j++ {
			if j != i {
				for k := 1; k < 6; k++ {
					if k != j && k != i {
						for l := 1; l < 6; l++ {
							if l != i && l != k && l != j {
								for a := 1; a < 6; a++ {
									if a != i && a != k && a != j && a != l {
										num = i*10000 + j*1000 + k*100 + l*10 + a
										nums = append(nums, num)
										temp++
									}
									num = 0
								}
							}
						}
					}
				}
			}
		}
	}

	return nums
} //Creates all 5 digit numbers and adds to an array

func digit3(a1 int, a2 int, a3 int) []int {
	num3 := []int{}

	num3 = append(num3, a1, a2, a3)

	return num3
} //Creates and updates unusable 3 digit numbers and adds to an array

func f3(v int) int {
	return (((v % 100000) / 10000) * 100) + ((v%10000)/1000)*10 + (v%1000)/100
} // Creates the unusable 3 digit numbers from the given number's first 3 digits

func s3(v int) int {
	return (((v % 10000) / 1000) * 100) + ((v%1000)/100)*10 + (v%100)/10
} // Creates the unusable 3 digit numbers from the given number's second 3 digits

func t3(v int) int {
	return (((v % 1000) / 100) * 100) + ((v%100)/10)*10 + (v%10)/1
} // Creates the unusable 3 digit numbers from the given number's third 3 digits

func check(v int, num3 []int) bool {
	a1 := f3(v)
	a2 := s3(v)
	a3 := t3(v)

	for _, x := range num3 {
		if a1 == x || a2 == x || a3 == x {
			return false
		}
	}
	return true
} //Checks the number is usable or not and return a boolean
