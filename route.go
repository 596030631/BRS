package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CodeSuccess CodeErrorParamFormat CodeErrorParamLess 接口参数校验公用代码
const CodeSuccess = 10000
const CodeErrorParamFormat = 10001
const CodeErrorParamLess = 10002
const CodeErrorDataBase = 10003

type BodyTip struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func login(w http.ResponseWriter, r *http.Request) {
	Login(w, r)
}
func register(w http.ResponseWriter, r *http.Request) {
	CreateUser(w, r)
}
func classesAdd(w http.ResponseWriter, r *http.Request) {
	AddClasses(w, r)
}
func classesDelete(w http.ResponseWriter, r *http.Request) {
	DeleteClasses(w, r)
}
func classesList(w http.ResponseWriter, r *http.Request) {
	ClassesQuery(w, r)
}
func materialAdd(w http.ResponseWriter, r *http.Request) {
	AddMaterial(w, r)
}
func materialList(w http.ResponseWriter, r *http.Request) {
	MaterialQuery(w, r)
}
func materialDelete(w http.ResponseWriter, r *http.Request) {
	DeleteMaterial(w, r)
}

func Listener() {
	http.HandleFunc("/user/login", login)
	http.HandleFunc("/user/register", register)
	http.HandleFunc("/classes/add", classesAdd)
	http.HandleFunc("/classes/delete", classesDelete)
	http.HandleFunc("/classes/list", classesList)
	http.HandleFunc("/material/add", materialAdd)
	http.HandleFunc("/material/list", materialList)
	http.HandleFunc("/material/delete", materialDelete)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	Fatal(err)
}

func BackTip(w http.ResponseWriter, code int, e string) {
	body := BodyTip{Code: code, Msg: e}
	fmt.Println(body)
	d, _ := json.Marshal(body)
	_, _ = w.Write(d)
}
