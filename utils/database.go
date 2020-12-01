package utils

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func checkfile() error {
	if IsExist(Dir+"/data/sqlite3.db") == false {
		f, err := os.Create(Dir + "/data/sqlite3.db")
		defer f.Close()
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func Dbinit() error {
	err := checkfile()
	if err != nil {
		return err
	}
	Db, err = sql.Open("sqlite3", Dir+"/data/sqlite3.db")
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}

	sqlinit := `
CREATE TABLE IF NOT EXISTS "scan_result" (
	"ip" VARCHAR(15) NOT NULL,
	"port" INT(5) NOT NULL,
	"time" INT(10) NOT NULL,
	"status" INT(1) default(0),
	"mag" INT(6) default(0),
	constraint pk_t2 primary key (ip,port)
);
`
	_, err = Db.Exec(sqlinit)
	if err != nil {
		return err
	}
	return nil
}

func Insert(ip string, port int, mag int, status int) error {
	stmt, err := Db.Prepare("replace into scan_result(ip, port, time, status, mag) VALUES (?,?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ip, port, time.Now().Unix(), status, mag)
	if err != nil {
		return err
	}
	return nil
}

// 值传递和引用传递
/*
	这里暂时还是没想好怎么写，就直接先limit 50000吧
*/
func Query(port int, status int, ctx context.Context) ([]string, error) {
	var results []string
	var result string
	rows, err := Db.QueryContext(ctx, "SELECT ip from scan_result where port="+strconv.Itoa(port)+" and status="+strconv.Itoa(status)+" limit 50000")
	if err != nil {
		return results, err
	}
	for rows.Next() {
		select {
		case <-ctx.Done():
			return results, nil
		default:
			err = rows.Scan(&result)
			if err != nil {
				return results, err
			}
			results = append(results, result)
		}
	}
	return results, nil
}
