package main

import (
	"fmt"
	"net/http"
	"io/iounit"
)

func httpPost() {
	resp, err := http.Post("",
	"application/x-www-form-urlencoded",
	strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//handle error
	}

	fmt.Println(string(body))
}

 func main() {
	 httpPost()
 }