package sql

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestDeleter(t *testing.T) {
	db, err := sql.Open("sqlite", "file:test.db?cache=shared&mode=memory")
	require.NoError(t, err)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		age INTEGER
	)`)
	require.NoError(t, err)

	res, err := db.ExecContext(ctx, `INSERT INTO users (first_name, last_name, age) VALUES (?, ?, ?)`, "John", "Doe", 30)
	require.NoError(t, err)
	res, err = db.ExecContext(ctx, `INSERT INTO users (first_name, last_name, age) VALUES (?, ?, ?)`, "Joey", "ErPang", 15)
	require.NoError(t, err)

	affected, err := res.RowsAffected()
	require.NoError(t, err)
	log.Println("受影响行数:", affected)
	id, err := res.LastInsertId()
	require.NoError(t, err)
	log.Println("最后插入的Id", id)

	//row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", 1)
	//require.NoError(t, row.Err())
	//tm := &TestModel{}
	//err = row.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age)
	//require.NoError(t, err)
	//log.Println(tm)

	rows, errors := db.QueryContext(ctx, "SELECT * FROM users")
	require.NoError(t, errors)
	for rows.Next() {
		tm := TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age)
		require.NoError(t, err)
		log.Println(tm)
	}
}

func TestPrepareStatement(t *testing.T) {
	db, err := sql.Open("sqlite", "file:test.db?cache=shared&mode=memory")
	require.NoError(t, err)
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		age INTEGER
	)`)
	require.NoError(t, err)

	res, err := db.ExecContext(ctx, `INSERT INTO users (first_name, last_name, age) VALUES (?, ?, ?)`, "John", "Doe", 30)
	require.NoError(t, err)
	res, err = db.ExecContext(ctx, `INSERT INTO users (first_name, last_name, age) VALUES (?, ?, ?)`, "Joey", "ErPang", 15)
	require.NoError(t, err)

	affected, err := res.RowsAffected()
	require.NoError(t, err)
	log.Println("受影响行数:", affected)
	id, err := res.LastInsertId()
	require.NoError(t, err)
	log.Println("最后插入的Id", id)

	//row := db.QueryRowContext(ctx, "SELECT * FROM users WHERE id=?", 1)
	//require.NoError(t, row.Err())
	//tm := &TestModel{}
	//err = row.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age)
	//require.NoError(t, err)
	//log.Println(tm)

	//痛点：硬查询（会导致sql查询连接查询上线），和关闭问题
	stmt, err := db.PrepareContext(ctx, "SELECT * FROM users WHERE id=?")
	require.NoError(t, err)

	rows, errors := stmt.QueryContext(ctx, "SELECT * FROM users")
	require.NoError(t, errors)
	for rows.Next() {
		tm := TestModel{}
		err = rows.Scan(&tm.Id, &tm.FirstName, &tm.LastName, &tm.Age)
		require.NoError(t, err)
		log.Println(tm)
	}
	cancel()
	stmt.Close()
}

type TestModel struct {
	Id        int64
	FirstName string
	Age       uint8
	LastName  string
}
