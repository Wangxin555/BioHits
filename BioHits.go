package BioHits

import (
	"fmt"
)

func BioHits() {

	// initiate parameters taken by BioHits
	/*
		var numPaper int
		var numWords int
		var keyWords string
		var searchOutputFile string
		var searchOutput bool
		var stopWordsLit string
		var ignoreWordsList string

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

	*/
	papers := FetchPaperInfo("COVID-19", 10)

	wordsFreq := AnalyzePaperInfo(papers, true, true,
		"../BioHits/stopwords/stopwords.txt", "../BioHits/stopwords/ignorewords.txt")

	wf := GetTopWords(wordsFreq, 20)

	//wf_clean := BioHits.WordTransform(wf)

	//topWords := GetTopWords(wf_clean, 30)
	//wordCounts := map[string]int{"k": 50, "important": 42, "noteworthy": 30, "meh": 3}
	SaveSearchResult("test.txt", papers)

	fmt.Println(wf)

	DrawWordCloud(wordsFreq, "test.png", "../BioHits/fonts/Roboto/Roboto-Black.ttf")
}
