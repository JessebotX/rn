package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing arguments.")
	}

	inputFile := os.Args[1]
	tags := os.Args[2:]

	sort.Strings(tags)

	// Format: date--title--author__tag1..._tagN
	var title, author, otherTags string

	for _, tag := range tags {
		if strings.HasPrefix(tag, "title=") {
			title = "--" + sanitize(strings.TrimLeft(tag, "title="))
			continue
		}

		if strings.HasPrefix(tag, "name=") {
			author = "--" + sanitize(strings.TrimLeft(tag, "name="))
			continue
		}

		if len(otherTags) <= 0 {
			otherTags = "_"
		}

		otherTags += "_" + sanitize(tag)
	}

	currentTime := time.Now()
	identifier := currentTime.Format("20060102T150405")

	result := identifier + title + author + otherTags + filepath.Ext(inputFile)
	err := os.Rename(inputFile, result)
	if err != nil {
		log.Fatal(err)
	}
}

func sanitize(text string) string {
	excludedPunctuationRegexp := regexp.MustCompile("[][{}!@#$%^&*()=+'\"?,.|;:~`‘’“”/]*")

	textRemovedPunctuation := excludedPunctuationRegexp.ReplaceAllString(text, "")
	textTrimmed := strings.Trim(textRemovedPunctuation, " ")
	textLowercased := strings.ToLower(textTrimmed)
	fillSpaces := strings.ReplaceAll(textLowercased, " ", "-")

	return fillSpaces
}
