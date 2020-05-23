package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var UserAgentList = []string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1"}

var rateLimiter = time.Tick(1000 * time.Millisecond)

//从上面列表中随机获取一个User-Agent
func GetRandomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return UserAgentList[r.Intn(len(UserAgentList))]
}

//Fetch 解析URL并将文本转成utf8编码
func Fetch(url ...string) ([]byte, error) {
	<-rateLimiter
	request, err := http.NewRequest("GET", url[0], nil)
	if err != nil {
		log.Print(err.Error())
	}
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Content-Type", "")
	//request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Referer", url[1])
	request.Header.Add("User-Agent", GetRandomUserAgent())
	//client.Jar = jar
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(request.URL, []*http.Cookie{
		&http.Cookie{Name: "bid", Value: string(RandUp(11)), HttpOnly: true},
	})
	client, err := NewHttpClient("")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}
	// 获取页面编码方式
	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	// 转成utf8编码
	utf8reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8reader)
}

func NewHttpClient(proxyAddr string) (*http.Client, error) {
	proxy := http.ProxyFromEnvironment
	if len(proxyAddr) > 0 {
		proxyURL, err := url.Parse(proxyAddr)
		if err != nil {
			return nil, err
		}
		proxy = http.ProxyURL(proxyURL)
	}
	netTransport := &http.Transport{
		//Proxy:                  http.ProxyURL(proxy),
		Proxy:                 proxy,
		MaxConnsPerHost:       10,
		ResponseHeaderTimeout: time.Second * 5,
	}
	netTransport.DisableKeepAlives = true
	return &http.Client{
		Transport: netTransport,
		Timeout:   time.Second * 5,
	}, nil
}

// 获取页面的encoding信息
func determinEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
