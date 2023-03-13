package controllers

import (
	"hash/fnv"
)

type StorageController struct {
	BaseController
}

func (c *StorageController) Test() {
	h := fnv.New64a()
	h.Write([]byte("18098922101@189.cn"))
	seed := h.Sum64()

	c.SuccessJson("", seed)
}
