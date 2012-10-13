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

var opt Options

func init() {
	opt = parseFlags()
}

func main() {
	response, postError := http.PostForm(opt.ServerAddr,
		url.Values{"input": {"hello"}, "lang": {opt.Lang}})
	if postError != nil {
		log.Fatalf("Failed to open %s: %s\n", opt.ServerAddr, err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(body))
}
