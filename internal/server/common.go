package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	client "k8s-test-backend/package"
	"time"
)

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

// CacheGet will return value of target key and value from cache, the return is one object
func CacheGet(ctx *gin.Context) {
	key := ctx.Param("key")
	value := ctx.Param("value")

	res := cacheClient.Find(key, value)

	logrus.Infof("Cache Get : k :[%s], v :[%s], res :[%s]", key, value, res)
	ctx.JSON(200, res)
}
