package main

import (
	"net/http"
	"bytes"
)

func main()  {
	var jsonStr = []byte(`{
				   "Site":["https://google.com","https://yahoo.com"],
   				   "SearchText":"Google"
			  }`);
	req, _ := http.NewRequest("POST", "http://localhost:8080/checkText", bytes.NewBuffer(jsonStr));
	client := &http.Client{}
	client.Do(req);
}
