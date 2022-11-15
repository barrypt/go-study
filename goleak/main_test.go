package main

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(t *testing.M) {
	//defer goleak.VerifyNone(t)
	//GetData()
	defer goleak.VerifyTestMain(t)
	GetData()
}