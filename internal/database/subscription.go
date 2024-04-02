package database

import (
	"strconv"

	"github.com/lankaiyun/company/global"
)

type Subscription struct {
	Id    string
	Emile string
	Time  string
}

func (s *Subscription) IsExist() bool {
	var temp string
	err := global.MySqlConn.QueryRow("SELECT COUNT(*) AS count FROM subscription WHERE email = ?", s.Emile).Scan(&temp)
	if err != nil {
		global.SugarLogger.Error("SelectArticleNum err:", err)
	}
	i, err := strconv.Atoi(temp)
	if err != nil {
		global.SugarLogger.Error("strconv.Atoi() 转换失败， err=", err)
	}
	if i >= 1 {
		return true
	} else {
		return false
	}
}

func (s *Subscription) Add() {
	_, err := global.MySqlConn.Exec("insert subscription set email=?, time=CURDATE()", s.Emile)
	if err != nil {
		global.SugarLogger.Error("Add err:", err)
	}
}

//func (a *Article) Select() Article {
//	stmt, err := global.MySqlConn.Prepare("select * from article where id=?")
//	var temp Article
//	defer stmt.Close()
//	err = stmt.QueryRow(a.Id).Scan(&temp.Id, &temp.Title, &temp.Category, &temp.Content, &temp.Time, &temp.Click)
//	if err != nil {
//		global.SugarLogger.Error("QueryRow err:", err)
//	}
//	return temp
//}
//
//func (a *Article) UpdateClick() {
//	_, err := global.MySqlConn.Exec("update article set click = click + 1 where id=?", a.Id)
//	if err != nil {
//		global.SugarLogger.Error("UpdateClick err:", err)
//	}
//}
//
//func (a *Article) SelectArticleNum() {
//	err := global.MySqlConn.QueryRow("select COUNT(*) AS articleNum from article where category=?",
//		a.Category).Scan(&a.ArticleNum)
//	if err != nil {
//		global.SugarLogger.Error("SelectArticleNum err:", err)
//	}
//}
//
//func (a *Article) SelectArticleNum2() {
//	err := global.MySqlConn.QueryRow("select COUNT(*) AS articleNum from article").Scan(&a.ArticleNum)
//	if err != nil {
//		global.SugarLogger.Error("SelectArticleNum err:", err)
//	}
//}
//
//func (a *Article) SelectAll() []Article {
//	var arr []Article
//	rows, err := global.MySqlConn.Query("select * from article where category=? order by id desc", a.Category)
//	if err != nil {
//		global.SugarLogger.Error("SelectAll Query err:", err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var temp Article
//		err = rows.Scan(&temp.Id, &temp.Title, &temp.Category, &temp.Content, &temp.Time, &temp.Click)
//		if err != nil {
//			fmt.Println("SelectAll Scan err:", err)
//		}
//		arr = append(arr, temp)
//	}
//	return arr
//}
//
//func (a *Article) Add() {
//	_, err := global.MySqlConn.Exec("insert article set title=?, category=?, content=?, time=CURDATE()",
//		a.Title, a.Category, a.Content)
//	if err != nil {
//		global.SugarLogger.Error("Add err:", err)
//	}
//}
//
//func (a *Article) Modify() {
//	_, err := global.MySqlConn.Exec("update article set title=?, content=?, time=CURDATE() where id=?", a.Title, a.Content, a.Id)
//	if err != nil {
//		global.SugarLogger.Error("Modify err:", err)
//	}
//}
