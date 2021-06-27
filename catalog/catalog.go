package catalog

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io/ioutil"
	"log"
)

func newDictFromFile(fileName string) map[string]string {
	dict := map[string]string{}
	b, err := ioutil.ReadFile(fmt.Sprintf("locales/%s/messages.json", fileName))
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &dict)
	return dict
}

func init() {
	dict := map[language.Tag]map[string]string{
		//language.AmericanEnglish: newDictFromFile("en-US"),
		language.Japanese: newDictFromFile("ja-JP"),
	}
	for i := range dict {
		for j := range dict[i] {
			if err := message.SetString(i, j, dict[i][j]); err != nil {
				log.Println(err)
			}
		}
	}
}
