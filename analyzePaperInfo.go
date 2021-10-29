package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// AnalyzePaperInfo takes in a list of papers and return top words based on word frequency
func AnalyzePaperInfo(papers []Paper, includeTitle bool, numHits int) []string {
	// split paragraphs into sentences
	paperSentences := make([]string, 0)
	if includeTitle {
		for i := range papers {
			processedAbstract := RemoveSpecialChar(papers[i].abstract)
			abstractSentences := strings.Split(processedAbstract, ".")
			paperSentences = append(paperSentences, abstractSentences...)
			processedTitle := RemoveSpecialChar(papers[i].title)
			paperSentences = append(paperSentences, processedTitle)
		}
	} else {
		for i := range papers {
			processedAbstract := RemoveSpecialChar(papers[i].abstract)
			abstractSentences := strings.Split(processedAbstract, ".")
			paperSentences = append(paperSentences, abstractSentences...)
		}
	}

	// read stop words from txt file
	stopWords := ReadStopWords("stopwords.txt")

	// turn sentences into strings
	infoWords := make([]string, 0)
	for _, sentence := range paperSentences {
		currentWords := strings.Split(sentence, " ")
		for _, word := range currentWords {
			cleanedword := CleanWord(word)
			cleanedword = strings.ToLower(cleanedword)
			if (!(StringInList(cleanedword, stopWords))) &&
				len([]rune(cleanedword)) != 0 {
				infoWords = append(infoWords, cleanedword)
			}
		}
	}

	return infoWords
}

// RemoveSpecialChar function removes \t and \n in a string
func RemoveSpecialChar(s string) string {
	tabChar := regexp.MustCompile(`\t`)
	newlineChar := regexp.MustCompile(`\n`)
	// word := "what \t is\n"
	s1 := tabChar.ReplaceAllString(s, "")
	s2 := newlineChar.ReplaceAllString(s1, "")
	return s2
}

// cleanWord function removes all white space in a string
// also remove "," and "." anchored at the beginning or the end of a string
// and remove ";" and ":" anchored at the beginning or the end of a string
func CleanWord(s string) string {
	wsChar := regexp.MustCompile(`\s`)
	commaChar := regexp.MustCompile(`^,|,$`)
	dotChar := regexp.MustCompile(`^\.|\.$`)
	semicolonChar := regexp.MustCompile(`^;|;$`)
	colonChar := regexp.MustCompile(`^:|:$`)
	s1 := wsChar.ReplaceAllString(s, "")
	s2 := commaChar.ReplaceAllString(s1, "")
	s3 := dotChar.ReplaceAllString(s2, "")
	s4 := semicolonChar.ReplaceAllString(s3, "")
	s5 := colonChar.ReplaceAllString(s4, "")
	return s5
}

// StringInList function returns whether a string is in a list
func StringInList(s string, list []string) bool {
	for _, element := range list {
		if element == s {
			return true
		}
	}
	return false
}

// read stop words from txt file
func ReadStopWords(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// build a list to store stop words
	stopWords := make([]string, 0)
	for scanner.Scan() {
		stopWords = append(stopWords, scanner.Text())
	}
	return stopWords
}

// get frequencies of a list of strings
func GetWordFreq(words []string) (wordFreq map[string]int) {
	wordFreq = make(map[string]int)
	if len(words) == 1 {
		wordFreq[words[0]] = 1
		return wordFreq
	} else if len(words) == 0 {
		panic("No word found!")
	} else {
		var recordedWord []string
		for _, s := range words {
			if StringInList(s, recordedWord) {
				wordFreq[s] += 1
			} else {
				wordFreq[s] = 1
				recordedWord = append(recordedWord, s)
			}
		}
		return wordFreq
	}
}

// DeepClean further process the words to get useful information
func DeepClean(words []string) []string {
	/*
		k := make([]string, 0)
	*/
	return k
}

//
func GetTopWords(wordFreq map[string]int, num int) []string {
	// create a list to store top words
	topWords := make([]string, num)
	values := make([]int, len(wordFreq))
	// find top words based on the word frequency
	for _, v := range wordFreq {
		values = append(values, v)
	}
	// sort values in increasing order
	sort.Ints(values)

	// get words with top frequencies
	topFreq := values[len(values)-num:]

	for w := range wordFreq {
		if IntinList(wordFreq[w], topFreq) {
			topWords = append(topWords, w)
		}
	}

	return topWords
}

// IntinList determines whether an int is in a list or not
func IntinList(i int, ints []int) bool {
	if len(ints) == 0 {
		return false
	} else {
		for item := range ints {
			if i == ints[item] {
				return true
			}
		}
		return false
	}
}
