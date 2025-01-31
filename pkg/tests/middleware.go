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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/spf13/viper"

	"github.com/go-sigma/sigma/pkg/types/enums"
)

// CIDatabase is the interface for the database in ci tests
type CIDatabase interface {
	// Init initializes the database or database file for ci tests
	Init() error
	// DeInit remove the database or database file for ci tests
	DeInit() error
}

type Factory interface {
	New() CIDatabase
}

var ciDatabaseFactories = make(map[string]Factory)

// RegisterCIDatabaseFactory registers a storage factory driver by name.
// If RegisterCIDatabaseFactory is called twice with the same name or if driver is nil, it panics.
func RegisterCIDatabaseFactory(name string, factory Factory) error {
	if _, ok := ciDatabaseFactories[name]; ok {
		return fmt.Errorf("ci database %q already registered", name)
	}
	ciDatabaseFactories[name] = factory
	return nil
}

// DB is the database for ci tests
var DB CIDatabase

func Initialize(t *testing.T) error {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	miniRedis := miniredis.RunT(t)
	viper.SetDefault("redis.url", "redis://"+miniRedis.Addr())

	typ := viper.GetString("ci.database.type")
	if typ == "" {
		typ = enums.DatabaseSqlite3.String()
	}

	factory, ok := ciDatabaseFactories[typ]
	if !ok {
		return fmt.Errorf("ci database %q not registered", typ)
	}
	DB = factory.New()

	return nil
}
