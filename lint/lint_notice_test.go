/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"io/ioutil"
	"testing"
)

func TestNoticeLinter_lint(t *testing.T) {
	l := NoticeLinter{
		File: "./__notice__/NOTICE",
	}
	err := l.lint()
	if err == nil {
		t.Fatalf("lint error")
	}
}

func TestNoticeLinter_fixed(t *testing.T) {
	// write test data
	ioutil.WriteFile("./__notice__/NOTICE_1", []byte("Copyright 2018-2021 The Apache Software Foundation"), 0644)

	// init notice linter
	l := NoticeLinter{
		File: "./__notice__/NOTICE_1",
	}

	// fixed notice
	err := l.fix()
	if err != nil {
		t.Fatalf("fixed error")
	}

	// check notice
	err = l.lint()
	if err != nil {
		t.Fatalf("lint error")
	}
}
