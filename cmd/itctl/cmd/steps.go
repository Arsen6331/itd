/*
 *	itd uses bluetooth low energy to communicate with InfiniTime devices
 *	Copyright (C) 2021 Arsen Musayelyan
 *
 *	This program is free software: you can redistribute it and/or modify
 *	it under the terms of the GNU General Public License as published by
 *	the Free Software Foundation, either version 3 of the License, or
 *	(at your option) any later version.
 *
 *	This program is distributed in the hope that it will be useful,
 *	but WITHOUT ANY WARRANTY; without even the implied warranty of
 *	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *	GNU General Public License for more details.
 *
 *	You should have received a copy of the GNU General Public License
 *	along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.arsenm.dev/itd/internal/types"
)

// steps.goCmd represents the steps.go command
var stepsCmd = &cobra.Command{
	Use:   "steps",
	Short: "Get step count from InfiniTime",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to itd UNIX socket
		conn, err := net.Dial("unix", viper.GetString("sockPath"))
		if err != nil {
			log.Fatal().Err(err).Msg("Error dialing socket. Is itd running?")
		}
		defer conn.Close()

		// Encode request into connection
		err = json.NewEncoder(conn).Encode(types.Request{
			Type: types.ReqTypeStepCount,
		})
		if err != nil {
			log.Fatal().Err(err).Msg("Error making request")
		}

		// Read one line from connection
		line, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			log.Fatal().Err(err).Msg("Error reading line from connection")
		}

		var res types.Response
		// Decode line into response
		err = json.Unmarshal(line, &res)
		if err != nil {
			log.Fatal().Err(err).Msg("Error decoding JSON data")
		}

		if res.Error {
			log.Fatal().Msg(res.Message)
		}

		// Print returned BPM
		fmt.Printf("%d Steps\n", int(res.Value.(float64)))
	},
}

func init() {
	getCmd.AddCommand(stepsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// steps.goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// steps.goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}