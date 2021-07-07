package main

import (
	"fmt"
	"strconv"
)

// variadic parameter
func getGrades(grades ...int) (sum int, avg int) {
	gradeLength := len(grades)

	for i := 0; i < gradeLength; i++ {
		sum += grades[i]
	}

	avg = sum / gradeLength

	return
}

func getGradeInfo(name string, grades ...int) string {
	sum, avg := getGrades(grades...) // variadic argument

	return "Nilai sdr/i " + name + " adalah:\n" + "Total: " + strconv.Itoa(sum) +  "\n" + "Rata-rata: " + strconv.Itoa(avg)
}


func main() {
	gradeInfo := getGradeInfo("Nurul", 70, 60, 80, 85)
	fmt.Println(gradeInfo)
}
