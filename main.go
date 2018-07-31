package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	jwtgo "github.com/dgrijalva/jwt-go"

	"github.com/ehlxr/go-utils/utils/log"
)

func main() {

	type jwtAuthenticator struct {
		keyFunc jwtgo.Keyfunc
	}

	var claims = &jwtgo.StandardClaims{}
	tokenStr := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZXJlcy5lbm5jbG91ZC5jbjphcHA6eTdnSU9hemlSIiwiY2VyZXMuZW5uY2xvdWQuY24vYXBwLmlkIjoieTdnSU9hemlSIiwiY2VyZXMuZW5uY2xvdWQuY24vbmFtZXNwYWNlIjoiZy1jZXJlcyJ9.JhSzDnUcCfenDFQkTudaAzLO2JJKaghTOPnHNT9bz4nysVFzSAD-wP4mIiQKTKGPYP4442QGbRtxocTZx-VTK7YkdEKh-QZDkpyyfNi7loTCdCDrcMUQHwK4w8zhZ8KzKOXQrmsYkMSO_kJ8FNKCpOpOeUS5zu-BN39MrgqwE5evFsE-9C-MhrsKzOxuLv5I_cF5AqNnfhHcdCdF7PhHEmXsWC8S_9ep21MxaPhXTspeZa56eZHylV5ddm-bj8WR4r_2OsBI0k1QRN_SZNh8j35eB-Ht3sReVBvYnAHyvGptB8kFTuN6fF-Lkxi-OhkxncAGpl0UpdA5gJ9U0zHaHEM18eE7yr2wQKZIPEWAvtP4xOVPK3GfCx6UAX5HU1Vp5OwSkIW-L6TIUfuGpTAfa36UyDpybXlQ6sU6kGbT5jTetffAjf3FLN4HbS61Mgj1QSjii2dUd2L3lT-jv1d2jSQJGtozL-sapRQ7o6F-IlIaRGmYV0AP7lhN7Pu-22SWseRBVYlkvdgcPXODm_WDmpxBVq77hAyJI2_ARAmbXRGfBDKmwD6kYD2YAvG8wAMiZApFazamwAIQKHmy0Y4tv8I2-r9YlOF5ri4vaZ36Uv65C9YaSL6ctbb25TwMHDVzwSaIYv-HhLMYxGxNBJxnOnIG-SHIE0f1rgYynJbL1sg"
	token, err := jwtgo.ParseWithClaims(tokenStr, claims, func(*jwtgo.Token) (interface{}, error) {

		f, err := os.Open("api-token-public.pem")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fd, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		return jwtgo.ParseRSAPublicKeyFromPEM(fd)

	})
	if err != nil {
		fmt.Print(err)
	}

	if !token.Valid {
		return
	}

	log.Debugf("jwt token received: %v", claims)

	// appKey@namespace
	temp := strings.Split(claims.Subject, "@")
	fmt.Println(temp[0])
}
