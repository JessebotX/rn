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
	var title, otherTags string
	author := ""

	for _, tag := range tags {
		if strings.HasPrefix(tag, "title=") {
			title = "--" + sanitize(strings.TrimPrefix(tag, "title="))
			continue
		}

		if strings.HasPrefix(tag, "name=") {
			authors := strings.Split(strings.TrimPrefix(tag, "name="), ",")

			for _, name := range authors {
				author += "_" + sanitize(name)
			}

			continue
		}

		otherTags += "_" + sanitize(tag)
	}

	if len(title) <= 0 {
		inputFileNoExtension := strings.TrimSuffix(inputFile, filepath.Ext(inputFile))
		title = "--" + sanitize(inputFileNoExtension)
	}

	currentTime := time.Now()
	identifier := currentTime.Format("20060102T150405")

	separator := ""
	if len(author) > 0 || len(otherTags) > 0 {
		separator = "_"
	}

	result := identifier + title + separator + author + otherTags + filepath.Ext(inputFile)
	err := os.Rename(inputFile, result)
	if err != nil {
		log.Fatal(err)
	}
}

func sanitize(text string) string {
	excludedPunctuationRegexp := regexp.MustCompile("[][{}!@#$%^&*()=+'\"?,.|;:~`‘’“”/]*")

	textRemovedPunctuation := excludedPunctuationRegexp.ReplaceAllString(text, "")
	textTrimmed := strings.TrimSpace(textRemovedPunctuation)
	textLowercased := strings.ToLower(textTrimmed)
	fillSpaces := strings.ReplaceAll(textLowercased, " ", "-")

	return fillSpaces
}
