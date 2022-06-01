// SPDX-License-Identifier: LGPL-2.1
// Copyright (C) 2021-2022 stu mark

package tools

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	DEBUG = 0
	INFO  = 200
	ERROR = 500
)

type Nixomosetools_logger struct {
	level int
}

func New_Nixomosetools_logger(level int) *Nixomosetools_logger {
	var logger Nixomosetools_logger
	logger.Set_level(level)
	return &logger
}
func (this *Nixomosetools_logger) Set_level(level int) {
	this.level = level
}

func (this *Nixomosetools_logger) json(msgtext string) string {
	var msg map[string]string = make(map[string]string)
	msg["msg"] = msgtext
	bytesout, err := json.Marshal(msg)
	if err != nil {
		return "{\"message\":\"unable to marshal log message into json, msg: " + msgtext + "\"}"
	}
	return string(bytesout)
}

func (this *Nixomosetools_logger) Debug(msg ...interface{}) {
	var fullmsg string = fmt.Sprint(msg...)
	if this == nil {
		// fmt.Fprintln(os.Stderr, this.json(fullmsg))
		log.Println(this.json(fullmsg))
		return
	}
	// when we have a real file to log to...
	if this.level <= DEBUG {
		// fmt.Fprintln(os.Stderr, this.json(fullmsg))
		log.Println(this.json(fullmsg))
	}
}

func (this *Nixomosetools_logger) Info(msg ...interface{}) {
	var fullmsg string = fmt.Sprint(msg...)
	if this == nil {
		log.Println(this.json(fullmsg))
		// fmt.Fprintln(os.Stderr, this.json(fullmsg))
		return
	}
	if this.level <= INFO {
		log.Println(this.json(fullmsg))
		// fmt.Fprintln(os.Stderr, this.json(fullmsg))
	}
}

func (this *Nixomosetools_logger) Error(msg ...interface{}) {
	// xxxz must add mutex here and around other log lines
	var fullmsg string = fmt.Sprint(msg...)
	if this == nil {
		log.Println(this.json(fullmsg))
		// fmt.Fprintln(os.Stderr, this.json(fullmsg))
		return
	}
	if this.level <= ERROR {
		log.Println(this.json(fullmsg))
		// fmt.Fprintln(os.Stderr, this.json(fullmsg))
	}
}
