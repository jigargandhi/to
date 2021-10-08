package url

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/browser"
	"gopkg.in/yaml.v2"
)

type Tags struct {
	Tags map[string]string
}

func Go(url, tag string) {
	tags := getTags()

	if val, ok := tags.Tags[url]; ok {

		finalUrl := strings.Replace(val, "{TAG}", tag, 1)
		browser.OpenURL(finalUrl)
	}
}

func getTags() *Tags {
	path, _ := homedir.Expand(("~\\.to\\shortcuts.yml"))
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

func Add(tag, url string) {

}
