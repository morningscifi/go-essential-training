package sqrt

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

type testCase struct {
	value    float64
	expected float64
}

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestMany(t *testing.T) {
	// Read test cases from "sqrt_case.csv" and check them
	csvData, err := os.ReadFile("sqrt_cases.csv")
	if err != nil {
		log.Fatalf("Reading file failed")
	}

	csvString := string(csvData)
	if err != nil {
		log.Fatalf("Bytes failed to convert to string")
	}

	r := csv.NewReader(strings.NewReader(csvString))
	testCases, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc[0]), func(t *testing.T) {
			val, err := strconv.ParseFloat(tc[0], 64)
			if err != nil {
				t.Fatal("Test value failed to convert from string")
			}
			out, err := Sqrt(val)
			if err != nil {
				t.Fatal("Sqrt produced error")
			}
			exp, err := strconv.ParseFloat(tc[1], 64)
			if !almostEqual(out, exp) {
				t.Fatalf("%f != %f", out, float64(exp))
			}
		})
	}
}
