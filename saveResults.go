package BioHits

import (
	"fmt"
	"os"
)

// export the search results in a txt file
func SaveSearchResult(filename string, papers []Paper) {
	outFile, err := os.Create(filename)
	if err != nil {
		panic("Sorry: cannot create the file!")
	}
	defer outFile.Close()

	// set header
	fmt.Fprintln(outFile, "PMID"+"\t"+"Title"+"\t"+"Abstract")

	for _, data := range papers {
		fmt.Fprintln(outFile, data.PMID+"\t"+data.title+"\t",
			data.abstract)
	}
	fmt.Println("Successfully write search results to a txt file!")
}

// SaveWords save words in a txt file
func SaveWords(words []string, filename string) {
	outFile, err := os.Create(filename)
	if err != nil {
		panic("Sorry: cannot create the file!")
	}
	defer outFile.Close()

	for _, w := range words {
		fmt.Fprintln(outFile, w)
	}
	fmt.Println("Successfully write words to a txt file!")
}
