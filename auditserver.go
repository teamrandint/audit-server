package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"seng468/auditserver/commands"
	"seng468/auditserver/log"
	"time"
)

func userCommandHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()

	v := &commands.UserCommand{
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

	v := &commands.QuoteServer{
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

	v := &commands.AccountTransaction{
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

	v := &commands.SystemEvent{
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

	v := &commands.ErrorEvent{
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

func dumpLogHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dumpfile := query.Get("filename")
	userLog := query.Get("username")

	file, err := os.Create(dumpfile)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("Dumping log to %v, with user set as %v",
		dumpfile, userLog)
	log.Write(file)
	file.Close()
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
	log := log.Log{}
	fmt.Print(log)

	http.HandleFunc("/userCommand", userCommandHandler)
	http.HandleFunc("/quoteServer", quoteServerHandler)
	http.HandleFunc("/accountTransaction", accountTransactionHandler)
	http.HandleFunc("/systemEvent", systemEventHandler)
	http.HandleFunc("/errorEvent", errorEventHandler)
	http.HandleFunc("/dumpLog", dumpLogHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
