package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const CodeErrorLogin = 11001

type User struct {
	Id        int    `json:"id"`
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Passwd    string `json:"passwd"`
	UserSex   string `json:"user_sex"`
	UserAge   int    `json:"user_age"`
	UserLevel int    `json:"user_level"`
}

func query(userId string) (*User, error) {
	query := fmt.Sprintf("SELECT * FROM user WHERE user_id IN ('%s') LIMIT 1", userId)
	fmt.Println(query)
	rows, err := Conn.Query(query)
	var user = new(User)
	if err == nil {
		if rows.Next() {
			err = rows.Scan(&user.Id, &user.UserId, &user.UserName, &user.Passwd, &user.UserSex, &user.UserAge, &user.UserLevel)
			fmt.Println(fmt.Sprintf("%+v", &user))
		}
	}
	return user, err
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		userId := r.Form.Get("user_id")
		if len(userId) > 0 {
			user, e := query(userId)
			if e != nil {
				_, _ = w.Write([]byte(e.Error()))
				return
			}
			passwd := r.Form.Get("passwd")
			if len(passwd) > 0 && passwd == user.Passwd {
				body := BodyUser{Code: CodeSuccess, Msg: "登录成功", User: user}
				fmt.Println(body)
				d, _ := json.Marshal(body)
				_, _ = w.Write(d)
			} else {
				body := BodyError{Code: CodeErrorLogin, Msg: "登录失败"}
				d, _ := json.Marshal(body)
				_, _ = w.Write(d)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			body := BodyError{Code: CodeErrorParamLess, Msg: "param less"}
			d, _ := json.Marshal(body)
			_, _ = w.Write(d)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		body := BodyError{Code: CodeErrorParamFormat, Msg: "param parse error"}
		d, _ := json.Marshal(body)
		_, _ = w.Write(d)
	}
}

func CreateUser() {

}
