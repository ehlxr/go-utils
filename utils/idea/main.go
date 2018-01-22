package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetOutput(os.Stdout)

	log.Println("************************************************************")
	log.Printf("** %-55s**", "JetBrains License Server")
	log.Printf("** %-55s**", "Please support genuine!!!")

	port := flag.Int("p", 21017, "port")
	host := flag.String("h", "0.0.0.0", "Bind TCP Address")
	flag.Parse()

	log.Printf("** listen on %-45s**", fmt.Sprintf("%s:%d...", *host, *port))

	addr := fmt.Sprintf("%s:%d", *host, *port)
	if strings.Contains(addr, "0.0.0.0") {
		addr = strings.Replace(addr, "0.0.0.0", "", 1)
		*host = strings.Replace(*host, "0.0.0.0", "127.0.0.1", 1)
	}

	log.Printf("** You can use %-43s**", fmt.Sprintf("http://%s:%d as license server", *host, *port))
	log.Println("************************************************************")

	routerBinding()
	err := http.ListenAndServe(addr, http.DefaultServeMux)
	if err != nil {
		log.Fatalln(err)
	}
}

func urlMatcher(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.ToLower(r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func routerBinding() {
	mux := http.NewServeMux()
	http.Handle("/", urlMatcher(mux))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/rpc/ping.action", ping)

	mux.HandleFunc("/rpc/obtainticket.action", obtainTicket)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is starting!"))
}

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	salt := r.URL.Query().Get("salt")
	xmlResponse := "<PingResponse><message></message><responseCode>OK</responseCode><salt>" + salt + "</salt></PingResponse>"
	xmlSignature, _ := signature(xmlResponse)
	w.Header().Add("Content-Type", "text/xml")
	w.Write([]byte("<!-- " + xmlSignature + " -->\n" + xmlResponse))
}

func obtainTicket(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL)
	//buildDate := r.URL.Query().Get("buildDate")
	//clientVersion := r.URL.Query().Get("clientVersion")
	//hostName := r.URL.Query().Get("hostName")
	//machineId := r.URL.Query().Get("machineId")
	//productCode := r.URL.Query().Get("productCode")
	//productFamilyId := r.URL.Query().Get("productFamilyId")
	salt := r.URL.Query().Get("salt")
	//secure := r.URL.Query().Get("secure")
	username := r.URL.Query().Get("userName")
	//version := r.URL.Query().Get("version")
	//versionNumber := r.URL.Query().Get("versionNumber")

	if salt == "" || username == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	prolongationPeriod := 607875500

	xmlResponse := "<ObtainTicketResponse><message></message><prolongationPeriod>" + strconv.Itoa(prolongationPeriod) + "</prolongationPeriod><responseCode>OK</responseCode><salt>" + salt + "</salt><ticketId>1</ticketId><ticketProperties>licensee=" + username + "\tlicenseType=0\t</ticketProperties></ObtainTicketResponse>"
	xmlSignature, _ := signature(xmlResponse)
	w.Header().Add("Content-Type", "text/xml")
	w.Write([]byte("<!-- " + xmlSignature + " -->\n" + xmlResponse))
}

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBALecq3BwAI4YJZwhJ+snnDFj3lF3DMqNPorV6y5ZKXCiCMqj8OeOmxk4YZW9aaV9
ckl/zlAOI0mpB3pDT+Xlj2sCAwEAAQJAW6/aVD05qbsZHMvZuS2Aa5FpNNj0BDlf38hOtkhDzz/h
kYb+EBYLLvldhgsD0OvRNy8yhz7EjaUqLCB0juIN4QIhAOeCQp+NXxfBmfdG/S+XbRUAdv8iHBl+
F6O2wr5fA2jzAiEAywlDfGIl6acnakPrmJE0IL8qvuO3FtsHBrpkUuOnXakCIQCqdr+XvADI/UTh
TuQepuErFayJMBSAsNe3NFsw0cUxAQIgGA5n7ZPfdBi3BdM4VeJWb87WrLlkVxPqeDSbcGrCyMkC
IFSs5JyXvFTreWt7IQjDssrKDRIPmALdNjvfETwlNJyY
-----END RSA PRIVATE KEY-----
`)

func signature(message string) (string, error) {
	pem, _ := pem.Decode(privateKey)
	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(pem.Bytes)

	hashedMessage := md5.Sum([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.MD5, hashedMessage[:])
	if err != nil {
		return "", err
	}
	hexSignature := hex.EncodeToString(signature)
	return hexSignature, nil
}
