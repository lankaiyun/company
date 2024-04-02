package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lankaiyun/company/global"
	"github.com/lankaiyun/company/internal/database"
	"github.com/lankaiyun/company/pkg/email"
	"net/http"
)

type Controller struct{}

func (ctl Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "main/index.html", nil)
}

func (ctl Controller) Contact(c *gin.Context) {
	t := c.Query("type")
	if t == "business" {
		c.HTML(http.StatusOK, "main/contact-business.html", nil)
	} else if t == "product" {
		c.HTML(http.StatusOK, "main/contact-product.html", nil)
	} else {
		c.HTML(http.StatusNotFound, "main/404.html", nil)
	}
}

func (ctl Controller) AddSubscription(c *gin.Context) {
	e := c.PostForm("email")
	s := &database.Subscription{Emile: e}
	if s.IsExist() {
		c.JSON(http.StatusOK, gin.H{"result": "0"})
	} else {
		s.Add()
		c.JSON(http.StatusOK, gin.H{"result": "1"})
	}
}

func (ctl Controller) EmailContact(c *gin.Context) {
	t := c.PostForm("type")
	name := c.PostForm("name")
	message := c.PostForm("message")
	contact := c.PostForm("contact")
	mailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	if t == "business" {
		err := mailer.SendMail("shangwu@lankaiyun.com", "商务合作", fmt.Sprintf("姓名：%s 联系方式：%s 内容：%s",
			name, contact, message))
		if err != nil {
			global.SugarLogger.Error("mailer.SendMail err:", err)
		}
	} else if t == "product" {
		err := mailer.SendMail("chanpin@lankaiyun.com", "商务合作", fmt.Sprintf("姓名：%s 联系方式：%s 内容：%s",
			name, contact, message))
		if err != nil {
			global.SugarLogger.Error("mailer.SendMail err:", err)
		}
	}
	c.JSON(http.StatusOK, nil)
}
