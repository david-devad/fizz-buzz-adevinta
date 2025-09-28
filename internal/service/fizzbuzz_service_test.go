package service

import (
	"reflect"
	"testing"
)

func TestGenerateFizzBuzz(t *testing.T) {
    res := GenerateFizzBuzz(3, 5, 31, "fizz", "buzz")
    expected := []string{
        "1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz",
        "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", 
		"buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", 
		"29", "fizzbuzz", "31",
    }
	
    if !reflect.DeepEqual(res, expected) {
        t.Errorf("GenerateFizzBuzz() = %v, want %v", res, expected)
    }
}