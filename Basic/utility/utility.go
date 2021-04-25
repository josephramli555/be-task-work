package utility

import (
	"fmt"
	"strings"
	// "unicode"
)

func ReverseUpperLower(input string) string {
	ans := ""
	for _, char := range input {
		if int(char) >= 65 && int(char) <= 90 {
			ans = ans + string((int(char) + 32))
		} else if (char) >= 97 && int(char) <= 122 {
			ans = ans + string((int(char) - 32))
		} else {
			ans = ans + string(char)
		}
	}
	return ans
}

type WordStats struct {
	WordCounters       map[string]int
	TotalWordCount     int
	TotalWordShowOnce  int
	HighestCountWords  []string
	SmallestCountWords []string
}

func newWordStats(WordCounters map[string]int, TotalWordCount int, TotalWordShowOnce int, HighestCountWords []string, SmallestCountWords []string) WordStats {
	return WordStats{
		WordCounters:       WordCounters,
		TotalWordCount:     TotalWordCount,
		TotalWordShowOnce:  TotalWordShowOnce,
		HighestCountWords:  HighestCountWords,
		SmallestCountWords: SmallestCountWords,
	}
}

func ProcessWord(input string) WordStats {
	a := strings.Split(input, " ")
	wordMap := mappingWord(a)
	var highestCountWords, smallestCountWords []string
	var totalWord, highCount, smallCount, TotalWordShowOnce int = 0, 0, 0, 0
	for key, value := range wordMap {
		if key == "and" {
			fmt.Println(totalWord)
			fmt.Println("Horee")
		}
		if totalWord == 0 {
			highCount = value
			smallCount = value
			highestCountWords = append(highestCountWords, key)
			smallestCountWords = append(smallestCountWords, key)
		}
		if totalWord != 0 {
			//update if there's word with highest count
			if key == "and" {
				fmt.Println(totalWord)
				fmt.Println("Horee")
			}
			if value > highCount {
				highCount = value
				highestCountWords = nil
				highestCountWords = append(highestCountWords, key)
			} else if value == highCount {
				highestCountWords = append(highestCountWords, key)
			}
			if value < smallCount {
				smallCount = value
				smallestCountWords = nil
				smallestCountWords = append(smallestCountWords, key)
			} else if value == smallCount {
				smallestCountWords = append(smallestCountWords, key)
			}
		}
		if value == 1 {
			TotalWordShowOnce++
		}
		totalWord += value
	}
	return newWordStats(wordMap, totalWord, TotalWordShowOnce, highestCountWords, smallestCountWords)
}

func mappingWord(words []string) map[string]int {
	ans := ""
	wordMap := make(map[string]int)
	for _, word := range words {
		ans = ""
		for _, char := range word {
			if (int(char) >= 65 && int(char) <= 90) || (int(char) >= 97 && int(char) <= 122) || (int(char) >= 48 && int(char) <= 57) {
				ans = ans + string(char)
			} else {
				break
			}
		}
		if ans == ""{
			continue
		}
		ans = strings.ToLower(ans)
		if _, exist := wordMap[ans]; exist {
			wordMap[ans] += 1
		} else {
			wordMap[ans] = 1
		}
	}

	return wordMap
}
