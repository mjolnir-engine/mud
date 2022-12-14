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

package data_sources

import (
	"github.com/mjolnir-engine/engine"
	"github.com/mjolnir-engine/engine/uid"
	"github.com/stretchr/testify/assert"
	"testing"

	engineTesting "github.com/mjolnir-engine/engine/testing"
)

type fakeAccount struct {
	Id uid.UID `bson:"_id,omitempty"`
}

func TestCreateAccountsDataSource(t *testing.T) {
	e := engineTesting.Setup(func(e *engine.Engine) {
		e.RegisterDataSource(CreateAccountsDataSource(e))
	})
	defer engineTesting.Teardown(e)

	id, err := e.SaveInDataSource("accounts", fakeAccount{})

	assert.NoError(t, err)
	assert.NotNil(t, id)

	err = e.DeleteFromDataSource("accounts", fakeAccount{
		Id: id,
	})

	assert.NoError(t, err)
}
