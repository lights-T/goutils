package single

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/lights-T/goutils"
	"github.com/lights-T/goutils/domain"
	"github.com/rs/zerolog"
)

type SqlServerDB struct {
	DB *goqu.Database
}

func NewSqlServer(dbConf *domain.MssqlConf) (*SqlServerDB, error) {
	if dbConf == nil {
		return nil, errors.New("Database config is empty ")
	}
	address := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", dbConf.Server, dbConf.UserName, dbConf.Password, dbConf.Port, dbConf.DbName)
	dialect := goqu.Dialect("sqlserver")
	msDb, err := sql.Open("sqlserver", address)
	if err != nil {
		return nil, err
	}
	sqlServerDB := dialect.DB(msDb)
	l := zerolog.New(os.Stdout)
	sqlServerDB.Logger(&l)
	if err = msDb.Ping(); err != nil {
		return nil, goutils.Errorf("Db connection failure, err: %s", err.Error())
	}
	return &SqlServerDB{DB: sqlServerDB}, nil
}

func (c *SqlServerDB) Query(exps string) (int64, error) {
	var count int64
	var err error
	var TableName string
	rows, err := c.DB.Query(fmt.Sprintf("select count(*) count from %s where (%s)", TableName, exps))
	if err != nil {
		return count, err
	}

	defer rows.Close()
	for rows.Next() {
		//要与sqlserver数据库字段顺序一致
		if err := rows.Scan(&count); err != nil {
			return count, err
		}
	}
	if rows.Err() != nil {
		return count, err
	}

	return count, nil
}
