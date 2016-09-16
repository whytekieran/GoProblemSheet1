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

	//SOLUTION TO PROBLEM 5
	var fourNotEntered bool = true
	var word string
	var letterChangeCnt int = 1
	var loopControlCnt int = 1
	var mainSwitchCnt int = 0
	var permutations = make([]string, 0)
	var resultWord string
	var loop = true

	//Keep asking user for a four letter word until one is entered
	for fourNotEntered {
		fmt.Print("Enter 4 letter word: ")
		fmt.Scanln(&word)

		if(len(word) == 4){
			fourNotEntered = false
		}
	}

	//stored word will be used each time we start of. We can make each 6 combinations from each letter
	//So for example if word starts with 'f' we can make 6 combinations and keep f at the start 6x4=24
	//The program will start with a letter then switch the remaining three letters after it 6 times.

	//Steps
	//1 Store original word
	//2 Get the last three characters
	//3 Switch them six times, 1st-2nd, 1st-3rd, 1st-2nd etc ..each time appending them onto the starting letter
	//4 After six switches we grab the original word again and switch 1st-2nd letters
	//5 repeat steps 2 and 3
	//6 Again after the six switches grab original word, this time switch 1st-3rd letters
	//7 repeat steps 2 and 3
	//8 Again after the six switches grab original word, this time switch 1st-4th letters
	//9 repeat steps 2 and 3
	//All Done :)
	//(Each word is appended to a string slice as we go along)

	storedWord := []rune(word)
	lastThreeLetters := storedWord[1:4]
	firstLetter := storedWord[0:1]

	for loop {

		if letterChangeCnt == 1 {
			permutations = append(permutations, string(storedWord))
		} else {

			if letterChangeCnt % 2 == 0{

				lastThreeLetters = reverseFirstTwoLetters(string(lastThreeLetters))
				resultWord = string(firstLetter)
				resultWord += string(lastThreeLetters)
				permutations = append(permutations, resultWord)

			} else if letterChangeCnt % 2 == 1 {
				lastThreeLetters = reverseFirstAndThirdLetters(string(lastThreeLetters))
				resultWord = string(firstLetter)
				resultWord += string(lastThreeLetters)
				permutations = append(permutations, resultWord)
			}

		}

		if letterChangeCnt == 6 && mainSwitchCnt == 0 {
			storedWord = reverseFirstTwoLetters(string(storedWord))
			lastThreeLetters = storedWord[1:4]
			firstLetter = storedWord[0:1]

			mainSwitchCnt++
			letterChangeCnt = 0
		} else if letterChangeCnt == 6 && mainSwitchCnt == 1 {
			storedWord = reverseFirstAndThirdLetters(string(storedWord))
			lastThreeLetters = storedWord[1:4]
			firstLetter = storedWord[0:1]

			mainSwitchCnt++
			letterChangeCnt = 0
		} else if letterChangeCnt == 6 && mainSwitchCnt == 2 {
			storedWord = reverseFirstAndFourthLetters(string(storedWord))
			lastThreeLetters = storedWord[1:4]
			firstLetter = storedWord[0:1]

			letterChangeCnt = 0
		}

		letterChangeCnt++
		loopControlCnt++

		if loopControlCnt == 25 {
			loop = false
		}
	}

	fmt.Printf("Here are all combinations of the word '%s'", word)
	fmt.Println()
	fmt.Println(permutations)

	//END SOLUTION TO PROBLEM 5
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

func reverseFirstTwoLetters(word string) []rune {

	var switcher rune
	data := []rune(word)

	switcher = data[0]
	data[0] = data[1]
	data[1] = switcher

	return data
}

func reverseFirstAndThirdLetters(word string) []rune {

	var switcher rune
	data := []rune(word)

	switcher = data[0]
	data[0] = data[2]
	data[2] = switcher

	return data
}

func reverseFirstAndFourthLetters(word string) []rune {

	var switcher rune
	data := []rune(word)

	switcher = data[0]
	data[0] = data[3]
	data[3] = switcher

	return data
}