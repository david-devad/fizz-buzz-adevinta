package service

import (
	"strconv"
)

func GenerateFizzBuzz(int1, int2, limit int, str1, str2 string) []string {
	result := make([]string, 0, limit)
    
	for i := 1; i <= limit; i++ {
        var fizzOrBuzz string
		
        switch {
			case i%int1 == 0 && i%int2 == 0:
				fizzOrBuzz = str1 + str2
			case i%int1 == 0:
				fizzOrBuzz = str1
			case i%int2 == 0:
				fizzOrBuzz = str2
			default:
				fizzOrBuzz = strconv.Itoa(i)
        }
        result = append(result, fizzOrBuzz)
    }

    return result
}