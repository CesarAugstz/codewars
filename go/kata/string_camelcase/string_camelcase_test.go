package string_camelcase_test

import (
	. "codewars/kata/string_camelcase"
	. "fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestStringCamelCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringCamelCase Suite")
}

func dotest(str, exp string) {
	Println("input:", str)
	var ans = ToCamelCase(str)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		dotest("", "")
		dotest("The_Stealth_Warrior", "TheStealthWarrior")
		dotest("the-Stealth-Warrior", "theStealthWarrior")
	})
})
