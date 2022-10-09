// SPDX-License-Identifier: LGPL-2.1
// Copyright (C) 2021-2022 stu mark

// Package tools has a package comment to satisfy the static checker, complete with capital P and everything.
package tools

import (
	"fmt"
)

// 1/19/22 replacing SUCCESS with nil. const SUCCESS_CODE int = 0

const DEFAULT_ERROR_CODE int = 1

type Ret interface {
	Error() error
	Get_errcode() int
	Get_errmsg() string
}

type RetInstance struct {
	logger    *Nixomosetools_logger
	errmsg    string // string to be logged
	logged    bool
	errorcode int
}

// verify that retinstance implements error
var _ error = &RetInstance{}
var _ error = (*RetInstance)(nil)

//  1/19/22 success has been replaced with nil
// func Success() Ret {
// 	var ret Ret
// 	ret.errmsg = ""
// 	ret.logged = true // we never log as we have no logger
// 	ret.errorcode = SUCCESS_CODE
// 	return ret
// }

func R(logger_in *Nixomosetools_logger, errmsg ...interface{}) RetInstance {
	var ret RetInstance = R_no_log(logger_in, errmsg...)
	ret.log()
	return ret
}

func R_no_log(logger_in *Nixomosetools_logger, errmsg ...interface{}) RetInstance {
	var ret RetInstance
	ret.logger = logger_in
	ret.errmsg = fmt.Sprint(errmsg...)
	ret.logged = false
	ret.errorcode = DEFAULT_ERROR_CODE
	return ret
}

func Error(logger_in *Nixomosetools_logger, errmsg ...interface{}) RetInstance {
	return R(logger_in, errmsg...)
}

func ErrorWithCodeNoLog(logger_in *Nixomosetools_logger, code int, errmsg ...interface{}) RetInstance {
	var ret = R_no_log(logger_in, errmsg...)
	ret.logged = true
	ret.errorcode = code
	return ret
}

func ErrorWithCode(logger_in *Nixomosetools_logger, code int, errmsg ...interface{}) RetInstance {
	var ret = R(logger_in, errmsg...)
	ret.errorcode = code
	return ret
}

func (this *RetInstance) log() {
	if this.logged {
		return
	}
	em := this.errmsg
	// if they specified an error code, log it too.
	if this.errorcode != DEFAULT_ERROR_CODE {
		em = fmt.Sprint(this.errmsg, " Errorcode: ", this.errorcode)
	}
	this.logger.Error(em)
	// linter says this is pointless, not sure why. this.logged = true
}

func (this *RetInstance) Get_errcode() int {
	return this.errorcode
}

func (this *RetInstance) Get_errmsg() string {
	return this.errmsg
}

func (this *RetInstance) Set_errcode(e int) {
	this.errorcode = e
}

func (this *RetInstance) Set_errmsg(m string) {
	this.errmsg = m
}

// to be an error as well
func (this *RetInstance) Error() string {
	return this.Get_errmsg()
}
