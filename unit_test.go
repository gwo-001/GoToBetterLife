package main

import (
	"GoToBetterLife/util"
	"testing"
)

func TestUnit(t *testing.T) {
	util.HasLength(1)

	util.GetNowDate()
}
