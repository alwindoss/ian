package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseURL             = "https://opentdb.com"
	randomQuestionsPath = "/api.php?amount=10"
	tokenRequestPath    = "/api_token.php?command=request"
)

func logError(err error) {
	log.Println(err)
}

func main() {
	fmt.Println("Hello")
	resp, err := http.Get(baseURL + randomQuestionsPath)
	if err != nil {
		logError(fmt.Errorf("something went wrong with downloading the questions: %v", err))
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logError(fmt.Errorf("unable to read the response: %v", err))
	}
	var respObj MixedQuestionResponse
	err = json.Unmarshal(data, &respObj)
	if err != nil {
		logError(fmt.Errorf("unable to parse the response: %v", err))
	}
	for i, result := range respObj.Results {
		fmt.Printf("Question #%d: %s\n", i+1, string(result.Question))
	}
}

func getQuestions() {

}
