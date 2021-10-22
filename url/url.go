package url

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/browser"
	"gopkg.in/yaml.v2"
)

var (
	defaultFilePath = "~\\.to\\shortcuts.yml"
)

type Tags struct {
	Tags map[string]string
}

func Go(tag_name, parameter string) {

	tags := getTagsFromFile(defaultFilePath)

	if val, ok := tags.Tags[tag_name]; ok {

		finalUrl := strings.Replace(val, "{TAG}", parameter, 1)
		err := browser.OpenURL(finalUrl)
		if err != nil {
			fmt.Printf("error opening browser %+v\n", err)
		}
	} else {
		fmt.Println("matching tags are:")
		for _, v := range selectTagValue(tags, tag_name) {
			fmt.Printf("\t%s:\t%s\n", v, tags.Tags[v])
		}
	}
}

func selectTagValue(tag *Tags, input string) []string {
	matching := make([]string, 0)

	for k, _ := range tag.Tags {
		if strings.Index(k, input) == 0 || levensteinDistance(input, k) <= 3 {
			matching = append(matching, k)
		}
	}
	return matching

}

func getTagsFromFile(filePath string) *Tags {
	path, _ := homedir.Expand(filePath)
	data, err := os.ReadFile(path)

	if err != nil {
		switch e := err.(type) {
		case *os.PathError:
			fmt.Printf("ERROR: %v\n", e)
			defaultText, _ := yaml.Marshal(&Tags{})
			os.WriteFile(path, defaultText, os.ModeAppend)
		default:
			fmt.Println("Unknown error occurred")
			return nil
		}
	}
	tags := &Tags{}
	err = yaml.Unmarshal(data, &tags)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return tags
}

func saveTags(tags *Tags, filePath string) {
	defaultText, _ := yaml.Marshal(tags)
	path, _ := homedir.Expand(filePath)
	os.WriteFile(path, defaultText, os.ModeAppend)
}

func Add(tag, url string) {
	existingTags := getTagsFromFile(defaultFilePath)

	if existing, ok := existingTags.Tags[tag]; ok {
		fmt.Printf("%v already exists with value %v, it will be replaced by %v", tag, existing, url)
	}

	existingTags.Tags[tag] = url
	saveTags(existingTags, defaultFilePath)

}

func levensteinDistance(first string, second string) int {
	firstLen, secondLen := len(first), len(second)

	if firstLen == 0 {
		return secondLen
	}
	if secondLen == 0 {
		return firstLen
	}

	firstLen = firstLen + 1
	secondLen = secondLen + 1
	space := make([][]int, firstLen)

	for i := 0; i < firstLen; i++ {
		space[i] = make([]int, secondLen)

		for j := 0; j < secondLen; j++ {
			space[i][j] = 0
		}
	}
	for i := 0; i < secondLen; i++ {
		space[0][i] = i
	}

	for j := 0; j < firstLen; j++ {
		space[j][0] = j
	}

	for i := 1; i < firstLen; i++ {
		for j := 1; j < secondLen; j++ {
			substitutionCost := 0
			if first[i-1] != second[j-1] {
				substitutionCost = 1
			}

			space[i][j] = min(1+space[i-1][j], 1+space[i][j-1], substitutionCost+space[i-1][j-1])

		}
	}

	return space[firstLen-1][secondLen-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
