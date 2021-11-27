package BioHits

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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
		// process abstract by removing useless white spaces
		abstractWords := strings.Split(data.abstract, " ")
		for i := range abstractWords {
			wsChar := regexp.MustCompile(`\s`)
			s1 := wsChar.ReplaceAllString(abstractWords[i], "")
			abstractWords[i] = RemoveSpecialChar(s1)
		}
		cleanedAbstract := strings.Join(abstractWords[:], " ")
		fmt.Fprintln(outFile, data.PMID+"\t"+data.title+"\t",
			cleanedAbstract)
	}
	fmt.Println("Successfully wrote search results to a txt file!")
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
	fmt.Println("Successfully wrote words to a txt file!")
}
