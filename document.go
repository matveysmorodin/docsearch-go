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

func (d *Document) CanStartProcessing() bool {
	return d.Status == "uploaded" && d.Name != "" && d.Size > 0
}

func (d *Document) StartProcessing() bool {
	if d.CanStartProcessing() {
		d.Status = "processing"
		return true
	}
	return false
}
