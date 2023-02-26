package main

import (
	"testing"
)

func TestDoActionWithLoadingAnimation(t *testing.T) {
	ret := doActionWithLoadingAnimation("Testing")
	if ret != true {
		t.Errorf("Failed, false value was returned")
	}
}