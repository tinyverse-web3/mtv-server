package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/form3tech-oss/jwt-go"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

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
		"/v0/user/getsssdata",
		"/v0/user/sendmail",
		"/v0/user/verifymail",
		"/v0/storage/test",
		"/v0/im/relays",
		"/v0/im/exchangeimpkey",
	}
	if !isContain(uris, uri) {
		var publicKey string
		tmp := c.Ctx.Request.Header["Public_key"]
		if tmp == nil {
			c.ErrorJson("600000", "public_key不能为空")
			return
		} else {
			publicKey = tmp[0]
		}
		logs.Info("public key = ", publicKey)

		var signature string
		tmp = c.Ctx.Request.Header["Sign"]
		if tmp == nil {
			c.ErrorJson("600000", "sign不能为空")
			return
		} else {
			signature = tmp[0]
		}
		logs.Info("sign = ", signature)

		var address string
		tmp = c.Ctx.Request.Header["Address"]
		if tmp == nil {
			c.ErrorJson("600000", "Address不能为空")
			return
		} else {
			address = tmp[0]
		}
		logs.Info("address = ", address)

		var data string
		method := c.Ctx.Request.Method
		logs.Info("method = ", method)
		switch method {
		case "GET":
			data = strings.Replace(uri, "/v0", "", 1)
			break
		case "POST":
			data = string(c.Ctx.Input.RequestBody)
		}
		logs.Info("data = ", data)

		match := sign(address, data, signature, publicKey)
		if !match {
			logs.Info("22222")
			c.ErrorJson("600000", "验签失败")
			return
		}

		o := orm.NewOrm()
		var user models.User
		user.PublicKey = publicKey
		err := o.Read(&user, "public_key")
		if err == nil { // 绑定邮箱时，user表中没有对应数据
			CurUser = user
		}

		// tmp := c.Ctx.Request.Header["Authorization"] // 格式：Bearer xxx
		// if tmp == nil {
		// 	logs.Info("00000")
		// 	c.ErrorJson("600000", "用户获取签名")
		// 	return
		// } else {
		// 	if len(strings.Split(tmp[0], " ")) != 2 {
		// 		logs.Info("22222")
		// 		c.ErrorJson("600000", "用户获取签名")
		// 		return
		// 	}
		// 	token := strings.Split(tmp[0], " ")[1]
		// 	o := orm.NewOrm()
		// 	var user models.User
		// 	user.Token = token
		// 	err := o.Read(&user, "token")
		// 	if err != nil {
		// 		logs.Info("11111")
		// 		c.ErrorJson("600000", "用户获取签名")
		// 		return
		// 	} else {
		// 		tokenUpdateTime := user.TokenUpdateTime
		// 		curTime := time.Now()
		// 		if curTime.Sub(tokenUpdateTime).Hours() > 168 { // token失效时间：7*24h
		// 			c.ErrorJson("600000", "签名已过期")
		// 			return
		// 		} else {
		// 			user.TokenUpdateTime = time.Now()
		// 			o.Update(&user)
		// 			CurUser = user
		// 		}
		// 	}
		// }
	}

	logs.Info("Prepare end")
}

func sign(address string, data string, signature string, publicKeyStr string) bool {
	hashData := solsha3.SoliditySHA3(
		[]string{"address", "string"},
		[]interface{}{
			address,
			data,
		},
	)
	logs.Info("sign hashData:", hexutil.Encode(hashData))

	publicKeyByte, err := hexutil.Decode(publicKeyStr)
	if err != nil {
		logs.Info(err)
		return false
	}

	signatureByte, _ := hexutil.Decode(signature)
	if signatureByte[64] != 0 && signatureByte[64] != 1 {
		signatureByte[64] -= 27
	}

	signatureNoRecoverID := signatureByte[:len(signatureByte)-1]
	result := crypto.VerifySignature(publicKeyByte, hashData, signatureNoRecoverID)

	return result
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
