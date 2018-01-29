package commands

import (
	"encoding/xml"
	"fmt"
)

// Command contains types of commands user can run
type Command interface {
	String() string
	Byte() []byte
}

type UserCommand struct {
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

type QuoteServer struct {
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

type AccountTransaction struct {
	XMLName        xml.Name `xml:"accountTransaction"`
	Timestamp      int64    `xml:"timestamp"`
	Server         string   `xml:"server"`
	TransactionNum string   `xml:"transactionNum"`
	Action         string   `xml:"action"`
	Username       string   `xml:"username,omitempty"`
	Funds          string   `xml:"funds,omitempty"`
}

type SystemEvent struct {
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

type ErrorEvent struct {
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

// String returns a string representation of userCommand
func (u *UserCommand) String() string {
	return string(u.Byte())
}

func (u *UserCommand) Byte() []byte {
	output, err := xml.MarshalIndent(u, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return output
}
