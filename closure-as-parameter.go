package main

import (
	"fmt"
	"strings"
)

func filter(data []string, callback func(string) bool) (result []string) {
	for _, val := range data {
		if filtered := callback(val); filtered {
			result = append(result, val)
		}
	}

	return
}

func main() {
	data := []string{"Pamungkas", "Bambang", "Winarko", "Jumadi"}

	fmt.Println("Aselinya:", data)

	dataContainsU := filter(data, func(val string) bool {
		return strings.Contains(strings.ToLower(val), "u")
	})

	fmt.Println("Contains u:", dataContainsU)

	filterContainsN := func(val string) bool {
		return strings.Contains(strings.ToLower(val), "n")
	}

	dataContainsN := filter(data, filterContainsN)

	fmt.Println("Contains n:", dataContainsN)
}