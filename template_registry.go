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

import "github.com/rs/zerolog"

type templateRegistry struct {
	templates map[string]Template
	logger    zerolog.Logger
	mud       *Mud
}

func newTemplateRegistry(m *Mud) *templateRegistry {
	return &templateRegistry{
		templates: make(map[string]Template),
		mud:       m,
		logger:    m.logger.With().Str("component", "template_registry").Logger(),
	}
}

//
//func (r *templateRegistry) register(t Template) {
//	r.logger.Info().Str("template", t.Name()).Msg("registering template")
//	r.templates[t.Name()] = t
//}
//
//func (r *templateRegistry) get(name string) (Template, error) {
//	t, ok := r.templates[name]
//
//	if !ok {
//		return nil, errors.TemplateNotFoundError{
//			Name: name,
//		}
//	}
//
//	return t, nil
//}
//
//// RegisterTemplate registers a template with the template registry. If a template with the same name already exists, it
//// will be overwritten.
//func (m *Mud) RegisterTemplate(t Template) {
//	m.templateRegistry.register(t)
//}
//
//func (m *Mud) RenderTemplate(name string, data interface{}) (string, error) {
//	template, err := m.templateRegistry.get(name)
//
//	if err != nil {
//		return "", err
//	}
//
//	th, err := theme_registry.GetTheme("default")
//
//	if err != nil {
//		return "", err
//	}
//
//	style := th.GetStyleFor(t.Style())
//	text, err := t.Render(data)
//
//	if err != nil {
//		return "", err
//	}
//
//	return style.Render(text), nil
//}
