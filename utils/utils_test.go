package utils

import (
	"fmt"
	"testing"
	"github.com/dgrijalva/jwt-go"
)

func TestParseClaims(t *testing.T) {
	claims := ParseClaims()
	claimsMap := claims.(jwt.MapClaims)
	fmt.Printf("Decoded claims : %v\n", claimsMap["foo"])
	if claimsMap["foo"] != "bar" {
		t.Fatal("Expected bar but got :", claimsMap["foo"])
	}
}
