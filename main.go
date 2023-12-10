package main

import "fmt"

func gradingStudents(grades []int) []int {
	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			continue
		}

		if grades[i] % 5 > 2 {
			grades[i] += 5 - (grades[i] % 5)
		}
	}

	return grades
}

func main() {
	// Read the grade from the user.
    // var grade int
    // fmt.Scanln(&grade)

    // // Round the grade.
    // roundedGrade := gradingStudents(grade)

    // Print the rounded grade.
    // fmt.Println("The rounded grade is:", roundedGrade)

	grades := []int{73, 67, 38, 33}

	fmt.Println(gradingStudents(grades))

	// for _, grade := range grades {
	// 	fmt.Println(gradingStudents(grade))
	// }
}