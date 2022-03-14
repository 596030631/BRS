package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

const CodeErrorLoginPasswd = 11001

const CodeErrorRegisterUserExist = 11011

type User struct {
	UserId    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Passwd    string `json:"passwd"`
	UserSex   string `json:"user_sex"`
	UserAge   int    `json:"user_age"`
	UserLevel int    `json:"user_level"`
	UserIcon  string `json:"user_icon"`
}

type BodyUser struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	User *User  `json:"user"`
}

func query(userId string) (*User, error) {
	query := fmt.Sprintf("SELECT * FROM user WHERE user_id IN ('%s') LIMIT 1", userId)
	fmt.Println(query)
	rows, err := Conn.Query(query)
	var user = new(User)
	if err == nil {
		if rows.Next() {
			err = rows.Scan(&user.UserId, &user.UserName, &user.Passwd, &user.UserSex, &user.UserAge, &user.UserLevel, &user.UserIcon)
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
			user, err := query(userId)
			if err == nil {
				passwd := r.Form.Get("passwd")
				if len(passwd) > 0 && passwd == user.Passwd {
					body := BodyUser{Code: CodeSuccess, Msg: "login successful", User: user}
					d, _ := json.Marshal(body)
					_, _ = w.Write(d)
				} else {
					BackTip(w, CodeErrorLoginPasswd, "login failure")
				}
			} else {
				BackTip(w, CodeErrorDataBase, err.Error())
			}
		} else {
			BackTip(w, CodeErrorParamLess, "param less")
		}
	} else {
		BackTip(w, CodeErrorParamFormat, "param parse error")
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		userId := r.Form.Get("user_id")
		userName := r.Form.Get("user_name")
		passwd := r.Form.Get("passwd")
		userSex := r.Form.Get("user_sex")
		userAge := r.Form.Get("user_age")
		userLevel := r.Form.Get("user_level")
		userIcon := r.Form.Get("user_icon")

		if len(userId) < 6 || len(userId) > 128 {
			BackTip(w, CodeErrorParamFormat, "user_id error")
		} else if match, _ := regexp.MatchString("[A-z|0-9]+", userId); !match {
			BackTip(w, CodeErrorParamFormat, "user_id error format")
		} else if len(passwd) < 6 || len(passwd) > 128 {
			BackTip(w, CodeErrorParamFormat, "password error")
		} else if match, _ := regexp.MatchString("[A-z|0-9]+", userId); !match {
			BackTip(w, CodeErrorParamFormat, "password error format")
		} else if len(userLevel) != 1 {
			BackTip(w, CodeErrorParamFormat, "user_level error")
		} else if match, _ := regexp.MatchString("[1-2]", userLevel); !match {
			BackTip(w, CodeErrorParamFormat, "user_level error, right in [1,2]")
		} else if len(userName) < 6 || len(userName) > 128 {
			BackTip(w, CodeErrorParamFormat, "user_name error")
		} else {
			userAgeInt := 0
			if len(userAge) > 0 && len(userAge) < 4 {
				if userAgeInt, err = strconv.Atoi(userAge); err != nil {
					BackTip(w, CodeErrorParamFormat, "userAge error")
					return
				}
			}

			var userLevelInt int
			if userLevelInt, err = strconv.Atoi(userLevel); err != nil {
				BackTip(w, CodeErrorParamFormat, "userLevel error")
				return
			}

			user := User{userId, userName, passwd, userSex, userAgeInt, userLevelInt, userIcon}
			prepare, err := Conn.Prepare(`INSERT user (user_id, user_name, passwd, user_sex,user_age,user_level,user_icon)  VALUES (?,?,?,?,?,?,?)`)
			if err == nil {
				exec, err := prepare.Exec(user.UserId, user.UserName, user.Passwd, user.UserSex, user.UserAge, user.UserLevel, user.UserIcon)
				if err == nil {
					eff, err := exec.RowsAffected()
					if err == nil {
						if eff == 1 {
							body := BodyUser{Code: CodeSuccess, Msg: "register successful", User: &user}
							d, _ := json.Marshal(body)
							_, _ = w.Write(d)
						} else {
							BackTip(w, CodeErrorDataBase, "insert error")

						}
					} else {
						BackTip(w, CodeErrorDataBase, err.Error())
					}
				} else if match, _ := regexp.MatchString("Error 1062: Duplicate entry .+ for key 'PRIMARY'", err.Error()); match {
					BackTip(w, CodeErrorRegisterUserExist, "user has exist!")
				} else {
					BackTip(w, CodeErrorDataBase, err.Error())
				}
			} else {
				BackTip(w, CodeErrorDataBase, err.Error())
			}
		}
	} else {
		BackTip(w, CodeErrorParamFormat, err.Error())
	}
}
