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
	"github.com/stretchr/testify/assert"
	"testing"
)

type testTemplate struct{}

func (t *testTemplate) Name() string {
	return "test"
}

func (t *testTemplate) Style() string {
	return "test"
}

func (t *testTemplate) Render(ctx interface{}) (string, error) {
	return "This is a test", nil
}

func TestMud_RegisterTemplate(t *testing.T) {
	m, _ := testSetup(t, func(m *Mud, _ *engine.Engine) {
		m.RegisterTemplate(&testTemplate{})
	})

	assert.Len(t, m.templateRegistry.templates, 1)
}
