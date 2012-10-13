// Copyright 2012, Braille Printer Team. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    braille-printer-client.go
 *  Author:      Leonardo YongUk Kim <dalinaum@gmail.com>, Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-10-13 12:44:45.828012 +0900 KST
 *  Description: Main source file in braille-printer-client
 */

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var options Options
var arguments []string 

func init() {
	options, arguments = parseFlags()
}

func braille() {
	var input string
	if len(arguments) > 1 {
		input = arguments[1]
	} else {
		input = "hello world"
	}
	response, postError := http.PostForm(options.ServerAddr,
		url.Values{"input": {input}, "lang": {options.Lang}})
	if postError != nil {
		log.Fatalf("Failed to open %s: %s\n", options.ServerAddr, postError)
	}
	defer response.Body.Close()

	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		log.Fatalf("Failed to read %s\n", readError)
	}
	fmt.Printf(string(body))
}

func main() {
	if len(arguments) == 0 {
		return
	}

	switch command := arguments[0]; command {
		case "braille":
			braille()
		default:
			return
	}
}
