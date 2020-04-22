package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	_"reflect"
)

//for json decode purpose, the variables must be defined to global variable
type TopicOBJ struct {
	TopicType string    `json:"type"`
	Topic     []string  `json:"topic"`
}

type SubPubOBJ struct {
	ClientID  int      `json:"client_id"`
	SubOBJ    []TopicOBJ  `json:"sub"`
	PubOBJ    []TopicOBJ  `json:"pub"`
}

type SubPubOBJs []SubPubOBJ

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("topics.json")
	check(err)
	//fmt.Println(reflect.TypeOf(data))
	//fmt.Print(string(data))

	//struct method
	fmt.Println("use struct method")
	var configuration SubPubOBJs
	err = json.Unmarshal([]byte(data), &configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Println(configuration)
	for _, id := range  configuration {
		fmt.Println(id.ClientID)
		for _, topicObj := range id.SubOBJ {
			fmt.Println(topicObj.TopicType)
			for _, topic := range topicObj.Topic {
				fmt.Println(topic)
			}
		}
	}

	fmt.Println()
	//map method
	fmt.Println("use map method")
	var jsonObj []map[string]interface{}
	//var subObj []map[string]interface{}
	err = json.Unmarshal([]byte(data), &jsonObj)
	if err != nil {
		fmt.Println("error:", err)
	}

	//number default use float64
	//you should always use interface{} to define when the top define using interface{}
	for i:=0; i<len(jsonObj); i++ {
		fmt.Println(jsonObj[i]["client_id"].(float64))
		subObj := jsonObj[i]["sub"].([]interface{})
		for _, p := range subObj {
			obj := p.(map[string]interface{})
			fmt.Println(obj["type"].(string))
			topics := obj["topic"].([]interface{})
			for _, topic := range topics {
				fmt.Println(topic.(string))
			}
		}
	}

}
