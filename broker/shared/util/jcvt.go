package util

import (
	"errors"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"

)

func NumericToFloat64(n pgtype.Numeric) (float64, error) {
	if !n.Valid {
		return 0, errors.New("numeric value not present")
	}

	float8, err := n.Float64Value()
	if err != nil {
		return 0, err
	}
	return float8.Float64, nil
}

func float64ToNumeric(f float64) (pgtype.Numeric, error) {
	// 創建一個新的 Numeric 值
	n := pgtype.Numeric{}

	// 使用 Set 方法設置 float64 值
	if err := n.Scan(f); err != nil {
		// 如果轉換有錯誤，返回錯誤
		return pgtype.Numeric{}, err
	}

	return n, nil
}

func StringToNumeric(value string) (pgtype.Numeric, error) {
	var result pgtype.Numeric
	err := result.Scan(value)
	return result, err
}

func Float64ToString(value float64) string {
	// 第二個參數指定格式，'f' 表示不使用指數形式
	// 第三個參數指定小數點後的位數
	// 第四個參數用於指定要轉換的浮點型號，64 表示 float64
	return strconv.FormatFloat(value, 'f', 2, 64)
}
