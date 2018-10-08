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
	"log"

	ws "github.com/gorilla/websocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/Kalachevskyi/go-ws-client/config"
)

// RootCMD represents the base command when called without any subcommands
func RootCMD() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
		Use:   "ws-client",
		Short: "The WebSocket command-line client",
		Long: `The WebSocket command-line client, written on Golang, 
intended for debugging applications that use WebSocket technology.`,
		Run: webSocket,
	}

	// Initialize flags for the root command
	flags := cmd.Flags()
	flags.StringP(config.FormatFlag, config.FormatFlagShort, "",
		"Format the webSocket response (Only json is available at the moment)")
	flags.StringP(config.DelimiterFlag, config.DelimiterFlagShort, config.DelimiterDefault,
		"Delimiter for a multi-line request.")
	flags.BoolP(config.MultiLineFlag, config.MultiLineFlagShort, false,
		"Multi-line request.")
	flags.StringP(config.URLFlag, config.URLFlagShort, "",
		"Connect to WebSocket host.")

	// Required parameters
	if err := cmd.MarkFlagRequired(config.URLFlag); err != nil {
		return nil, err
	}

	// AutomaticEnv has Viper check ENV variables for all.
	// keys set in config, default & flags
	viper.AutomaticEnv()
	return
}

// webSocket establish connection to the WebSocket server
func webSocket(cmd *cobra.Command, args []string) {
	// Initializing zap logger
	loggerConfig := config.NewProductionConfig()
	zLog, err := loggerConfig.Build()
	if err != nil {
		log.Println(err)
		return
	}

	// Bind cobra flags to viper
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		zLog.Error(err.Error())
		return
	}

	// Loading config
	conf := &config.Config{}
	if err = viper.Unmarshal(conf); err != nil {
		zLog.Error(err.Error())
		return
	}

	url := cmd.Flag(config.URLFlag).Value.String()
	con, _, err := ws.DefaultDialer.Dial(url, nil)
	if err != nil {
		zLog.Error(err.Error())
		return
	}

	defer func() {
		if err := con.Close(); err != nil {
			zLog.Error(err.Error())
		}
	}()

	uc, err := getUseCase(cmd.Flag(config.FormatFlag).Value.String(), con, *conf)
	if err != nil {
		zLog.Error(err.Error())
		return
	}

	if err := uc.Connect(); err != nil {
		zLog.Error(err.Error())
	}
	return
}
