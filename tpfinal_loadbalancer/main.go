package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"

	"go.uber.org/ratelimit"
	. "go.uber.org/ratelimit"
)

type node struct {
	Ip       string `json:"ip"`
	Port     int    `json:"port"`
	Protocol string `json:protocol`
}

type loadBalancerConfig struct {
	Ip               string `json:"ip"`
	Port             int    `json:"port"`
	DefaultAlgorithm string `json:"default_algorithm"`
	RateLimit        int    `json:"rate_limit"`
	Nodes            []node `json:"nodes"`
}

type LoadBalancer struct {
	conf      loadBalancerConfig
	algorithm LBAlgorithm
	rl        Limiter
}

type LBAlgorithm interface {
	FindNode() string
}

type RandomAlgorithm struct {
	conf loadBalancerConfig
}

func (a RandomAlgorithm) FindNode() string {
	var targetIndex = rand.Intn(len(a.conf.Nodes))
	node := a.conf.Nodes[targetIndex]
	fmt.Println(node)
	return node.Protocol + "://" + node.Ip + ":" + strconv.Itoa(node.Port)
}

type RoundRobinAlgorithm struct {
	conf      loadBalancerConfig
	lastIndex *int
}

func (a RoundRobinAlgorithm) FindNode() string {
	node := a.conf.Nodes[*a.lastIndex]
	*a.lastIndex += 1
	if *a.lastIndex == len(a.conf.Nodes) {
		*a.lastIndex = 0
	}
	fmt.Println(node)
	return node.Protocol + "://" + node.Ip + ":" + strconv.Itoa(node.Port)
}

func (lb *LoadBalancer) init() {
	lb.rl = ratelimit.New(lb.conf.RateLimit) // per second

	if lb.conf.DefaultAlgorithm == "random" {
		log.Println("Using random algorithm")
		lb.algorithm = RandomAlgorithm{conf: lb.conf}
	} else if lb.conf.DefaultAlgorithm == "round-robin" {
		log.Println("Using round-robin algorithm")
		lastIndex := 0
		lb.algorithm = RoundRobinAlgorithm{conf: lb.conf, lastIndex: &lastIndex}
	} else {
		log.Println("Using random algorithm")
		lb.algorithm = RandomAlgorithm{conf: lb.conf}
	}
}

func (lb *LoadBalancer) handle(res http.ResponseWriter, req *http.Request) {
	lb.rl.Take()

	// parse the url
	log.Printf("%v,", lb.conf)
	url, _ := url.Parse(lb.algorithm.FindNode())

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func loadConfig(filePath string) loadBalancerConfig {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var conf loadBalancerConfig
	json.Unmarshal(byteValue, &conf)

	return conf
}

// func LB2Handler(res http.ResponseWriter, req *http.Request) {

// 	// parse the url
// 	url, _ := url.Parse("http://localhost:8082")

// 	// create the reverse proxy
// 	proxy := httputil.NewSingleHostReverseProxy(url)

// 	// Update the headers to allow for SSL redirection
// 	req.URL.Host = url.Host
// 	req.URL.Scheme = url.Scheme
// 	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
// 	req.Host = url.Host

// 	// Note that ServeHttp is non blocking and uses a go routine under the hood
// 	proxy.ServeHTTP(res, req)
// 	//https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b
// }

// func LBHandler(w http.ResponseWriter, r *http.Request) {
// 	// log.Printf("Path: %s", html.EscapeString(r.URL.Path))
// 	// log.Printf("Method: %s", html.EscapeString(r.Method))
// 	// log.Printf("Host: %s", html.EscapeString(r.Host))
// 	// log.Printf("Protocol: %s", html.EscapeString(r.Proto))
// 	// body, _ := ioutil.ReadAll(r.Body)
// 	// log.Printf("Body: %v", body)
// 	log.Printf("Headers: %v", r.Header)

// 	client := &http.Client{}
// 	req, _ := http.NewRequest(html.EscapeString(r.Method), "http://localhost:8082"+html.EscapeString(r.URL.Path), r.Body)
// 	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
// 	for k, v := range r.Header {
// 		log.Printf("Header: %s %s", k, v)
// 		req.Header.Add(k, v[0])
// 	}
// 	resp, _ := client.Do(req)
// 	log.Printf("Respuesta: %v", resp)
// 	body, _ := ioutil.ReadAll(resp.Body)

// 	for k, v := range resp.Header {
// 		w.Header().Set(k, v[0])
// 	}

// 	fmt.Fprint(w, body)

// 	log.Println("===========================================================================")
// }

func main() {
	rand.Seed(time.Now().UnixNano())

	//https://productos-p6pdsjmljq-uc.a.run.app/dia/productos/1
	loadBalancerConfig := loadConfig("config.json")
	//log.Printf("%v,", loadBalancerConfig)
	lb := LoadBalancer{conf: loadBalancerConfig}
	lb.init()

	http.HandleFunc("/", lb.handle)
	log.Fatal(http.ListenAndServe(loadBalancerConfig.Ip+":"+strconv.Itoa(loadBalancerConfig.Port), nil))
}
