package icpquery

import (
	"net/http"
	"strings"
	"time"

	"fmt"
	"math/rand"

	"github.com/PuerkitoBio/goquery"
)

type Icp struct {
	IcpNumber string `json:"icp_number"`
	IcpName   string `json:"icp_name"`
	Attr      string `json:"attr"`
	Date      string `json:"date"`
}

func ICPQuery(url string) (*Icp, error) {
	url = "https://icp.chinaz.com/" + url
	icp := new(Icp)
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return icp, err
	}
	request.Header.Add("user-agent", RandomUserAgent())

	resp, err := client.Do(request)
	if err != nil {
		return icp, err
	}
	defer resp.Body.Close()
	gp, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return icp, err
	}

	gp.Find("#first > li").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			icp.IcpName = strings.TrimSpace(s.Find("p").Text())
		}

		if i == 1 {
			icp.Attr = strings.TrimSpace(s.Find("p").Text())
		}

		if i == 2 {
			icp.IcpNumber = strings.TrimSpace(s.Find("p > font").Text())
		}

		if i == 6 {
			icp.Date = strings.TrimSpace(s.Find("p").Text())
		}

	})

	if icp.IcpName == "" {
		return icp, fmt.Errorf("没有查询到备案信息")
	}
	return icp, nil
}

var uaGens = []func() string{
	genFirefoxUA,
	genChromeUA,
}

// RandomUserAgent generates a random browser user agent on every request
func RandomUserAgent() string {
	rand.Seed(time.Now().Unix())
	return uaGens[rand.Intn(len(uaGens))]()
}

var ffVersions = []float32{
	58.0,
	57.0,
	56.0,
	52.0,
	48.0,
	40.0,
	35.0,
}

var chromeVersions = []string{
	"65.0.3325.146",
	"64.0.3282.0",
	"41.0.2228.0",
	"40.0.2214.93",
	"37.0.2062.124",
}

var osStrings = []string{
	"Macintosh; Intel Mac OS X 10_10",
	"Windows NT 10.0",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",
	"X11; Linux x86_64",
}

func genFirefoxUA() string {
	version := ffVersions[rand.Intn(len(ffVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

func genChromeUA() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}
