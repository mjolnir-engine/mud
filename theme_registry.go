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
	"github.com/mjolnir-engine/mud/errors"
	"github.com/rs/zerolog"
)

type themeRegistry struct {
	themes map[string]Theme
	mud    *Mud
	logger zerolog.Logger
}

func newThemeRegistry(mud *Mud) *themeRegistry {
	return &themeRegistry{
		themes: make(map[string]Theme),
		mud:    mud,
		logger: mud.logger.With().Str("component", "themeRegistry").Logger(),
	}
}

func (r *themeRegistry) register(theme Theme) {
	r.logger.Info().Str("theme", theme.Name()).Msg("registering theme")
	r.themes[theme.Name()] = theme
}

func (r *themeRegistry) get(name string) (Theme, error) {
	r.logger.Debug().Str("theme", name).Msg("getting theme")
	theme, ok := r.themes[name]

	if !ok {
		return nil, errors.ThemeNotFoundError{
			Name: name,
		}
	}

	return theme, nil
}

// RegisterTheme registers a theme with the theme registry. If a theme with the same name already exists, it will be
// overwritten.
func (m *Mud) RegisterTheme(theme Theme) {
	m.themeRegistry.register(theme)
}

// GetTheme returns a theme from the theme registry. If the theme does not exist, an error is returned.
func (m *Mud) GetTheme(name string) (Theme, error) {
	return m.themeRegistry.get(name)
}
