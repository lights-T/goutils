package master

import (
	"fmt"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlserver"
	//_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	//_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gotest.tools/assert"
)

var (
	masterConf = &Conf{
		InstanceName: "mssql_rw",
		DriverName:   "sqlserver",
		DataSource:   fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", "192.168.0.210", "sa", "ROot^123", 1433, "demo"),
	}
	slaverConf = &Conf{
		InstanceName: "mssql_r",
		DriverName:   "sqlserver",
		DataSource:   fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", "192.168.0.210", "sa", "ROot^123", 1433, "demo"),
		IsSlave:      true,
	}
	curInstance = "mssql_rw"
)

func TestNew(t *testing.T) {
	cs := []*Conf{masterConf, slaverConf}
	_, err := New(cs)
	assert.NilError(t, err)
}

func TestDB_GetInstance(t *testing.T) {
	cs := []*Conf{masterConf}
	db, err := New(cs)
	assert.NilError(t, err)

	d := db.GetInstance("")
	t.Log(d)
}

type Reminders struct {
	Title       string
	Description string
	Alias       string
}

func TestNew2(t *testing.T) {
	cs := []*Conf{masterConf, slaverConf}
	db, err := New(cs)
	assert.NilError(t, err)

	d := db.GetInstance(curInstance)
	t.Log(d)

	u := &Reminders{
		Title:       "123",
		Description: "123",
		Alias:       "123",
	}
	r, err := d.Insert("Reminders").
		Prepared(true).
		Cols("title,description,alias").
		Rows(u).
		//OnConflict(goqu.DoUpdate("title", goqu.Record{"title": u.Age})).
		Executor().Exec()
	assert.NilError(t, err)
	t.Log(r)
}

func TestNew3(t *testing.T) {
	cs := []*Conf{masterConf, slaverConf}
	db, err := New(cs)
	assert.NilError(t, err)

	d := db.GetInstance(curInstance)
	t.Log(d)

	u := &Reminders{
		Title:       "123",
		Description: "123",
		Alias:       "123",
	}

	err = d.WithTx(func(tx *goqu.TxDatabase) error {
		_, err := tx.Insert("t_user").
			Prepared(true).
			Rows(u).
			OnConflict(goqu.DoUpdate("name", goqu.Record{"age": u.Title})).
			Executor().Exec()
		return err
	})
	assert.NilError(t, err)
}
