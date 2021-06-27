/* This file is part of SplatStatsGo.
 *
 * SplatStatsGo is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * SplatStatsGo is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with SplatStatsGo.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/routing"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io/ioutil"
	"log"
	"net/http"
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

var dict = map[language.Tag]map[string]string{
	//language.AmericanEnglish: newDictFromFile("en-US"),
	language.Japanese: newDictFromFile("ja-JP"),
}

func init() {
	for i := range dict {
		for j := range dict[i] {
			if err := message.SetString(i, j, dict[i][j]); err != nil {
				log.Println(err)
			}
		}
	}
}

func main() {
	router := routing.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
