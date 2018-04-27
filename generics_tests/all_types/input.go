// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import "./lib"

type T int

type newType struct {
	a T
	b lib.T
	c *T
	d (T)
	e chan T
	f <-chan T
	g chan<- T
	h []T
	i [10]T
	j map[T]T
	k func(T, T) (T, T)
	l interface {
		f(T)
	}
	m struct {
		T
		a T
	}
}

func f(...T) {
}
