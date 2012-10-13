// Copyright 2012, Homin Lee. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    options.go
 *  Author:      Homin Lee <homin.lee@suapapa.net>
 *  Created:     2012-10-13 12:44:45.83294 +0900 KST
 *  Description: Option parsing for braille-printer-client
 */

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// TODO Customize exported (capitalized) variables, types, and functions.

var (
	CmdHelpUsage string // Custom usage string.
	CmdHelpFoot  string // Printed after help.
)

// A struct that holds braille-printer-client's parsed command line flags.
type Options struct {
	Verbose    bool
	ServerAddr string
	Lang string
}

//  Create a flag.FlagSet to parse the braille-printer-client's flags.
func SetupFlags(opt *Options) *flag.FlagSet {
	fs := flag.NewFlagSet("braille-printer-client", flag.ExitOnError)
	fs.BoolVar(&opt.Verbose, "v", false, "Verbose program output.")
	fs.StringVar(&opt.ServerAddr, "a", "http://localhost:8080/braille",
		"Address of braille-printer (server)")
	fs.StringVar(&opt.Lang, "l", "ko", "Braille language {ko|en}")
	return setupUsage(fs)
}

// Check the braille-printer-client's flags and arguments for acceptable values.
// When an error is encountered, panic, exit with a non-zero status, or override
// the error.
func VerifyFlags(opt *Options, fs *flag.FlagSet) {
	switch opt.Lang {
	case "ko":
	case "en":
	default:
		log.Fatalf("Unknown lang, %s! Use one of ko or en.", opt.Lang)

	}
}

/**************************/
/* Do not edit below here */
/**************************/

//  Print a help message to standard error. See CmdHelpUsage and CmdHelpFoot.
func PrintHelp() { SetupFlags(&Options{}).Usage() }

//  Hook up CmdHelpUsage and CmdHelpFoot with flag defaults to function flag.Usage.
func setupUsage(fs *flag.FlagSet) *flag.FlagSet {
	printNonEmpty := func(s string) {
		if s != "" {
			fmt.Fprintf(os.Stderr, "%s\n", s)
		}
	}
	fs.Usage = func() {
		printNonEmpty(CmdHelpUsage)
		fs.PrintDefaults()
		printNonEmpty(CmdHelpFoot)
	}
	return fs
}

//  Parse the flags, validate them, and post-process (e.g. Initialize more complex structs).
func parseFlags() Options {
	var opt Options
	fs := SetupFlags(&opt)
	fs.Parse(os.Args[1:])
	VerifyFlags(&opt, fs)
	// Process the verified Options...
	return opt
}