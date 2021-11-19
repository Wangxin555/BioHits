package BioHits

/*
func main() {

	fileName := "data.csv"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Can not create a file!")
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("pubmed.ncbi.nlm.nih.gov"),
	)

	c.OnHTML(".docsum-content", func(e *colly.HTMLElement) {

		writer.Write([]string{
			e.ChildText("a"),
			e.ChildAttr("a", "href"),
		})
	})

	for i := 0; i < 100; i++ {
		fmt.Printf("Scraping page : %d\n", i)

		c.Visit("https://pubmed.ncbi.nlm.nih.gov/trending/?page=" + strconv.Itoa(i))

	}

	log.Print("Complete\n")
	log.Println(c)
}

*/
