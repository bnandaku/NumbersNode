package main

import (
	advanced_db "advanced-db"
	"fmt"
)

type sampleType struct {

	Score int
	Number string

}


func main() {

	sample := sampleType{
		Score: 10,
		Number:"3017447928",
	}
	node := advanced_db.NewNode()
	node.Insert(sample, "3017447928")
	sample.Score = 11
	sample.Number = "3017447922"
	node.Insert(sample, "3017447922")

	sample.Score = 13
	sample.Number = "3017447923"

	node.Insert(sample, "3017447923")

	fmt.Println(node.Search("3017447928"))
	sample.Number = "3017447928"
	sample.Score = 20
	node.Update(sample, "3017447928")
	fmt.Println(node.Search("3017447928"))
	node.Delete("3017447928")
	fmt.Println(node.Search("3017447928"))
	fmt.Println(node.Search("3017447928"))
	fmt.Println(node.Search("3017447927"))
	fmt.Println(node.Search("3017447922"))
	fmt.Println(node.Search("3017447923"))
	fmt.Println(node.Search("1234567890"))
}
