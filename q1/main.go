package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	testValue := []string{"3.9.5", "4.3.11", "8.1.2", "4.3.6", "4.5.6"}
	//testValue := []string{"1.0.1", "2.14.15"}

	// j - first element, i - second element
	sort.SliceStable(testValue, func(secondElement, firstElement int) bool {
		v1 := testValue[firstElement]
		v2 := testValue[secondElement]

		vs1 := strings.Split(v1, ".")
		vs2 := strings.Split(v2, ".")

		// Convert the string value to integer
		// and split them up into major, minor and patch
		// version. To make it easy to compare
		major1, _ := strconv.Atoi(vs1[0])
		minor1, _ := strconv.Atoi(vs1[1])
		patch1, _ := strconv.Atoi(vs1[2])

		major2, _ := strconv.Atoi(vs2[0])
		minor2, _ := strconv.Atoi(vs2[1])
		patch2, _ := strconv.Atoi(vs2[2])

		// since SliceStable(..) is like Less(..) we do comparison
		// the other way around. Compared the first value element
		// with second value element
		if major1 < major2 {
			return true
		}

		// major version is the same, we need to check minor
		if major1 == major2 {
			// minor version is the same, we need to check patch
			if minor1 == minor2 {

				// patch the same then just return true, does not make
				// any difference
				if patch1 == patch2 {
					return true
				}

				// patch value for first element is less return true
				if patch1 < patch2 {
					return true
				}

				// otherwise return false
				return false
			}

			// minor version value from first element is less
			// return true
			if minor1 < minor2 {
				return true
			}

			// otherwise return false
			return false
		}

		if major1 > major2 {
			return false
		}

		return true

	})

	fmt.Println(testValue)
}
