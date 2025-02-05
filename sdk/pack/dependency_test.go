// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pack

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestDependency_Validate(t *testing.T) {
	testCases := []struct {
		inputDependency          *Dependency
		expectedOutputDependency *Dependency
		expectError              bool
		name                     string
	}{
		{
			inputDependency: &Dependency{
				Name:   "example",
				Source: "git://example.com/example",
			},
			expectedOutputDependency: &Dependency{
				Name:    "example",
				Source:  "git://example.com/example",
				Enabled: pointerOf(true),
			},
			name: "nil enabled input",
		},
		{
			inputDependency: &Dependency{
				Name:    "example",
				Source:  "git://example.com/example",
				Enabled: pointerOf(false),
			},
			expectedOutputDependency: &Dependency{
				Name:    "example",
				Source:  "git://example.com/example",
				Enabled: pointerOf(false),
			},
			name: "false enabled input",
		},
		{
			inputDependency: &Dependency{
				Name:    "example",
				Source:  "git://example.com/example",
				Enabled: pointerOf(true),
			},
			expectedOutputDependency: &Dependency{
				Name:    "example",
				Source:  "git://example.com/example",
				Enabled: pointerOf(true),
			},
			name: "false enabled input",
		},
	}

	for _, tc := range testCases {
		err := tc.inputDependency.validate()
		if tc.expectError {
			must.NotNil(t, err, must.Sprint(tc.name))
		} else {
			must.Nil(t, err, must.Sprint(tc.name))
		}
		must.Eq(t, tc.expectedOutputDependency, tc.inputDependency, must.Sprint(tc.name))
	}
}
