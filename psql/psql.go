package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

/*
create table test (
	id serial primary key,
	descn varchar(100),
	remarks varchar(150),
	updated_on timestamp not null
);

insert into test (descn, remarks, updated_on) values ('Hello again', 'Hello There!!', current_timestamp);

select * from test
*/
func main() {
	uid := "wdom"
	// pwd := ""
	host := "192.168.56.42"
	port := 5432
	dbname := "wdom"

	// spr := ""
	// if pwd != "" {
	// 	spr = ":"
	// }
	// url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", uid, pwd, host, port, dbname)
	// url := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v", host, port, dbname, uid, pwd)
	url := fmt.Sprintf("host=%v port=%v dbname=%v user=%v", host, port, dbname, uid)

	fmt.Println("URL: ", url)

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select descn, remarks, updated_on from test where descn = $1", "Hello again")
	if err != nil {
		log.Fatalf("Unable to query database: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var dsc, rmk string
		var tms time.Time
		err := rows.Scan(&dsc, &rmk, tms)
		if err != nil {
			log.Fatalf("Unable to read row: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v | %v | %v \n", dsc, rmk, tms)
	}
}
