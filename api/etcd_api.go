package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/liuqianhong6007/etcd_backend/internal"
)

func init() {
	internal.AddRoute(internal.Routes{
		{
			Method:  http.MethodGet,
			Path:    "/etcd/get",
			Handler: getEtcd,
		},
		{
			Method:  http.MethodPost,
			Path:    "/etcd/add",
			Handler: putEtcd,
		},
		{
			Method:  http.MethodPost,
			Path:    "/etcd/delete",
			Handler: deleteEtcd,
		},
	})
}

func getEtcd(c *gin.Context) {
	key := c.Query("key")
	val, err := internal.Get(key)
	internal.CheckValue(c, err, "get etcd value error")

	internal.SuccessJsonRsp(c, val)
}

type PutEtcdParam struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func putEtcd(c *gin.Context) {
	var param PutEtcdParam
	err := c.BindJSON(&param)
	internal.CheckValue(c, err, "param format error")
	internal.CheckValue(c, param.Key != "", "param[key] is null")
	internal.CheckValue(c, param.Val != "", "param[val] is null")

	err = internal.Put(param.Key, param.Val)
	internal.CheckValue(c, err, "put etcd value error")

	internal.SuccessJsonRsp(c, nil)
}

type DelEtcdParam struct {
	Keys []string `json:"keys"`
}

func deleteEtcd(c *gin.Context) {
	var param DelEtcdParam
	err := c.BindJSON(&param)
	internal.CheckValue(c, err, "param format error")
	internal.CheckValue(c, len(param.Keys) != 0, "param[keys] is null")

	for _, key := range param.Keys {
		err := internal.Delete(key)
		internal.CheckValue(c, err, "delete etcd value error")
	}

	internal.SuccessJsonRsp(c, nil)
}
