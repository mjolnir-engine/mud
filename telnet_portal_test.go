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
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestTelnetPortal_CreatesListener(t *testing.T) {
	e := engineTesting.Setup(func(e *engine.Engine) {
		e.RegisterPlugin(New(&Configuration{
			Telnet: &TelnetConfiguration{
				Host: "localhost",
				Port: 4000,
			},
		}))
	}, "telnet")
	defer e.Stop()

	connected := make(chan bool)

	go func() {
		con, err := net.Dial("tcp", "localhost:4000")
		defer func() { _ = con.Close() }()

		if err != nil {
			t.Error(err)
		}

		if con == nil {
			t.Error("Connection is nil")
		}

		buf := make([]byte, 1024)

		_, err = con.Read(buf)
		connected <- true
	}()

	assert.True(t, <-connected)
}

func TestTelnetPortal_SendsDataToSession(t *testing.T) {
	_, e := testSetup(t, func(m *Mud, e *engine.Engine) {
	})

	defer e.Stop()

	controller, err := e.GetController("test")

	assert.NoError(t, err)

	con, err := net.Dial("tcp", "localhost:4000")
	defer func() { _ = con.Close() }()

	assert.Equal(t, "Welcome to the test server!", <-controller.(testController).GetSentDataChannel())
}

func TestTelnetPortal_ConnectionCanReceiveData(t *testing.T) {
	_, e := testSetup(t, func(m *Mud, e *engine.Engine) {
	})

	defer e.Stop()

	controller, err := e.GetController("test")

	assert.NoError(t, err)

	con, err := net.Dial("tcp", "localhost:4000")
	defer func() { _ = con.Close() }()

	if err != nil {
		t.Error(err)
	}

	if con == nil {
		t.Error("Connection is nil")
	}

	buf := make([]byte, 1024)

	_, err = con.Write([]byte("test"))

	<-controller.(testController).GetSentDataChannel()

	n, err := con.Read(buf)

	assert.Equal(t, "Mjolnir MUD Engine\r\n", string(buf[:n]))
}
