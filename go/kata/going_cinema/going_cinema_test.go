package going_cinema_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
  . "codewars/kata/going_cinema"
	"testing"
)

func TestMovie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Movie Suite")
}

func dotest(card, ticket int, perc float64, exp int) {
	var ans = Movie(card, ticket, perc)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests Movie", func() {
	fmt.Println("****** Test Movie ******")

	It("should handle basic cases", func() {
		dotest(500, 15, 0.9, 43)
		dotest(100, 10, 0.95, 24)
		dotest(0, 10, 0.95, 2)
	})
})
