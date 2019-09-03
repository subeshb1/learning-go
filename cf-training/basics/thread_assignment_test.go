package main

import "testing"

func TestSum(t *testing.T) {

	m := Math{}

	result := m.Square(4)

	//There are no built-in assert methods!

	if result != 16 {

		//t.Fail() marks this test failed without a message and continues execution
		//t.FailNow() marks this test failed and exits this test immediately
		//t.Fatalf calls FailNow _and_ logs an error message

		t.Errorf("Expected 16 but found %d", result) //Calls Fail() and logs an error message
	}

}

