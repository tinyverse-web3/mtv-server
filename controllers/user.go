package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Name            string            `json:"name"`
	Password        string            `json:"password"`
	Email           string            `json:"email"`
	ConfirmCode     string            `json:"confirmCode"`
	GuardianSssData string            `json:"guardianSssData"`
	QuestionSssData string            `json:"questionSssData"`
	Ipns            string            `json:"ipns"`
	DbAddress       string            `json:"dbAddress"`
	Questions       []models.Question `json:"questions"`
	ImgCid          string            `json:"imgCid"`
	SafeLevel       int               `json:"safeLevel"`
}

// @Title UpdateSafeLevel
// @Description 更新用户安全等级。当安全等级低于当前用户的等级，则不更新。(需验签)
// @Param safeLevel body int true "安全等级"
// @Success 200 {object} controllers.RespJson
// @router /updatesafelevel [post]
func (c *UserController) UpdateSafeLevel() {
	CurUser := c.CurUser()

	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)
	safeLevel := user.SafeLevel

	o := orm.NewOrm()
	user.Id = CurUser.Id
	err := o.Read(&user, "id")
	if err != nil {
		c.ErrorJson("400000", "获取数据失败")
		return
	}
	if safeLevel < user.SafeLevel {
		logs.Info("不需更新安全等级")
		c.SuccessJson("", "")
	}

	user.SafeLevel = safeLevel
	_, err = o.Update(&user, "safe_level")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "更新安全等级失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title SavePassword
// @Description 保存用户密码(需验签)
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
// @Description 获取用户密码(不需验签)
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /getpassword [post]
func (c *UserController) GetPassword() {
	var tmpUser UserInfo
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &tmpUser)

	verify, msg := verifyEmailAndConfirmCode(tmpUser.Email, tmpUser.ConfirmCode)
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	o := orm.NewOrm()
	hashEmail := crypto.Md5(tmpUser.Email)
	var password string
	var user models.User
	user.Email = hashEmail
	err := o.Read(&user, "email")
	if err != nil {
		var guardian models.Guardian
		guardian.Account = hashEmail
		err = o.Read(&guardian, "account")
		if err != nil {
			c.ErrorJson("400000", "获取数据失败")
			return
		} else {
			user.Id = guardian.UserId
			err = o.Read(&user, "id")
			if err != nil {
				c.ErrorJson("400000", "获取数据失败")
				return
			}
		}
	}
	password = user.Password
	c.SuccessJson("", password)
}

type GuardianSssInfo struct {
	SssData   string            `json:"sssData"`
	Guardians []models.Guardian `json:"guardians"`
}

// @Title GetSssData4Guardian
// @Description 获取分片数据(守护者备份)和守护者列表(不需验签)
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /getsssdata4guardian [post]
func (c *UserController) GetSssData4Guardian() {
	var tmpUser UserInfo
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &tmpUser)

	verify, msg := verifyEmailAndConfirmCode(tmpUser.Email, tmpUser.ConfirmCode)
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	hashEmail := crypto.Md5(tmpUser.Email)
	o := orm.NewOrm()
	var guardian models.Guardian
	guardian.Account = hashEmail
	err := o.Read(&guardian, "account")
	if err != nil {
		logs.Error("守护者不存在")
		c.ErrorJson("400000", "获取分片数据失败")
		return
	}

	var user models.User
	user.Id = guardian.UserId
	err = o.Read(&user, "id")
	if err != nil {
		logs.Error("主账号不存在")
		c.ErrorJson("400000", "获取分片数据失败")
		return
	}
	if user.GuardianSssData == "" {
		c.ErrorJson("400000", "分片数据为空")
		return
	}
	key, _ := config.String("crypto")
	deKey := crypto.DecryptBase64(key)
	deData := crypto.DecryptAES(user.GuardianSssData, deKey)

	var info GuardianSssInfo
	info.SssData = deData

	var data []models.Guardian

	tmp := new(models.Guardian)
	qt := orm.NewOrm().QueryTable(tmp)
	qt.Filter("user_id", user.Id).All(&data, "type", "account", "accountMask")
	info.Guardians = data

	c.SuccessJson("", info)
}

// @Title GetSssData4Question
// @Description 获取分片数据(智能隐私备份，不需验签)
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /getsssdata4question [post]
func (c *UserController) GetSssData4Question() {
	var tmpUser UserInfo
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &tmpUser)

	verify, msg := verifyEmailAndConfirmCode(tmpUser.Email, tmpUser.ConfirmCode)
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	o := orm.NewOrm()
	var user models.User
	user.Email = crypto.Md5(tmpUser.Email)
	err := o.Read(&user, "email")
	if err != nil {
		c.ErrorJson("400000", "获取分片数据失败")
		return
	}
	if user.QuestionSssData == "" {
		c.ErrorJson("400000", "分片数据为空")
		return
	}

	key, _ := config.String("crypto")
	deKey := crypto.DecryptBase64(key)
	deData := crypto.DecryptAES(user.QuestionSssData, deKey)

	var userInfo UserInfo
	var data []models.Question
	question := new(models.Question)
	qt := orm.NewOrm().QueryTable(question)
	qt.Filter("user_id", user.Id).All(&data, "Id", "type", "Content", "title")
	userInfo.Questions = data
	userInfo.Email = user.Email
	userInfo.QuestionSssData = deData

	c.SuccessJson("", userInfo)
}

