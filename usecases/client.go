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

// Package usecases implements the business logic of the application
package usecases

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
	ws "github.com/gorilla/websocket"
	"gitlab.com/Kalachevskyi/go-ws-client/config"
)

//WSClient represents an interface for webSocket connections
type WSClient interface {
	Connect() error
}

//responser - represents an interface for the different format of response
type responser interface {
	Print([]byte) error
}

//Connection - represents an interface of connection
type Connection interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
}

// NewClient returning a new default WebSocket Client instance
func NewClient(con Connection, resp responser, conf config.Config) *Client {
	client := &Client{
		con:         con,
		conf:        conf,
		resp:        resp,
		done:        make(chan struct{}),
		err:         make(chan error),
		nextRequest: make(chan struct{}),
	}

	if conf.MultiLine {
		client.readStdIn = client.readStdInMultiLine
		return client
	}

	client.readStdIn = client.readStdInSingleLine
	return client
}

//Client - represents structure to WebSocket client
type Client struct {
	conf        config.Config
	con         Connection
	nextRequest chan struct{}
	done        chan struct{}
	err         chan error
	resp        responser
	readStdIn
}

//Connect - process ws connection
func (c *Client) Connect() error {
	go c.processResponse()
	go c.processStdIn()
	c.nextRequest <- struct{}{}

	err := <-c.err
	close(c.done)
	return err
}

// processResponse - processing WebSocket response
func (c *Client) processResponse() {
	for {
		select {
		case <-c.done:
			return
		default:
			_, resp, err := c.con.ReadMessage()
			if err != nil {
				c.err <- err
				return
			}

			err = c.resp.Print(resp)
			if err != nil {
				c.err <- err
				return
			}
			c.nextRequest <- struct{}{}
		}
	}
}

// processStdIn - processing incoming data from the terminal
func (c *Client) processStdIn() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-c.done:
			return
		case <-c.nextRequest:
			if _, err := color.New(color.FgGreen).Print(config.InputIndicator); err != nil {
				c.err <- err
				return
			}

			if err := c.con.WriteMessage(ws.TextMessage, c.readStdIn(scanner)); err != nil {
				c.err <- err
				return
			}

			if err := scanner.Err(); err != nil {
				c.err <- err
				return
			}
		}
	}
}

type readStdIn func(s *bufio.Scanner) []byte

func (c *Client) readStdInMultiLine(s *bufio.Scanner) []byte {
	var lines []byte
	for s.Scan() {
		line := s.Bytes()
		lines = append(lines, line...)
		if strings.HasSuffix(string(line), c.conf.Delimiter) {
			break
		}
	}

	lines = lines[:len(lines)-1]
	return lines
}

func (c *Client) readStdInSingleLine(s *bufio.Scanner) []byte {
	s.Scan()
	return s.Bytes()
}
