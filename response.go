package main

// MixedQuestionResponse gets the response from api
type MixedQuestionResponse struct {
	ResponseCode int `json:"response_code"`
	Results      []struct {
		Category         string   `json:"category"`
		Type             string   `json:"type"`
		Difficulty       string   `json:"difficulty"`
		Question         string   `json:"question"`
		CorrectAnswer    string   `json:"correct_answer"`
		IncorrectAnswers []string `json:"incorrect_answers"`
	} `json:"results"`
}

// TokenResponse holds the response for the token request
type TokenResponse struct {
	ResponseCode    int    `json:"response_code"`
	ResponseMessage string `json:"response_message"`
	Token           string `json:"token"`
}

// CategoryResponse holds the response retured for the category request call
type CategoryResponse struct {
	TriviaCategories []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"trivia_categories"`
}

// Category holds the structure of category
type Category struct {
	ID   int
	Name string
}
