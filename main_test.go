package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMetersForCustomer(t *testing.T) {
	testCases := []struct {
		name     string
		customer string
		expected []Building
	}{
		{
			name:     "Test Case 1",
			customer: "Aquaflow",
			expected: []Building{
				{
					Name:     "Treatment Plant A",
					Customer: "Aquaflow",
					SerialID: "1111-1111-1111",
				},
				{
					Name:     "Treatment Plant B",
					Customer: "Aquaflow",
					SerialID: "1111-1111-2222",
				},
			},
		},
		{
			name:     "Test Case 2",
			customer: "Albers Facilities Management",
			expected: []Building{
				{
					Name:     "Student Halls",
					Customer: "Albers Facilities Management",
					SerialID: "1111-1111-3333",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/getMetersForCustomer?customer="+tc.customer, nil)
			if err != nil {
				t.Fatalf("failed to create http request: %v", err)
			}

			rr := httptest.NewRecorder()
			getMetersForCustomer(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("test handler returned wrong response code: got: %v want %v", rr.Code, http.StatusOK)
			}

			var result []Building
			err = json.Unmarshal(rr.Body.Bytes(), &result)
			if err != nil {
				t.Fatalf("error decoding json resp: %v", err)
			}

			if !equalBuildings(result, tc.expected) {
				t.Errorf("test returned unexpected result: got: %v, want: %v", result, tc.expected)
			}
		})
	}
}

// helper to compare slices because reflect.DeepEqual is annoying & I want to stick with core packages
// for this exercise
func equalBuildings(a, b []Building) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Name != b[i].Name || a[i].Customer != b[i].Customer || a[i].SerialID != b[i].SerialID {
			return false
		}
	}
	return true
}
