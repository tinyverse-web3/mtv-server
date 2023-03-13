package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"github.com/form3tech-oss/jwt-go"

	"strings"
	"time"

	"mtv/models"
)

type RespJson struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var langTypes []string
var CurUser models.User

func init() {
	// Initialize language type list.
	tmp, _ := config.String("lang_types")
	langTypes = strings.Split(tmp, "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		logs.Info("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			logs.Error("Fail to set message file:", err)
			return
		}
	}
}

type BaseController struct {
	beego.Controller
	i18n.Locale
}

func (c *BaseController) Prepare() {
	logs.Info("Prepare start")
	c.setLang() // 国际化
	// logs.Info(c.Tr("hi")) // 国际化示例

	uri := c.Ctx.Request.RequestURI
	logs.Info(uri)
	uris := []string{
		"/v0/user/login",
		"/v0/user/sendmail",
		"/v0/storage/test",
		"/v0/im/relays",
		"/v0/im/exchangeimpkey",
	}
	if !isContain(uris, uri) {
		tmp := c.Ctx.Request.Header["Authorization"] // 格式：Bearer xxx
		if tmp == nil {
			logs.Info("00000")
			c.ErrorJson("400000", "用户未登录")
			return
		} else {
			if len(strings.Split(tmp[0], " ")) != 2 {
				logs.Info("22222")
				c.ErrorJson("400000", "用户未登录")
				return
			}
			token := strings.Split(tmp[0], " ")[1]
			o := orm.NewOrm()
			var user models.User
			user.Token = token
			err := o.Read(&user, "token")
			if err != nil {
				logs.Info("11111")
				c.ErrorJson("400000", "用户未登录")
				return
			} else {
				tokenUpdateTime := user.TokenUpdateTime
				curTime := time.Now()
				if curTime.Sub(tokenUpdateTime).Hours() > 168 { // token失效时间：7*24h
					c.ErrorJson("400000", "登录已过期")
					return
				} else {
					user.TokenUpdateTime = time.Now()
					o.Update(&user)
					CurUser = user
				}
			}
		}
	}

	logs.Info("Prepare end")
}

func (c *BaseController) setLang() {
	c.Lang = "" // This field is from i18n.Locale.

	al := c.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			c.Lang = al
		}
	}

	if len(c.Lang) == 0 {
		c.Lang = "en-US"
	}

	c.Data["Lang"] = c.Lang
}

// func (c *BaseController) CurUser() models.User {
// 	tmp := c.Ctx.Request.Header["Authorization"]
// 	token := strings.Split(tmp[0], " ")[1]
// 	o := orm.NewOrm()

// 	var user models.User
// 	user.Token = token
// 	o.Read(&user, "token")
// 	return user
// }

func (c *BaseController) SuccessJson(msg string, data interface{}) {
	if msg == "" {
		msg = "success"
	}
	res := RespJson{
		"000000", msg, data,
	}
	c.Data["json"] = res
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) ErrorJson(code string, msg string) {
	res := RespJson{
		code, msg, "",
	}
	c.Data["json"] = res
	c.ServeJSON()
	c.StopRun()
}

func isContain(items []string, item string) bool {
	contain := false
	for _, eachItem := range items {
		if strings.Index(eachItem, item) != -1 {
			contain = true
			break
		}
	}

	return contain
}

// 生成token
func createToken(email string) string {
	appKey := "1qazXSW@3edc"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": email,
		"exp":  time.Now().Add(24 * time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenStr, _ := token.SignedString([]byte(appKey))
	return tokenStr
}
