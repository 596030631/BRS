package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type Borrow struct {
	OrderId     string `json:"order_id"`
	Status      string `json:"order_status"`
	UserId      string `json:"user_id"`
	MaterialId  string `json:"material_id"`
	BorrowCount string `json:"borrow_count"`
	BorrowDate  string `json:"borrow_date"`
	ReturnDate  string `json:"return_date"`
	ReturnCount string `json:"return_count"`
	Remake      string `json:"remake"`
}

type BodyBorrows struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
	Borrow []Borrow `json:"data"`
}

func BorrowMaterial(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		BackTip(w, CodeErrorParamFormat, err.Error())
		return
	}
	flag := r.Form.Get("type")
	switch flag {
	case "borrow":
		fmt.Println("borrow")
		borrowMaterial(w, r)
	case "back":
		fmt.Println("back")
		backMaterial(w, r)
	default:
		BackTip(w, CodeErrorParamLess, "less type in (borrow or back)")
		return
	}
}

func borrowMaterial(w http.ResponseWriter, r *http.Request) {
	userId := r.Form.Get("uid")
	if len(userId) == 0 {
		BackTip(w, CodeErrorParamLess, "less param uid")
		return
	}
	materialId := r.Form.Get("mid")
	if len(materialId) == 0 {
		BackTip(w, CodeErrorParamLess, "less param mid")
		return
	}
	count := r.Form.Get("count")
	if len(count) == 0 {
		BackTip(w, CodeErrorParamLess, "less param count")
		return
	}
	if match, _ := regexp.MatchString("^[0-9]+", count); !match {
		BackTip(w, CodeErrorParamFormat, "err param count")
		return
	}

	tx, err := Conn.Begin()
	if err != nil {
		BackTip(w, CodeErrorDataBase, "事务失败")
		return
	}

	stmt, err := tx.Prepare(`UPDATE material SET borrow = borrow + ? WHERE mid in (?) AND borrow + ? <= material.count`)
	if err != nil {
		rollBack(w, tx, "sql error")
		return
	}
	exec, err := stmt.Exec(count, materialId, count)
	if err != nil {
		rollBack(w, tx, "sql exec error")
		return
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		rollBack(w, tx, err.Error())
		return
	}
	if rowsAffected == 1 {
		prepare, err := tx.Prepare(`INSERT borrow (order_id, order_status, user_id, material_id, borrow_count, borrow_date, return_date, return_count,remake)  VALUES (?,?,?,?,?,?,?,?)`)
		if err != nil {
			rollBack(w, tx, err.Error())
			return
		}
		exec, err := prepare.Exec(RandString(16), "borrow", userId, materialId, count, time.Now(), time.Now(), 0, "")
		if err != nil {
			rollBack(w, tx, err.Error())
			return
		}
		affected, err := exec.RowsAffected()
		if err != nil {
			rollBack(w, tx, "新增借单失败")
			return
		}
		if affected == 1 {
			err = tx.Commit()
			if err == nil {
				BackTip(w, CodeSuccess, "Insert Successful")
			} else {
				BackTip(w, CodeErrorDataBase, err.Error())
			}
		} else {
			rollBack(w, tx, "修改失败")
		}
	} else {
		rollBack(w, tx, "访问失败，可能原因有物料号错误、库存不足")
	}
}

