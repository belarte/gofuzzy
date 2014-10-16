package engine

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	knowledgeBase KnowledgeBase
)

func Init() {
	knowledgeBase = NewKnowledgeBase()
}

func Open(name string) error {
	xmlFile, err := os.Open("../resources/" + name)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	b, err2 := ioutil.ReadAll(xmlFile)

	if err2 != nil {
		return err2
	}

	var base BaseDefinition
	if err3 := xml.Unmarshal(b, &base); err3 != nil {
		return err3
	}

	for _, f := range base.Functions {
		generateFunction(f)
	}
	for _, r := range base.Rules {
		generateRule(r)
	}

	return nil
}

func generateFunction(def FunctionDefinition) {
	paramsAsString := strings.Fields(def.Parameters)
	paramsAsFloat := make([]float64, len(paramsAsString))

	for i, s := range paramsAsString {
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			paramsAsFloat[i] = f
		}
	}

	if f, err := NewMembershipFunction(def.Attribute, def.Shape, paramsAsFloat); err == nil {
		knowledgeBase.AddFunction(def.Name, f)
	}
}

func generateRule(def RuleDefinition) {
	if rule, err := Parse(def.Definition); err == nil {
		knowledgeBase.AddRule(def.Name, rule)
	}
}

type BaseDefinition struct {
	XMLName   xml.Name             `xml:"base"`
	Functions []FunctionDefinition `xml:"function"`
	Rules     []RuleDefinition     `xml:"rule"`
}

type FunctionDefinition struct {
	XMLName    xml.Name `xml:"function"`
	Name       string   `xml:"name,attr"`
	Attribute  string   `xml:"attribute,attr"`
	Shape      string   `xml:"type,attr"`
	Parameters string   `xml:"parameters,attr"`
}

type RuleDefinition struct {
	XMLName    xml.Name `xml:"rule"`
	Name       string   `xml:"name,attr"`
	Definition string   `xml:"definition,attr"`
}
