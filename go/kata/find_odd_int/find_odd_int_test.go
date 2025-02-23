package find_odd_int_test

import (
  . "codewars/kata/find_odd_int"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestFindOddInt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FindOddInt Suite")
}

var _ = Describe("Basic tests", func() {
	arr := [...][]int{
		{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
		{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
		{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
		{10},
		{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
		{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
	}
	sol := [...]int{5, -1, 5, 10, 10, 1}
	for i, v := range arr {
		cap_v := v
		cap_i := i
		It(fmt.Sprintf("Testing input %v", v), func() { Expect(FindOdd(cap_v)).To(Equal(sol[cap_i])) })
	}
})
