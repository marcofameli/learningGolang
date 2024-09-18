package main

import (
	"io"
	"net/http"
)

func main() { // metodo 1
	//c := http.Client{}
	//jsonVar := bytes.NewBuffer([]byte(`{"name": "John"}`))
	//resp, err := c.Post("https://google.com.br", "application/json", jsonVar)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//io.CopyBuffer(os.Stdout, resp.Body, nil)

	c := http.Client{}
	req, err := http.NewRequest("GET", "https://google.com.br", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	println(string(body))
	// PAREI NA AULA 21

}
