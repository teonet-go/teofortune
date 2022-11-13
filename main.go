// Copyright 2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Teonet fortune microservice. This is simple Teonet micriservice application
// which return linux `fortune` messages. The [fortune](https://linux.die.net/man/6/fortune)
// application should be installed first.
package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"

	"github.com/teonet-go/teomon"
	"github.com/teonet-go/teonet"
)

const (
	appShort   = "teofortune"
	appName    = "Teonet micriservice application which return linux fortune"
	appLong    = ""
	appVersion = "0.0.4"
)

var appStartTime = time.Now()
var monitor string

func main() {

	// Application logo
	teonet.Logo(appName, appVersion)

	// Parse application command line parameters
	appShortParam := flag.String("name", appShort, "application short name")
	port := flag.Int("p", 0, "local port")
	stat := flag.Bool("stat", false, "show trudp statistic")
	hotkey := flag.Bool("hotkey", false, "run hotkey meny")
	loglevel := flag.String("loglevel", "NONE", "log level")
	logfilter := flag.String("logfilter", "", "log filter")
	flag.StringVar(&monitor, "monitor", "", "monitor address")
	flag.Parse()

	// Create Teonet connector
	teo, err := teonet.New(
		*appShortParam, *port, teonet.Stat(*stat), teonet.Hotkey(*hotkey),
		*loglevel, teonet.Logfilter(*logfilter),
	)
	if err != nil {
		return
	}
	fmt.Println("teonet address:", teo.Address())

	// Connect to Teonet
	err = teo.Connect()
	if err != nil {
		fmt.Println("can't connect to Teonet, error:", err)
		return
	}
	fmt.Println("connected to teonet")

	// Connect to monitor if it set in parameters
	if len(monitor) > 0 {
		teomon.Connect(teo, monitor, teomon.Metric{
			AppName:      appName,
			AppShort:     appShort,
			AppVersion:   appVersion,
			TeoVersion:   teonet.Version,
			AppStartTime: appStartTime,
		})
		fmt.Println("connected to monitor")
	}

	// Create new API, add API commands and start API reader
	api := teo.NewAPI(appName, appShort, appLong, appVersion)
	commands(api)
	teo.AddReader(api.Reader())
	fmt.Printf("this app API description:\n\n%s\n\n", api.Help())

	select {}
}

// commands describe this service API commands
func commands(api *teonet.API) {
	api.Add(
		// Get fortune message; return message string
		func(cmdApi teonet.APInterface) teonet.APInterface {
			var name = "forta"
			cmdApi = teonet.MakeAPI2().
				SetCmd(api.Cmd(129)).              // Command number cmd = 129
				SetName(name).                     // Command name
				SetShort("get 'fortune' message"). // Short description
				SetUsage("").                      // Usage (input parameter)
				SetReturn("<string>").             // Return (output parameters)
				// Command reader (execute when command received)
				SetReader(func(c *teonet.Channel, p *teonet.Packet, data []byte) bool {
					data = []byte(fortune())
					api.SendAnswer(cmdApi, c, data, p)
					return true
				}).SetAnswerMode(teonet.DataAnswer)
			return cmdApi
		}(teonet.APIData{}),

		// Get fortune message; return ID and message string
		func(cmdApi teonet.APInterface) teonet.APInterface {
			var name = "fortb"
			cmdApi = teonet.MakeAPI2().
				SetCmd(api.CmdNext()).             // Command number cmd = 130
				SetName(name).                     // Command name
				SetShort("get 'fortune' message"). // Short description
				SetUsage("").                      // Usage (input parameter)
				SetReturn("<id uint32><string>").  // Return (output parameters)
				// Command reader (execute when command received)
				SetReader(func(c *teonet.Channel, p *teonet.Packet, data []byte) bool {
					data = []byte(fortune())
					api.SendAnswer(cmdApi, c, data, p)
					return true
				}).SetAnswerMode(teonet.PacketIDAnswer | teonet.DataAnswer)
			return cmdApi
		}(teonet.APIData{}),
	)
}

// fortune return linux fortune message
func fortune() string {
	out, err := exec.Command("fortune").Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}
