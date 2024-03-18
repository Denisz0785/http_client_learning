package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	request, err := http.NewRequest(http.MethodGet,
		"http://localhost:8080/ti", nil)
	if err != nil {
		fmt.Println("Ошибка формирования запроса", err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Ошибка отправки запроса", err)
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка чтения ответа", err)
		return
	}
	fmt.Println(string(body))

	// for {
	// 	response, err := http.Post("http://localhost:8080/", "text/plain", nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	body, err := io.ReadAll(response.Body)
	// 	response.Body.Close()
	// 	if err != nil {
	// 		fmt.Println("Error of reading", err)
	// 		break
	// 	}
	// 	fmt.Println(string(body))
	// 	time.Sleep(5 * time.Second)
	// }
}
