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
	"github.com/charmbracelet/lipgloss"
	"github.com/mjolnir-engine/engine"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testTheme struct{}

func (t *testTheme) Name() string {
	return "test"
}

func (t *testTheme) GetStyle(name string) lipgloss.Style {
	return lipgloss.NewStyle()
}

func TestMud_RegisterThemeAndGetTheme(t *testing.T) {
	m, e := testSetup(t, func(m *Mud, e *engine.Engine) {
		m.RegisterTheme(&testTheme{})
	})
	defer e.Stop()

	theme, err := m.GetTheme("test")

	assert.NoError(t, err)

	assert.Equal(t, "test", theme.Name())
}

func TestMud_RenderWithTheme(t *testing.T) {
	m, e := testSetup(t, func(m *Mud, e *engine.Engine) {
		m.RegisterTheme(&testTheme{})
		m.RegisterTemplate(&testTemplate{})
	})
	defer e.Stop()

	result, err := m.RenderWithTheme("test", "test", nil)

	assert.NoError(t, err)

	assert.Equal(t, "This is a test", result)
}