// @Title SaveSssData4Guardian
// @Description 保存分片数据(守护者备份，需验签)
// @Param guardianSssData body string true "分片数据"
// @Success 200 {object} controllers.RespJson
// @router /savesssdata4guardian [post]
func (c *UserController) SaveSssData4Guardian() {
	CurUser := c.CurUser()

	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	o := orm.NewOrm()
	user.Id = CurUser.Id

	key, _ := config.String("crypto")
	deKey := crypto.DecryptBase64(key)
	enData := crypto.EncryptAES(user.GuardianSssData, deKey)
	user.GuardianSssData = enData

	_, err := o.Update(&user, "guardian_sss_data")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "保存分片数据(守护者备份)失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title SaveSssData4Question
// @Description 保存分片数据(智能隐私备份，需验签)
// @Param questionSssData body string true "分片数据"
// @Success 200 {object} controllers.RespJson
// @router /savesssdata4question [post]
func (c *UserController) SaveSssData4Question() {
	CurUser := c.CurUser()

	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)

	o := orm.NewOrm()
	user.Id = CurUser.Id
	key, _ := config.String("crypto")
	deKey := crypto.DecryptBase64(key)
	enData := crypto.EncryptAES(user.QuestionSssData, deKey)
	user.QuestionSssData = enData

	_, err := o.Update(&user, "question_sss_data")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "保存分片数据(智能隐私备份)失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title GetImPubKeyList
// @Description 获取IM公钥列表(除当前用户，需验签)
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
// @Description 获取当前用户信息(需验签)
// @Success 200 {object} controllers.RespJson
// @router /getuserinfo [get]
func (c *UserController) GetUserInfo() {
	var user UserInfo
	curUser := c.CurUser()

	if user.QuestionSssData != "" {
		key, _ := config.String("crypto")
		deKey := crypto.DecryptBase64(key)
		ct := crypto.DecryptAES(user.QuestionSssData, deKey)
		user.QuestionSssData = ct
	}
	if user.GuardianSssData != "" {
		key, _ := config.String("crypto")
		deKey := crypto.DecryptBase64(key)
		ct := crypto.DecryptAES(user.GuardianSssData, deKey)
		user.GuardianSssData = ct
	}

	// walletAddress := CurUser.Address
	// ipns, _ := utils.GetDFSPath(walletAddress)
	// user.Ipns = ipns
	user.Ipns = curUser.Ipns

	user.DbAddress = curUser.DbAddress
	user.Name = curUser.Name
	user.Email = curUser.Email // hash值
	user.SafeLevel = curUser.SafeLevel

	ipfsGateWay, _ := config.String("ipfs_gate_way")
	user.ImgCid = ipfsGateWay + "/" + curUser.ImgCid

	c.SuccessJson("", user)
}

// @Title UpdateImPkey
// @Description 更新当前用户IM公钥(需验签)
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

// @Title UpdateName
// @Description 更新用户名称(需验签)
// @Param name body string false "用户名称"
// @Success 200 {object} controllers.RespJson
// @router /updatename [post]
func (c *UserController) UpdateName() {
	CurUser := c.CurUser()
	logs.Info("cur user = ", CurUser)
	var user models.User
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &user)
	logs.Info("user = ", user)
	o := orm.NewOrm()

	name := strings.TrimSpace(user.Name)
	if name == "" {
		c.ErrorJson("400000", "用户名不能为空")
	}

	if CurUser.Name != "" {
		c.ErrorJson("400000", "用户名只能免费修改一次。")
		return
	}

	if name != CurUser.Name {
		var data []models.User
		qt := orm.NewOrm().QueryTable(user)
		qt.Filter("name", name).Exclude("id__in", CurUser.Id).All(&data)
		if len(data) > 0 {
			c.ErrorJson("400000", "用户名已存在。")
			return
		}
	}

	user.Id = CurUser.Id
	user.Name = name
	_, err := o.Update(&user, "name")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "更新用户名称失败")
	} else {
		c.SuccessJson("", "")
	}
}

// @Title ModifyUser
// @Description 更新当前用户信息(需验签)
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

