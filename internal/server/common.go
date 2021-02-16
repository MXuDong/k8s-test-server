package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	client "k8s-test-backend/package"
	"time"
)

/*
`common.go` support some common request handler. All the request mapping follow RestFul.
The method contain:
- Get
- Get-list
- Post
- Put
- Patch
- Delete
- Watch
*/

//CommonGet will set now time to response body
func CommonGet(ctx *gin.Context) {
	logrus.Infof("Common get request %s", time.Now())
	ctx.JSON(200, time.Now())
}

// CommonPost will set request body to response body
func CommonPost(ctx *gin.Context) {
	bytes, _ := ioutil.ReadAll(ctx.Request.Body)
	logrus.Infof("Common post %s", string(bytes))
	ctx.String(200, string(bytes))
}

// CommonDelete will set 204 to response http code, delete without result
func CommonDelete(ctx *gin.Context) {
	bytes, _ := ioutil.ReadAll(ctx.Request.Body)
	logrus.Infof("Common delete %s", string(bytes))
	ctx.JSON(204, nil)
}

// CommonPut will set now time to response body
func CommonPut(ctx *gin.Context) {
	logrus.Infof("Common Put request %s", time.Now())
	ctx.JSON(200, time.Now())
}

// CommonPatch will set now time to response body
func CommonPatch(ctx *gin.Context) {
	logrus.Infof("Common Patch request %s", time.Now())
	ctx.JSON(200, time.Now())
}

// ========================================================= cache about

var cacheClient = client.InitCacheClient()

// CachePost will create value of input, and the input must a json object
func CachePost(ctx *gin.Context) {
	value, _ := ioutil.ReadAll(ctx.Request.Body)
	jsonObject := map[string]interface{}{}
	err := json.Unmarshal(value, &jsonObject)
	if err != nil {
		ctx.JSON(500, err)
	}
	cacheClient.Add(jsonObject)

	logrus.Infof("Cache post %s", value)
	ctx.JSON(200, value)
}

func CacheList(ctx *gin.Context) {
	logrus.Infof("Cache list")
	ctx.JSON(200, cacheClient.Values)
}

// CacheGet will return value of target key and value from cache, the return is one object
func CacheGet(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")

	res := cacheClient.Find(key, value)

	logrus.Infof("Cache Get : k :[%s], v :[%s], res :[%s]", key, value, res)
	ctx.JSON(200, res)
}

// CacheDelete will delete target value for key and value, and set response to 204
func CacheDelete(ctx *gin.Context) {
	k := ctx.Param("key")
	v := ctx.Param("value")

	cacheClient.Delete(k, v)

	logrus.Infof("Cache Delete : k :[%s], v :[%s]", k, v)

	ctx.JSON(204, nil)
}

// CachePut will put the target object, and if some field not set, it will clear.
func CachePut(ctx *gin.Context) {
	k := ctx.Param("key")
	v := ctx.Param("value")

	value, _ := ioutil.ReadAll(ctx.Request.Body)
	if len(value) == 0 {
		ctx.String(400, "Update value not find")
		return
	}

	if cacheClient.Find(k, v) == nil {
		ctx.String(400, "Target value can't found.")
		return
	}

	// cover for all field
	cacheClient.Delete(k, v)
	jsonObject := map[string]interface{}{}
	err := json.Unmarshal(value, &jsonObject)
	if err != nil {
		ctx.JSON(500, err)
	}
	cacheClient.Add(jsonObject)
	logrus.Infof("Cache Put : k :[%s], v :[%s], res :[%s]", k, v, jsonObject)
	ctx.JSON(200, jsonObject)
}

// CachePatch will update field which input set, if some field not set, will keep value.
func CachePatch(ctx *gin.Context) {
	k := ctx.Param("key")
	v := ctx.Param("value")

	value, _ := ioutil.ReadAll(ctx.Request.Body)
	if len(value) == 0 {
		ctx.String(400, "Update not find")
		return
	}

	if temp := cacheClient.Find(k, v); temp == nil {
		ctx.String(400, "Target value can't found.")
		return
	} else {
		// do update
		jsonObject := map[string]interface{}{}
		err := json.Unmarshal(value, &jsonObject)
		if err != nil {
			ctx.JSON(500, err)
		}

		for ki, vi := range jsonObject {
			temp[ki] = vi
		}
		cacheClient.Delete(k, v)
		cacheClient.Add(temp)
		ctx.JSON(200, temp)
		logrus.Infof("Cache Patch : k :[%s], v :[%s], res :[%s]", k, v, temp)
	}
}
