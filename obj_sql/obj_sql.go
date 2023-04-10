package obj_sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/enums"
	"log"
	"runtime/debug"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	dbURI := "host=localhost port=5432 user=postgres password=testPass dbname=postgres sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}
}

func readObjWithId(id int64, objectTable string) *sql.Row {
	return db.QueryRow("SELECT * FROM " + objectTable + " WHERE pk = " + fmt.Sprint(id) + ";")
}

func readObjWithUserSplatnet(user int64, splatnet int64, objectTable string) *sql.Row {
	return db.QueryRow("SELECT * FROM " + objectTable + " WHERE user_id = " + fmt.Sprint(user) + " AND splatnet_number = " + fmt.Sprint(splatnet) + " LIMIT 1;")
}

func ReadKeyArrayKey(keyField, objectTable, sort, sortColumn string) ([]int64, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM " + objectTable + " ORDER BY " + sortColumn + " " + sort + ";")
	if err := row.Scan(&count); err != nil {
		log.Println("SELECT COUNT(*) FROM " + objectTable + " ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}
	rows, err := db.Query("SELECT " + keyField + " FROM " + objectTable + " ORDER BY " + sortColumn + " " + sort + ";")
	if err != nil {
		log.Println("SELECT " + keyField + " FROM " + objectTable + " ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}
	arr := make([]int64, count)
	i := 0
	for rows.Next() {
		if err := rows.Scan(&arr[i]); err != nil {
			log.Println("SELECT " + keyField + " FROM " + objectTable + " ORDER BY " + sortColumn + " " + sort + ";")
			debug.PrintStack()
			return nil, fmt.Errorf("%v: %w", err, rows.Close())
		}
		i += 1
	}
	if err := rows.Close(); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return arr, nil
}

func ReadKeyArrayKeyCondition(values []interface{}, conditionColumns []string, keyField, objectTable, sort, sortColumn string) ([]int64, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionColumns, values) + " ORDER BY " + sortColumn + " " + sort + ";")
	if err := row.Scan(&count); err != nil {
		log.Println("SELECT COUNT(*) FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionColumns, values) + " ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}
	rows, err := db.Query("SELECT " + keyField + " FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionColumns, values) + " ORDER BY " + sortColumn + " " + sort + ";")
	if err != nil {
		log.Println("SELECT " + keyField + " FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionColumns, values) + " ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}
	arr := make([]int64, count)
	i := 0
	for rows.Next() {
		if err := rows.Scan(&arr[i]); err != nil {
			log.Println("SELECT " + keyField + " FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionColumns, values) + " ORDER BY " + sortColumn + " " + sort + ";")
			debug.PrintStack()
			return nil, fmt.Errorf("%v: %w", err, rows.Close())
		}
		i += 1
	}
	if err := rows.Close(); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return arr, nil
}

func ReadKeyArrayWithKey(startKey, startKeyField, endKeyField, objectTable, sort, sortColumn string) ([]int64, error) {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM " + objectTable + " WHERE " + startKeyField + " = '" + startKey + "' ORDER BY " + sortColumn + " " + sort + ";").Scan(&count); err != nil {
		log.Println("SELECT COUNT(*) FROM " + objectTable + " WHERE " + startKeyField + " = '" + fmt.Sprint(startKey) + "' ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}

	rows, err := db.Query("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = '" + startKey + "' ORDER BY " + sortColumn + " " + sort + ";")
	if err != nil {
		log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = '" + startKey + "' ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}
	arr := make([]int64, count)
	i := 0
	for rows.Next() {
		if err := rows.Scan(&arr[i]); err != nil {
			log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = '" + startKey + "' ORDER BY " + sortColumn + " " + sort + ";")
			debug.PrintStack()
			return nil, fmt.Errorf("%v: %w", err, rows.Close())
		}
		i += 1
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return arr, nil
}

func ReadKeyArrayWithKeyTable(startKey interface{}, startKeyField, tableName, tableColumnName, endKeyField, objectTable string) ([]int64, error) {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM " + objectTable + " WHERE " + startKeyField + " = " + processValue(startKey) + "AND " + tableColumnName + " = '" + tableName + "';").Scan(&count); err != nil {
		log.Println("SELECT COUNT(*) FROM " + objectTable + " WHERE " + startKeyField + " = " + processValue(startKey) + "AND " + tableColumnName + " = '" + tableName + "';")
		debug.PrintStack()
		return nil, err
	}
	rows, err := db.Query("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = " + processValue(startKey) + "AND " + tableColumnName + " = '" + tableName + "';")
	if err != nil {
		log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = " + processValue(startKey) + "AND " + tableColumnName + " = '" + tableName + "';")
		debug.PrintStack()
		return nil, err
	}
	arr := make([]int64, count)
	i := 0
	for rows.Next() {
		if err := rows.Scan(&arr[i]); err != nil {
			log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = " + processValue(startKey) + "AND " + tableColumnName + " = '" + tableName + "';")
			debug.PrintStack()
			return nil, fmt.Errorf("%v: %w", err, rows.Close())
		}
		i += 1
	}
	if err := rows.Close(); err != nil {
		log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = " + processValue(startKey) + "AND " + tableColumnName + " = '" + tableName + "';")
		debug.PrintStack()
		return nil, err
	}
	return arr, nil
}

func ReadValuesWithKey(key interface{}, keyField, objectTable string, fields []string) *sql.Row {
	return db.QueryRow("SELECT " + createColumnListString(fields) + " FROM " + objectTable + " WHERE " + keyField + " = " + processValue(key) + ";")
}

func ReadKeyArrayWithCondition(condition []interface{}, conditionField []string, endKeyField, objectTable, sort, sortColumn string) ([]int64, error) {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionField, condition) + " ORDER BY " + sortColumn + " " + sort + ";").Scan(&count); err != nil {
		log.Println("SELECT COUNT(*) FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionField, condition) + " ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}
	rows, err := db.Query("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionField, condition) + " ORDER BY " + sortColumn + " " + sort + ";")
	if err != nil {
		log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionField, condition) + " ORDER BY " + sortColumn + " " + sort + ";")
		debug.PrintStack()
		return nil, err
	}
	arr := make([]int64, count)
	i := 0
	for rows.Next() {
		if err := rows.Scan(&arr[i]); err != nil {
			log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + interweaveColumnsValuesCondition(conditionField, condition) + " ORDER BY " + sortColumn + " " + sort + ";")
			debug.PrintStack()
			return nil, fmt.Errorf("%v: %w", err, rows.Close())
		}
		i += 1
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return arr, nil
}

func readKeyArrayWithKeyCondition(startKey int64, condition interface{}, startKeyField, conditionField, endKeyField, objectTable string) ([]int64, error) {
	var count int
	if err := db.QueryRow("SELECT COUNT(*) FROM " + objectTable + " WHERE " + startKeyField + " = " + fmt.Sprint(startKey) + " AND " + conditionField + " = " + processValue(condition) + ";").Scan(&count); err != nil {
		log.Println("SELECT COUNT(*) FROM " + objectTable + " WHERE " + startKeyField + " = " + fmt.Sprint(startKey) + " AND " + conditionField + " = " + processValue(condition) + ";")
		debug.PrintStack()
		return nil, err
	}
	rows, err := db.Query("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = " + fmt.Sprint(startKey) + " AND " + conditionField + " = " + processValue(condition) + ";")
	if err != nil {
		log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = " + fmt.Sprint(startKey) + " AND " + conditionField + " = " + processValue(condition) + ";")
		debug.PrintStack()
		return nil, err
	}
	arr := make([]int64, count)
	i := 0
	for rows.Next() {
		if err := rows.Scan(&arr[i]); err != nil {
			log.Println("SELECT " + endKeyField + " FROM " + objectTable + " WHERE " + startKeyField + " = " + fmt.Sprint(startKey) + " AND " + conditionField + " = " + processValue(condition) + ";")
			debug.PrintStack()
			return nil, fmt.Errorf("%v: %w", err, rows.Close())
		}
		i += 1
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	return arr, nil
}

func writeIntoTable(objectTable string, columnList []string, values []interface{}) error {
	if _, err := db.Exec("INSERT INTO " + objectTable + " (" + createColumnListString(columnList) + ") VALUES(" + fmt.Sprint(processValues(values)...) + ");"); err != nil {
		log.Println("INSERT INTO " + objectTable + " (" + createColumnListString(columnList) + ") VALUES(" + fmt.Sprint(processValues(values)...) + ");")
		debug.PrintStack()
		return err
	}
	return nil
}

func writeIntoTableGetPk(objectTable string, columnList []string, values []interface{}) (*int64, error) {
	result, err := db.Exec("INSERT INTO " + objectTable + " (" + createColumnListString(columnList) + ") VALUES(" + fmt.Sprint(processValues(values)...) + ");")
	if err != nil {
		log.Println("INSERT INTO " + objectTable + " (" + createColumnListString(columnList) + ") VALUES(" + fmt.Sprint(processValues(values)...) + ");")
		debug.PrintStack()
		return nil, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println("INSERT INTO " + objectTable + " (" + createColumnListString(columnList) + ") VALUES(" + fmt.Sprint(processValues(values)...) + ");")
		debug.PrintStack()
		return nil, err
	}
	return &lastId, nil
}

func createColumnListString(columnList []string) string {
	columnListString := ""
	for i := range columnList {
		columnListString += "`" + columnList[i] + "`, "
	}
	return columnListString[0 : len(columnListString)-2]
}

func processValues(values []interface{}) []interface{} {
	resultingValues := make([]interface{}, len(values))
	for i := range values {
		resultingValues[i] = processValue(values[i])
		if i <= len(values)-2 {
			resultingValues[i] = resultingValues[i].(string) + ", "
		}
	}
	return resultingValues
}

func processValue(value interface{}) string {
	switch value.(type) {
	case string:
		return "'" + strings.Replace(value.(string), "'", "''", -1) + "'"
	case enums.GenderEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.GenderEnum)), "'", "''", -1)) + "'"
	case enums.SpeciesEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SpeciesEnum)), "'", "''", -1)) + "'"
	case enums.BattleWeaponEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.BattleWeaponEnum)), "'", "''", -1)) + "'"
	case enums.SalmonWeaponEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonWeaponEnum)), "'", "''", -1)) + "'"
	case enums.SalmonSplatnetScheduleStageImageEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonSplatnetScheduleStageImageEnum)), "'", "''", -1)) + "'"
	case enums.SalmonStageEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonStageEnum)), "'", "''", -1)) + "'"
	case enums.SalmonWeaponScheduleEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonWeaponScheduleEnum)), "'", "''", -1)) + "'"
	case enums.SalmonWeaponScheduleSpecialEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonWeaponScheduleSpecialEnum)), "'", "''", -1)) + "'"
	case enums.BattleStatinkWeaponEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.BattleStatinkWeaponEnum)), "'", "''", -1)) + "'"
	case enums.SalmonWaterLevel:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonWaterLevel)), "'", "''", -1)) + "'"
	case enums.SalmonEvent:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonEvent)), "'", "''", -1)) + "'"
	case enums.SalmonSpecial:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.SalmonSpecial)), "'", "''", -1)) + "'"
	case enums.Lobby:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.Lobby)), "'", "''", -1)) + "'"
	case enums.Rule:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.Rule)), "'", "''", -1)) + "'"
	case enums.FailureReasonEnum:
		return "'" + fmt.Sprint(strings.Replace(fmt.Sprint(value.(enums.FailureReasonEnum)), "'", "''", -1)) + "'"
	case time.Time:
		return "'" + value.(time.Time).Format("2006-01-02 15:04:05") + "'"
	case bool:
		if value.(bool) {
			return "1"
		} else {
			return "0"
		}
	case nil:
		return "NULL"
	case *int64:
		if value.(*int64) == nil {
			return "NULL"
		} else {
			return fmt.Sprint(*value.(*int64))
		}
	case *int:
		if value.(*int) == nil {
			return "NULL"
		} else {
			return fmt.Sprint(*value.(*int))
		}
	case *float64:
		if value.(*float64) == nil {
			return "NULL"
		} else {
			return fmt.Sprint(*value.(*float64))
		}
	case *string:
		if value.(*string) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(*value.(*string), "'", "''", -1) + "'"
		}
	case *enums.GenderEnum:
		if value.(*enums.GenderEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.GenderEnum))), "'", "''", -1) + "'"
		}
	case *enums.SpeciesEnum:
		if value.(*enums.SpeciesEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.SpeciesEnum))), "'", "''", -1) + "'"
		}
	case *enums.BattleWeaponEnum:
		if value.(*enums.BattleWeaponEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.BattleWeaponEnum))), "'", "''", -1) + "'"
		}
	case *enums.SalmonSpecial:
		if value.(*enums.SalmonSpecial) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.SalmonSpecial))), "'", "''", -1) + "'"
		}
	case *enums.SalmonEvent:
		if value.(*enums.SalmonEvent) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.SalmonEvent))), "'", "''", -1) + "'"
		}
	case *enums.SalmonWaterLevel:
		if value.(*enums.SalmonWaterLevel) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.SalmonWaterLevel))), "'", "''", -1) + "'"
		}
	case *enums.FailureReasonEnum:
		if value.(*enums.FailureReasonEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.FailureReasonEnum))), "'", "''", -1) + "'"
		}
	case *enums.SalmonWeaponEnum:
		if value.(*enums.SalmonWeaponEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.SalmonWeaponEnum))), "'", "''", -1) + "'"
		}
	case *enums.SalmonWeaponScheduleEnum:
		if value.(*enums.SalmonWeaponScheduleEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.SalmonWeaponScheduleEnum))), "'", "''", -1) + "'"
		}
	case *enums.SalmonWeaponScheduleSpecialEnum:
		if value.(*enums.SalmonWeaponScheduleSpecialEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.SalmonWeaponScheduleSpecialEnum))), "'", "''", -1) + "'"
		}
	case *enums.BattleStatinkWeaponEnum:
		if value.(*enums.BattleStatinkWeaponEnum) == nil {
			return "NULL"
		} else {
			return "'" + strings.Replace(fmt.Sprint(*(value.(*enums.BattleStatinkWeaponEnum))), "'", "''", -1) + "'"
		}
	case *time.Time:
		if value.(*time.Time) == nil {
			return "NULL"
		} else {
			return "'" + (*value.(*time.Time)).Format("2006-01-02 15:04:05") + "'"
		}
	default:
		return fmt.Sprint(value)
	}
}

