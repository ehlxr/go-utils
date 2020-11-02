package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// type jwtAuthenticator struct {
	// 	keyFunc jwtgo.Keyfunc
	// }

	// var claims = &jwtgo.StandardClaims{}
	// tokenStr := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjZXJlcy5lbm5jbG91ZC5jbjphcHA6eTdnSU9hemlSIiwiY2VyZXMuZW5uY2xvdWQuY24vYXBwLmlkIjoieTdnSU9hemlSIiwiY2VyZXMuZW5uY2xvdWQuY24vbmFtZXNwYWNlIjoiZy1jZXJlcyJ9.JhSzDnUcCfenDFQkTudaAzLO2JJKaghTOPnHNT9bz4nysVFzSAD-wP4mIiQKTKGPYP4442QGbRtxocTZx-VTK7YkdEKh-QZDkpyyfNi7loTCdCDrcMUQHwK4w8zhZ8KzKOXQrmsYkMSO_kJ8FNKCpOpOeUS5zu-BN39MrgqwE5evFsE-9C-MhrsKzOxuLv5I_cF5AqNnfhHcdCdF7PhHEmXsWC8S_9ep21MxaPhXTspeZa56eZHylV5ddm-bj8WR4r_2OsBI0k1QRN_SZNh8j35eB-Ht3sReVBvYnAHyvGptB8kFTuN6fF-Lkxi-OhkxncAGpl0UpdA5gJ9U0zHaHEM18eE7yr2wQKZIPEWAvtP4xOVPK3GfCx6UAX5HU1Vp5OwSkIW-L6TIUfuGpTAfa36UyDpybXlQ6sU6kGbT5jTetffAjf3FLN4HbS61Mgj1QSjii2dUd2L3lT-jv1d2jSQJGtozL-sapRQ7o6F-IlIaRGmYV0AP7lhN7Pu-22SWseRBVYlkvdgcPXODm_WDmpxBVq77hAyJI2_ARAmbXRGfBDKmwD6kYD2YAvG8wAMiZApFazamwAIQKHmy0Y4tv8I2-r9YlOF5ri4vaZ36Uv65C9YaSL6ctbb25TwMHDVzwSaIYv-HhLMYxGxNBJxnOnIG-SHIE0f1rgYynJbL1sg"
	// token, err := jwtgo.ParseWithClaims(tokenStr, claims, func(*jwtgo.Token) (interface{}, error) {

	// 	f, err := os.Open("api-token-public.pem")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer f.Close()
	// 	fd, err := ioutil.ReadAll(f)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	return jwtgo.ParseRSAPublicKeyFromPEM(fd)

	// })
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// if !token.Valid {
	// 	return
	// }

	// log.Debugf("jwt token received: %v", claims)

	// // appKey@namespace
	// temp := strings.Split(claims.Subject, "@")
	// fmt.Println(temp[0])

	// sf, _ := util.NewWorker(1)
	// fmt.Println(uint64(sf.GetId()))
	//log.Trace("Hello %s!", "World") // YYYY/MM/DD 12:34:56 [TRACE] Hello World!
	//log.Info("Hello %s!", "World")  // YYYY/MM/DD 12:34:56 [ INFO] Hello World!
	//log.Warn("Hello %s!", "World")  // YYYY/MM/DD 12:34:56 [ WARN] Hello World!
	//log.Error("Hello %s!", "World")
	//// Graceful stopping all loggers before exiting the program.
	//log.Stop()

	// user := User{"root", "root@ddd.com", 1, false}
	var user User = User{
		username:    "root",
		email:       "root@ddd.com",
		signInCount: 1,
		active:      false,
	}
	fmt.Println(user)

	fmt.Println("Hello World!")

	a := "https://cdn.jsdelivr.net/gh/0vo/oss/images/kafka-logo.png 400 %}\n"
	a1 := strings.Split(a, " ")[0]

	fmt.Println(a1)                                                             //输出为：www.waylau.com/golang-strings-split-get-url/
	fmt.Println(strings.Index(a, "https://cdn.jsdelivr.net/gh/0vo/oss/images")) //输出为：www.waylau.com
}

type User struct {
	username    string
	email       string
	signInCount uint64
	active      bool
}

func revertFileByLine() {
	file, err := os.Open("/Users/ehlxr/Desktop/o")
	if err != nil {
		println(err)
	}
	defer file.Close()

	var strs []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			println(err)
		}
		if err == io.EOF {
			break
		}
		strs = append(strs, line)
	}

	f, err := os.Create("/Users/ehlxr/Desktop/n.txt")
	if err != nil {
		println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i := len(strs) - 1; i >= 0; i-- {
		_, _ = fmt.Fprint(w, strs[i])
	}
	_ = w.Flush()
}

//func init() {
//	err := log.NewConsole()
//	if err != nil {
//		panic("unable to create new logger: " + err.Error())
//	}
//
//	err = log.NewFileWithName("base", log.FileConfig{
//		Level:    log.LevelInfo,
//		Filename: "clog.log",
//	})
//	if err != nil {
//		panic("unable to create new logger: " + err.Error())
//	}
//
//	err = log.NewFileWithName("err", log.FileConfig{
//		Level:    log.LevelError,
//		Filename: "clog-err.log",
//	})
//	if err != nil {
//		panic("unable to create new logger: " + err.Error())
//	}
//}
