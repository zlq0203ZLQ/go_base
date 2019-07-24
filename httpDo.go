package main 

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST",
	"http://www.01happy.com/demo/accept.php",
	strings.NewReader("name=cjb")	
	)

	if err != nil {

	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//handle err
	}
	fmt.Println(string(body))
}