package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sajari/regression"
)

// import needed packages

func RunReg(data [][]string) (regression.Regression, regression.Regression, regression.Regression, regression.Regression) {
	var r1 regression.Regression
	r1.SetObserved("yvar1")
	r1.SetVar(0, "xvar1")
	var r2 regression.Regression
	r2.SetObserved("yvar2")
	r2.SetVar(0, "xvar2")
	var r3 regression.Regression
	r3.SetObserved("yvar3")
	r3.SetVar(0, "xvar3")
	var r4 regression.Regression
	r4.SetObserved("yvar4")
	r4.SetVar(0, "xvar4")

	// Loop of records in the CSV, adding the training data to the regression value.
	for i, record := range data {
		// Skip the header.
		if i == 0 {
			continue
		}
		yvar1, _ := strconv.ParseFloat(record[4], 64)
		xvar1, _ := strconv.ParseFloat(record[0], 64)
		r1.Train(
			regression.DataPoint(yvar1, []float64{xvar1}),
		)

		yvar2, _ := strconv.ParseFloat(record[5], 64)
		xvar2, _ := strconv.ParseFloat(record[1], 64)
		r2.Train(
			regression.DataPoint(yvar2, []float64{xvar2}),
		)

		yvar3, _ := strconv.ParseFloat(record[6], 64)
		xvar3, _ := strconv.ParseFloat(record[2], 64)
		r3.Train(
			regression.DataPoint(yvar3, []float64{xvar3}),
		)

		yvar4, _ := strconv.ParseFloat(record[7], 64)
		xvar4, _ := strconv.ParseFloat(record[3], 64)
		r4.Train(
			regression.DataPoint(yvar4, []float64{xvar4}),
		)

	}
	// Train/fit the regression model.
	r1.Run()
	r2.Run()
	r3.Run()
	r4.Run()
	return r1, r2, r3, r4
}

func main() {

	f, err := os.Open("anscombe.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// we create a new csv reader specifying
	// the number of columns it has
	testdata := csv.NewReader(f)
	testdata.FieldsPerRecord = 8
	// we read all the records
	records, err := testdata.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	r1, r2, r3, r4 := RunReg(records)

	// Output the trained model parameters.
	fmt.Printf("\nRegression Formula:\n%v\n\n", r1.Formula)
	fmt.Printf("\nRegression Formula:\n%v\n\n", r2.Formula)
	fmt.Printf("\nRegression Formula:\n%v\n\n", r3.Formula)
	fmt.Printf("\nRegression Formula:\n%v\n\n", r4.Formula)

}
