// SPDX-License-Identifier: LGPL-2.1
// Copyright (C) 2021-2022 stu mark

package tools

// Package hexdump provides utility functions to display binary slices as
// hex and printable ASCII.
// made a copy so I could remove the color because it's not configurable.

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Dump the byte slice to a human-readable hex dump using the default
// configuration.
func Dump(buf []byte) string { return defaultConfig.Dump(buf) }

// Config allows customizing the dump configuration.
type Config struct {
	// Number of bytes from the input buffer to print in a single row. The default
	// is 32.
	Width int
}

type dumpState struct {
	Config
	rowIndex    int
	maxRowWidth int
}

func (this *dumpState) dump(out io.Writer, buf []byte) {
	N := this.Width
	for i := 0; i*N < len(buf); i++ {
		a, b := i*N, (i+1)*N
		if b > len(buf) {
			b = len(buf)
		}
		row := buf[a:b]
		hex, ascii := printable(row)

		if len(row) < this.maxRowWidth {
			padding := this.maxRowWidth*2 + this.maxRowWidth/4 - len(row)*2 - len(row)/4
			hex += strings.Repeat(" ", padding)
		}
		this.maxRowWidth = len(row)

		fmt.Fprintf(out, "%5d: %s | %s\n", this.rowIndex*N, hex, ascii)
		this.rowIndex++
	}
}

func (this Config) newDumpState() *dumpState {
	s := &dumpState{Config: this}
	if s.Width == 0 {
		s.Width = kDefaultWidth
	}
	return s
}

// Dump converts the byte slice to a human-readable hex dump.
func (this Config) Dump(data []byte) string {
	var out bytes.Buffer
	this.newDumpState().dump(&out, data)
	return out.String()

}

// Read will read from the input io.Reader and write human-readable, formatted
// hexdumps (with color annotations) to the output. The entire input reader is
// consumed. Any errors other than io.EOF are returned.
func (this Config) Stream(in io.Reader, out io.Writer) error {
	s := this.newDumpState()
	buf := make([]byte, 1*this.Width)
	for {
		n, err := io.ReadFull(in, buf)
		s.dump(out, buf[:n])
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

const kDefaultWidth = 32

var defaultConfig = Config{kDefaultWidth}

func printable(data []byte) (hex, ascii string) {
	s := string(data)
	for i := 0; i < len(s); i++ {

		if s[i] < 32 || s[i] >= 127 {
			ascii += "░"
			hex += fmt.Sprintf("%02x", s[i])
		} else {
			ascii += string(s[i])
			hex += fmt.Sprintf("%02x", s[i])
		}
		if i%4 == 3 {
			hex += " "
		}
	}
	return hex, ascii
}
