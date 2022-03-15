package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

const CodeErrorMaterialExist = 11011

type Material struct {
	Mid  string `json:"mid"`
	Cid  string `json:"cid"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type BodyMaterials struct {
	Code      int        `json:"code"`
	Msg       string     `json:"msg"`
	Materials []Material `json:"materials"`
}

type BodyMaterial struct {
	Code     int       `json:"code"`
	Msg      string    `json:"msg"`
	Material *Material `json:"material"`
}

func MaterialQuery(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		BackTip(w, CodeErrorParamFormat, err.Error())
		return
	}
	cid := r.Form.Get("cid")
	var rows *sql.Rows
	var err error
	if len(cid) == 0 || cid == "all" {
		rows, err = Conn.Query(`SELECT * FROM material`)
	} else {
		rows, err = Conn.Query(`SELECT * FROM material WHERE cid in (?)`, cid)
	}

	var data []Material
	if err == nil {
		for rows.Next() {
			var c = Material{}
			err = rows.Scan(&c.Mid, &c.Cid, &c.Name, &c.Icon)
			fmt.Println(fmt.Sprintf("%+v", &c))
			data = append(data, c)
		}
		body := BodyMaterials{Code: CodeSuccess, Msg: "query successful", Materials: data}
		d, _ := json.Marshal(body)
		_, _ = w.Write(d)
	} else {
		BackTip(w, CodeErrorDataBase, err.Error())
	}
}

func AddMaterial(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		uId := r.Form.Get("mid")
		pId := r.Form.Get("cid")
		name := r.Form.Get("name")
		icon := r.Form.Get("icon")
		material := Material{uId, pId, name, icon}
		prepare, err := Conn.Prepare(`INSERT material (mid, cid, name, icon)  VALUES (?,?,?,?)`)
		if err == nil {
			exec, err := prepare.Exec(material.Mid, material.Cid, material.Name, material.Icon)
			if err == nil {
				eff, err := exec.RowsAffected()
				if err == nil {
					if eff == 1 {
						body := BodyMaterial{Code: CodeSuccess, Msg: "add successful", Material: &material}
						d, _ := json.Marshal(body)
						_, _ = w.Write(d)
					} else {
						BackTip(w, CodeErrorDataBase, "insert failure")
					}
				} else {
					BackTip(w, CodeErrorDataBase, err.Error())
				}
			} else if match, _ := regexp.MatchString("Error 1062: Duplicate entry .+ for key 'PRIMARY'", err.Error()); match {
				BackTip(w, CodeErrorRegisterUserExist, "user has exist!")
			} else {
				BackTip(w, CodeErrorDataBase, err.Error())
			}
		}
	}
}

func DeleteMaterial(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil {
		mid := r.Form.Get("mid")
		if mid == "" {
			BackTip(w, CodeErrorParamLess, "cid not found")
			return
		}
		prepare, err := Conn.Prepare(`DELETE FROM material WHERE mid in (?)`)
		if err == nil {
			exec, err := prepare.Exec(mid)
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
