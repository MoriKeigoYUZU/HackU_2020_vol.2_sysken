package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Task struct {
	TaskId       int
	UserId       int
	SubjectId    int
	Name         string
	Limit        string
	EndFlag      int
	RegisterDate string
	EndDate      string
}

type Tasks []*Task

func SelectGettingTodo(token string) (Tasks, error) {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	var rowUserId string
	if err := db.QueryRow("SELECT user_id FROM users WHERE token = ?", token).Scan(&rowUserId); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM tasks WHERE user_id = ? ", rowUserId)
	//rows, err := db.Query("SELECT * FROM tasks WHERE user_id = ?", user_id)
	if err != nil {
		return nil, err
	}
	return convertToTodos(rows)
}

func convertToTodos(rows *sql.Rows) (Tasks, error) {
	var tasks Tasks
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.TaskId, &task.UserId, &task.SubjectId, &task.Name, &task.Limit, &task.EndFlag, &task.RegisterDate, &task.EndDate); err != nil {
			log.Println(errors.New("scan failed"))
			return nil, err
		}
		//要素を追加
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		log.Println(errors.New("rows scan failed"))
		return nil, err
	}
	return tasks, nil
}

//tokenをuser_idに変換する
func SelectUserId(token string) int {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	var user_id string
	// userテーブルへのレコードの登録を行うSQLを入力する
	// 単レコード取得の書き方がわからない
	// todo 書く
	if err := db.QueryRow("SELECT user_id FROM users WHERE token = ? ", token).Scan(&user_id); err != nil {
		log.Fatal(err)
	}
	fmt.Println(user_id, token)
	var i int
	i, _ = strconv.Atoi(user_id)
	return i
}

func InsertTodo(record *Task) error {
	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	// userテーブルへのレコードの登録を行うSQLを入力する
	stmt, err := db.Prepare("INSERT INTO tasks (user_id, subject_id, name, time_limite, end_flag , end_date ) VALUES ( ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.UserId, record.SubjectId, record.Name, record.Limit, 0, "0")
	return err
}

func InsertTodoEnd(record *Task) error {
	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	// userテーブルへのレコードの登録を行うSQLを入力する
	// "UPDATE tbl_users set password = ? where name = ? "
	stmt, err := db.Prepare("UPDATE tasks SET end_flag = 1 WHERE user_id = ? AND subject_id = ? AND name = ? AND time_limite = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.UserId, record.SubjectId, record.Name, record.Limit)
	return err
}
