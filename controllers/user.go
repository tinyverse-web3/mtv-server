package controllers

import (
	"encoding/json"
	"fmt"
	"mtv/models"
	"mtv/utils"
	"mtv/utils/crypto"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	var user models.User
	// err := c.BindJSON(&user)
	// if err != nil {
	// 	c.ErrorJson("400000", "参数不能为空")
	// 	return
	// }
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	logs.Info(user)
	email := user.Email
	if strings.TrimSpace(email) == "" {
		c.ErrorJson("400000", "邮箱地址不能为空")
		return
	}
	if !utils.IsEmail(email) {
		c.ErrorJson("400000", "邮箱地址错误")
		return
	}

	confirmCode := user.ConfirmCode
	if strings.TrimSpace(confirmCode) == "" {
		c.ErrorJson("400000", "验证码不能为空")
		return
	}
	o := orm.NewOrm()

	err := o.Read(&user, "email")
	if err != nil {
		c.ErrorJson("400000", email+"不存在")
		return
	}
	// oriConfirmCode := user.ConfirmCode

	// user.ConfirmCode = confirmCode
	err = o.Read(&user, "email", "confirm_code")
	if err != nil {
		c.ErrorJson("400000", "验证码错误")
		return
	}

	// 判断验证码是否过期
	confirmCodeUpdateTime := user.ConfirmCodeUpdateTime
	curTime := time.Now()
	if curTime.Sub(confirmCodeUpdateTime).Seconds() > 60 { // 验证码1分钟失效
		c.ErrorJson("400000", "验证码已过期")
		return
	}

	name := user.Name
	if name == "" { // 如果未设置name，则随机生成name，格式为：mtv_6位随机数字
		var tmpUser models.User
		for true {
			name = "mtv_" + utils.RandomNum(6)
			tmpUser.Name = name
			err := o.Read(&tmpUser, "name")
			if err == orm.ErrNoRows {
				user.Name = name
				break
			}
		}

	}

	// user.ConfirmCode = oriConfirmCode
	user.ConfirmCodeUpdateTime = time.Now()
	user.Status = 1 // 已验证

	token := createToken(user.Email)
	user.Token = token
	user.TokenUpdateTime = time.Now()

	o.Update(&user)

	c.SuccessJson("", token)
}

func (c *UserController) GetImPubKeyList() {
	var data []models.User

	user := new(models.User)
	qt := orm.NewOrm().QueryTable(user)

	email := c.GetString("email")
	if email != "" {
		qt = qt.Filter("email", email)
	}

	qt.Exclude("email__in", CurUser.Email).All(&data, "Email", "NostrPublicKey")
	c.SuccessJson("", data)
}

type UserInfo struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	SssData   string `json:"sssData"`
	Ipns      string `json:"ipns"`
	DbAddress string `json:"dbAddress"`
}

func (c *UserController) GetUserInfo() {
	var user UserInfo

	if CurUser.SssData != "" {
		key, _ := config.String("crypto")
		deKey := crypto.DecryptBase64(key)
		ct := crypto.DecryptAES(CurUser.SssData, deKey)
		user.SssData = ct
	}

	walletAddress := CurUser.Address
	ipns, _ := utils.GetDFSPath(walletAddress)
	user.Ipns = ipns

	user.DbAddress = CurUser.DbAddress
	user.Name = CurUser.Name
	user.Email = CurUser.Email
	c.SuccessJson("", user)
}

