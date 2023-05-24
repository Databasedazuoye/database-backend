package filter

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
)

// http 请求拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 签名校验逻辑，逻辑调用自己的签名逻辑即可
		// checkSign := signature.CheckSign(c)
		// if !checkSign {
		// 	c.Abort()
		// 	return
		// }
		request := c.Request
		_ = request.ParseForm()
		// 获取参数值
		req, ok := getReqParam(c, request)
		if ok {
			return
		}
		// 获取响应结果
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Request = request
		c.Writer = blw
		// 执行下一个
		c.Next()
		// 获取 响应

	}
}

func getReqParam(c *gin.Context, request *http.Request) (string, bool) {
	req := request.Form.Encode()
	// if(len(strings.Trim(req,"")) <= 0)
	if !util.IsBlank(req) {
		return req, false
	}
	// 获取body的值，后面的请求参数消失，暂时注释
	//body := request.Body
	//data, err := ioutil.ReadAll(body)
	data, err := c.GetRawData()
	//if err != nil {
	//	// 执行下一个
	//	c.Next()
	//	return "", true
	//}
	//req = string(data)
	return "", false
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
