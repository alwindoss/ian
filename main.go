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
	categoryRequestPath = "/api_category.php"
)

func logError(err error) {
	log.Println(err)
}

func main() {
	fmt.Println("Categories:")
	for _, category := range getCategories() {
		fmt.Printf("ID: %d\t%s\n", category.ID, category.Name)
	}
}

func getQuestions() {
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

func getCategories() []Category {
	resp, err := http.Get(baseURL + categoryRequestPath)
	if err != nil {
		logError(fmt.Errorf("something went wrong with downloading the questions: %v", err))
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logError(fmt.Errorf("unable to read the response: %v", err))
	}
	var categResp CategoryResponse
	err = json.Unmarshal(data, &categResp)
	if err != nil {
		logError(fmt.Errorf("unable to parse the response: %v", err))
	}
	var categories []Category
	for _, category := range categResp.TriviaCategories {
		cat := Category{
			ID:   category.ID,
			Name: category.Name,
		}
		categories = append(categories, cat)
	}
	return categories
}
