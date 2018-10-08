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

// Package printer is an implementation of response handlers
package printer

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"gitlab.com/Kalachevskyi/go-ws-client/config"
	"gitlab.com/Kalachevskyi/go-ws-client/jsonformation"
)

//NewResponseJSON - returns new ResponseJSON instance
func NewResponseJSON() *ResponseJSON {
	return &ResponseJSON{}
}

//ResponseJSON - represents the json output response
type ResponseJSON struct{}

//Print - printing with JSON type formatting
func (*ResponseJSON) Print(resp []byte) error {
	resp, err := jsonformation.Format(resp)
	if err != nil {
		return errors.New("can't format json response")
	}

	if _, err := color.New(color.FgRed).Print(config.OutputIndicator); err != nil {
		return err
	}

	if _, err := fmt.Println(string(resp)); err != nil {
		return err
	}
	return nil
}
