package main

import (
	"fmt"
)

func main(){

	//UNCOMMENT SOLUTIONS 1 AT A TIME TO RUN

	//SOLUTION TO PROBLEM 1
	/*var sum int
	  var i int

	for i = 0; i <= 1000; i++ { //loop from 0 to 1000
		if i % 3 == 0 || i % 5 == 0 { //if number is multiple of 3 or 5

			sum += i; //add it to the sum

		}
	}

	fmt.Printf("Sum is %d", sum)*/

	//END SOLUTION TO PROBLEM 1

	//SOLUTION TO PROBLEM 2

	/*var primeCount int = 0
	var i int = 1
	var j int
	var prime10001 int
	var keepLooping bool = true

	for keepLooping { //continues loop
		i++
		if (i % i == 0) && (i % 1 == 0){ //check for prime number candidate

			if i == 2 {
				//increment prime number count as 2 is prime, its the only exception for even numbers
				primeCount++
			} else { //otherwise we must check

				//check if prime number candidate can be divided by any number between 2 and the number below it
				for j = 2; j < i; j++ {

					if i % j == 0{//if any of those numbers can its not a prime number
						break
					} else if j == i-1  { //if we have checked all numbers it is
						primeCount++ //increment prime number count

						if primeCount == 10001 { //when it reaches 10001 we have what we need
							prime10001 = i
							keepLooping = false
						}
					} else {
						continue
					}

				}
			}
		}
	}

	fmt.Printf("The 10001st prime number is %d", prime10001)*/
	//END SOLUTION TO PROBLEM 2

	//SOLUTION TO PROBLEM 3
	//Uncomment reverse function below func main()
	//We use runes in the reverse function to reverse a string. A rune slice takes each character
	//as an individual element so we can loop through them
	//Runes can convert characters to their ascii value, like so:
	/*var h rune = 'a'
	fmt.Printf("%d", h)
	fmt.Println()
	fmt.Println()


	word1 := "GoProgramming"
	reversed1 := reverse(word1)
	fmt.Println(word1)
	fmt.Println(reversed1)

	word2 := "Rune"
	reversed2 := reverse(word2)
	fmt.Println(word2)
	fmt.Println(reversed2)*/
	//END SOLUTION TO PROBLEM 3

	//SOLUTION TO PROBLEM 4
	/*var resolved int = 0
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	//keep looping until resolved is 1
	for resolved != 1 {

		if number % 2 == 0 {
			number = number / 2
			if number == 1 { //when number is 1 we will always get repeat of 4 2 1, 4, 2, 1 so just stop
				resolved = 1
			}

		} else if number % 2 == 1 {
			number = (number * 3) + 1
			fmt.Println("odd")
			fmt.Println(number)
		}
	}

	fmt.Println("All done")*/
	//END SOLUTION TO PROBLEM 4
}

//reverse function used to solve problem 3
/*func reverse(word string) string {

	// Convert string to rune slice.
	data := []rune(word)
	result := []rune{}

	// Add runes in reverse order.
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}

	// Return new string.
	return string(result)
}*/

