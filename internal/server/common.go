package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	httpClient "k8s-test-backend/pkg/client"
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
	logrus.Infof("Common get request %v", time.Now())
	httpClient.GetResponse(ctx, time.Now())
}

// CommonPost will set request body to response body
func CommonPost(ctx *gin.Context) {
	bytes, _ := ioutil.ReadAll(ctx.Request.Body)
	logrus.Infof("Common post %v", string(bytes))
	httpClient.PostResponse(ctx, string(bytes))
}

// CommonDelete will set 204 to response http code, delete without result
func CommonDelete(ctx *gin.Context) {
	bytes, _ := ioutil.ReadAll(ctx.Request.Body)
	logrus.Infof("Common delete %v", string(bytes))
	httpClient.DeleteResponse(ctx)
}

// CommonPut will set now time to response body
func CommonPut(ctx *gin.Context) {
	logrus.Infof("Common Put request %v", time.Now())
	httpClient.PutResponse(ctx, time.Now())
}

// CommonPatch will set now time to response body
func CommonPatch(ctx *gin.Context) {
	logrus.Infof("Common Patch request %v", time.Now())
	httpClient.PatchResponse(ctx, time.Now())
}

// ========================================================= cache about

var cacheClient = httpClient.InitCacheClient()

// CachePost will create value of input, and the input must a json object
func CachePost(ctx *gin.Context) {
	value, _ := ioutil.ReadAll(ctx.Request.Body)
	jsonObject := map[string]interface{}{}
	err := json.Unmarshal(value, &jsonObject)
	if err != nil {
		ctx.JSON(500, err)
	}
	cacheClient.Add(jsonObject)

	logrus.Infof("Cache post %v", value)
	httpClient.PostResponse(ctx, value)
}

func CacheList(ctx *gin.Context) {
	logrus.Infof("Cache list")
	httpClient.GetResponse(ctx, cacheClient.Values)
}

// CacheGet will return value of target key and value from cache, the return is one object
func CacheGet(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")

	res := cacheClient.Find(key, value)

	logrus.Infof("Cache Get : k :[%v], v :[%v], res :[%v]", key, value, res)
	httpClient.GetResponse(ctx, res)
}

// CacheDelete will delete target value for key and value, and set response to 204
func CacheDelete(ctx *gin.Context) {
	k := ctx.Param("key")
	v := ctx.Param("value")
	cacheClient.Delete(k, v)
	logrus.Infof("Cache Delete : k :[%v], v :[%v]", k, v)
	httpClient.DeleteResponse(ctx)
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
	logrus.Infof("Cache Put : k :[%v], v :[%v], res :[%v]", k, v, jsonObject)
	httpClient.PutResponse(ctx, jsonObject)
}

func CacheClean(ctx *gin.Context) {
	cacheClient.Clean()
	logrus.Infof("Cache clean on [%v]", time.Now())
	httpClient.DeleteResponse(ctx)
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
		httpClient.PatchResponse(ctx, temp)
		logrus.Infof("Cache Patch : k :[%v], v :[%v], res :[%v]", k, v, temp)
	}
}
