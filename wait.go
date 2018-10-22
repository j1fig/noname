package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func requestTime(stopId int) *http.Response {
	email := url.PathEscape(os.Getenv("NONAME_EMAIL"))
	client := &http.Client{}

	baseUrl := "http://m.carris.pt/pt/tempo-espera-email/"
	queryParam := fmt.Sprintf("paragem=%v", stopId)
	url := fmt.Sprintf("%v?%v", baseUrl, queryParam)
	submitData := fmt.Sprintf("email=%v&my_request=this_is_my_submit", email)
	data := bytes.NewBufferString(submitData)

	req, _ := http.NewRequest("POST", url, data)

	// spoofing here
	req.Header.Add("Origin", "http://m.carris.pt")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	req.Header.Add("Referer", url)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "en,en-US;q=0.9,pt;q=0.8,it;q=0.7")

	response, _ := client.Do(req)
	return response
}
