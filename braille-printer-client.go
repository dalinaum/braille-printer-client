// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    braille-printer-client.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
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
	resp2, err := http.PostForm(opt.ServerAddr,
		url.Values{"input": {"hello"}, "lang": {opt.Lang}})
	if err != nil {
		log.Fatalf("Failed to open %s: %s\n", opt.ServerAddr, err)
	}

	defer resp2.Body.Close()
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Printf(string(body2))
}
