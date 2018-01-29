package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func userCommandHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()

	type userCommand struct {
		XMLName        xml.Name `xml:"userCommand"`
		Timestamp      int64    `xml:"timestamp"`
		Server         string   `xml:"server"`
		TransactionNum string   `xml:"transactionNum"`
		Command        string   `xml:"command"`
		Username       string   `xml:"username,omitempty"`
		StockSymbol    string   `xml:"stockSymbol,omitempty"`
		Filename       string   `xml:"filename,omitempty"`
		Funds          string   `xml:"funds,omitempty"`
	}

	v := &userCommand{
		Timestamp:      timestamp,
		Server:         query.Get("server"),
		TransactionNum: query.Get("transactionNum"),
		Command:        query.Get("command"),
		Username:       query.Get("username"),
		StockSymbol:    query.Get("stockSymbol"),
		Filename:       query.Get("filename"),
		Funds:          query.Get("funds"),
	}

	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func quoteServerHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()

	type quoteServer struct {
		XMLName         xml.Name `xml:"quoteServer"`
		Timestamp       int64    `xml:"timestamp"`
		Server          string   `xml:"server"`
		TransactionNum  string   `xml:"transactionNum"`
		Price           string   `xml:"price"`
		StockSymbol     string   `xml:"stockSymbol"`
		Username        string   `xml:"username"`
		QuoteServerTime string   `xml:"quoteServerTime"`
		CryptoKey       string   `xml:"cryptoKey"`
	}

	v := &quoteServer{
		Timestamp:       timestamp,
		Server:          query.Get("server"),
		TransactionNum:  query.Get("transactionNum"),
		Username:        query.Get("username"),
		StockSymbol:     query.Get("stockSymbol"),
		Price:           query.Get("price"),
		QuoteServerTime: query.Get("quoteServerTime"),
		CryptoKey:       query.Get("cryptoKey"),
	}

	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func accountTransactionHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()

	type accountTransaction struct {
		XMLName        xml.Name `xml:"accountTransaction"`
		Timestamp      int64    `xml:"timestamp"`
		Server         string   `xml:"server"`
		TransactionNum string   `xml:"transactionNum"`
		Action         string   `xml:"action"`
		Username       string   `xml:"username,omitempty"`
		Funds          string   `xml:"funds,omitempty"`
	}

	v := &accountTransaction{
		Timestamp:      timestamp,
		Server:         query.Get("server"),
		TransactionNum: query.Get("transactionNum"),
		Action:         query.Get("username"),
		Username:       query.Get("stockSymbol"),
		Funds:          query.Get("price"),
	}

	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func systemEventHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()

	type systemEvent struct {
		XMLName        xml.Name `xml:"systemEvent"`
		Timestamp      int64    `xml:"timestamp"`
		Server         string   `xml:"server"`
		TransactionNum string   `xml:"transactionNum"`
		Command        string   `xml:"command"`
		Username       string   `xml:"username,omitempty"`
		StockSymbol    string   `xml:"stockSymbol,omitempty"`
		Filename       string   `xml:"filename,omitempty"`
		Funds          string   `xml:"funds,omitempty"`
	}

	v := &systemEvent{
		Timestamp:      timestamp,
		Server:         query.Get("server"),
		TransactionNum: query.Get("transactionNum"),
		Command:        query.Get("command"),
		Username:       query.Get("username"),
		StockSymbol:    query.Get("stockSymbol"),
		Filename:       query.Get("filename"),
		Funds:          query.Get("funds"),
	}

	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func errorEventHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()

	type errorEvent struct {
		XMLName        xml.Name `xml:"errorEvent"`
		Timestamp      int64    `xml:"timestamp"`
		Server         string   `xml:"server"`
		TransactionNum string   `xml:"transactionNum"`
		Command        string   `xml:"command"`
		Username       string   `xml:"username,omitempty"`
		StockSymbol    string   `xml:"stockSymbol,omitempty"`
		Filename       string   `xml:"filename,omitempty"`
		Funds          string   `xml:"funds,omitempty"`
		ErrorMessage   string   `xml:"errorMessage,omitempty"`
	}

	v := &errorEvent{
		Timestamp:      timestamp,
		Server:         query.Get("server"),
		TransactionNum: query.Get("transactionNum"),
		Command:        query.Get("command"),
		Username:       query.Get("username"),
		StockSymbol:    query.Get("stockSymbol"),
		Filename:       query.Get("filename"),
		Funds:          query.Get("funds"),
		ErrorMessage:   query.Get("errorMessage"),
	}

	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func writeEncodedStruct(v, w http.ResponseWriter) {
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func main() {
	http.HandleFunc("/userCommand", userCommandHandler)
	http.HandleFunc("/quoteServer", quoteServerHandler)
	http.HandleFunc("/accountTransaction", accountTransactionHandler)
	http.HandleFunc("/systemEvent", systemEventHandler)
	http.HandleFunc("/errorEvent", errorEventHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
