package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/yanzay/tbot/model"
)

// GetAuth user on telegra.ph
func GetAuth(user model.User) (interface{}, error) {

	uri := "https://api.telegra.ph/createAccount?short_name=%s&author_name=%s"

	author := fmt.Sprintf("%s%s", user.FirstName, user.LastName)

	if len([]rune(user.UserName)) > 0 {
		author = user.UserName
	}

	uri = fmt.Sprintf(uri, user.FirstName, author)

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("GET auth error: %v", err)
	}
	defer resp.Body.Close()

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("JSON decode error when try to auth: %v\n", err)
	}

	return data, err
}

// CreatePage on telegra.ph
func CreatePage(page Page) (interface{}, error) {

	uri := "https://api.telegra.ph/createPage?access_token=%s&title=%s&author_name=%s&content=%s&return_content=%t"

	uri = fmt.Sprintf(uri, page.AccessToken, page.Title, page.AuthorName, page.Content, page.ReturnContent)

	fmt.Println(uri)

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("GET create page error: %v\n", err)
	}
	defer resp.Body.Close()

	var data interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("JSON decode error when try to create page: %v\n", err)
	}

	return data, err
}

// PrepageContent for telegra.ph API
func PrepareContent(content string) string {

	var data []ContentBlock

	c := strings.Split(content, "\n")

	for _, p := range c {

		p = strings.TrimSpace(p)
		p = strings.Replace(p, " ", "+", -1)

		if len([]rune(p)) > 0 {

			data = append(data, ContentBlock{
				Tag:      "p",
				Children: []string{p},
			})
		}
	}

	result, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Can't encode content json data: %v\n", err)
	}

	return string(result)
}
