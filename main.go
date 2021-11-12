package main

import (
	"flag"
	"fmt"
)

func main() {

	// initiate parameters taken by BioHits
	var numPaper int
	var keyWords string
	var searchOutputFile string
	var searchOutput bool

	// numPaper refer to the number of paper the user wants to consider
	flag.IntVar(&numPaper, "numPaper", 100, "Number of paper to be considered")
	// keyWords are the words that user wants to search
	// multiple words should use , to separate, such as gene,cancer
	flag.StringVar(&keyWords, "keyWords", "trending", "Key words for searching")
	flag.StringVar(&searchOutputFile, "searchFileName", "searchOutput.txt",
		"A txt file to store the search output")
	flag.BoolVar(&searchOutput, "exportSearchOutput", false,
		"Whether to write the search output to txt file")
	flag.Parse()

	papers := FetchPaperInfo(keyWords, numPaper)

	words := AnalyzePaperInfo(papers, true, true, 10)

	wf := GetWordFreq(words)

	wf_clean := HandlePlural(wf)

	topWords := GetTopWords(wf_clean, 30)

	if searchOutput {
		SaveSearchResult(searchOutputFile, papers)
	}

	fmt.Println(topWords)

}
