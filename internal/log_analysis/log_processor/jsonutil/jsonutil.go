package jsonutil

/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	jsoniter "github.com/json-iterator/go"
)

func NewEncoderNamingStrategy(translate func(string) string) jsoniter.Extension {
	return &encoderNamingStrategy{
		translate: translate,
	}
}

type encoderNamingStrategy struct {
	jsoniter.DummyExtension
	translate func(string) string
}

// UpdateStructDescription implements jsoniter.Extension.
// It rewrites field names using the provided translate function
func (extension *encoderNamingStrategy) UpdateStructDescriptor(desc *jsoniter.StructDescriptor) {
	for _, binding := range desc.Fields {
		for i, name := range binding.ToNames {
			if name := extension.translate(name); name != "" {
				binding.ToNames[i] = name
			}
		}
	}
}

func UnquoteJSON(data []byte) []byte {
	if len(data) > 1 && data[0] == '"' {
		data = data[1:]
		if n := len(data) - 1; 0 <= n && n < len(data) && data[n] == '"' {
			return data[:n]
		}
	}
	return data
}

