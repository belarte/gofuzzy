package engine

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	knowledgeBase KnowledgeBase
)

func Init() {
	knowledgeBase = NewKnowledgeBase()
}

func Open(name string) {
	xmlFile, err := os.Open("../resources/test.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	b, err2 := ioutil.ReadAll(xmlFile)

	if err2 != nil {
		fmt.Println("Error reading file:", err2)
	}

	var base BaseDefinition
	if err3 := xml.Unmarshal(b, &base); err3 != nil {
		fmt.Println("Error unmarshaling:", err3)
	}

	for i, f := range base.Functions {
		fmt.Println(i, f.Name)
	}
	for i, r := range base.Rules {
		fmt.Println(i, r.Name, r.Definition)
	}
}

type BaseDefinition struct {
	XMLName   xml.Name             `xml:"base"`
	Functions []FunctionDefinition `xml:"function"`
	Rules     []RuleDefinition     `xml:"rule"`
}

type FunctionDefinition struct {
	XMLName xml.Name `xml:"function"`
	Name    string   `xml:"name,attr"`
	Shape   string   `xml:"type,attr"`
}

type RuleDefinition struct {
	XMLName    xml.Name `xml:"rule"`
	Name       string   `xml:"name,attr"`
	Definition string   `xml:"definition,attr"`
}
