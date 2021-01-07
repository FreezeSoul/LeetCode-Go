package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://leetcode.com/api/problems/all/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var result []StatStatusPairs
	var lpa LeetCodeProblemAll
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &lpa)
	if err != nil {
		fmt.Println(err)
		return
	}
	result = lpa.StatStatusPairs
	//fmt.Println(result)
	res, _ := json.Marshal(result)
	write(res)
	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
}

func write(content []byte) {
	file, err := os.OpenFile("leetcode_problem", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write file successful")
}
