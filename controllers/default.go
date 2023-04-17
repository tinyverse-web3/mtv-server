package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"strings"

	"mtv/models"
	"mtv/utils"
)

type RespJson struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var langTypes []string

// var CurUser models.User

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

	if !isContain(uri) {
		var publicKey string
		tmp := c.Ctx.Request.Header["Public_key"]
		if tmp == nil {
			c.ErrorJson("600000", "public_key不能为空")
			return
		} else {
			publicKey = tmp[0]
		}
		logs.Info("public key = ", publicKey)
		o := orm.NewOrm()

		var user models.User
		user.PublicKey = publicKey
		err := o.Read(&user, "public_key")
		if err == orm.ErrNoRows {
			name := generateUserName()
			user = models.User{Status: 1, PublicKey: publicKey, Name: name}
			_, err = o.Insert(&user)
		}

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
			if data == "" {
				data = strings.Replace(uri, "/v0", "", 1)
			}
			break
		}
		logs.Info("data = ", data)

		match := sign(address, data, signature, publicKey)
		if !match {
			logs.Info("验签失败")
			c.ErrorJson("600000", "验签失败")
			return
		}
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

func generateUserName() (name string) {
	o := orm.NewOrm()
	var tmpUser models.User
	for true {
		name = "mtv_" + utils.RandomNum(6)
		tmpUser.Name = name
		err := o.Read(&tmpUser, "name")
		if err == orm.ErrNoRows {
			break
		}
	}
	return
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

func (c *BaseController) CurUser() models.User {
	var publicKey string
	tmp := c.Ctx.Request.Header["Public_key"]
	publicKey = tmp[0]

	o := orm.NewOrm()

	var user models.User
	user.PublicKey = publicKey
	o.Read(&user, "public_key")
	return user
}

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

func isContain(item string) bool {
	uris := []string{
		"/v0/auth/checksign",

		"/v0/user/sendmail4verifycode",
		"/v0/user/getpassword",
		"/v0/user/getsssdata4guardian",
		"/v0/user/getsssdata4question",
		"/v0/user/uploadimg",

		"/v0/storage/test", // for test

		"/v0/im/relays",
		"/v0/im/exchangeimpkey",
	}

	contain := false
	for _, eachItem := range uris {
		if strings.Index(item, eachItem) != -1 {
			contain = true
			break
		}
	}

	return contain
}

// // 生成token
// func createToken(email string) string {
// 	appKey := "1qazXSW@3edc"
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user": email,
// 		"exp":  time.Now().Add(24 * time.Hour * time.Duration(1)).Unix(),
// 		"iat":  time.Now().Unix(),
// 	})
// 	tokenStr, _ := token.SignedString([]byte(appKey))
// 	return tokenStr
// }
