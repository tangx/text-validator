package main

import (
	"fmt"
	"testing"
)

func TestHasSensetiveWord(t *testing.T) {

	for _, msg := range []string{
		"这里什么都没有",
		"可以给我一本禁书吗？",
		"哪里有卖昏药的？",
	} {
		fmt.Println(msg, "=>", IsValidateText(msg))
	}

}
