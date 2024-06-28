package sqlserver

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

func CheckFiledIsExist(DB *goqu.Database, tableName, filed string) (bool, error) {
	isExist := false
	var count int64
	var err error
	rows, err := DB.Query(
		fmt.Sprintf("select count(*) count from INFORMATION_SCHEMA.COLUMNS where (TABLE_NAME = '%s' and COLUMN_NAME = '%s')",
			tableName, filed))
	if err != nil {
		return isExist, err
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return isExist, err
		}
	}
	if rows.Err() != nil {
		return isExist, err
	}

	if count > 0 {
		isExist = true
	}

	return isExist, nil
}
