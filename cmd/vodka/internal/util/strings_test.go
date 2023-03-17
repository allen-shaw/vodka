package util

import (
	"fmt"
	"testing"

	"github.com/stvp/assert"
)

func TestSnakeCase(t *testing.T) {
	s := "thisIsATestCase"
	sc := SnakeCase(s)
	fmt.Println(sc)
	assert.Equal(t, "this_is_a_test_case", sc)
}

func TestCamelCase(t *testing.T) {
	s1 := "this_is_a_test_case"
	c1 := CamelCase(s1)

	s2 := "ThisIsATestCase"
	c2 := CamelCase(s2)

	s3 := "HTTP_handler"
	c3 := CamelCase(s3)

	fmt.Println(c1, c2, c3)

}

func TestPascalCase(t *testing.T) {
	s1 := "this_is_a_test_case"
	c1 := PascalCase(s1)

	s2 := "ThisIsATestCase"
	c2 := PascalCase(s2)

	s3 := "HTTP_handler"
	c3 := PascalCase(s3)

	fmt.Println(c1, c2, c3)
}
