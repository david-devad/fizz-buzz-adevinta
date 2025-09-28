package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFizzBuzzHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=31&str1=fizz&str2=buzz", nil)
    if err != nil {
        t.Fatal(err)
    }

    nRec := httptest.NewRecorder()
    FizzBuzzHandler(nRec, req)

    if nRec.Code != http.StatusOK {
        t.Errorf("Expected status %v, got %v", http.StatusOK, nRec.Code)
    }

    expected := `["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31"]`
    if nRec.Body.String() != expected+"\n" {
        t.Errorf("Expected body %v, got %v", expected, nRec.Body.String())
    }
}

func TestFizzBuzzHandlerWithLimit100(t *testing.T) {
    req, err := http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=100&str1=fizz&str2=buzz", nil)
    if err != nil {
        t.Fatal(err)
    }

    nRec := httptest.NewRecorder()
    FizzBuzzHandler(nRec, req)

    if nRec.Code != http.StatusOK {
        t.Errorf("Expected status %v, got %v", http.StatusOK, nRec.Code)
    }

    expected := `["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz","fizz","52","53","fizz","buzz","56","fizz","58","59","fizzbuzz","61","62","fizz","64","buzz","fizz","67","68","fizz","buzz","71","fizz","73","74","fizzbuzz","76","77","fizz","79","buzz","fizz","82","83","fizz","buzz","86","fizz","88","89","fizzbuzz","91","92","fizz","94","buzz","fizz","97","98","fizz","buzz"]`
    if nRec.Body.String() != expected+"\n" {
        t.Errorf("Expected body %v, got %v", expected, nRec.Body.String())
    }
}

func TestFizzBuzzHandlerInvalidIntegerParams(t *testing.T) {
    req, err := http.NewRequest("GET", "/fizzbuzz?int1=abc&int2=5&limit=15&str1=fizz&str2=buzz", nil)
    if err != nil {
        t.Fatal(err)
    }

    nRec := httptest.NewRecorder()
    FizzBuzzHandler(nRec, req)

    if nRec.Code != http.StatusBadRequest {
        t.Errorf("Expected status %v, got %v", http.StatusBadRequest, nRec.Code)
    }
}

func TestFizzBuzzHandlerInvalidStringEmptyParams(t *testing.T) {
    req, err := http.NewRequest("GET", "/fizzbuzz?int1=abc&int2=5&limit=15&str1=fizz&str2=", nil)
    if err != nil {
        t.Fatal(err)
    }

    nRec := httptest.NewRecorder()
    FizzBuzzHandler(nRec, req)

    if nRec.Code != http.StatusBadRequest {
        t.Errorf("Expected status %v, got %v", http.StatusBadRequest, nRec.Code)
    }
}

func TestFizzBuzzHandlerInvalidLimitParams(t *testing.T) {
    req, err := http.NewRequest("GET", "/fizzbuzz?int1=abc&int2=5&limit=1001&str1=fizz&str2=", nil)
    if err != nil {
        t.Fatal(err)
    }

    nRec := httptest.NewRecorder()
    FizzBuzzHandler(nRec, req)

    if nRec.Code != http.StatusBadRequest {
        t.Errorf("Expected status %v, got %v", http.StatusBadRequest, nRec.Code)
    }
}