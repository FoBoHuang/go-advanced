package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type activity struct {
	Id         int64
	ActivityId int32
	Name       string
	ST         string
	ET         string
}

func initSql() (*sql.DB, error) {
	// 创建数据库连接抽象
	db, err := sql.Open("mysql", "root:980414@tcp(127.0.0.1:3306)/jike_go?charset=utf8")
	if err != nil {
		return db, errors.Wrap(err, "fail to prepares the database abstraction for later use")
	}
	err = db.Ping()
	if err != nil {
		return db, errors.Wrap(err, "fail to connect mysql")
	}
	return db, err
}

func queryData(db *sql.DB, sqlStr string) (activity, error) {
	var ac activity
	rows, err := db.Query(sqlStr)
	if err != nil {
		return ac, errors.Wrap(err, "db query fail")
	}
	//我们先检查错误，只有在没有错误的情况下才调用rows.Close()，以避免出现运行时恐慌。
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ac.Id, &ac.ActivityId, &ac.Name, &ac.ST, &ac.ET)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				err = errors.Wrap(err, "is sql.ErrNoRows")
			default:
				err = errors.Wrap(err, "scan error")
			}
		}
	}
	err = rows.Err()
	if err != nil {
		err = errors.Wrap(err, "sql rows error")
	}
	return ac, err
}

func queryData1(db *sql.DB, sqlStr string, id int64) (activity, error) {
	var ac activity
	err := db.QueryRow(sqlStr, id).Scan(&ac.Id, &ac.ActivityId, &ac.Name, &ac.ST, &ac.ET)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = errors.Wrap(err, "is sql.ErrNoRows")
		default:
			err = errors.Wrap(err, "scan error")
		}
	}
	return ac, err
}
