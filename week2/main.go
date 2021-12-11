package main

import "fmt"

func main() {
	db, errSql := initSql()
	if errSql != nil {
		fmt.Printf("%+v", errSql, "\n")
	}
	defer db.Close()
	ac, errQuery := queryData(db, "select * from activity where id = 1")
	fmt.Println("activity data is :", ac)
	if errQuery != nil {
		fmt.Printf("%+v", errQuery, "\n")
	}
}
