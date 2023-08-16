package services_test

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name   string
		score  int
		expect string
	}

	cases := []testCase{
		{name: "a", score: 80, expect: "A"},
		{name: "b", score: 70, expect: "B"},
		{name: "c", score: 60, expect: "C"},
		{name: "d", score: 50, expect: "D"},
		{name: "f", score: 0, expect: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)

			assert.Equal(t, c.expect, grade)
			// if grade != c.expect {
			// 	t.Errorf("got %v c.expect %v", grade, c.expect)
			// }
		})

	}

}

func BenchmarkCheckGrade(b *testing.B) {

	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output: A
}
