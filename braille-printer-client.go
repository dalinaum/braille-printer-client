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
	"fmt"
	"log"
	"strconv"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BRAILLE_PATH    = "/braille"
	PRINTQ_ADD_PATH = "/printq/add"
)

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

func braille() {
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

func printqAdd() {
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
	fmt.Printf("OK: printq-add\n");
}

func main() {
	if len(arguments) == 0 {
		return
	}

	command := arguments[0]
	switch command {
	case "braille":
		braille()
	case "printq-add":
		printqAdd()
	default:
		return
	}
}
