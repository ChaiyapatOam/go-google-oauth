package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/chaiyapatoam/go-google-oauth/config"
	modelResponse "github.com/chaiyapatoam/go-google-oauth/model/response"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     config.GetEnv("GOOGLE_CLIENTID"),
		ClientSecret: config.GetEnv("GOOGLE_CLIENTSECRET"),
		RedirectURL:  config.GetEnv("GOOGLE_REDIRECT"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	return conf
}

func GetToken(c *fiber.Ctx) *oauth2.Token {
	token, err := GoogleConfig().Exchange(c.Context(), c.FormValue("code"))
	if err != nil {
		panic(err)
	}
	return token
}

func GetProfile(token string) modelResponse.GoogleResponse {
	reqUrl, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		panic(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)

	res := &http.Request{
		Method: "GET",
		URL:    reqUrl,
		Header: map[string][]string{"Authorization": {ptoken}},
	}

	req, err := http.DefaultClient.Do(res)
	if err != nil {
		fmt.Println("e", err)
		panic(err)
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data modelResponse.GoogleResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data

}
