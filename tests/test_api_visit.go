/*
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"context"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/obj_sql"
	"net/http"
	"time"
)

func main() {
	shifts, err := obj_sql.ReadKeyArrayKey("splatnet_number", "two_salmon_shift", "desc", "play_time")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	for i := range shifts {
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://localhost:8080/api/two_salmon/13/%d", shifts[i]), nil)
		if err != nil {
			panic(err)
		}
		if _, err := client.Do(req); err != nil {
			panic(err)
		}
	}
	battles, err := obj_sql.ReadKeyArrayKey("splatnet_number", "two_battles_battle", "desc", "time")
	if err != nil {
		panic(err)
	}
	for i := range battles {
		ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
		req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("http://localhost:8080/api/two_battles/13/%d", battles[i]), nil)
		if err != nil {
			panic(err)
		}
		if _, err := client.Do(req); err != nil {
			panic(err)
		}
	}
}