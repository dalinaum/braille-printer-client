// Copyright 2012, Braille Printer Team. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	// "fmt"
	// "github.com/suapapa/go-serial/serial"
	// "github.com/huin/goserial"
	// "log"
)

/*  Filename:    braille-printer-client.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Description: Code for control braille printer hw
 */

func DoPrint(port string, braille string) {
	// spo := serial.OpenOptions{
	// 	PortName: port,
	// 	BaudRate: 9600,
	// 	DataBits: 8,
	// 	StopBits: 1,
	// 	InterCharacterTimeout: 10,
	// 	MinimumReadSize: 4, // TODO: need this?
	// }

	// sp, err := serial.Open(spo)
	// if err != nil {
	// 	log.Fatalf("Fail to open serial port: %v\n", err)
	// }
	// defer sp.Close()

	// for _, bc := range braille {
	// 	fmt.Println(bc)
	// 	sp.Write([]byte{byte(bc & 0xFF),})
	// }

}
