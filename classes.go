package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Classes struct {
	Id   string `json:"cid"`
	Name string `json:"name"`
	Pid  string `json:"pid"`
}

type BodyClasses struct {
	Code    int       `json:"code"`
	Msg     string    `json:"msg"`
	Classes []Classes `json:"classes"`
}

type BodySingleClasses struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Classes *Classes `json:"single_classes"`
}

func ClassesQuery(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		BackTip(w, CodeErrorParamFormat, err.Error())
		return
	}
	pid := r.Form.Get("pid")
	var rows *sql.Rows
	var err error

	if len(pid) == 1 {
		rows, err = Conn.Query(`SELECT * FROM classes WHERE pid in (?)`, pid)
	} else {
		rows, err = Conn.Query(`SELECT * FROM classes`)
	}
	var data []Classes
	if err == nil {
		for rows.Next() {
			var c = Classes{}
			err = rows.Scan(&c.Id, &c.Name, &c.Pid)
			fmt.Println(fmt.Sprintf("%+v", &c))
			data = append(data, c)
		}
		body := BodyClasses{Code: CodeSuccess, Msg: "query successful", Classes: data}
		d, _ := json.Marshal(body)
		_, _ = w.Write(d)
	} else {
		BackTip(w, CodeErrorDataBase, err.Error())
	}
}

func AddClasses(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		id := r.Form.Get("cid")
		name := r.Form.Get("name")
		pid := r.Form.Get("pid")
		if id == "" {
			BackTip(w, CodeErrorParamLess, "cid not find")
			return
		}
		if name == "" {
			BackTip(w, CodeErrorParamLess, "name not find")
			return
		}
		if pid == "" {
			BackTip(w, CodeErrorParamLess, "pid not find")
			return
		}

		classes := Classes{id, name, pid}
		prepare, err := Conn.Prepare(`INSERT classes (cid, name, pid)  VALUES (?,?,?)`)
		if err == nil {
			exec, err := prepare.Exec(classes.Id, classes.Name, classes.Pid)
			if err == nil {
				eff, err := exec.RowsAffected()
				if err == nil {
					if eff == 1 {
						body := BodySingleClasses{Code: CodeSuccess, Msg: "add successful", Classes: &classes}
						d, _ := json.Marshal(body)
						_, _ = w.Write(d)
					} else {
						BackTip(w, CodeErrorDataBase, "insert error")
					}
				} else {
					BackTip(w, CodeErrorDataBase, err.Error())
				}
			} else if match, _ := regexp.MatchString("Error 1062: Duplicate entry .+ for key 'PRIMARY'", err.Error()); match {
				BackTip(w, CodeErrorRegisterUserExist, "classes has exist!")
			} else {
				BackTip(w, CodeErrorDataBase, err.Error())
			}
		}
	}
}

func DeleteClasses(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		cid := r.Form.Get("cid")
		if cid == "" {
			BackTip(w, CodeErrorParamLess, "cid not found")
			return
		}
		prepare, err := Conn.Prepare(`DELETE FROM classes WHERE cid in (?)`)
		if err == nil {
			exec, err := prepare.Exec(cid)
			if err == nil {
				eff, err := exec.RowsAffected()
				if err == nil {
					if eff == 1 {
						body := BodyTip{Code: CodeSuccess, Msg: "delete successful"}
						d, _ := json.Marshal(body)
						_, _ = w.Write(d)
					} else {
						BackTip(w, CodeErrorDataBase, "delete error")
					}
				} else {
					BackTip(w, CodeErrorDataBase, err.Error())
				}
			} else {
				BackTip(w, CodeErrorDataBase, err.Error())
			}
		}
	}
}
