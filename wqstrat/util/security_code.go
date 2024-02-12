package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

const (
	SecurityCodeUrl       = "/oauth2/tokenP"
	RemoveSecurityCodeUrl = "/oauth2/revokeP"
)

type OAuthSecurityCodeRequest struct {
	RESTAuth
	GrantType string `json:"grant_type"`
}

type OAuthRemoveSecurityCodeRequest struct {
	RESTAuth
	Token string `json:"token"`
}

type OAuthSecurityCodeResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	ExpiresAt   string `json:"access_token_token_expired"`
}

type OAuthRemoveSecurityCodeResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *KISClient) SetOAuthSecurityCode() error {
	if c.isOAuthKeyAvailable() {
		color.Green("already has security code")
		return nil
	}

	result := OAuthSecurityCodeResponse{}

	// Create request information
	body := OAuthSecurityCodeRequest{
		GrantType: "client_credentials",
		RESTAuth:  c.UserInfoREST,
	}
	bstr, err := json.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		"POST",
		whereToRequest(c.isTest, SecurityCodeUrl),
		bytes.NewReader(bstr),
	)
	if err != nil {
		return err
	}

	// Set header
	request.Header.Set("content-type", "application/json; utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("failed to get security code: %v", err)
	}
	defer response.Body.Close()

	// Parse the response and register security code to the client
	bytes, _ := io.ReadAll(response.Body)
	if err = json.Unmarshal(bytes, &result); err != nil {
		return errors.New(fmt.Sprintf("failed to register security code: %v", err))
	} else {
		c.setSecurityCode(result)
		return nil
	}
}

func (c *KISClient) RemoveOAuthSecuritCode() error {
	if c.OAuthKey == "" {
		// No OAuthKey to remove
		return nil
	}

	result := OAuthRemoveSecurityCodeResponse{}

	// Create request information
	body := OAuthRemoveSecurityCodeRequest{
		Token:    c.OAuthKey,
		RESTAuth: c.UserInfoREST,
	}
	bstr, err := json.Marshal(body)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		"POST",
		whereToRequest(c.isTest, RemoveSecurityCodeUrl),
		bytes.NewReader(bstr),
	)
	if err != nil {
		return err
	}
	request.Header.Set("content-type", "application/json; utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("failed to remove security code: %v", err)
	}
	defer response.Body.Close()

	// Parse the result
	bytes, _ := io.ReadAll(response.Body)
	if err = json.Unmarshal(bytes, &result); err != nil {
		return errors.New(fmt.Sprintf("failed to get appropriate response removing security code: %v", err))
	} else {
		if result.Code != 200 {
			return errors.New(fmt.Sprintf("failed to remove security code: %s(%v)", result.Message, result.Code))
		}
		// Re-Initialize the OAuthKey and OAuthKeyExpire token
		color.Red("security code removed at %v", time.Now())
		c.OAuthKey = ""
		c.OAuthKeyExpire = time.Time{}
		return nil
	}
}

func (c *KISClient) setSecurityCode(secCd OAuthSecurityCodeResponse) {
	// secCd.ExpiresIn - 10. Give 10 seconds slack
	expireDue := time.Now().Add(time.Second * time.Duration(secCd.ExpiresIn-10))

	// Set OAuth keys and its expiration date
	c.OAuthKeyExpire = expireDue
	c.OAuthKey = secCd.AccessToken
}

func (c *KISClient) isOAuthKeyAvailable() bool {
	now := time.Now()

	// No OAuth key requested in the first place. Emit error
	if c.OAuthKey == "" {
		return false
	}

	if now.After(c.OAuthKeyExpire) {
		return false
	} else {
		return true
	}
}

func (c *KISClient) getBearerAuthorization() string {
	return fmt.Sprintf("Bearer %s", c.OAuthKey)
}
