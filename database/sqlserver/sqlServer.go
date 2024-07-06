package sqlserver

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/lights-T/goutils"
)

//CheckFiledIsExist 检查数据表中字段是否存在
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

//AuthCreateDb 自动创建数据库，已存在不会继续处理
func AuthCreateDb(ctx context.Context, masterAddress, targetDbName string) error {
	dialect := goqu.Dialect("sqlserver")
	msDb, err := sql.Open("sqlserver", masterAddress)
	if err != nil {
		return goutils.Errorf("Failed to open db, err: %s", err.Error())
	}
	masterDB := dialect.DB(msDb)
	if err = msDb.Ping(); err != nil {
		return goutils.Errorf("Db connection failure, err: %s", err.Error())
	}

	self := &struct {
		DBId      int64  `db:"dbid" json:"dbId,omitempty" goqu:"pk,skipinsert,skipupdate"`
		Name      string `db:"name" json:"name,omitempty" goqu:"defaultifempty"`
		Sid       string `db:"sid" json:"sid,omitempty" goqu:"defaultifempty"`
		Mode      string `db:"mode" json:"mode,omitempty" goqu:"defaultifempty"`
		Status    string `db:"status" json:"status,omitempty" goqu:"defaultifempty"`
		Status2   string `db:"status2" json:"status2,omitempty" goqu:"defaultifempty"`
		Crdate    string `db:"crdate" json:"crdate,omitempty" goqu:"defaultifempty"`
		Reserved  string `db:"reserved" json:"reserved,omitempty" goqu:"defaultifempty"`
		Category  int    `db:"category" json:"category,omitempty" goqu:"defaultifempty"`
		Cmptlevel int    `db:"cmptlevel" json:"cmptlevel,omitempty" goqu:"defaultifempty"`
		Filename  string `db:"filename" json:"filename,omitempty" goqu:"defaultifempty"`
		Version   int    `db:"version" json:"version,omitempty" goqu:"defaultifempty"`
	}{}
	//ColumnFields := []interface{}{"name"}
	targetBD := targetDbName
	exps := exp.NewExpressionList(exp.AndType)
	exps = exps.Append(goqu.I("name").Eq(targetBD))
	if _, err := masterDB.From(goqu.T("sysdatabases").Schema("sys")).
		Prepared(true).
		Select("name").
		Where(exps).
		ScanStructContext(ctx, self); err != nil {
		return goutils.Errorf("Failed to get db, err: %s", err.Error())
	}
	if self == nil || len(self.Name) == 0 {
		if _, err := masterDB.Exec(fmt.Sprintf("CREATE DATABASE %s", targetBD)); err != nil {
			return goutils.Errorf("Failed to create db, err: %s", err.Error())
		}
	} else {
		fmt.Println("The database already exists.")
	}
	return nil
}
