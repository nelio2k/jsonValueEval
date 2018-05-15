package jsonValueEval

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"testing"
)

func generateRandomData() ([][]byte, int, error) {
	//	data := make([][]byte, 1)
	//	totalBytes, err := genRandomUsers(32534059803498589, data)
	//	if err != nil {
	//		return nil, 0, err
	//	}
	//	fmt.Printf("TotalBytes: %v TotalEntries: %v\n", totalBytes, len(data))

	mbsToGenerate := 100
	avgBytesOfOneRecord := 77
	rowsToGenerate := mbsToGenerate * 1048576 / avgBytesOfOneRecord
	data := make([][]byte, rowsToGenerate)
	totalBytes, err := genRandomUsers(32534059803498589, data)
	if err != nil {
		return nil, 0, err
	}
	fmt.Printf("MBs To Generate: %v TotalBytes: %v TotalEntries: %v\n", mbsToGenerate, totalBytes, len(data))
	return data, totalBytes, nil
}

func BenchmarkSimpleMatcher(b *testing.B) {
	//	b.SetParallelism(4)
	m := NewFlexibleMatcher()
	data, totalBytes, err := generateRandomData()
	if err != nil || len(data) == 0 {
		b.Fatalf("Matcher error: %s", err)
	}

	// Expression reformatted:
	// firstName=='Neil' || (age<50 && isActive==true)
	expression, err := govaluate.NewEvaluableExpression("firstName == 'Neil' || (age < 50 && isActive == true)")
	if err != nil {
		b.Fatal("NewEvaluableExpression Error: %s", err)
		return
	}

	// Pre-make parameters and re-use
	//	parameters := make(map[string]interface{}, 3)
	parameters := NewParameterArray(3)
	b.SetBytes(int64(totalBytes))
	b.ResetTimer()

	for j := 0; j < b.N; j++ {
		//		garbageTracker.Start()
		for i := 0; i < len(data); i++ {
			_, err := m.Match(data[i], expression, *parameters)

			if err != nil {
				b.Fatalf("Matcher error: %s", err)
			}
		}
	}
}

func TestParserImpl(t *testing.T) {
	fmt.Println("====================== TestParserImpl Started ==========================")
	m := NewFlexibleMatcher()
	data, _, _ := generateRandomData()

	// Expression reformatted:
	// firstName=='Neil' || (age<50 && isActive==true)
	expression, err := govaluate.NewEvaluableExpression("firstName == 'Neil' || (age < 50 && isActive == true)")
	if err != nil {
		t.Fatal("NewEvaluableExpression Error: %s", err)
		return
	}

	// Pre-make parameters and re-use
	parameters := make(map[string]interface{}, 3)

	for _, oneData := range data {
		m.Match(oneData, expression, parameters)
	}

	fmt.Println("====================== TestParserImpl Ended   ==========================")
}
