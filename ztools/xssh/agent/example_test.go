// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package agent_test

import (
	"log"
	"net"
	"os"

	"gopkg.in/eniac/zgrab.v0/ztools/xssh"
	"gopkg.in/eniac/zgrab.v0/ztools/xssh/agent"
)

func ExampleClientAgent() {
	// ssh-agent has a UNIX socket under $SSH_AUTH_SOCK
	socket := os.Getenv("SSH_AUTH_SOCK")
	conn, err := net.Dial("unix", socket)
	if err != nil {
		log.Fatalf("net.Dial: %v", err)
	}
	agentClient := agent.NewClient(conn)
	config := &xssh.ClientConfig{
		User: "username",
		Auth: []xssh.AuthMethod{
			// Use a callback rather than PublicKeys
			// so we only consult the agent once the remote server
			// wants it.
			xssh.PublicKeysCallback(agentClient.Signers),
		},
	}

	sshc, err := xssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatalf("Dial: %v", err)
	}
	// .. use sshc
	sshc.Close()
}
