package data

import (
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RequestData struct {
	Host        map[string]string `json:"host"`
	Environment map[string]string `json:"env"`
	Headers     map[string]string `json:"headers"`
	QueryParams map[string]string `json:"query"`
	Caller      Caller            `json:"caller"`
	Body        string            `json:"body"`
}

func (req *RequestData) GetHeaders(h http.Header) {
	if req.Headers == nil {
		req.Headers = make(map[string]string)
	}
	for k, v := range h {
		req.Headers[k] = v[0]
	}
}

type Caller struct {
	IP string `json:"ip"`
}

var localIP net.IP
var calls int

func GetCallData(c echo.Context) RequestData {

	calls++

	data := RequestData{}

	addEnvironmentInfo(&data)
	addHostInfo(&data)

	data.GetHeaders(c.Request().Header)

	data.Caller = Caller{
		IP: c.RealIP(),
	}

	return data
}

func addEnvironmentInfo(req *RequestData) {
	env := map[string]string{}
	req.Environment = env

	env["name"] = os.Getenv("INSTANCE_NAME")
	env["calls"] = strconv.Itoa(calls)
}

func addHostInfo(req *RequestData) {

	if localIP == nil {
		localIP = getOutboundIP()
	}

	host := map[string]string{}
	req.Host = host
	host["hostname"], _ = os.Hostname()
	host["ip"] = localIP.String()
	host["os"] = runtime.GOOS
	host["architecture"] = runtime.GOARCH
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
