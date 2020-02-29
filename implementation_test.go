package lab1

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPostfixToPrefix(c *C) {
	var res string
	var err error
	// Difficult example
	res, _ = PostfixToPrefix("5 4 3 * + 12 4 / - 1 -")
	c.Assert(res, Equals, "- - + 5 * 4 3 / 12 4 1")
	// Simple examples
	res, _ = PostfixToPrefix("4 3 * 2 -")
	c.Assert(res, Equals, "- * 4 3 2")
	// Error case "must be non empty string"
	_, err = PostfixToPrefix("")
	c.Assert(err, ErrorMatches, "must be non empty string")
	// Error case "after operand must be whitespace"
	_, err = PostfixToPrefix("4 3 * 2-")
	c.Assert(err, ErrorMatches, "after operand must be whitespace")
	// Error case "after operator must be whitespace"
	_, err = PostfixToPrefix("4 3 *2 -")
	c.Assert(err, ErrorMatches, "after operator must be whitespace")
	// Error case "first char must be operand not operator"
	_, err = PostfixToPrefix("+ 4 3 * 2 -")
	c.Assert(err, ErrorMatches, "first char must be operand not operator")
	// Error case "first char must be operand not whitespace"
	_, err = PostfixToPrefix(" 5 4 +")
	c.Assert(err, ErrorMatches, "first char must be operand not whitespace")
	// Error case "after whitespace must be operand or operator"
	_, err = PostfixToPrefix("5  4")
	c.Assert(err, ErrorMatches, "after whitespace must be operand or operator")
	// Error case "last char must be operand or operator"
	_, err = PostfixToPrefix("5 + 4 ")
	c.Assert(err, ErrorMatches, "last char must be operand or operator")
	// Error case "count of operators must be less operands then one"
	_, err = PostfixToPrefix("5 4 + + + + +")
	c.Assert(err, ErrorMatches, "count of operators must be less operands then one")
	// Letters case
	res, _ = PostfixToPrefix("ab c +")
	c.Assert(res, Equals, "+ ab c")
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("5 4 3 *+ 12 4 /- 1-")
	fmt.Println(res)
}
