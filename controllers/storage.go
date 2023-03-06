package controllers

import (
	"fmt"
	"mtv/utils/crypto"

	"github.com/beego/beego/v2/core/config"
)

type StorageController struct {
	BaseController
}

func (c *StorageController) Test() {
	pt := "0x8420A24D450Da2dAB02C21ccEEd78C71a04E0005"
	key, _ := config.String("crypto")
	fmt.Println("key = ", key)
	deKey := crypto.DecryptBase64(key)
	fmt.Println("deKey = ", deKey)

	// e := crypto.EncryptAES([]byte(deKey), pt)
	// ct := crypto.DecryptAES([]byte(deKey), e)
	e := crypto.EncryptAES(pt, deKey)
	ct := crypto.DecryptAES(e, deKey)

	c.SuccessJson("", ct)
}
