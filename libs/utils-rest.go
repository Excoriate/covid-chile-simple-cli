package libs

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPHeaders struct {
	Key string
	Value string
}

type HttpRequestArguments struct {
	Headers HTTPHeaders
	uri string
	method string
}

/**
Make a HTTP GET request
 */
func ExecHttpRequest(url string, method string) string {
	
	req, _ := http.NewRequest(method, url, nil)
	
	// TODO: replace later as secrets
	req.Header.Add("x-rapidapi-host", "chile-coronapi1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "f255b0bc76msh76c5f8695aae921p10ccfbjsn2dc36c8dabde")
	
	res, err := http.DefaultClient.Do(req)
	
	if err != nil{
		log.Fatal(err)
	}
	
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var bodyString = string(body)
	
	return bodyString
	
}
