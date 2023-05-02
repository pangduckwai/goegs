package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

/*
create table test (
 id serial primary key,
 descn varchar(100),
 remarks varchar(150),
 updated_on timestamp not null
);

CREATE TABLE tests (
 m INTEGER NOT NULL,
 n INTEGER NOT NULL,
 o VARCHAR,
 t TIMESTAMP NOT NULL,
 PRIMARY KEY (m, n)
);

insert into tests (m, n, o, t) values (0, 0, 'A', current_timestamp);
insert into tests (m, n, o, t) values (0, 1, 'A', current_timestamp);
insert into tests (m, n, o, t) values (0, 2, 'A', current_timestamp);
*/

// SQL_TEST_SEL select tests
const SQL_TEST_SEL = "SELECT n, o, t FROM tests WHERE m=$1"

// SQL_TESTS_UPS upsert tests
const SQL_TESTS_UPS = "WITH t AS (\n" +
	" INSERT INTO tests (m, n, o, t)\n" +
	" VALUES ($1, $2, $3, current_timestamp)\n" +
	" ON CONFLICT ON CONSTRAINT tests_pkey\n" +
	" DO UPDATE SET\n" +
	"  o = tests.o || EXCLUDED.o,\n" +
	"  t = EXCLUDED.t\n" +
	"  RETURNING xmax)\n" +
	" SELECT SUM(CASE WHEN NOT xmax = 0 THEN 1 ELSE 0 END) AS upd FROM t"

func run() int {
	uid := "paul"
	host := "localhost"
	port := 5432
	dbname := "wdom"
	// dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v pool_min_conns=%v pool_max_conns=%v password=s3crEt", host, port, dbname, uid, 5, 10)
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v pool_min_conns=%v pool_max_conns=%v", host, port, dbname, uid, 5, 10)

	fmt.Printf("URL: %v\n\n", dsn)

	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return 1
	}
	defer conn.Close()

	var m, n int32
	var o string
	var t time.Time
	var u int

	// Update
	err = conn.QueryRow(context.Background(), SQL_TESTS_UPS, 0, 2, ".").Scan(&u)
	if err != nil {
		log.Fatalf("Insert to database failed: %v\n", err)
		return 1
	}
	fmt.Printf("Insert to database: %v\n\n", u)

	// Insert
	err = conn.QueryRow(context.Background(), SQL_TESTS_UPS, 0, 3, "A").Scan(&u)
	if err != nil {
		log.Fatalf("Insert to database failed: %v\n", err)
		return 1
	}
	fmt.Printf("Insert to database: %v\n\n", u)

	// Select
	time.Sleep(1 * time.Second)

	m = 0
	rows, err := conn.Query(context.Background(), SQL_TEST_SEL, m)
	if err != nil {
		log.Fatalf("Unable to query database: %v\n", err)
		return 1
	}
	count := 0
	for rows.Next() {
		count++
		err := rows.Scan(&n, &o, &t)
		if err != nil {
			log.Fatalf("Unable to read row: %v\n", err)
			return 1
		}
		fmt.Printf("%3v> | %5v | %5v | %-36v | %v\n", count, m, n, t, o)
	}
	if count == 0 {
		log.Println("No row selected")
		return 0
	}

	// err = conn.QueryRow(context.Background(), "select * from trees where v=$1", "WYZ").Scan(&vrn, &tms)
	// if err != nil && err != pgx.ErrNoRows {
	// 	log.Fatalf("Unable to query database: %v\n", err)
	// 	os.Exit(1)
	// }
	// fmt.Printf("%-30v | %v \n", vrn, tms)

	return 0
}

func main() {
	x := run()
	time.Sleep(1 * time.Second)
	os.Exit(x)
}
