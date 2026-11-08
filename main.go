package main

import "fmt"

func main() {
	doc := Document{
		ID:     1,
		Name:   "contract.pdf",
		Size:   4096,
		Status: "uploaded",
	}
	fmt.Println(doc.Name, doc.Status)
	ChangeStatus(&doc, "processing")
	fmt.Println(doc.Status)
}
