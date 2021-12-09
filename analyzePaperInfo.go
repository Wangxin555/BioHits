package BioHits

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

// AnalyzePaperInfo takes in a list of papers and returns a cleaned word map frequency
// papers is a list of paper structs gotten from FetchPaperInfo function
// includeTitle stands for whether to include paper title when analysis
// ignoreWord controls whether ignore additional words other than stop words
// ignoreWordsList and stopWordsList is the directory to ignore words and stop words
func AnalyzePaperInfo(papers []Paper, includeTitle, ignoreWord bool,
	stopWordsList, ignoreWordsList string) map[string]int {
	// split paragraphs into sentences
	paperSentences := make([]string, 0)
	// user can include the title in analysis
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
	// user can use the stopwords.txt provided by BioHits,
	// or use a customized txt file as an input
	stopWords := ReadWordsList(stopWordsList)

	// turn sentences into strings
	infoWords := make([]string, 0)
	for _, sentence := range paperSentences {
		currentWords := strings.Split(sentence, " ")
		// clean every word before creating a map
		// including removing all special symbols in each word,
		// transform uppercase letter into lowercase letters
		// and discard words in stopwords.txt
		for _, word := range currentWords {
			cleanedword := CleanWord(word)
			cleanedword = strings.ToLower(cleanedword)
			if (!(StringInList(cleanedword, stopWords))) &&
				len([]rune(cleanedword)) != 0 {
				infoWords = append(infoWords, cleanedword)
			}
		}
	}
	// wash words deeply by removing ignored words and words with no letters
	infoWords = DeepClean(infoWords, ignoreWord, ignoreWordsList)

	wordFrequency := GetWordFreq(infoWords)
	// handle plurals and adverbs in wordTransform function
	wordFrequencyClean := WordTransform(wordFrequency)

	return wordFrequencyClean
}

// RemoveSpecialChar function removes \t and \n in a string
func RemoveSpecialChar(s string) string {
	tabChar := regexp.MustCompile(`\t`)
	newlineChar := regexp.MustCompile(`\n`)
	s1 := tabChar.ReplaceAllString(s, "")
	s2 := newlineChar.ReplaceAllString(s1, "")
	return s2
}

// cleanWord function removes all special characters in a string
// including "," and "." anchored at the beginning or the end of a string,
// ";" and ":" anchored at the beginning or the end of a string
// and any form of parentheses and brackets
func CleanWord(s string) string {
	wsChar := regexp.MustCompile(`\s`)
	commaChar := regexp.MustCompile(`^,|,$`)
	dotChar := regexp.MustCompile(`^\.|\.$`)
	semicolonChar := regexp.MustCompile(`^;|;$`)
	colonChar := regexp.MustCompile(`^:|:$`)
	parenthesesLeft := regexp.MustCompile(`^\(|\($`)
	parenthesesRight := regexp.MustCompile(`^\)|\)$`)
	quotationChar := regexp.MustCompile(`^\"|\"$`)
	apostropheChar := regexp.MustCompile(`^\'|\'$`)
	bracketLeft := regexp.MustCompile(`^\[|\[$`)
	bracketRight := regexp.MustCompile(`^\]|\]$`)
	bparenthesesLeft := regexp.MustCompile(`^\{|\{$`)
	bparenthesesRight := regexp.MustCompile(`^\}|\}$`)

	// remove special character
	s1 := wsChar.ReplaceAllString(s, "")
	s2 := commaChar.ReplaceAllString(s1, "")
	s3 := dotChar.ReplaceAllString(s2, "")
	s4 := semicolonChar.ReplaceAllString(s3, "")
	s5 := colonChar.ReplaceAllString(s4, "")
	s6 := parenthesesLeft.ReplaceAllString(s5, "")
	s7 := parenthesesRight.ReplaceAllString(s6, "")
	s8 := quotationChar.ReplaceAllString(s7, "")
	s9 := apostropheChar.ReplaceAllString(s8, "")
	s10 := bracketLeft.ReplaceAllString(s9, "")
	s11 := bracketRight.ReplaceAllString(s10, "")
	s12 := bparenthesesLeft.ReplaceAllString(s11, "")
	s13 := bparenthesesRight.ReplaceAllString(s12, "")
	return s13
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

// read words list from txt file and store in a list
func ReadWordsList(filename string) []string {
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

// get word frequencies of a list of strings
func GetWordFreq(oriWords []string) (wordFreq map[string]int) {
	// remove all empty words in words list
	words := RemoveEmptyString(oriWords)
	// create a map to store words and their frequencies
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
// including removing words with no letters or listed in ignoreWordsList
func DeepClean(words []string, ignoreWords bool, ignoreWordsList string) []string {
	// remove string that does not contain any character from a to z
	letters := regexp.MustCompile(`[a-z]`)
	cleanedWords := words

	// read words to be ignored
	if ignoreWords {
		ignoreWords := ReadWordsList(ignoreWordsList)
		for i := range words {
			stringWithLetter := letters.FindAllString(words[i], -1)
			// set a word to empty if it has no letters in it
			// or it shows in ignoredWordList
			if len(stringWithLetter) == 0 || StringInList(words[i], ignoreWords) {
				cleanedWords[i] = ""
			}
		}
	} else {
		for i := range words {
			stringWithLetter := letters.FindAllString(words[i], -1)
			if len(stringWithLetter) == 0 {
				cleanedWords[i] = ""
			}
		}
	}

	return cleanedWords
}

// GetTopWords returns the top num words in a word map
func GetTopWords(wordFreq map[string]int, num int) []string {
	// warn if the number of words in word map is less than the number that user requires
	if len(wordFreq)-1 < num {
		fmt.Printf("Warning: You want top %d words but only %d words detected.\n",
			num, len(wordFreq)-1)
		fmt.Printf("So we will return you %d top words instead.", len(wordFreq)-1)
	}
	// create a list to store top words
	topWords := make([]string, 0)
	values := make([]int, 0)
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

// RemoveEmptyString removes empty strings "" for a given list of strings
func RemoveEmptyString(words []string) []string {
	var cleanedWords []string
	for _, word := range words {
		if word != "" {
			cleanedWords = append(cleanedWords, word)
		}
	}
	return cleanedWords
}

// GetKeys function returns keys of a map
func GetKeys(m map[string]int) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
