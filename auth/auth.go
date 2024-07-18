// TODO: Log any errors

// Authorization module
package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type authData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token"`
}

// Retrieves access and refresh tokens
func GetSessionToken(code string, codeVerifier string) authData {

	endpoint := "https://accounts.ea.com/connect/token"

	formData := url.Values{
		"token_format":  {"JWS"},
		"grant_type":    {"authorization_code"},
		"client_id":     {"JUNO_PC_CLIENT"},
		"client_secret": {clientSecret},
		"redirect_uri":  {redirectUri},
		"code":          {code},
		"code_verifier": {codeVerifier}}

	/*client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	reqBody := strings.NewReader(formData.Encode())

	req, err := http.NewRequest("POST", endpoint, reqBody)
	req.Header = headers

	resp, err := client.Do(req)*/

	resp, err := http.PostForm(endpoint, formData)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//var respBody []byte
	//_, err = resp.Body.Read(respBody)
	respBody, err := io.ReadAll(resp.Body)
	var tokens authData

	// Parse JSON and check for errors
	if err = json.Unmarshal(respBody, &tokens); err != nil {
		panic(err)
	}

	return tokens
}
