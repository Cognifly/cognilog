//Package cognilog Copyright 2016 Cognifly and Contributors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package cognilog

import (
	"fmt"
	"log"
	"strings"
)

var (
	green   = string([]byte{27, 91, 57, 50, 59, 52, 48, 109})
	white   = string([]byte{27, 91, 57, 55, 59, 52, 48, 109})
	yellow  = string([]byte{27, 91, 57, 51, 59, 52, 48, 109})
	red     = string([]byte{27, 91, 57, 49, 59, 52, 48, 109})
	blue    = string([]byte{27, 91, 57, 52, 59, 52, 48, 109})
	magenta = string([]byte{27, 91, 57, 53, 59, 52, 48, 109})
	cyan    = string([]byte{27, 91, 57, 54, 59, 52, 48, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

// return the byte colour of c.
func colour(c string) string {
	switch c {
	case "green":
		return green
	case "white":
		return white
	case "yellow":
		return yellow
	case "red":
		return red
	case "blue":
		return blue
	case "magenta":
		return magenta
	case "cyan":
		return cyan
	default:
		return reset
	}
}

// Log glues Prefix E.G 2016/07/22 13:37:08 to to msg before printing to stdout.
func Log(c string, msg interface{}) {
	fmt.Print(colour(c))
	log.Printf("%v %v\n", msg, reset)
}

// LogINFO glues Prefix E.G 2016/07/22 13:37:08 and [infor] to msg before printing
// to stdout.
func LogINFO(c string, i, msg interface{}) {
	fmt.Print(colour(c))
	info := fmt.Sprintf("%v", i)
	log.Printf("[%s] %v %v\n", strings.ToUpper(info), msg, reset)
}

// Fatal glues Prefix E.G 2016/07/22 13:37:08 to to msg before printing to stdout.
func Fatal(c string, msg interface{}) {
	fmt.Print(colour(c))
	log.Fatalf("%v %v\n", msg, reset)
}

// FatalINFO glues Prefix E.G 2016/07/22 13:37:08 and [infor] to msg before printing
// to stdout.
func FatalINFO(c string, i, msg interface{}) {
	fmt.Print(colour(c))
	info := fmt.Sprintf("%v", i)
	log.Fatalf("[%s] %v %v\n", strings.ToUpper(info), msg, reset)
}

// Prtln prints the msg stdout
func Prtln(c string, msg interface{}) {
	fmt.Print(colour(c))
	fmt.Printf("%v %v\n", msg, reset)
}

// PrtlnINFO glues Prefix [infor] to msg before printing to stdout.
func PrtlnINFO(c string, i, msg interface{}) {
	fmt.Print(colour(c))
	info := fmt.Sprintf("%v", i)
	fmt.Printf("[%s] %v %v\n", strings.ToUpper(info), msg, reset)
}
