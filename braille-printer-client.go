// Copyright 2012, Braille Printer Team. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    braille-printer-client.go
 *  Author:      Leonardo YongUk Kim <dalinaum@gmail.com>, 
 *               Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-10-13 12:44:45.828012 +0900 KST
 *  Description: Main source file in braille-printer-client
 */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	BRAILLE_PATH    = "/braille"
	PRINTQ_ADD_PATH = "/printq/add"
	PRINTQ_LIST     = "/printq/list"
	PRINTQ_ITEM     = "/printq/item"
	PRINTQ_UPDATE   = "/printq/update"
)

type PrintqItem struct {
	Qid  int
	Type string
}

type Item struct {
	Origin string
	Result string
}

var options Options
var arguments []string

func init() {
	options, arguments = parseFlags()
}

func parseStatusCode(statusString string) (int, error) {
	tokenizedString := strings.Split(statusString, " ")
	statusCode, conversionError := strconv.Atoi(tokenizedString[0])
	if conversionError != nil {
		return 0, fmt.Errorf("Failed to conversion %s", statusString)
	}
	return statusCode, nil
}

func handleBraille() {
	var input string
	if len(arguments) > 1 {
		input = arguments[1]
	} else {
		input = "hello world"
	}

	requestUri := options.ServerAddr + BRAILLE_PATH
	response, postError := http.PostForm(requestUri,
		url.Values{"input": {input}, "lang": {options.Lang},
			"format": {options.Format}})
	if postError != nil {
		log.Fatalf("Failed to open %s: %s\n", requestUri, postError)
	}
	defer response.Body.Close()

	statusCode, conversionError := parseStatusCode(response.Status)
	if conversionError != nil {
		log.Fatalf("Failed to conversion: %s\n", conversionError)
	}
	if statusCode != 200 {
		log.Fatalf("Status code is not 200\n")
	}

	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		log.Fatalf("Failed to read %s\n", readError)
	}
	fmt.Printf("%s\n", string(body))
}

func handlePrintqAdd() {
	var input string
	if len(arguments) > 1 {
		input = arguments[1]
	} else {
		input = "hello world"
	}

	requestUri := options.ServerAddr + PRINTQ_ADD_PATH
	response, postError := http.PostForm(requestUri,
		url.Values{"input": {input}, "lang": {options.Lang}})
	if postError != nil {
		log.Fatalf("Failed to open %s: %s\n", requestUri, postError)
	}
	defer response.Body.Close()

	statusCode, conversionError := parseStatusCode(response.Status)
	if conversionError != nil {
		log.Fatalf("Failed to conversion: %s\n", conversionError)
	}
	if statusCode != 200 {
		log.Fatalf("Status code is not 200\n")
	}
	fmt.Printf("OK: printq-add\n")
}

func handlePrintqList() {
	requestUri := options.ServerAddr + PRINTQ_LIST + "?type=" + options.Type
	response, getError := http.Get(requestUri)
	if getError != nil {
		log.Fatalf("Failed to open %s: %s\n", requestUri, getError)
	}
	defer response.Body.Close()

	statusCode, conversionError := parseStatusCode(response.Status)
	if conversionError != nil {
		log.Fatalf("Failed to conversion: %s\n", conversionError)
	}
	if statusCode != 200 {
		log.Fatalf("Status code is not 200\n")
	}

	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		log.Fatalf("Failed to read %s\n", readError)
	}

	var printqList []PrintqItem
	unmarshalError := json.Unmarshal(body, &printqList)
	if unmarshalError != nil {
		log.Fatalf("Failed to unmarshalError %s\n", unmarshalError)
	}

	for k, v := range printqList {
		fmt.Printf("[%2d] qid: %2d type: %s\n", k+1, v.Qid, v.Type)
	}
}

func handlePrintqItem() {
	var qid string
	if len(arguments) > 1 {
		qid = arguments[1]
	} else {
		log.Fatalf("qid is necessary to print out.");
	}
	
	requestUri := options.ServerAddr + PRINTQ_ITEM + "?qid=" + qid +
		"&format=" + options.Format
	response, getError := http.Get(requestUri)
	if getError != nil {
		log.Fatalf("Failed to open %s: %s\n", requestUri, getError)
	}
	defer response.Body.Close()

	statusCode, conversionError := parseStatusCode(response.Status)
	if conversionError != nil {
		log.Fatalf("Failed to conversion: %s\n", conversionError)
	}
	if statusCode != 200 {
		log.Fatalf("Status code is not 200\n")
	}

	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		log.Fatalf("Failed to read %s\n", readError)
	}

	var item Item
	unmarshalError := json.Unmarshal(body, &item)
	if unmarshalError != nil {
		log.Fatalf("Failed to unmarshalError %s\n", unmarshalError)
	}

	fmt.Printf("[qid: %d]\n -from: %s\n -to: %s\n", qid, item.Origin,
		item.Result)
}

func handlePrintqUpdate() {
	var qid string
	if len(arguments) > 1 {
		qid = arguments[1]
	} else {
		log.Fatalf("qid is necessary to print out.");
	}

	requestUri := options.ServerAddr + PRINTQ_UPDATE + "?qid=" + qid +
		"&status=" + strconv.Itoa(options.Status)
	response, postError := http.PostForm(requestUri, url.Values {})
	if postError != nil {
		log.Fatalf("Failed to open %s: %s\n", requestUri, postError)
	}
	defer response.Body.Close()

	statusCode, conversionError := parseStatusCode(response.Status)
	if conversionError != nil {
		log.Fatalf("Failed to conversion: %s\n", conversionError)
	}
	if statusCode != 200 {
		log.Fatalf("Status code is not 200\n")
	}
	fmt.Printf("OK: printq-update\n")
}

func main() {
	if len(arguments) == 0 {
		return
	}

	command := arguments[0]
	switch command {
	case "braille":
		handleBraille()
	case "printq-add":
		handlePrintqAdd()
	case "printq-list":
		handlePrintqList()
	case "printq-item":
		handlePrintqItem()
	case "printq-update":
		handlePrintqUpdate()
	default:
		fmt.Printf("...\n")
	}
}
