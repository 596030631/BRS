package main

import (
	"net/http"
)

// CodeSuccess CodeErrorParamFormat CodeErrorParamLess 接口参数校验公用代码
const CodeSuccess = 10000
const CodeErrorParamFormat = 10001
const CodeErrorParamLess = 10002

type BodyError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type BodyUser struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	User *User  `json:"user"`
}

func login(w http.ResponseWriter, r *http.Request) {
	Login(w, r)
}

func Listener() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	Fatal(err)
}
