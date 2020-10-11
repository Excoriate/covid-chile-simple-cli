package libs

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPHeaders struct {
	Key   string
	Value string
}

type HTTPRequestArgs struct {
	Headers []HTTPHeaders
	Uri     string
	Method  string
}

/**
Make a HTTP GET request
*/
func ExecHttpRequest(args HTTPRequestArgs) string {
	
	var method = ""
	if args.Method == "" {
		method = "GET"
	} else {
		method = args.Method
	}
	
	req, _ := http.NewRequest(method, args.Uri, nil)
	
	for index := 0; index < len(args.Headers); index++ {
		 req.Header.Add(args.Headers[index].Key, args.Headers[index].Value)
	}
	
	res, err := http.DefaultClient.Do(req)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var bodyString = string(body)
	
	return bodyString
	
}
