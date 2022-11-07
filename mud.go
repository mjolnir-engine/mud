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
	"github.com/mjolnir-engine/mud/data_sources"
	"github.com/rs/zerolog"
)

type Mud struct {
	Engine           *engine.Engine
	config           *Configuration
	logger           zerolog.Logger
	portal           Portal
	templateRegistry *templateRegistry
	themeRegistry    *themeRegistry
}

func (m *Mud) Name() string {
	return "mud"
}

func (m *Mud) Init(e *engine.Engine) error {
	m.Engine = e
	m.logger = e.Logger().With().Str("plugin", m.Name()).Logger()
	m.templateRegistry = newTemplateRegistry(m)
	m.themeRegistry = newThemeRegistry(m)

	e.RegisterService("telnet")

	return nil
}

func (m *Mud) Start(e *engine.Engine) error {
	e.RegisterDataSource(data_sources.CreateAccountsDataSource(e))

	if e.GetService() == "telnet" {
		m.portal = newTelnetPortal(m)
		m.portal.Start()
	}

	return nil
}

func (m *Mud) Stop(e *engine.Engine) error {
	return nil
}

// New creates a new instance of the Mud plugin.
func New(config *Configuration) *Mud {
	return &Mud{
		config: config,
	}
}
