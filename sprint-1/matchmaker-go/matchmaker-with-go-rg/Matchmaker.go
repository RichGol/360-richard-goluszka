/*Rich Goluszka
 *Eric Pogue
 *CPSC 360-1: MatchMaker with Go
 *1/26/2021*/
package main

import (
	"fmt"
	"math"
)

/*declare constants for:
 *-question target responses
 *-question weights
 *-exit statement thresholds*/

const qTar1 = 4
const qTar2 = 2
const qTar3 = 5
const qTar4 = 4
const qTar5 = 5

const qWt1 = 2
const qWt2 = 1
const qWt3 = 3
const qWt4 = 2
const qWt5 = 3

const maxCond = 90
const midCond = 75
const minCond = 65

func greeting(msg string) {
	//a function to format messages to the user in a standardized style
	border := "================================================================================"
	dispMsg := ""

	//center text to display in a nice UI
	padding := len(border) - len(msg)
	for i := 0; i < padding/2; i++ {
		dispMsg += " "
	}
	dispMsg += msg

	//display formatted message
	fmt.Println(dispMsg)
	fmt.Println(border)
}

func validate(answer int, err error) (int, error) {
	//a function to validate user input against a range and against type
	if err != nil { //user input was not an integer
		var clearline string
		fmt.Scanf("%s\n", &clearline) //clear input stream, reset error
		err = nil
		fmt.Println("\nThe responses are meant to be a number value")
	} else if answer < 1 || answer > 5 { //user input outside expected range
		fmt.Println("\nThe questions are meant to be rated from 1 to 5")
	} else {
		fmt.Println("\nI have no clue what you broke or how it broke")
	}

	fmt.Print("Please correct your response: ")
	_, err = fmt.Scanf("%d\n", &answer)

	return answer, err
}

func askQuestion(question string, target, weight int) int {
	//a function to simplify the process of asking questions to the user
	//initialize variable to hold user response
	var answer int = 0

	//Prompt Question
	fmt.Print(question)

	//get response and strip newline from STDIN
	_, err := fmt.Scanf("%d\n", &answer)

	//validate user input
	for answer < 1 || answer > 5 || err != nil {
		answer, err = validate(answer, err)
	}

	//calculate + display question compatability score
	answer = int(math.Abs(float64(answer - target)))
	fmt.Println("Compatability Score for Question:", answer)

	//calculate + display weighted question value
	answer *= weight
	fmt.Println("Weighted Score for Question:", answer)

	return answer
}

func main() {
	//display intro greeting
	fmt.Println("")
	greeting("MatchMaker with Go")
	greeting("Answer five questions to determine compatability")
	greeting("Rate each statement from 1 (Strongly Disagree) to 5 (Strongly Agree)")

	//initialize variable to track total weighted compatability score
	var totScore int = 0

	//question 1
	fmt.Println("")
	greeting("Question 1")
	totScore += askQuestion("Winter is the best season: ", qTar1, qWt1)

	//question 2
	fmt.Println("")
	greeting("Question 2")
	totScore += askQuestion("Movies are overrated: ", qTar2, qWt2)

	//question 3
	fmt.Println("")
	greeting("Question 3")
	totScore += askQuestion("Coding is fun: ", qTar3, qWt3)

	//question 4
	fmt.Println("")
	greeting("Question 4")
	totScore += askQuestion("Reading is more fun than Math: ", qTar4, qWt4)

	//question 5
	fmt.Println("")
	greeting("Question 5")
	totScore += askQuestion("Programs are not a great way to build relationships: ", qTar5, qWt5)

	//calculate + display final score
	totScore = 100 - totScore
	fmt.Println("")
	fmt.Printf("Overall Weighted Compatability Score: %d%%\n\n", totScore)

	//max possible score is 100
	//min possible score is 61
	if totScore >= maxCond {
		fmt.Println("Just one more question: Should we go to the movies or grab lunch?")
	} else if totScore >= midCond {
		fmt.Println("Let's just be facebook friends, cool?")
	} else if totScore >= minCond {
		fmt.Println("Our friendship threw a NullPointerException ... oops")
	} else {
		fmt.Println("How did you even score that low?")
	}
}
