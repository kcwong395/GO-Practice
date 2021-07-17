package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
	=== RUN   Test1
	Generated code: 694561
	--- PASS: Test1 (0.02s)
*/
func Test1(t *testing.T) {
	rdb := Init()

	v := Verifier{rdb: rdb}
	code := v.GenerateCode("example@example.com")
	fmt.Printf("Generated code: %s \n", code)

	// this shall pass
	res := v.Verification("example@example.com", code)
	if res != true {
		t.Errorf("Expected to match")
	}
}

/*
	=== RUN   Test2
	--- PASS: Test2 (0.00s)
*/
func Test2(t *testing.T) {
	rdb := Init()
	v := Verifier{rdb: rdb}

	// this shall fail
	res := v.Verification("example@example.com", "1000000000")
	if res != false {
		t.Errorf("Expected to fail")
	}
}

/*
	=== RUN   Test3
	694561
	Code expired or wrong email
	--- PASS: Test3 (5.01s)
*/
func Test3(t *testing.T) {

	rdb := Init()
	v := Verifier{rdb: rdb}

	code, _ := rdb.Get(context.Background(), "example@example.com").Result()
	fmt.Println(code)

	time.Sleep(5 * time.Second)

	// this shall fail
	res := v.Verification("example@example.com", code)
	if res != false {
		t.Errorf("Expected to fail")
	}
}
