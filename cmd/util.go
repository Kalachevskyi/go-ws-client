// Copyright © 2018 Volodymyr Kalachevskyi <v.kalachevskyi@gmail.com>
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

// Package cmd is a command line interface
package cmd

import (
	"fmt"

	ws "github.com/gorilla/websocket"
	"gitlab.com/Kalachevskyi/go-ws-client/config"
	"gitlab.com/Kalachevskyi/go-ws-client/printer"
	"gitlab.com/Kalachevskyi/go-ws-client/usecases"
)

//ucJSON - usecase for json formation
const ucJSON = "json"

func getUseCase(key string, con *ws.Conn, conf config.Config) (uc usecases.WSClient, err error) {
	switch key {
	case "":
		uc = usecases.NewClient(con, printer.NewResponse(), conf)
		return
	case ucJSON:
		uc = usecases.NewClient(con, printer.NewResponseJSON(), conf)
		return
	default:
		return nil, fmt.Errorf("can't find a usecase for the flag: %s", key)
	}
}
