package BioHits

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// a paper struct stores information of a scientific publication in PubMed,
// including PMID, title and abstract of the paper
type Paper struct {
	PMID     string
	title    string
	abstract string
}

// GetAbstract takes the PMID of an article and return its abstract
func GetAbstract(PMID string) string {
	// PMID here is in "/000000/" format, which is directly obtained from website
	link := "https://pubmed.ncbi.nlm.nih.gov" + PMID
	c := colly.NewCollector(
		colly.AllowedDomains("pubmed.ncbi.nlm.nih.gov"),
	)

	var abstract string
	// use proper selector to crawl abstract
	c.OnHTML("#enc-abstract", func(e *colly.HTMLElement) {
		abstract = e.ChildText("p")
	})

	c.Visit(link)
	return abstract
}

// FetchPaperInfo is the main function to crawl data and store in paper objects
func FetchPaperInfo(keywords string, numPaper int) []Paper {

	var numPage int
	var Papers = make([]Paper, 0)

	// build a collector for crawling data in webpage including search results
	contentCollector := colly.NewCollector(
		colly.AllowedDomains("pubmed.ncbi.nlm.nih.gov"),
	)

	// build another collector for crawling data in webpage including search results
	// which is only for obtaining the number of results PubMed has found
	numResultCollector := colly.NewCollector(
		colly.AllowedDomains("pubmed.ncbi.nlm.nih.gov"),
	)

	// contentCollector crawls the title and PMID and store the information in paper object
	// with abstract gotten from GetAbstract function
	contentCollector.OnHTML(".docsum-content", func(e *colly.HTMLElement) {

		paperTitle := e.ChildText("a")
		paperPMIDString := e.ChildAttr("a", "href")

		// extract PMID numbers from PMID string using regex
		regularMatch := regexp.MustCompile("[0-9]+")
		paperPMID := regularMatch.FindAllString(paperPMIDString, -1)[0]
		paperAbstract := GetAbstract(paperPMIDString)

		paper := Paper{
			PMID:     paperPMID,
			title:    paperTitle,
			abstract: paperAbstract,
		}

		Papers = append(Papers, paper)
	})

	// because there are two elements under the div.results-amount
	// to avoid visiting twice for same information we jump over the second visit
	count := 1
	numResultCollector.OnHTML(".results-amount", func(e *colly.HTMLElement) {

		if count == 1 {
			// obtain the number of results searched
			numResultSting := e.ChildText("span")
			reg, _ := regexp.Compile("[^0-9]+")
			processedString := reg.ReplaceAllString(numResultSting, "")
			numResults, err := strconv.Atoi(processedString)

			if err != nil {
				// handle error
				fmt.Println(err)
			}

			// set number of paper as the number of results and give a message
			// if the number of results is less than the number user requires
			if numResults < numPaper {
				fmt.Println("The number of paper you input is less than the number of results.")
				fmt.Printf("BioHits will consider the total number of results"+
					" (%d) papers instead.\n", numResults)
				numPaper = numResults
			}

			// set number of papers equal to the nearest integer multiple of 10
			// if numPaper is not an integer multiple of 10
			if numPaper%10 != 0 {
				numPage = int(math.Floor(float64(numPaper)/10) + 1)

				fmt.Println("The number of paper you input is not an integer multiple of 10.")
				fmt.Printf("BioHits will search %d papers for you instead.\n", numPage*10)
			} else {
				// number of pages is equal to the value devided by 10
				// if it is an integer multiple of 10
				numPage = numPaper / 10
			}
		}
		count += 1
	})

	// process keywords parameter
	if keywords == "trending" {
		numResultCollector.Visit("https://pubmed.ncbi.nlm.nih.gov/trending/")

		// visit different pages each time
		for i := 1; i < numPage+1; i++ {
			fmt.Printf("Scraping page : %d\n", i)
			contentCollector.Visit("https://pubmed.ncbi.nlm.nih.gov/trending/?sort=date&page=" +
				strconv.Itoa(i))
		}
		//log.Print(contentCollector)

	} else {

		// split words
		words := strings.Split(keywords, ",")

		// joint words with special characters
		// that can be directed used to visit the webpage
		joinedWords := strings.Join(words, "%20")
		numResultCollector.Visit("https://pubmed.ncbi.nlm.nih.gov/?term=" + joinedWords)

		// visit different pages each time
		for i := 1; i < numPage+1; i++ {
			fmt.Printf("Scraping page : %d\n", i)
			contentCollector.Visit("https://pubmed.ncbi.nlm.nih.gov/?term=" +
				joinedWords + "&sort=date&page=" + strconv.Itoa(i))
		}
		//log.Print(contentCollector)
	}

	//log.Print("Complete fetching data!\n")
	fmt.Println("Complete fetching data!")
	return Papers
}
