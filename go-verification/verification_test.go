package main

import (
	"context"
	"log"
	"testing"
	"time"
)

/*
	=== RUN   Test1
	Generated code: 694561
	--- PASS: Test1 (0.02s)
*/
func Test1(t *testing.T) {
	rdb, err := Init()
	if err != nil {
		log.Panicln(err)
	}

	v := Verifier{rdb: rdb}
	code, err := v.GenerateCode("example@example.com")
	if err != nil {
		t.Errorf("Failed to generate code")
	}

	isValid := v.Verification("example@example.com", code)
	if !isValid {
		t.Errorf("Expected to pass")
	}
}

/*
	=== RUN   Test2
	--- PASS: Test2 (0.00s)
*/
func Test2(t *testing.T) {
	rdb, err := Init()
	if err != nil {
		log.Panicln(err)
	}

	v := Verifier{rdb: rdb}

	// this shall fail
	isValid := v.Verification("example@example.com", "1000000000")
	if isValid {
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
	rdb, err := Init()
	if err != nil {
		log.Panicln(err)
	}

	v := Verifier{rdb: rdb}

	code, _ := rdb.Get(context.Background(), "example@example.com").Result()

	// wait until expire
	time.Sleep(5 * time.Second)

	// this shall fail
	isValid := v.Verification("example@example.com", code)
	if isValid {
		t.Errorf("Expected to fail")
	}
}
