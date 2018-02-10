package main

import (
	"bytes"
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
	fmt.Printf("Received userCommand at %v\n", timestamp)

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

	eventlog.Insert(v)
	w.Write(v.Byte())
}

func quoteServerHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()
	fmt.Printf("Received quoteServer at %v\n", timestamp)

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

	eventlog.Insert(v)
	w.Write(v.Byte())
}

func accountTransactionHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()
	fmt.Printf("Received accountTransaction at %v\n", timestamp)

	v := &commands.AccountTransaction{
		Timestamp:      timestamp,
		Server:         query.Get("server"),
		TransactionNum: query.Get("transactionNum"),
		Action:         query.Get("username"),
		Username:       query.Get("stockSymbol"),
		Funds:          query.Get("price"),
	}

	eventlog.Insert(v)
	fmt.Print(eventlog)
	w.Write(v.Byte())
}

func systemEventHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()
	fmt.Printf("Received systemEvent at %v\n", timestamp)

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

	eventlog.Insert(v)
	w.Write(v.Byte())
}

func errorEventHandler(w http.ResponseWriter, r *http.Request) {
	timestamp := makeTimestamp()
	query := r.URL.Query()
	fmt.Printf("Received errorEvent at %v\n", timestamp)

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

	eventlog.Insert(v)
	w.Write(v.Byte())
}

func dumpLogHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dumpfile := query.Get("filename")
	userLog := query.Get("username")
	dumpfileB := string(bytes.Trim([]byte(dumpfile), "\x00"))
	if dumpfileB != "./test.log" {
		panic(fmt.Sprintf("Names not equal %q ./test.log\n len=%v", dumpfile, len(dumpfile)))
	}

	file, err := os.Create(string(dumpfileB))
	if err != nil {
		fmt.Printf("error: %v %v\n", err, file)
	}

	fmt.Printf("Dumping log to %v, with user set as %v",
		dumpfile, userLog)
	eventlog.Write(file)
	file.Close()
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

var eventlog = log.Log{}

func main() {
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
