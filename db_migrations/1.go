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
	"database/sql"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/enums"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"runtime/debug"
	"strings"
	"time"
)

var db *sql.DB

func initClient() error {
	var err error
	db, err = sql.Open("mysql", "splatstats_go:KsBOaGwt0uHDslA62SapA0kvGJYYzxfbsixtOkPcvfSfeWhCsh3kL3pctHxPpiE0tERoGLk81LsXHRCaMtybTJoDRcPPz27aJPcFnvqWjLcLFCvzxzDF3VMtI6iCJMyc@/db?parseTime=true")
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.PingContext(ctx)
}

func updateTable(objectTable string, columnList []string, values []interface{}, column, value string) error {
	if db == nil {
		if err := initClient(); err != nil {
			return err
		}
	}
	_, err := db.Exec("UPDATE " + objectTable + " SET " + interweaveColumnsValues(columnList, values) + " WHERE " + column + " = '" + strings.Replace(value, "'", "''", -1) + "';")
	if err != nil {
		log.Println("UPDATE " + objectTable + " SET " + interweaveColumnsValues(columnList, values) + " WHERE " + column + " = '" + strings.Replace(value, "'", "''", -1) + "';")
		debug.PrintStack()
		return err
	}
	return nil
}

func interweaveColumnsValues(columnList []string, values []interface{}) string {
	resultString := ""
	for i := range columnList {
		resultString += columnList[i] + " = "
		switch values[i].(type) {
		case string:
			resultString += "'" + strings.Replace(values[i].(string), "'", "''", 0) + "'"
		case bool:
			if values[i].(bool) {
				resultString += "1"
			} else {
				resultString += "0"
			}
		case nil:
			resultString += "NULL"
		default:
			resultString += strings.Replace(fmt.Sprint(values[i]), "'", "''", -1)
		}
		if i <= len(values)-2 {
			resultString += ", "
		}
	}
	return resultString
}

func main() {
	if err := initClient(); err != nil {
		panic(err)
	}

	columns := []string{"player_w1_weapon", "player_w2_weapon", "player_w3_weapon", "teammate0_w1_weapon",
		"teammate0_w2_weapon", "teammate0_w3_weapon", "teammate1_w1_weapon", "teammate1_w2_weapon",
		"teammate1_w3_weapon", "teammate2_w1_weapon", "teammate2_w2_weapon", "teammate2_w3_weapon",
	}
	weapons := enums.GetShiftWeapons()
	for i := range columns {
		columnList := []string{columns[i]}
		for j := range weapons {
			if err := updateTable("two_salmon_shift",
				columnList,
				[]interface{}{interface{}(string(weapons[j]))},
				columns[i],
				weapons[j].GetDisplay(message.NewPrinter(language.AmericanEnglish)),
			); err != nil {
				log.Println(err)
			}
		}
	}
	_, err := db.Exec("ALTER TABLE `two_salmon_shift` \n    CHANGE COLUMN `player_w1_weapon` `player_w1_weapon` VARCHAR(5) NOT NULL ,\n    CHANGE COLUMN `player_w2_weapon` `player_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `player_w3_weapon` `player_w3_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate0_w1_weapon` `teammate0_w1_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate0_w2_weapon` `teammate0_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate0_w3_weapon` `teammate0_w3_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate1_w1_weapon` `teammate1_w1_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate1_w2_weapon` `teammate1_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate1_w3_weapon` `teammate1_w3_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate2_w1_weapon` `teammate2_w1_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate2_w2_weapon` `teammate2_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\n    CHANGE COLUMN `teammate2_w3_weapon` `teammate2_w3_weapon` VARCHAR(5) NULL DEFAULT NULL")
	if err != nil {
		log.Println("ALTER TABLE `two_salmon_shift` \nCHANGE COLUMN `player_w1_weapon` `player_w1_weapon` VARCHAR(5) NOT NULL ,\nCHANGE COLUMN `player_w2_weapon` `player_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `player_w3_weapon` `player_w3_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate0_w1_weapon` `teammate0_w1_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate0_w2_weapon` `teammate0_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate0_w3_weapon` `teammate0_w3_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate1_w1_weapon` `teammate1_w1_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate1_w2_weapon` `teammate1_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate1_w3_weapon` `teammate1_w3_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate2_w1_weapon` `teammate2_w1_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate2_w2_weapon` `teammate2_w2_weapon` VARCHAR(5) NULL DEFAULT NULL ,\nCHANGE COLUMN `teammate2_w3_weapon` `teammate2_w3_weapon` VARCHAR(5) NULL DEFAULT NULL;")
		panic(err)
	}
}
