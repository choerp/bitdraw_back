package data_test

import (
	"fmt"
	"go_workspace/hello/data"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestUserData(t *testing.T) {
	// When given Stock Data from User. We make average data

	//given
	testDoc1 := `[
		{
		"index": 0,
		"lowPrice": 20,
		"highPrice": 50
		},
		{
		"index": 1,
		"lowPrice": 30,
		"highPrice": 60
		},
		{
		"index": 2,
		"lowPrice": 10,
		"highPrice": 80
		}
	]`

	expectedTestDoc1 := &[]data.UserData{
		{Index: 0, LowPrice: 20, HighPrice: 50},
		{Index: 1, LowPrice: 30, HighPrice: 60},
		{Index: 2, LowPrice: 10, HighPrice: 80},
	}

	var expectedTestDoc1array [3]data.UserData
	copy(expectedTestDoc1array[:], *expectedTestDoc1)

	//when
	actualTestDoc1 := data.CreatedUserData(testDoc1)

	var actualTestDoc1array [3]data.UserData
	copy(actualTestDoc1array[:], *actualTestDoc1)

	//Then
	assertEqual(t, actualTestDoc1array, expectedTestDoc1array, "")
}
