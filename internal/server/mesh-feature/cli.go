package mesh_feature

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"k8s-test-backend/conf"
	"k8s-test-backend/pkg/client"
	"k8s-test-backend/pkg/common"
	"net/http"
	"strings"
)

func RegisterRoute(r *gin.RouterGroup, methods []string, name, host, mode string) {
	for _, methodItem := range methods {
		switch methodItem {
		case http.MethodGet:
			r.GET(ProxyModeSwitch(name, host, mode))
		case http.MethodPost:
			r.POST(ProxyModeSwitch(name, host, mode))
		case http.MethodPut:
			r.PUT(ProxyModeSwitch(name, host, mode))
		case http.MethodPatch:
			r.PATCH(ProxyModeSwitch(name, host, mode))
		case http.MethodDelete:
			r.DELETE(ProxyModeSwitch(name, host, mode))
		}
	}
}

func ProxyModeSwitch(name, host, mode string) (string, func(ctx *gin.Context)) {
	path := host
	if mode == conf.MapperModeDirectly {
		path = "/" + name
	} else if mode == conf.MapperModeHostReplace {
		path = "/" + name + "/*path"
	}

	if path == host {
		return path, emptyFunction()
	}

	return path, generatorHandle(host)
}

func emptyFunction() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(204, nil)
	}
}

func generatorHandle(host string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		logrus.Infof("Proxy method %s, to %s", ctx.Request.Method, host)
		paths := ctx.Param("path")
		request, err := CopyRequest(ctx, host+paths)
		if err != nil {
			ctx.JSON(400, err)
			common.Error(400, ctx, err)
			return
		}
		innerProcess(request, ctx)
	}
}

// deal with request and parse response
func innerProcess(request *http.Request, ctx *gin.Context) {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	resHeader, resBody, bodyStr, err := ParseResponse(response)
	if err != nil {
		client.BaseResponse(200, ctx, bodyStr)
		return
	}
	client.BaseResponse(200, ctx, client.HeaderResponse{
		Header: resHeader,
		Body:   resBody,
	})
}

// get response from target value
func ParseResponse(response *http.Response) (map[string][]string, map[string]interface{}, string, error) {
	var responseJsonBody map[string]interface{}
	readRes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil, string(readRes), err
	}
	err = json.Unmarshal(readRes, &responseJsonBody)
	return response.Header, responseJsonBody, string(readRes), err
}

// copy the request form ctx
func CopyRequest(ctx *gin.Context, targetRul string) (*http.Request, error) {
	request, err := http.NewRequest(ctx.Request.Method, targetRul, ctx.Request.Body)
	if err != nil {
		return request, nil
	}
	// copy header
	for headerKey, headerValue := range ctx.Request.Header {
		request.Header.Set(headerKey, strings.Join(headerValue, ","))
	}
	// copy params
	request.URL.RawQuery = ctx.Request.URL.Query().Encode()
	return request, nil
}

func MeshInfo(ctx *gin.Context) {
	if !conf.ApplicationConfig.EnableServerFeature {
		ctx.JSON(200, "the service mesh is disable now")
		return
	}

	ctx.JSON(200, struct {
		Status bool        `json:"mesh_status"`
		List   interface{} `json:"list"`
	}{
		Status: conf.ApplicationConfig.EnableServerFeature,
		List:   conf.ApplicationConfig.ServiceMeshMapper,
	})
}
