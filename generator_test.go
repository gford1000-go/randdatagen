package main

import (
	"strings"
	"testing"
)

func TestGenerator(t *testing.T) {

	var builder strings.Builder
	var recordCount = 5

	g := NewGenerator(recordCount, "GB")

	err := g.Create(&builder)
	if err != nil {
		t.Fatalf("unexpected error - got %v", err)
	}

	// Remove the last \n prior to splitting
	s := builder.String()
	records := strings.Split(s[:len(s)-1], "\n")

	if len(records) != recordCount {
		t.Fatalf("Invalid number of records: expected %v, got %v", recordCount, len(records))
	}

	for i, record := range records {
		if n := len(strings.Split(record, ",")); n != 11 {
			t.Fatalf("Invalid number of columns in record %v: expected %v, got %v", i, 11, n)
		}
	}
}
