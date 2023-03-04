package controllers

type StorageController struct {
	BaseController
}

func (c *StorageController) Test() {
	c.SuccessJson("", "")
}
