package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	red    = color.Red
	green  = color.Green
	yellow = color.Yellow
	bold   = color.New(color.Bold).SprintFunc()

	config = &Config{}

	Version = "1.1.0"
)

func init() {
	header()

	// read config file
	data, err := readFile("config.json")
	if err != nil {
		red(bold("Error: " + err.Error()))
		os.Exit(1)
	}

	// unmarshal json
	err = jsonToStruct(data, config)
	if err != nil {
		red(bold("Error: " + err.Error()))
		os.Exit(1)
	}

	if checkConfig() {
		red(bold("Please fill all the fields in config.json"))
		os.Exit(1)
	}
}

func main() {
	mobile := flag.String("to", "", "Mobile number to send SMS, eg: +919999999999")
	message := flag.String("message", "", "Message to send, eg: 'Hello World'")
	log := flag.Bool("log", false, "Log messages sent, if this flag is set, mobile and message flags are ignored")
	version := flag.Bool("version", false, "Print version and exit")

	flag.Parse()

	if *version {
		yellow(bold("Version: " + Version))
		return
	}

	if *log {
		response, err := seeMessages()
		if err != nil {
			red(bold("Error: " + err.Error()))
		}

		// unmarshal json
		res := &Response{}
		err = json.Unmarshal([]byte(response), res)
		if err != nil {
			red(bold("Error: " + err.Error()))
		}

		// print messages
		for _, v := range res.Messages {
			var c = color.Yellow

			if v.Status == "delivered" {
				c = green
			} else if v.Status == "undelivered" {
				c = red
			}

			c(bold("From: " + v.From))
			c(bold("To: " + v.To))
			c(bold("Status: " + v.Status))
			c(bold("Body: " + v.Body))
			c(bold("Date Created: " + v.DateCreated))
			c(bold("--------------------------------------------------"))
		}

		return
	}

	// check if mobile number and message are empty
	if *mobile == "" || *message == "" {
		yellow(bold("Mobile number or message is empty"))
		return
	}

	// send sms
	_, err := sendSMS(*mobile, *message)
	if err != nil {
		red(bold(err))
		return
	}

	green(bold("Message sent successfully"))
}

func sendSMS(mobile, message string) (body any, error error) {
	text := "To=" + mobile + "&Body=" + message + "&MessagingServiceSid=" + config.MessagingServiceSid
	client := &http.Client{}
	var data = strings.NewReader(text)
	req, err := http.NewRequest("POST", config.URL, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(config.AccountSid, config.AuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText, nil
}

func header() {
	text := "4pSM4pSs4pSQ4pSM4pSA4pSQ4pSM4pSA4pSQ4pSM4pSs4pSQ4pSM4pSA4" +
		"pSQ4pSs4pSA4pSQ4pSs4pSM4pSA4pSQ4pSM4pSA4pSQ4pSs4pSA4pSQCuKUgu" +
		"KUguKUguKUnOKUpCDilJTilIDilJDilILilILilILilJzilKQg4pSc4pSs4pS" +
		"Y4pSC4pSM4pSA4pSY4pSc4pSkIOKUnOKUrOKUmArilLQg4pS04pSU4pSA4pSY" +
		"4pSU4pSA4pSY4pS0IOKUtOKUlOKUgOKUmOKUtOKUlOKUgOKUtOKUlOKUgOKUm" +
		"OKUlOKUgOKUmOKUtOKUlOKUgA=="

	// decode base64
	decoded, _ := base64.StdEncoding.DecodeString(text)
	green(bold(string(decoded)))
}

// see messages sent
func seeMessages() (data string, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.URL, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(config.AccountSid, config.AuthToken)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), nil
}

// function open the file and read the content
func readFile(filename string) (content string, err error) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// read file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func jsonToStruct(data string, v interface{}) (err error) {
	err = json.Unmarshal([]byte(data), &v)
	if err != nil {
		return err
	}

	return nil
}

// check config struct is empty or not
func checkConfig() (empty bool) {
	if config.AccountSid == "" || config.AuthToken == "" || config.MessagingServiceSid == "" {
		return true
	}

	return false
}

type Response struct {
	FirstPageURI    string      `json:"first_page_uri"`
	End             int         `json:"end"`
	PreviousPageURI interface{} `json:"previous_page_uri"`
	Messages        []Message   `json:"messages"`
	URI             string      `json:"uri"`
	PageSize        int         `json:"page_size"`
	Start           int         `json:"start"`
	NextPageURI     string      `json:"next_page_uri"`
	Page            int         `json:"page"`
}

type Message struct {
	Body                string      `json:"body"`
	NumSegments         string      `json:"num_segments"`
	Direction           string      `json:"direction"`
	From                string      `json:"from"`
	DateUpdated         string      `json:"date_updated"`
	Price               string      `json:"price"`
	ErrorMessage        interface{} `json:"error_message"`
	URI                 string      `json:"uri"`
	AccountSid          string      `json:"account_sid"`
	NumMedia            string      `json:"num_media"`
	To                  string      `json:"to"`
	DateCreated         string      `json:"date_created"`
	Status              string      `json:"status"`
	Sid                 string      `json:"sid"`
	DateSent            string      `json:"date_sent"`
	MessagingServiceSid string      `json:"messaging_service_sid"`
	ErrorCode           interface{} `json:"error_code"`
	PriceUnit           string      `json:"price_unit"`
	APIVersion          string      `json:"api_version"`
	SubresourceUris     struct {
		Media    string `json:"media"`
		Feedback string `json:"feedback"`
	} `json:"subresource_uris"`
}

type Config struct {
	URL                 string `json:"url"`
	AccountSid          string `json:"account_sid"`
	AuthToken           string `json:"auth_token"`
	MessagingServiceSid string `json:"messaging_service_sid"`
}
