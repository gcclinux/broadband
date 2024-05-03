package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/showwin/speedtest-go/speedtest"
)

func GetOutboundIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func GetBroadbandSpeed() (latency, downloadspeed, uploadspeed float64, isp string) {

	var speedtestClient = speedtest.New()

	serverList, _ := speedtestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})
	location := targets.String()
	parts := strings.SplitN(location, " ", 3)

	if len(parts) >= 3 {
		isp = parts[2]
	}
	for _, s := range targets {
		s.PingTest(nil)
		s.DownloadTest()
		s.UploadTest()

		latency = s.Latency.Seconds()
		downloadspeed = s.DLSpeed
		uploadspeed = s.ULSpeed
		s.Context.Reset()
	}
	return latency, downloadspeed, uploadspeed, isp
}
