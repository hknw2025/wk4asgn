# Week 4 Asignment - Stats in Go

This program takes data in an input CSV and preforms regression analysis on the underlying data. It is meant to be compared to the results of similar Python and R programs to make sure Go handles regression aanalysis correctly. 

## How to use this tool:
* After cloning this git repository you can 
`go run main.go

## Example:

  {"xvar1", "xvar2", "xvar3", "xvar4", "yvar1", "yvar2", "yvar3", "yvar4},
        {"10", "10.5", "8", "20", "5", "12", "1", "10"},


If this data is contained in "anscombe.csv", this program will read the data in and preform regression analysis on the similarly numbered columns. So for example there will be a regression preformed with xvar1 as an independent variable and yvar1 as the dependent. 


This function and the accompanying R and Python functions show that regression analysis can be carried out in Go and recieves the same results as in other languages. Additionally, the time I got whe running the benchmark test in this repository, 0.767 seconds was only around 0.3 seconds slower than the the time I got when running the R runtime test in the R program in this repository, 0.440 seconds. Based on these results and mangement's desire to limit the number of languages used by coders in the business, I recommend that Go should be used for statistical analysis in the business. 





















