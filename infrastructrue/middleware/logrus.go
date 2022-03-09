package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	//配置项
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	return handle
}

func handle(c *gin.Context) {
	startTime := time.Now()
	ip := c.ClientIP()
	path := c.Request.URL.Path
	method := c.Request.Method
	param := c.Request.URL.RawQuery
	header := c.Request.Header
	body, _ := c.GetRawData()
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	//body := c.Request.Body

	logrus.Infof("request|%d|ip:%s|%s|%s|%s|%s|%s", &c, ip, method, path, param, body, header)
	c.Next()

	blw := bodyLogWriter{
		bodyBuf:        bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}

	delay := time.Now().Sub(startTime)
	status := c.Writer.Status()
	c.Writer = blw
	strBody := strings.Trim(blw.bodyBuf.String(), "\n")
	if len(strBody) > 512 {
		strBody = strBody[:(512 - 1)]
	}

	logrus.Infof("response|%d|%s|%d|%s", &c, delay, status, strBody)

	//logrus.Infof()
	//c.
	//请求后

}
