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

	log.Printf("%v,", lb.conf)
	url, _ := url.Parse(lb.algorithm.FindNode())

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

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

func main() {
	rand.Seed(time.Now().UnixNano())

	loadBalancerConfig := loadConfig("config.json")
	lb := LoadBalancer{conf: loadBalancerConfig}
	lb.init()

	http.HandleFunc("/", lb.handle)
	log.Fatal(http.ListenAndServe(loadBalancerConfig.Ip+":"+strconv.Itoa(loadBalancerConfig.Port), nil))
}
