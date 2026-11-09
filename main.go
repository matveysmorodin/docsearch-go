package main

import "fmt"

func main() {

	documents := []Document{
		{
			ID:     1,
			Name:   "contract.pdf",
			Size:   4096,
			Status: "uploaded",
		},
		{
			ID:     2,
			Name:   "empty.pdf",
			Size:   0,
			Status: "uploaded",
		},
		{
			ID:     3,
			Name:   "",
			Size:   100,
			Status: "uploaded",
		},
		{
			ID:     4,
			Name:   "report.pdf",
			Size:   2000,
			Status: "processed",
		},
	}

	for i := range documents {
		fmt.Println(documents[i].ID, documents[i].StartProcessing(), documents[i].Status)
	}

}
