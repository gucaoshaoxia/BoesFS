/*
 * Copyright (c) 2019, 2020 Oracle and/or its affiliates. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"github.com/fnproject/cli/common"
	"github.com/urfave/cli"
)

// GetCommand returns get cli.command
func GetCommand() cli.Command {
	return cli.Command{
		Name:         "get",
		Aliases:      []string{"g"},
		Usage:        "\tGet an object to retrieve its information",
		Category:     "MANAGEMENT COMMANDS",
		Hidden:       false,
		ArgsUsage:    "<subcommand>",
		Description:  "This command gets a 'call', 'configuration' or 'log' to retrieve information for an object ('app' or 'function').",
		Subcommands:  GetCommands(GetCmds),
		BashComplete: common.DefaultBashComplete,
	}
}