func (c *UserController) ModifyUser() {
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)
	logs.Info("user = ", user)
	o := orm.NewOrm()

	name := strings.TrimSpace(user.Name)
	if name != "" {
		if name != CurUser.Name {
			var data []models.User
			qt := orm.NewOrm().QueryTable(user)
			qt.Filter("name", name).Exclude("id__in", CurUser.Id).All(&data)
			if len(data) > 0 {
				c.ErrorJson("400000", "用户名已存在。")
				return
			} else {
				CurUser.Name = name
			}
		}
	}

	sssData := user.SssData
	if sssData != "" {
		key, _ := config.String("crypto")
		logs.Info("key = ", key)
		deKey := crypto.DecryptBase64(key)
		logs.Info("deKey = ", deKey)
		logs.Info("sssData = ", sssData)
		enData := crypto.EncryptAES(sssData, deKey)
		logs.Info("enData = ", enData)
		CurUser.SssData = enData
		// CurUser.SssData = sssData
	}

	nostrPublicKey := user.NostrPublicKey
	if nostrPublicKey != "" {
		CurUser.NostrPublicKey = nostrPublicKey
	}

	publicKey := user.PublicKey
	if publicKey != "" {
		CurUser.PublicKey = publicKey
	}

	walletAddress := user.Address
	if walletAddress != "" {
		CurUser.Address = walletAddress
	}

	sign := user.Sign
	if sign != "" {
		CurUser.Sign = sign
		logs.Info("sign = ", sign)
	}

	ipns := user.Ipns
	if ipns != "" {
		logs.Info("ipns = ", ipns)
		if ipns != CurUser.Ipns {
			CurUser.Ipns = ipns
			if walletAddress == "" {
				walletAddress = CurUser.Address
			}
			logs.Info("wallet address = ", walletAddress)
			success, err := utils.SetDFSPath(walletAddress, ipns, user.Sign) // 更新钱包中用户的dfs地址
			if !success {
				logs.Error(err)
				c.ErrorJson("400000", err.Error())
				return
			}
		}
	}

	dbAddress := user.DbAddress
	if dbAddress != "" {
		CurUser.DbAddress = dbAddress
	}

	_, err := o.Update(&CurUser)
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "更新用户信息失败")
	} else {
		c.SuccessJson("", "")
	}
}

func (c *UserController) SendMail() {
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	email := user.Email
	if strings.TrimSpace(email) == "" {
		c.ErrorJson("400000", "邮箱地址不能为空")
		return
	}
	if !utils.IsEmail(email) {
		c.ErrorJson("400000", "邮箱地址错误")
		return
	}

	o := orm.NewOrm()
	tmpUser := models.User{Email: email}
	status := -1
	err := o.Read(&tmpUser, "email")
	if err != orm.ErrNoRows {
		status = tmpUser.Status
	}

	// confirmCode := utils.RandomNum(6)
	confirmCode := "123456"

	if status == -1 { // email不存在
		user = models.User{Email: email, ConfirmCode: confirmCode, Status: 0, ConfirmCodeUpdateTime: time.Now()}
		_, err := o.Insert(&user)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "Send mail faild!")
			return
		}
	} else {
		// 如果email存在，判断发送email的频率
		confirmCodeUpdateTime := tmpUser.ConfirmCodeUpdateTime
		curTime := time.Now()
		if curTime.Sub(confirmCodeUpdateTime).Seconds() < 60 {
			c.ErrorJson("400000", "验证码已发送，请查看邮箱。")
			return
		}
	}

	subject := "发送验证码"
	message := `
		<p> Hi %s,</p>
		<p style="text-indent:2em">Thanks for using MTV.</p>
		<p style="text-indent:2em"></p>
		<p style="text-indent:2em">Grab the confirmation code below and enter it on the screen.</P>
		<h2 style="background:#eee;text-align:center;padding:8px;border-radius:5px">%s</h2>
		<p style="text-indent:2em"> - The MTV Team</p>
	`

	success := utils.Send(email, subject, fmt.Sprintf(message, strings.Split(email, "@")[0], confirmCode))

	if success {
		user = models.User{Email: email}
		if o.Read(&user, "email") == nil {
			user.ConfirmCode = confirmCode
			user.ConfirmCodeUpdateTime = time.Now()
			_, err := o.Update(&user)
			if err != nil {
				logs.Error(err)
				c.ErrorJson("400000", "Send mail faild!")
				return
			}
		}
		c.SuccessJson("Send mail success!", "")
	} else {
		c.ErrorJson("400000", "Send mail faild!")
	}
}

func checkUserStatus(email string) (status int) {
	o := orm.NewOrm()
	user := models.User{Email: email}

	err := o.Read(&user, "email")
	if err == orm.ErrNoRows {
		status = -1
	} else {
		status = user.Status
	}
	return status
}
