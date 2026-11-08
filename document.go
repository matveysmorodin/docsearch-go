package main

type Document struct {
	ID     int
	Name   string
	Size   int64
	Status string
}

func ChangeStatus(doc *Document, newStatus string) {
	doc.Status = newStatus
}

func (d Document) IsProcessing() bool {
	return d.Status == "processing"
}

func (d *Document) MarkProcessed() {
	d.Status = "processed"
}
