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

	tags := getTags(defaultFilePath)

	if val, ok := tags.Tags[tag_name]; ok {

		finalUrl := strings.Replace(val, "{TAG}", parameter, 1)
		err := browser.OpenURL(finalUrl)
		if err != nil {
			fmt.Printf("error opening browser %+v\n", err)
		}
	}
}

func getTags(filePath string) *Tags {
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
	existingTags := getTags(defaultFilePath)

	if existing, ok := existingTags.Tags[tag]; ok {
		fmt.Printf("%v already exists with value %v, it will be replaced by %v", tag, existing, url)
	}

	existingTags.Tags[tag] = url
	saveTags(existingTags, defaultFilePath)

}
