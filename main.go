package main

import (
	"assignment-3/config"
	"assignment-3/controller"
	"assignment-3/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func init() {
	// initialize connect to db
	// config.Connect()
	config.StartDB()
}

func main() {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				schedulerProcess()
			}
		}
	}()

	r := echo.New()
	r.POST("/environment", controller.AddDataEnvironment)

	r.Logger.Fatal(r.Start(":9000"))
}

func schedulerProcess() {
	url := "http://localhost:9000/environment"
	environment := models.Environment{
		Water: getRandomNumber(),
		Wind:  getRandomNumber(),
	}
	json_data, err := json.Marshal(environment)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	resp.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(resp)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body == nil {
		panic(errors.New("Err, response body is nil"))
	}

	if res.StatusCode != http.StatusOK {
		panic(res.Status)
	}

	printEnvirontment(res, &environment)

	defer res.Body.Close()
}

func printEnvirontment(res *http.Response, env *models.Environment) {
	var response models.Response
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := json.Marshal(map[string]int{
		"water": env.Water,
		"wind":  env.Wind,
	})

	if err != nil {
		log.Fatal(err)
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, resp, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	statusWater, statusWind := checkStatusEnvironment(env)

	fmt.Println(prettyJSON.String())
	fmt.Println("status water :", statusWater)
	fmt.Printf("status wind : %v\n\n", statusWind)
}

func checkStatusEnvironment(env *models.Environment) (string, string) {
	var statusWater string
	var statusWind string

	switch {
	case env.Water <= 5:
		statusWater = "aman"
	case env.Water >= 6 && env.Water <= 8:
		statusWater = "siaga"
	case env.Water > 8:
		statusWater = "bahaya"
	default:
		statusWater = "data tidak ditemukan"
	}

	switch {
	case env.Wind <= 6:
		statusWind = "aman"
	case env.Wind >= 7 && env.Wind <= 15:
		statusWind = "siaga"
	case env.Wind > 15:
		statusWind = "bahaya"
	default:
		statusWind = "data tidak ditemukan"
	}

	return statusWater, statusWind
}

func getRandomNumber() int {
	rand.NewSource(time.Now().UnixNano())
	min := 1
	max := 100
	return (rand.Intn(max-min+1) + min)
}
