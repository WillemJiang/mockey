//go:build !go1.21
// +build !go1.21

/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package stw

import (
	_ "unsafe"
)

func newSTWCtx() ctx {
	return &stwCtx{}
}

type stwCtx struct{}

func (ctx *stwCtx) StopTheWorld() {
	stopTheWorld("mockey")
}

func (ctx *stwCtx) StartTheWorld() {
	startTheWorld()
}

//go:linkname stopTheWorld runtime.stopTheWorld
func stopTheWorld(reason string)

//go:linkname startTheWorld runtime.startTheWorld
func startTheWorld()
