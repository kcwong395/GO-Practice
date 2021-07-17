package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"strconv"
	"time"
)

type Verifier struct {
	rdb *Client
}

func addPaddingZero(code string) string {
	for len(code) < 6 {
		code = "0" + code
	}
	return code
}

// Generate a random 6 digits number and map to an emailAddr
func (v *Verifier) GenerateCode(emailAddr string) string {
	// generate random number with seed
	rand.Seed(time.Now().UnixNano())
	// the generated code ranges 000000 - 999999
	code := addPaddingZero(strconv.Itoa(rand.Intn(1000000)))

	// the key is only valid for 3 seconds
	v.rdb.Set(context.Background(), emailAddr, code, 3*time.Second)
	return code
}

func (v *Verifier) Verification(emailAddr string, claimedCode string) bool {
	code, err := v.rdb.Get(context.Background(), emailAddr).Result()
	if err != redis.Nil && err == nil {
		return code == claimedCode
	}
	fmt.Println("Code expired or wrong email")
	return false
}