// @Title BindMail
// @Description 绑定邮箱(需验签，且如果未设置name，则随机生成name，格式为：mtv_6位随机数字)
// @Param public_key header string true "public key"
// @Param email body string true "Email"
// @Param confirmCode body string true "验证码"
// @Success 200 {object} controllers.RespJson
// @router /bindmail [post]
func (c *UserController) BindMail() {
	var tmpUser UserInfo

	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &tmpUser)

	logs.Info(tmpUser)
	email := strings.TrimSpace(tmpUser.Email)

	publicKey := c.Ctx.Request.Header["Public_key"][0]
	// 验证public key是否为空
	if publicKey == "" {
		c.ErrorJson("400000", "public_key不能为空")
		return
	}

	verify, msg := verifyEmailAndConfirmCode(tmpUser.Email, tmpUser.ConfirmCode)
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	hashEmail := crypto.Md5(email)
	o := orm.NewOrm()

	var user models.User
	user.Email = hashEmail
	err := o.Read(&user, "email")
	if err == nil {
		c.ErrorJson("400000", email+"已绑定。")
		return
	}

	user.PublicKey = publicKey
	err = o.Read(&user, "public_key")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "Bind mail faild!")
		return
	} else { // 在验签时已经insert user，所以此处只需添加守护者
		user.Email = hashEmail
		_, err := o.Update(&user)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "Bind mail faild!")
			return
		}

		// 生成默认守护者
		var guardian models.Guardian
		guardian = models.Guardian{UserId: user.Id, Account: hashEmail, AccountMask: utils.Mask(email), Type: "email"}
		_, err = o.Insert(&guardian)
		if err != nil {
			logs.Error(err)
			c.ErrorJson("400000", "Bind mail faild!")
			return
		} else {
			c.SuccessJson("", "")
			return
		}
	}
}

type VerifyCodeInfo struct {
	Email string `json:"email"`
}

// @Title SendMail4VerifyCode
// @Description 发送验证码(不需验签)
// @Param email body string true "email"
// @Success 200 {object} controllers.RespJson
// @router /sendmail4verifycode [post]
func (c *UserController) SendMail4VerifyCode() {
	var info VerifyCodeInfo
	body := c.Ctx.Input.RequestBody
	json.Unmarshal(body, &info)

	email := info.Email
	verify, msg := verifyEmail(email) // email不需转为hash
	if !verify {
		c.ErrorJson("400000", msg)
		return
	}

	hashEmail := crypto.Md5(email)

	// verifyCode := utils.RandomNum(6)
	verifyCode := "123456" // TODO:for test

	// 判断发送验证码的频率
	tmpVerifyCode, _ := utils.GetStr(hashEmail)

	if tmpVerifyCode != "" {
		c.ErrorJson("400000", "验证码已发送，请查看邮箱。")
		return
	}

	utils.SetStr(hashEmail, verifyCode, 1*time.Minute)
	subject := "发送验证码"
	message := `
		<p> Hi %s,</p>
		<p style="text-indent:2em">Thanks for using MTV.</p>
		<p style="text-indent:2em"></p>
		<p style="text-indent:2em">Grab the confirmation code below and enter it on the screen.</P>
		<h2 style="background:#eee;text-align:center;padding:8px;border-radius:5px">%s</h2>
		<p style="text-indent:2em"> - The MTV Team</p>
	`

	success := utils.Send(email, subject, fmt.Sprintf(message, strings.Split(email, "@")[0], verifyCode))

	if success {
		c.SuccessJson("", "")
	} else {
		c.ErrorJson("400000", "Send verification code faild!")
	}
}

// @Title UploadImg
// @Description 上传头像(需验签)
// @Param file body string true "图片文件"
// @Success 200 {object} controllers.RespJson
// @router /uploadimg [post]
func (c *UserController) UploadImg() {
	curUser := c.CurUser()

	file, header, err := c.GetFile("file")
	if err != nil {
		logs.Error("999999999999999999")
		c.ErrorJson("400000", err.Error())
	}
	fileName := header.Filename
	logs.Info("file name = ", fileName)
	logs.Info("file = ", file)

	defer file.Close()
	// c.SaveToFile("file", "upload/"+fileName)

	dataBytes, err := ioutil.ReadAll(file)
	dataStr := string(dataBytes[:])

	ipfsUrl := getIpfsUrl()
	ipfs := utils.NewIpfs(ipfsUrl)
	hash, _ := ipfs.Add(dataStr, true)

	curUser.ImgCid = hash
	o := orm.NewOrm()
	_, err = o.Update(&curUser, "img_cid")
	if err != nil {
		logs.Error(err)
		c.ErrorJson("400000", "上传头像失败")
	} else {
		ipfsGateWay, _ := config.String("ipfs_gate_way")
		url := ipfsGateWay + "/" + hash
		c.SuccessJson("", url)
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

func verifyEmailAndConfirmCode(email, verifyCode string) (bool, string) {
	var msg string
	flag := true

	email = strings.TrimSpace(email)
	verifyCode = strings.TrimSpace(verifyCode)

	verify, msg := verifyEmail(email) // email不需转为hash
	if !verify {
		return flag, msg
	}

	verify, msg = verifyConfirmCode(verifyCode)
	if !verify {
		return flag, msg
	}

	hashEmail := crypto.Md5(email)
	tmpVerifyCode, _ := utils.GetStr(hashEmail)
	// 判断验证码是否过期
	if tmpVerifyCode == "" { // 验证码1分钟失效
		msg = "验证码已过期"
		flag = false
		return flag, msg
	}
	// 判断mail + confirm code是否匹配
	if tmpVerifyCode != verifyCode {
		msg = "验证码错误"
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

func getIpfsUrl() (ipfs string) {
	ipfs, _ = config.String("ipfs_url")
	return
}
