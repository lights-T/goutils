package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

func AutoCreateSqlServerDb(ctx context.Context, targetDbName, masterAddress string) error {
	dialect := goqu.Dialect("sqlserver")
	msDb, err := sql.Open("sqlserver", masterAddress)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to open db, err: %s", err.Error()))
	}
	masterDB := dialect.DB(msDb)
	if err = msDb.Ping(); err != nil {
		return errors.New(fmt.Sprintf("Db connection failure, err: %s", err.Error()))
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
		return errors.New(fmt.Sprintf("Failed to get db, err: %s", err.Error()))
	}
	if self == nil || len(self.Name) == 0 {
		if _, err := masterDB.Exec(fmt.Sprintf("CREATE DATABASE %s", targetBD)); err != nil {
			return errors.New(fmt.Sprintf("Failed to create db, err: %s", err.Error()))
		}
	} else {
		fmt.Println("The database already exists.")
	}
	return nil
}
