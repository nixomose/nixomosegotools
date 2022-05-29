// SPDX-License-Identifier: LGPL-2.1
// Copyright (C) 2021-2022 stu mark

package main

import (
	"container/list"
	"fmt"
	"github.com/nixomose/nixomosegotools/tools"
)

func main() {

	var log *tools.Nixomosetools_logger = tools.New_Nixomosetools_logger(tools.DEBUG)

	log.Debug("this is a test log")

	var l *list.List = list.New()
	l.PushBack("item 1")
	l.PushBack(2)

	fmt.Println("item as pos 0: ", tools.Getlistitematpos(l, 0))
	fmt.Println("item as pos 1: ", tools.Getlistitematpos(l, 1))

	fmt.Println("max of 4 and 5: ", tools.Maxint(4, 5))

	tools.ErrorWithCode(log, 749, "this is error 749")

}
