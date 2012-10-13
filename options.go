// Copyright 2012, Braille Printer Team. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*  Filename:    options.go
 *  Author:      Leonardo YongUkKim <dalinaum@gmail.com>, 
 *               Homin Lee <homin.lee@suapapa.net>
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
	Lang       string
	Format     string
	Type       string
}

//  Create a flag.FlagSet to parse the braille-printer-client's flags.
func SetupFlags(opt *Options) *flag.FlagSet {
	fs := flag.NewFlagSet("braille-printer-client", flag.ExitOnError)
	fs.BoolVar(&opt.Verbose, "v", false, "Verbose program output.")
	fs.StringVar(&opt.ServerAddr, "a", "http://braille-printer.appspot.com",
		"Address of braille-printer (server)")
	fs.StringVar(&opt.Lang, "l", "ko", "Braille language {ko|en}")
	fs.StringVar(&opt.Format, "f", "svg", "Format to print out {svg|text}")
	fs.StringVar(&opt.Type, "t", "all",
		"Type to display out {all|label|paper}")
	return setupUsage(fs)
}

// Check the braille-printer-client's flags and arguments for acceptable values.
// When an error is encountered, panic, exit with a non-zero status, or override
// the error.
func VerifyFlags(opt *Options, fs *flag.FlagSet) {
	switch opt.Lang {
	case "ko", "en":
	default:
		log.Fatalf("Unknown lang, %s! Use one of ko, en.", opt.Lang)
	}

	switch opt.Format {
	case "text", "svg":
	default:
		log.Fatalf("Unkonw format, %s! Use one of text or svg.",
			opt.Format)
	}

	switch opt.Type {
	case "label", "paper", "all":
	default:
		log.Fatalf("Unkonw type, %s! Use one of label, paper or all.",
			opt.Type)
	}
}

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
func parseFlags() (Options, []string) {
	var opt Options
	fs := SetupFlags(&opt)
	fs.Parse(os.Args[1:])
	VerifyFlags(&opt, fs)
	// Process the verified Options...
	return opt, fs.Args()
}
