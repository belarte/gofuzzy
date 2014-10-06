package engine

import (
	"regexp"
	"strings"
)

var (
	tokens         []string
	currentIndex   int
	lastIdentifier string
)

func Parse(input string) Function {
	tokens = make([]string, 0)
	tokens = strings.Fields(input)

	currentIndex = 0

	return nil
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

func readExpression() Function {
	return nil
}

func readAndExpression() Function {
	return nil
}

func readNegationExpression() Function {
	return nil
}

func readValueExpression() Function {
	return nil
}
