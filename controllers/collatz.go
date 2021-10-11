package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func CollatzDisplayAction(c *gin.Context) {
	c.HTML(200, "collatz.html", gin.H{})
}

func CalcNextStep(n int64) int64 {
	// calc next step of Collatz procedure
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}
}

func CalcCollatz(n int64) []int64 {
	// calc Collatz procedure
	sequence := []int64{n}
	a_i := n
	for {
		a_i = CalcNextStep(a_i)
		sequence = append(sequence, a_i)
		if a_i == 1 {
			return sequence
		}
	}
}

func GetCollatzSequence(n int64) string {
	sequence := CalcCollatz(n)
	sequence_display := strconv.FormatInt(sequence[0], 10)
	for i := 1; i < len(sequence); i++ {
		sequence_display += " -> "
		sequence_display += strconv.FormatInt(sequence[i], 10)
	}
	return sequence_display
}
