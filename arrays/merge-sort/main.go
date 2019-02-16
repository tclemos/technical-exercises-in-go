package main

import (
	"log"
	"os"

	arraytools "github.com/tclemos/technical-exercises-in-go/array-tools"
	mergesort "github.com/tclemos/technical-exercises-in-go/arrays/merge-sort/lib"
)

func main() {

	numbersToSort := arraytools.ToIntSlice(os.Args[1:])

	sortedNumbers := mergesort.AscendingSort(numbersToSort)

	log.Println(sortedNumbers)
}
