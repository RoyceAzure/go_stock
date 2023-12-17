package util

import (
	"database/sql"
	"strconv"
	"time"
)

// StringToSqlNiStr 將str轉換成sql.NullInt64
func StringToSqlNiStr(str string) sql.NullString {
	return sql.NullString{String: str, Valid: true}
}

// TimeToSqlNiTime 將Time轉換成sql.NullTime
func TimeToSqlNiTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}

func IntToSqlNiInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

func IntToSqlNiString(i int64) sql.NullString {
	str := strconv.FormatInt(i, 10)
	return sql.NullString{String: str, Valid: true}
}

// sql.Null 就有一個valid 欄位  所以可以根據這個欄位判斷是否轉換成功  不需要額外回傳error
// StringToSqlNiInt64 將str轉換成sql.NullInt64
// 零值字串或轉換失敗將會回傳 Valid為false的sql.NullInt64
func StringToSqlNiInt64(str string) sql.NullInt64 {
	if str == "" {
		return sql.NullInt64{Valid: false}
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: i, Valid: true}
}

// StringToSqlNiTime 將str轉換成sql.NullTime
// str參數是timestamp格式，需要是毫秒級的timestamp
// 零值字串或轉換失敗將會回傳 Valid為false的sql.NullTime
func StringToSqlNiTime(str string) sql.NullTime {
	if str == "" {
		return sql.NullTime{Valid: false}
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return sql.NullTime{Valid: false}
	}
	t := time.Unix(i, 0)
	return sql.NullTime{Time: t, Valid: true}
}

func StringToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