func updateTable(objectTable string, columnList []string, values []interface{}, pk int64) error {
	_, err := db.Exec("UPDATE " + objectTable + " SET " + interweaveColumnsValues(columnList, values) + " WHERE pk = " + fmt.Sprint(pk) + ";")
	if err != nil {
		log.Println("UPDATE " + objectTable + " SET " + interweaveColumnsValues(columnList, values) + " WHERE pk = " + fmt.Sprint(pk) + ";")
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
			resultString += "'" + strings.Replace(values[i].(string), "'", "''", -1) + "'"
		case bool:
			if values[i].(bool) {
				resultString += "1"
			} else {
				resultString += "0"
			}
		case nil:
			resultString += "NULL"
		default:
			resultString += fmt.Sprint(values[i])
		}
		if i <= len(values)-2 {
			resultString += ", "
		}
	}
	return resultString
}

func interweaveColumnsValuesCondition(columnList []string, values []interface{}) string {
	resultString := ""
	for i := range columnList {
		switch values[i].(type) {
		case time.Time:
			if strings.Contains(columnList[i], "from") {
				resultString += fmt.Sprintf("%s >= %d", columnList[i][0:len(columnList[i])-4], values[i].(time.Time).Unix())
			} else if strings.Contains(columnList[i], "to") {
				resultString += fmt.Sprintf("%s <= %d", columnList[i][0:len(columnList[i])-2], values[i].(time.Time).Unix())
			}
		case string:
			if values[i].(string) == "any" {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(values[i].(string), "'", "''", -1) + "'"
			}
		case enums.Rule:
			if values[i].(enums.Rule) == enums.AllRules {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.Rule)), "'", "''", -1) + "'"
			}
		case enums.Lobby:
			if values[i].(enums.Lobby) == enums.AnyLobby {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.Lobby)), "'", "''", -1) + "'"
			}
		case enums.Rank:
			if values[i].(enums.Rank) == enums.AnyRank {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.Rank)), "'", "''", -1) + "'"
			}
		case enums.BattleWeaponEnum:
			if values[i].(enums.BattleWeaponEnum) == enums.AnyWeapon {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.BattleWeaponEnum)), "'", "''", -1) + "'"
			}
		case enums.BattleStage:
			if values[i].(enums.BattleStage) == enums.AnyStage {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.BattleStage)), "'", "''", -1) + "'"
			}
		case enums.SalmonStageEnum:
			if values[i].(enums.SalmonStageEnum) == enums.AnySalmonStage {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.SalmonStageEnum)), "'", "''", -1) + "'"
			}
		case enums.SalmonSpecial:
			if values[i].(enums.SalmonSpecial) == enums.AnySalmonSpecial {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.SalmonSpecial)), "'", "''", -1) + "'"
			}
		case enums.FailureReasonEnum:
			if values[i].(enums.FailureReasonEnum) == enums.AnyFailureReason {
				resultString += "true"
			} else {
				resultString += columnList[i] + " = "
				resultString += "'" + strings.Replace(fmt.Sprint(values[i].(enums.FailureReasonEnum)), "'", "''", -1) + "'"
			}
		case enums.TrinaryBool:
			if values[i].(enums.TrinaryBool) == enums.AnyBool {
				resultString += "true"
			} else if values[i].(enums.TrinaryBool) == enums.BoolT {
				resultString += columnList[i] + " = " + "true"
			} else if values[i].(enums.TrinaryBool) == enums.BoolF {
				resultString += columnList[i] + " = " + "false"
			}
		default:
			resultString += columnList[i] + " = "
			resultString += processValue(values[i])
		}
		if i <= len(values)-2 {
			resultString += " AND "
		}
	}
	return resultString
}
