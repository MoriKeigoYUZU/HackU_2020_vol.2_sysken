package model

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	User_id      int
	Name         string
	Password     string
	Token        string
	Feed_num     int
	Rgister_date string
}

type Users []*User

func InsertUserr(record *User) error {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}
	// userテーブルへのレコードの登録を行うSQLを入力する
	stmt, err := db.Prepare("INSERT INTO users (name, password, token , feed_num ) VALUES (?, ?, ?, ?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.Name, record.Password, record.Token, record.Feed_num)
	return err
}

func SelectUser(name, password string) (Users, error) {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}

	// userテーブルへのレコードの登録を行うSQLを入力する
	// 単レコード取得の書き方がわからない
	// todo 書く
	rows, err := db.Query("SELECT token , feed_num FROM users WHERE name = ? AND password = ?", name, password)
	if err != nil {
		log.Fatal(err)
	}

	return convertToUser(rows)
}

func UpdateFeed(token string, flag, feedNum int) {

	db, err := sql.Open("mysql", "root:root@tcp(hacku_db:3306)/raise_todo")
	if err != nil {
		panic(err.Error())
	}

	updateAfter := 0
	if flag == 0 {
		updateAfter = feedNum - 1
	} else {
		updateAfter = feedNum + 1
	}

	// UPDATE user SET point = 100 WHERE name = "山田";
	_, err = db.Exec("UPDATE users SET feed_num = ? WHERE token = ?", updateAfter, token)
	if err != nil {
		log.Fatal(err)
	}
	//return convertToUserFeed(rows)
}

func convertToUser(rows *sql.Rows) (Users, error) {
	var users Users
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Token, &user.Feed_num); err != nil {
			log.Println(errors.New("scan failed"))
			return nil, err
		}
		//要素を追加
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		log.Println(errors.New("rows scan failed"))
		return nil, err
	}
	return users, nil
}
