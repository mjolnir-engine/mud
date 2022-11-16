/*
 * Copyright (c) 2022 eightfivefour llc. All rights reserved.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
 * documentation files (the "Software"), to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
 * WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 * OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package mud

import (
	"github.com/mjolnir-engine/engine"
	engineTesting "github.com/mjolnir-engine/engine/testing"
	"testing"
)

type testCon struct {
	SentDataChannel chan string
}

func (c testCon) Name() string {
	return "test"
}

func (c testCon) Start(context *engine.ControllerContext) error {
	return context.Engine.SendToSession(context.SessionId, []byte("Welcome to the test server!"))
}

func (c testCon) Stop(context *engine.ControllerContext) error {
	return nil
}

func (c testCon) Resume(context *engine.ControllerContext) error {
	return nil
}

func (c testCon) HandleInput(context *engine.ControllerContext, input string) error {
	go func() {
		c.SentDataChannel <- input
	}()

	return nil
}

func (c testCon) GetSentDataChannel() chan string {
	return c.SentDataChannel
}

type testController interface {
	GetSentDataChannel() chan string
}

func testSetup(t *testing.T, cb func(m *Mud, e *engine.Engine)) (*Mud, *engine.Engine) {
	var m *Mud

	e := engineTesting.Setup(func(e *engine.Engine) {
		m = New(&Configuration{
			Telnet: &TelnetConfiguration{
				Host: "localhost",
				Port: 4000,
			},
		})

		e.RegisterPlugin(m)
		e.RegisterController(&testCon{
			SentDataChannel: make(chan string),
		})

		cb(m, e)
	}, "telnet")

	return m, e
}
