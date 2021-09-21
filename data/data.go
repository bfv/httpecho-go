package data

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type RequestData struct {
	Host        map[string]string `json:"host"`
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

func GetCallData(c echo.Context) RequestData {
	data := RequestData{}

	addHostInfo(&data)

	data.GetHeaders(c.Request().Header)

	data.Caller = Caller{
		IP: c.RealIP(),
	}

	return data
}

func addHostInfo(req *RequestData) {

	host := map[string]string{}
	req.Host = host
	host["hostname"], _ = os.Hostname()

	// conn, _ := net.Dial("ip:icmp", "google.com")
	// host["ip"] = conn.LocalAddr().String()

}
