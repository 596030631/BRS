package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

const CodeErrorMaterialExist = 11011

type Material struct {
	Uid  string `json:"uid"`
	Pid  string `json:"pid"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type BodyMaterial struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	Material *Material `json:"material"`
}

func materialQuery(userId string) (*User, error) {
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

func AddMaterial(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		uId := r.Form.Get("uid")
		pId := r.Form.Get("pid")
		name := r.Form.Get("name")
		icon := r.Form.Get("icon")
		material := Material{uId, pId, name, icon}
		prepare, err := Conn.Prepare(`INSERT material (uid, pid, name, icon)  VALUES (?,?,?,?)`)
		if err == nil {
			exec, err := prepare.Exec(material.Uid, material.Pid, material.Name, material.Icon)
			if err == nil {
				eff, err := exec.RowsAffected()
				if err == nil {
					if eff == 1 {
						body := BodyMaterial{Code: CodeSuccess, Msg: "add successful", Material: &material}
						d, _ := json.Marshal(body)
						_, _ = w.Write(d)
					} else {
						BackError(w, CodeErrorDataBase, "insert failure")
					}
				} else {
					BackError(w, CodeErrorDataBase, err.Error())
				}
			} else if match, _ := regexp.MatchString("Error 1062: Duplicate entry .+ for key 'PRIMARY'", err.Error()); match {
				BackError(w, CodeErrorRegisterUserExist, "user has exist!")
			} else {
				BackError(w, CodeErrorDataBase, err.Error())
			}
		}
	}
}
