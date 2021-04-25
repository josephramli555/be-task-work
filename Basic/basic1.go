package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Student struct {
	studentID  int
	name       string
	finalScore int
	grade      string
}

func printMenu(students []Student, passed int, failed int) {
	fmt.Println("=========================================================")
	fmt.Printf("%-6v %-12v %-15v %-15v %-10v \n", "No.", "Student ID", "Name", "Final Score", "Grade")
	fmt.Println("=========================================================")
	for idx, student := range students {
		fmt.Printf("%-6v %-12v %-15v %-15v %-10v \n", idx+1, student.studentID, student.name, student.finalScore, student.grade)
	}
	fmt.Println("=========================================================")
	fmt.Printf("%-30v: %d\n", "Number of Students", len(students))
	fmt.Printf("%-30v: %d\n", "Number of Passing Students", passed)
	fmt.Printf("%-30v: %d\n", "Number of Failed Students", failed)
}

func newStudent(id int, name string, finalScore int, grade string) Student {
	return Student{
		studentID:  id,
		name:       name,
		finalScore: finalScore,
		grade:      grade,
	}
}

func inputStudent() (Student, bool) {
	var id, finalScore int
	scoreMap := map[string]float32{"mid term": 0, "semester": 0, "attendance": 0}
	var inputData, name, grade string
	var passed bool
	var regex, _ = regexp.Compile("[a-zA-Z][a-zA-Z ]+")
	for {
		fmt.Printf("%-30v: ", "Input Student ID: ")
		fmt.Scanf("%s", &inputData)

		num, err := strconv.Atoi(inputData)
		if err != nil {
			fmt.Println("Student ID must be number!")
		} else {
			id = num
			break
		}
	}

	for {
		fmt.Printf("%-30v: ", "Input Student Name: ")
		fmt.Scanf("%s", &inputData)
		if !regex.MatchString(inputData) {
			fmt.Println("Name can only consists of Alphabet and space and minimum 2 letter")
		} else {
			name = inputData
			break
		}
	}
	for key, _ := range scoreMap {
		for {
			order := fmt.Sprintf("Input %s score: ", key)
			fmt.Printf("%-30v: ", order)
			fmt.Scanf("%s", &inputData)

			num, err := strconv.Atoi(inputData)
			if err != nil {
				fmt.Println("Score must be number!")
			} else if num < 0 || num > 100 {
				fmt.Println("Score must be between 0-100")
			} else {
				scoreMap[key] = float32(num)
				break
			}
		}
	}
	finalScore = int((0.2 * scoreMap["attendance"]) + (0.4 * scoreMap["mid term"]) + (0.4 * scoreMap["semester"]))

	switch {
	case finalScore >= 85 && finalScore <= 100:
		grade = "A"
	case finalScore >= 76 && finalScore <= 84:
		grade = "B"
	case finalScore >= 61 && finalScore <= 75:
		grade = "C"
	case finalScore >= 46 && finalScore <= 60:
		grade = "D"
	case finalScore >= 0 && finalScore <= 45:
		grade = "E"
	}

	if grade == "E" || grade == "D" {
		passed = false
	} else {
		passed = true
	}
	return newStudent(id, name, finalScore, grade), passed

}

func main() {
	students := []Student{}
	var passed, failed int = 0, 0
	var input string
	for {
		printMenu(students, passed, failed)
		for {
			fmt.Printf("Enter number of Inputted Student : ")
			fmt.Scanf("%s", &input)
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Input Number Only!")
			} else {
				for i := 0; i < num; i++ {
					student, pass := inputStudent()
					if pass {
						passed++
					} else {
						failed++
					}
					students = append(students, student)
					fmt.Println("Student Succesfully Inserted")
				}
				break
			}
		}
	}

}
