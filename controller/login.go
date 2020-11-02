package controller

import (
	"fmt"
	"google.golang.org/api/oauth2/v2"
	"net/http"
)

var httpClient = &http.Client{}

func VerifyIdToken(token string)(*oauth2.Tokeninfo, error){
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(token)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	fmt.Println(tokenInfo.UserId)
	fmt.Println(tokenInfo.Email)
	fmt.Println(tokenInfo.Audience)
	fmt.Println(tokenInfo.VerifiedEmail)
	fmt.Println(tokenInfo.IssuedTo)
	return tokenInfo, nil
}