func backMaterial(w http.ResponseWriter, r *http.Request) {
	orderId := r.Form.Get("order_id")
	if len(orderId) == 0 {
		BackTip(w, CodeErrorParamLess, "less param order_id")
		return
	}
	count := r.Form.Get("count")
	if len(count) == 0 {
		BackTip(w, CodeErrorParamLess, "less param count")
		return
	}
	if match, _ := regexp.MatchString("^[0-9]+", count); !match {
		BackTip(w, CodeErrorParamFormat, "err param count")
		return
	}
	row := Conn.QueryRow(`SELECT * FROM borrow WHERE order_id in (?) LIMIT 1`, orderId)
	order := Borrow{}
	err := row.Scan(&order.OrderId, &order.Status, &order.UserId, &order.MaterialId, &order.BorrowCount, &order.BorrowDate, &order.ReturnDate, &order.ReturnCount, &order.Remake)
	if err != nil {
		BackTip(w, CodeErrorDataBase, err.Error())
		return
	}

	if i, err := strconv.Atoi(count); i <= 0 || err != nil {
		BackTip(w, CodeErrorParamFormat, "count error")
		return
	}

	if order.BorrowCount < count+order.ReturnCount {
		BackTip(w, CodeErrorParamFormat, "借还数量不匹配")
		return
	}

	if order.BorrowCount == order.ReturnCount+count {
		order.Status = "return"
	}

	tx, err := Conn.Begin()
	if err != nil {
		BackTip(w, CodeErrorDataBase, "事务失败")
		return
	}
	stmt, err := tx.Prepare(`UPDATE material SET borrow = borrow - ? WHERE mid in (?)`)
	if err != nil {
		rollBack(w, tx, "sql error")
		return
	}
	exec, err := stmt.Exec(count, order.MaterialId)
	if err != nil {
		rollBack(w, tx, err.Error())
		return
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		rollBack(w, tx, err.Error())
		return
	}
	if rowsAffected == 1 {
		prepare, err := tx.Prepare(`UPDATE borrow SET order_status = ?, return_date = ? , return_count = ? , remake = ? WHERE order_id in (?)`)
		if err != nil {
			rollBack(w, tx, err.Error())
			return
		}
		exec, err := prepare.Exec(order.Status, time.Now(), count, "", order.OrderId)
		if err != nil {
			rollBack(w, tx, err.Error())
			return
		}
		affected, err := exec.RowsAffected()
		if err != nil {
			rollBack(w, tx, "返回失败")
			return
		}
		if affected == 1 {
			err = tx.Commit()
			if err == nil {
				BackTip(w, CodeSuccess, "Update Successful")
			} else {
				BackTip(w, CodeErrorDataBase, err.Error())
			}
		} else {
			rollBack(w, tx, "修改失败")
		}
	} else {
		rollBack(w, tx, "访问失败,可能物料不存在")
	}
}

func BorrowOrderList(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		BackTip(w, CodeErrorParamFormat, err.Error())
		return
	}
	userId := r.Form.Get("uid")
	if len(userId) == 0 {
		BackTip(w, CodeErrorParamLess, "less param user_id")
		return
	}
	status := r.Form.Get("status")
	if len(status) == 0 {
		BackTip(w, CodeErrorParamLess, "less param status")
		return
	}
	var rows *sql.Rows
	var err error
	if status == "all" {
		rows, err = Conn.Query(`SELECT * FROM borrow WHERE user_id in (?)`, userId)
	} else {
		rows, err = Conn.Query(`SELECT * FROM borrow WHERE user_id in (?) AND order_status in (?)`, userId, status)
	}
	if err != nil {
		BackTip(w, CodeErrorDataBase, err.Error())
		return
	}
	var data []Borrow
	for rows.Next() {
		var order = Borrow{}
		err = rows.Scan(&order.OrderId, &order.Status, &order.UserId, &order.MaterialId, &order.BorrowCount, &order.BorrowDate, &order.ReturnDate, &order.ReturnCount, &order.Remake)
		fmt.Println(fmt.Sprintf("%+v", &order))
		data = append(data, order)
	}
	body := BodyBorrows{Code: CodeSuccess, Msg: "query successful", Borrow: data}
	d, _ := json.Marshal(body)
	_, _ = w.Write(d)
}

func rollBack(w http.ResponseWriter, tx *sql.Tx, msg string) {
	_ = tx.Rollback()
	_ = tx.Commit()
	BackTip(w, CodeErrorDataBase, msg)
	return
}
