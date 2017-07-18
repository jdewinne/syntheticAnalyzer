package main

import (
	"encoding/xml"
	"fmt"
    "io/ioutil"
)

type Query struct {
	// Have to specify where to find episodes since this
	// doesn't match the xml tags of the data that needs to go into it
	TypeList []Type `xml:"type"`
}

type Type struct {
	Name string `xml:"type,attr"`
}

func (t Type) String() string {
	return fmt.Sprintf("+ %s", t.Name)
}
    
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    b, err := ioutil.ReadFile("/data/src/main/resources/synthetic.xml")
    check(err)

	var q Query
	xml.Unmarshal(b, &q)

	for _, xltype := range q.TypeList {
		fmt.Printf("\t%s\n", xltype)
	}
}
