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
	"database/sql"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid"

	"github.com/go-sigma/sigma/pkg/configs"
	"github.com/go-sigma/sigma/pkg/dal"
	"github.com/go-sigma/sigma/pkg/types/enums"
)

func init() {
	err := RegisterCIDatabaseFactory("mysql", &mysqlFactory{})
	if err != nil {
		panic(err)
	}
}

type mysqlFactory struct{}

var _ Factory = &mysqlFactory{}

func (mysqlFactory) New() CIDatabase {
	return &mysqlCIDatabase{}
}

type mysqlCIDatabase struct {
	dbname string
}

var _ CIDatabase = &mysqlCIDatabase{}

// Init sets the default values for the database configuration in ci tests
func (d *mysqlCIDatabase) Init() error {
	db, err := sql.Open("mysql", "root:sigma@tcp(127.0.0.1:3306)/")
	if err != nil {
		return err
	}

	d.dbname = gonanoid.MustGenerate("abcdefghijklmnopqrstuvwxyz", 6)
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", d.dbname))
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}

	err = dal.Initialize(configs.Configuration{
		Database: configs.ConfigurationDatabase{
			Type: enums.DatabaseMysql,
			Mysql: configs.ConfigurationDatabaseMysql{
				Host:     "127.0.0.1",
				Port:     3306,
				User:     "root",
				Password: "sigma",
				DBName:   d.dbname,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// DeInit remove the database or database file for ci tests
func (d *mysqlCIDatabase) DeInit() error {
	db, err := sql.Open("mysql", "root:sigma@tcp(127.0.0.1:3306)/")
	if err != nil {
		return err
	}
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s", d.dbname))
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}
