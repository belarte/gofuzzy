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
)

func Parse(input string) (Expression, error) {
	re1 := regexp.MustCompile("[(]")
	re2 := regexp.MustCompile("[)]")
	input = re1.ReplaceAllString(input, " ( ")
	input = re2.ReplaceAllString(input, " ) ")

	tokens = make([]string, 0)
	tokens = strings.Fields(input)

	currentIndex = 0
	lastIdentifier = ""

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
	if currentIndex < len(tokens) && tokens[currentIndex] == s {
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

func readExpression() (Expression, error) {
	result, err := readAndExpression()

	for readKeyword("or") && err == nil {
		right, err2 := readAndExpression()
		result = OrExpressionBuilder(result, right)
		err = err2
	}

	return result, err
}

func readAndExpression() (Expression, error) {
	result, err := readNegationExpression()

	for readKeyword("and") && err == nil {
		right, err2 := readNegationExpression()
		result = AndExpressionBuilder(result, right)
		err = err2
	}

	return result, err
}

func readNegationExpression() (Expression, error) {
	if readKeyword("not") {
		result, err := readNegationExpression()
		return NegationExpressionBuilder(result), err
	} else {
		return readValueExpression()
	}
}

func readValueExpression() (Expression, error) {
	switch {
	case readKeyword("("):
		result, err := readExpression()
		if !readKeyword(")") {
			return nil, errors.New("Missing closing parenthesis")
		}
		return result, err
	case readIdentifier():
		_, check := knowledgeBase.Function(lastIdentifier)
		if !check {
			return nil, errors.New("Unknown function: " + lastIdentifier)
		}
		result := ValueExpressionBuilder(lastIdentifier)
		return result, nil
	}

	return nil, errors.New("An argument is missing")
}
