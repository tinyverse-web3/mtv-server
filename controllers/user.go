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

// User API
type UserController struct {
	BaseController
}

type UserInfo struct {
	Name      string            `json:"name"`
	Password  string            `json:"password"`
	Email     string            `json:"email"`
	SssData   string            `json:"sssData"`
	Ipns      string            `json:"ipns"`
	DbAddress string            `json:"dbAddress"`
	Questions []models.Question `json:"questions"`
	ImgCid    string            `json:"imgCid"`
	SafeLevel int               `json:"safeLevel"`
}

// @Title UpdateSafeLevel
// @Description 更新用户安全等级
// @Param safeLevel body int true "安全等级"
// @Success 200 {object} controllers.RespJson
// @router /updatesafelevel [post]
func (c *UserController) UpdateSafeLevel() {
	CurUser := c.CurUser()

	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	o := orm.NewOrm()
	user.Id = CurUser.Id
	_, err := o.Update(&user, "safe_level")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "更新安全等级失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title UpdateImgCid
// @Description 更新用户头像
// @Param imgCid body string true "图片上传至IPFS上的CID"
// @Success 200 {object} controllers.RespJson
// @router /updateimgcid [post]
func (c *UserController) UpdateImgCid() {
	CurUser := c.CurUser()

	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	o := orm.NewOrm()
	user.Id = CurUser.Id
	_, err := o.Update(&user, "img_cid")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "更新头像失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title SavePassword
// @Description 保存用户密码
// @Param password body string true "密码"
// @Success 200 {object} controllers.RespJson
// @router /savepassword [post]
func (c *UserController) SavePassword() {
	CurUser := c.CurUser()

	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	o := orm.NewOrm()
	user.Id = CurUser.Id
	_, err := o.Update(&user, "password")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "保存用户密码失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title GetPassword
// @Description 获取用户密码
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /getpassword [post]
func (c *UserController) GetPassword() {
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	verify, msg := verifyEmailAndConfirmCode(user)
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	o := orm.NewOrm()
	hashEmail := crypto.Md5(user.Email)
	user.Email = hashEmail
	err := o.Read(&user, "email")
	if err != nil {
		c.ErrorJson("400000", "获取数据失败")
		return
	}

	c.SuccessJson("", user.Password)
}

// @Title VerifyMail
// @Description 验证邮箱
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /verifymail [post]
func (c *UserController) VerifyMail() {
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	logs.Info(user)
	flag, msg := verifyEmailAndConfirmCode(user)
	if !flag {
		c.ErrorJson("400000", msg)
		return
	}

	o := orm.NewOrm()
	hashEmail := crypto.Md5(user.Email)
	user.Email = hashEmail
	o.Read(&user, "email")

	var userInfo UserInfo

	if user.SssData != "" {
		key, _ := config.String("crypto")
		deKey := crypto.DecryptBase64(key)
		ct := crypto.DecryptAES(user.SssData, deKey)
		userInfo.SssData = ct
	}

	var data []models.Question
	question := new(models.Question)
	qt := orm.NewOrm().QueryTable(question)
	qt.Filter("user_id", user.Id).All(&data, "Id", "Content")
	userInfo.Questions = data

	c.SuccessJson("", userInfo)
}

// @Title GetSssData
// @Description 获取用户的分片数据
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /getsssdata [post]
func (c *UserController) GetSssData() {
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	verify, msg := verifyEmailAndConfirmCode(user)
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	o := orm.NewOrm()
	hashEmail := crypto.Md5(user.Email)
	user.Email = hashEmail
	err := o.Read(&user, "email")
	if err != nil {
		c.ErrorJson("400000", "获取数据失败")
		return
	}

	c.SuccessJson("", user.SssData)
}

// @Title BindMail
// @Description 绑定邮箱(如果未设置name，则随机生成name，格式为：mtv_6位随机数字)
// @Param public_key header string true "public key"
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /bindmail [post]
func (c *UserController) BindMail() {
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	logs.Info(user)
	email := strings.TrimSpace(user.Email)

	var publicKey string
	tmp := c.Ctx.Request.Header["Public_key"]
	if tmp == nil {
		c.ErrorJson("600000", "public_key不能为空")
		return
	} else {
		publicKey = tmp[0]
	}

	verify, msg := verifyEmailAndConfirmCode(user)
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	// 验证public key是否为空
	if publicKey == "" {
		c.ErrorJson("400000", "Public Key不能为空")
		return
	}

	o := orm.NewOrm()
	hashEmail := crypto.Md5(email)
	// email存在，判断public key与数据库中的数据是否一致
	user.Email = hashEmail
	o.Read(&user, "email") // verifyEmailAndConfirmCode方法已经验证邮箱是否存在，所以此处不需再做异常处理

	if user.PublicKey == publicKey {
		c.SuccessJson("", "")
		return
	}

	if user.PublicKey != "" && publicKey != user.PublicKey {
		c.ErrorJson("400000", "绑定失败：邮箱已绑定")
		return
	}

	// public key存在，判断email与数据库中的数据是否一致
	user.PublicKey = publicKey
	err := o.Read(&user, "public_key")
	if err == nil {
		if email != user.Email {
			logs.Info("public key 已存在")
			c.ErrorJson("400000", "绑定失败：邮箱已绑定")
			return
		}
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

	user.Status = 1 // 已验证
	o.Update(&user)

	c.SuccessJson("", "")
}

// @Title GetImPubKeyList
// @Description 获取IM公钥列表(除当前用户)
// @Param email query string false "Email"
// @Success 200 {object} controllers.RespJson
// @router /getimpubkeylist [get]
func (c *UserController) GetImPubKeyList() {
	CurUser := c.CurUser()
	var data []models.User

	user := new(models.User)
	qt := orm.NewOrm().QueryTable(user)

	email := c.GetString("email")
	if email != "" {
		hashEmail := crypto.Md5(email)
		qt = qt.Filter("email", hashEmail)
	}

	hashEmailForCurUser := crypto.Md5(CurUser.Email)
	qt.Exclude("email__in", hashEmailForCurUser).All(&data, "Email", "NostrPublicKey")
	c.SuccessJson("", data)
}

// @Title GetUserInfo
// @Description 获取当前用户信息
// @Success 200 {object} controllers.RespJson
// @router /getuserinfo [get]
func (c *UserController) GetUserInfo() {
	var user UserInfo
	CurUser := c.CurUser()

	if CurUser.SssData != "" {
		key, _ := config.String("crypto")
		deKey := crypto.DecryptBase64(key)
		ct := crypto.DecryptAES(CurUser.SssData, deKey)
		user.SssData = ct
	}

	// walletAddress := CurUser.Address
	// ipns, _ := utils.GetDFSPath(walletAddress)
	// user.Ipns = ipns
	user.Ipns = CurUser.Ipns

	user.DbAddress = CurUser.DbAddress
	user.Name = CurUser.Name
	user.Email = CurUser.Email // hash值
	user.SafeLevel = CurUser.SafeLevel
	user.ImgCid = CurUser.ImgCid
	c.SuccessJson("", user)
}

// @Title UpdateImPkey
// @Description 更新当前用户IM公钥
// @Param nostrPublicKey body string true "Email"
// @Success 200 {object} controllers.RespJson
// @router /updateimpkey [post]
func (c *UserController) UpdateImPkey() {
	CurUser := c.CurUser()
	logs.Info("cur user = ", CurUser)
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)
	logs.Info("user = ", user)

	o := orm.NewOrm()
	user.Id = CurUser.Id

	_, err := o.Update(&user, "nostr_public_key")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "更新聊天公钥失败")
	} else {
		c.SuccessJson("", "")
	}

}

// @Title ModifyUser
// @Description 更新当前用户信息
// @Param name body string false "用户名称"
// @Param sssData body string false "分片数据"
// @Param publicKey body string false "公钥"
// @Param address body string false "钱包地址"
// @Param sign body string false "签名"
// @Param ipns body string false "IPNS"
// @Param dbAddress body string false "数据库地址"
// @Success 200 {object} controllers.RespJson
// @router /modifyuser [post]
func (c *UserController) ModifyUser() {
	CurUser := c.CurUser()
	logs.Info("cur user = ", CurUser)
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
		deKey := crypto.DecryptBase64(key)
		enData := crypto.EncryptAES(sssData, deKey)
		CurUser.SssData = enData
		// CurUser.SssData = sssData
	}

	// nostrPublicKey := user.NostrPublicKey
	// if nostrPublicKey != "" {
	// 	CurUser.NostrPublicKey = nostrPublicKey
	// }

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
		if ipns != CurUser.Ipns {
			if walletAddress == "" {
				walletAddress = CurUser.Address
			}

			success, err := utils.SetDFSPath(walletAddress, ipns, user.Sign) // 更新钱包中用户的dfs地址
			if !success {
				logs.Error(err)
				c.ErrorJson("400000", err.Error())
				return
			}

		}
		CurUser.Ipns = ipns
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

// @Title SendMail
// @Description 发送验证码
// @Param email body string true "Email"
// @Success 200 {object} controllers.RespJson
// @router /sendmail [post]
func (c *UserController) SendMail() {
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	email := user.Email

	verify, msg := verifyEmail(email) // email不需转为hash
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	hashEmail := crypto.Md5(email)
	o := orm.NewOrm()
	tmpUser := models.User{Email: hashEmail}
	status := -1
	err := o.Read(&tmpUser, "email")
	if err != orm.ErrNoRows {
		status = tmpUser.Status
	}

	// confirmCode := utils.RandomNum(6)
	confirmCode := "123456" // TODO:for test

	if status == -1 { // email不存在
		user = models.User{Email: hashEmail, ConfirmCode: confirmCode, Status: 0, ConfirmCodeUpdateTime: time.Now()}
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
		user = models.User{Email: hashEmail}
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
	hashEmail := crypto.Md5(email)
	o := orm.NewOrm()
	user := models.User{Email: hashEmail}

	err := o.Read(&user, "email")
	if err == orm.ErrNoRows {
		status = -1
	} else {
		status = user.Status
	}
	return status
}

func verifyEmailAndConfirmCode(user models.User) (bool, string) {
	var msg string
	flag := true

	email := strings.TrimSpace(user.Email)
	confirmCode := strings.TrimSpace(user.ConfirmCode)

	verify, msg := verifyEmail(email) // email不需转为hash
	if !verify {
		return flag, msg
	}

	verify, msg = verifyConfirmCode(confirmCode)
	if !verify {
		return flag, msg
	}

	hashEmail := crypto.Md5(email)

	// 判断mail是否存在
	o := orm.NewOrm()
	user.Email = hashEmail
	err := o.Read(&user, "email")
	if err != nil {
		msg = email + "不存在"
		flag = false
		return flag, msg
	}

	// 判断mail + confirm code是否匹配
	user.ConfirmCode = confirmCode
	err = o.Read(&user, "email", "confirm_code")
	if err != nil {
		msg = "验证码错误"
		flag = false
		return flag, msg
	}

	// 判断验证码是否过期
	confirmCodeUpdateTime := user.ConfirmCodeUpdateTime
	curTime := time.Now()
	if curTime.Sub(confirmCodeUpdateTime).Seconds() > 60 { // 验证码1分钟失效
		msg = "验证码已过期"
		flag = false
		return flag, msg
	}

	return flag, msg
}

func verifyEmail(email string) (bool, string) {
	var msg string
	flag := true

	if email == "" {
		msg = "邮箱地址不能为空"
		flag = false
	} else if !utils.IsEmail(email) {
		msg = "邮箱地址格式错误"
		flag = false
	}
	return flag, msg
}

func verifyConfirmCode(confirmCode string) (bool, string) {
	var msg string
	flag := true
	if confirmCode == "" {
		msg = "验证码不能为空"
		flag = false
	}
	return flag, msg
}
