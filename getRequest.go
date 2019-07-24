package main 

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetData() {
	client := &http.Client{}
	resp, err := client.Get("https://www.easy-mock.com/mock/5c1b10665b48fa38e4318844/example/mock")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func main() {
	GetData()
}