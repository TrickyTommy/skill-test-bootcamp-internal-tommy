package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/*type Items struct{
 Items []Item json:"items"
}*/

type Item struct {
	Id        int       `json:"inventory_id"`
	Name      string    ` json:"name"`
	Type      string    `json:"type"`
	Tags      []string  `json:"tags"`
	Purchased int64     `json:"purchased_at"`
	Placement Placement `json:"placement"`
}

type Placement struct {
	Room_id int    `json:"room_id"`
	Name    string `json:"name"`
}

func (i Item) toString() string {
	bytes, err := json.Marshal(i)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

func getIteminMeetingRoom() {
	var index []int
	items := make([]Item, 5)
	raw, err := ioutil.ReadFile("items.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &items)
	for i, it := range items {

		if it.Placement.Name == "Meeting Room" {
			fmt.Printf("found item in meeting room at index %d \n", i+1)
			index = append(index, i)
			fmt.Println(it.toString())
		}

	}
}

func getElectronic() {

	items := make([]Item, 5)
	raw, err := ioutil.ReadFile("items.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &items)
	for i, it := range items {
		if it.Type == "electronic" {
			fmt.Printf("found Electronic at index %d \n", i+1)
			fmt.Println(it.toString())
		}

	}
}
func getFurniture() {
	items := make([]Item, 5)
	raw, err := ioutil.ReadFile("items.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &items)
	for i, it := range items {
		if it.Type == "furniture" {
			fmt.Printf("found furniture at index %d \n", i+1)
			fmt.Println(it.toString())
		}

	}
}

func getPurchasedon16Jan2020() {
	items := make([]Item, 5)
	raw, err := ioutil.ReadFile("items.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &items)
	for i, it := range items {

		then := time.Date(2020, 1, 16, 20, 34, 58, 651387237, time.UTC)
		timeT := time.Unix(it.Purchased, 0)

		if ((timeT.Year() == then.Year()) && (timeT.Month() == then.Month()) && (timeT.Day() == then.Day())) == true {
			fmt.Printf("found item purchased on 16 january 2020 at index %d \n", i+1)
			fmt.Println(it.toString())
		}

	}

}

func getBrownItem() {
	items := make([]Item, 5)
	raw, err := ioutil.ReadFile("items.json")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &items)
	for i, it := range items {
		for _, tag := range it.Tags {
			if tag == "brown" {
				fmt.Printf("found items with brown color at index %d \n", i+1)
				fmt.Println(it.toString())
			}
		}

	}
}

func main() {

	fmt.Println("++++++++++++++++++++++++ NOMOR 1 ++++++++++++++++++++++++")
	getIteminMeetingRoom()
	fmt.Println("++++++++++++++++++++++++ NOMOR 2 ++++++++++++++++++++++++")
	getElectronic()
	fmt.Println("++++++++++++++++++++++++ NOMOR 3 ++++++++++++++++++++++++")
	getFurniture()
	fmt.Println("++++++++++++++++++++++++ NOMOR 4 ++++++++++++++++++++++++")
	getPurchasedon16Jan2020()
	fmt.Println("++++++++++++++++++++++++ NOMOR 5 ++++++++++++++++++++++++")
	getBrownItem()

	time.Sleep(60 * time.Second)

}
