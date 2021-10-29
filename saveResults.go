package main

import (
	"fmt"
	"os"
)

// export the search results in a txt file
func SaveSearchResult(filename string, papers []Paper) {
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Sorry: couldn’t create the file!")
	}
	defer outFile.Close()

	fmt.Fprintln(outFile, "PMID"+"\t"+"Title"+"\t"+"Abstract")

	for _, data := range papers {
		fmt.Fprintln(outFile, data.PMID+"\t"+data.title+"\t",
			data.abstract)
	}
	fmt.Println("Successfully write to txt file!")
}
