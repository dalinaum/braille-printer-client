// Copyright 2012, Braille Printer Team. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    main_test.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-10-13 12:44:45.829081 +0900 KST
 *  Description: Main test file for braille-printer-client
 */

import (
	"log"
	"testing"
)

func TestBrailleprinterclient(t *testing.T) {

}

func TestSerialCommunication(t *testing.T) {
	log.Println("Enter TestSerialCommunication")
	DoPrint("/dev/ttyUSB0", "Hello world")
	log.Println("exit TestSerialCommunication")
}
