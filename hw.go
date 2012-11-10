// Copyright 2012, Braille Printer Team. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	// "fmt"
	// "github.com/suapapa/go-serial/serial"
	"github.com/huin/goserial"
	"log"
)

/*  Filename:    braille-printer-client.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Description: Code for control braille printer hw
 */

func DoPrint(port string, braille string) {
	log.Printf("Opening %s...\n", port)
        c := &goserial.Config{Name: port, Baud: 9600}
        s, err := goserial.OpenPort(c)
        if err != nil {
		log.Fatalf("Failed to open %s: %s\n", port, err)
        }

        // n, err := s.Write([]byte("test"))
        // if err != nil {
        //         log.Fatal(err)
        // }

	log.Printf("Reading %s...\n", port)
        buf := make([]byte, 256)
        n, err := s.Read(buf)
        if err != nil {
		log.Fatalf("Failed to read from serial: %s, %d", err, n)
        }
	log.Printf("Read %d: %s", n, string(buf[:n]))
}
