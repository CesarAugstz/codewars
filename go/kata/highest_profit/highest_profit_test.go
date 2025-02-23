package highest_profit_test

import (
	. "codewars/kata/highest_profit"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHighestProfit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HighestProfit Suite")
}

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(MinMax([]int{1, 2, 3, 4, 5})).To(Equal([2]int{1, 5}))
		Expect(MinMax([]int{2334454, 5})).To(Equal([2]int{5, 2334454}))
	})
})
