package test

import (
	"reflect"
	"testing"
	"time"
	"todoapp/model"
	"todoapp/sample"
)

func TestSampleData(t *testing.T) {
	sampleData := sample.Sample{}

	got, err := sampleData.GetSampleData()
	if err != nil {
		t.Error(err)
	}

	expectedData := []model.Todo{
		{
			ID:       1,
			Labels:   "Clean dishes",
			Comments: "Don't forget the garbage too..",
			DueDate:  time.Date(2021, 10, 28, 7, 40, 0, 0, time.UTC),
		},
		{
			ID:       2,
			Labels:   "Send kids to school",
			Comments: "Don't forget the Umbrella..",
			DueDate:  time.Date(2021, 10, 28, 8, 10, 0, 0, time.UTC),
		},
		{
			ID:       3,
			Labels:   "Go to work",
			Comments: "Do some small-talks with colleagues before starting the work..",
			DueDate:  time.Date(2021, 10, 28, 8, 30, 0, 0, time.UTC),
		},
	}

	if !reflect.DeepEqual(got, expectedData) {
		t.Fatalf("Want: %v, Got: %v\n", expectedData, got)
	}
}

func TestIds(t *testing.T) {

	ids := [3]int{1, 2, 3}

	sampleData := sample.Sample{}

	got, err := sampleData.GetSampleData()
	if err != nil {
		t.Error(err)
	}

	var result []int
	for _, tc := range got {
		result = append(result, tc.ID)
	}

	if len(ids) != len(result) {
		t.Fatalf("Test failed — Expected %d ids, got %d", len(ids), len(result))
	}
}
