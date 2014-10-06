package engine

import (
	"errors"
	"regexp"
	"strings"
)

var (
	tokens         []string
	currentIndex   int
	lastIdentifier string
	andOperator    string
	orOperator     string
)

func Parse(input string) (Function, error) {
	tokens = make([]string, 0)
	tokens = strings.Fields(input)

	currentIndex = 0
	lastIdentifier = ""
	andOperator = "min"
	orOperator = "max"

	if len(tokens) == 0 {
		return nil, errors.New("Empty input")
	}

	result, err := readExpression()

	if currentIndex != len(tokens) {
		return nil, errors.New("Malformed expression, currentIndex!=len(tokens)")
	}

	return result, err
}

func readKeyword(s string) bool {
	if tokens[currentIndex] == s {
		currentIndex++
		return true
	}
	return false
}

func readIdentifier() bool {
	token := tokens[currentIndex]
	r := regexp.MustCompile("^[A-Za-z][A-Za-z0-9_-]*$")

	if r.MatchString(token) && token != "and" && token != "or" && token != "not" {
		lastIdentifier = token
		currentIndex++
		return true
	}
	return false
}

func readExpression() (Function, error) {
	result, err := readAndExpression()

	for readKeyword("or") && err != nil {
		right, err2 := readAndExpression()
		result = BinaryExpressionBuilder(result, right, NewOperator(orOperator))
		err = err2
	}

	return result, err
}

func readAndExpression() (Function, error) {
	result, err := readNegationExpression()

	for readKeyword("and") && err == nil {
		right, err2 := readNegationExpression()
		result = BinaryExpressionBuilder(result, right, NewOperator(andOperator))
		err = err2
	}

	return result, err
}

func readNegationExpression() (Function, error) {
	if readKeyword("not") {
		result, err := readNegationExpression()
		return NegationExpressionBuilder(result), err
	} else {
		return readValueExpression()
	}
}

func readValueExpression() (Function, error) {
	switch {
	case readKeyword("("):
		result, err := readExpression()
		if !readKeyword(")") {
			return nil, errors.New("Missing closing parenthesis")
		}
		return result, err
	case readIdentifier():
		result := ValueExpressionBuilder(nil)
		return result, nil
	}

	return nil, errors.New("An argument is missing")
}
