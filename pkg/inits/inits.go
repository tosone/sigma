// Copyright 2023 sigma
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package inits

import "github.com/go-sigma/sigma/pkg/configs"

var inits = make(map[string]func(configs.Configuration) error)

// Initialize runs all registered inits.
func Initialize(config configs.Configuration) error {
	for _, init := range inits {
		err := init(config)
		if err != nil {
			return err
		}
	}
	return nil
}
