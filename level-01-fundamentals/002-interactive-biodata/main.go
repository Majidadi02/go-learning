package main

import "fmt"

func main() {
	var myName string
	var myUniversity string
	var myCourse string
	var myLanguage string
	var myCareer string

	fmt.Print("What is your name?")
	fmt.Scan(&myName)
	fmt.Print("Where do you study?")
	fmt.Scan(&myUniversity)
	fmt.Print("What is your course?")
	fmt.Scan(&myCourse)
	fmt.Print("What is your preferred programming language?")
	fmt.Scan(&myLanguage)
	fmt.Print("What is your career goal?")
	fmt.Scan(&myCareer)

	fmt.Println("====================================")
	fmt.Println("       STUDENT PROFILE")
	fmt.Println("====================================")
	fmt.Println("Your name is: ", myName)
	fmt.Println("You study at: ", myUniversity)
	fmt.Println("Your course is: ", myCourse)
	fmt.Println("Your preferred programming language is: ", myLanguage)
	fmt.Println("Your career goal is: ", myCareer)
	fmt.Println("Thank you for using this Program!")
}
