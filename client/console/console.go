// Wiregost - Post-Exploitation & Implant Framework
// Copyright © 2020 Para
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package console

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lmorg/readline"

	"github.com/maxlandon/wiregost/client/assets"
	"github.com/maxlandon/wiregost/client/completers"
	"github.com/maxlandon/wiregost/client/connection"
	"github.com/maxlandon/wiregost/client/console/util"
	client "github.com/maxlandon/wiregost/proto/v1/gen/go/client"
	dbpb "github.com/maxlandon/wiregost/proto/v1/gen/go/db"
	ghostpb "github.com/maxlandon/wiregost/proto/v1/gen/go/ghost"
	modulepb "github.com/maxlandon/wiregost/proto/v1/gen/go/module"
)

var (
	// Console - The client console object
	Console = newConsole()
)

// Console - Central object of the client UI
type console struct {
	ClientID uuid.UUID             // Unique identifier for this console
	User     *dbpb.User            // User information sent back after auth
	Shell    *readline.Instance    // Console readline input
	Config   *client.ConsoleConfig // Console configuration
	Module   *modulepb.Module      // Module currently on stack
	Ghost    *ghostpb.Ghost        // Current ghost implant
	Ghosts   int
	Jobs     int
}

// newConsole - Instantiates a console with some default behavior
func newConsole() *console {

	console := &console{
		Shell:  readline.NewInstance(),
		Module: &modulepb.Module{}, // Avoid nil dereference
		Ghost:  &ghostpb.Ghost{},   // Avoid nil dereference
	}

	return console
}

// Connect - The console loads the server configuration, connects to it and attempts user authentication
func (c *console) Connect() (err error) {

	// Load server connection configuration (check files in ~/.wiregost first, then binary)
	c.Config = assets.LoadConsoleConfig()

	// Connect to server via TLS
	conn, err := connection.ConnectTLS()

	// Authenticate (5 tries)
	var cli client.ConnectionRPCClient
	cli, c.User, c.ClientID = connection.Authenticate(conn)

	// Receive various infos sent by server when authenticated (ClientID, messages, users, version information, etc)
	c.GetConnectionInfo(cli)

	// Print banner, user and client/server version information
	c.PrintBanner()

	// Register all gRPC clients with the connection
	connection.RegisterRPCClients(conn)

	// Listen for incoming server/implant events
	c.StartEventListener(conn)

	return nil
}

// Setup - Setup various elements of the console.
func (c *console) Setup() {

	// Console configuration (from server first, ~/.wiregost second)
	c.Config = assets.LoadConsoleConfig()

	// Prompt
	c.SetPrompt()

	// Completion, Hints & Syntax
	c.Shell.TabCompleter = completers.TabCompleter
	c.Shell.HintText = completers.HintCompleter
	c.Shell.SyntaxHighlighter = completers.SyntaxHighlighter

	// Env
	util.LoadClientEnv()

	// Commands

	// Share context
}

// GetConnectionInfo - Get all information necessary to console upon connection
func (c *console) GetConnectionInfo(cli client.ConnectionRPCClient) {

	// Version Request
	// cli.GetVersion(ctx, in, opts)

	// Set fields
}

// ShareContext - The console exposes its context to other packages
func (c *console) ShareContext() {

}

// Start - Start the console
func (c *console) Start() {

	// Connect to server and authenticate
	c.Connect()

	// Setup console
	c.Setup()

	// Input loop
	for {
		// Recompute prompt each time
		c.Refresh()

		// Readline
		// line, _ := c.Readline()

		// Split & sanitize
		// sanitized, empty := Sanitize(line)

		// Process tokens

		// Execute the command input

		// Reset variables for command options (go-flags)
	}
}

// Refresh - Computes prompt and current context
func (c *console) Refresh() {

}

// Readline - Add an empty line between input line and output
func (c *console) Readline() (line string, err error) {
	line, err = c.Shell.Readline()
	fmt.Println()
	return
}

// Exit - Kill the current client console
func (c *console) Exit() {

}

// Sanitize - Trims spaces and other unwished elements from the input line
func Sanitize(line string) (sanitized []string, empty bool) {

	// Trim border spaces

	// Catch eventual empty items

	return
}

// PrintBanner - Print Wiregost banner, client & server information, etc.
func (c *console) PrintBanner() {

}
