package main

import "testing"

func TestDocument_StartProcessing(t *testing.T) {
	tests := []struct {
		name       string
		doc        Document
		wantResult bool
		wantStatus string
		repeat     bool
	}{
		{
			name: "Valid document",
			doc: Document{
				ID:     1,
				Name:   "contract.pdf",
				Size:   1024,
				Status: "uploaded",
			},

			wantResult: true,
			wantStatus: "processing",
			repeat:     false,
		},
		{
			name: "EmptyName",
			doc: Document{
				ID:     2,
				Name:   "",
				Size:   1024,
				Status: "uploaded",
			},
			wantResult: false,
			wantStatus: "uploaded",
			repeat:     false,
		}, {
			name: "ZeroSize",
			doc: Document{
				ID:     3,
				Name:   "contract.pdf",
				Size:   0,
				Status: "uploaded",
			},
			wantResult: false,
			wantStatus: "uploaded",
			repeat:     false,
		},
		{
			name: "AlreadyProcessed",
			doc: Document{
				ID:     4,
				Name:   "contract.pdf",
				Size:   1024,
				Status: "processed",
			},
			wantResult: false,
			wantStatus: "processed",
			repeat:     false,
		}, {
			name: "NegativeSize",
			doc: Document{
				ID:     5,
				Name:   "contract.pdf",
				Size:   -100,
				Status: "uploaded",
			},
			wantResult: false,
			wantStatus: "uploaded",
			repeat:     false,
		}, {
			name: "RepeatedStart",
			doc: Document{
				ID:     6,
				Name:   "contract.pdf",
				Size:   1024,
				Status: "uploaded",
			},
			wantResult: false,
			wantStatus: "processing",
			repeat:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.repeat {
				firstResult := tt.doc.StartProcessing()

				if !firstResult {
					t.Errorf("Первый запуск: ожидали true, получили false")
				}
			}

			result := tt.doc.StartProcessing()

			if result != tt.wantResult {
				t.Errorf("Ожидали результат %t, получили %t",
					tt.wantResult, result)
			}

			if tt.doc.Status != tt.wantStatus {
				t.Errorf("Ожидали статус %s, получили %s",
					tt.wantStatus, tt.doc.Status)
			}
		})
	}
}
