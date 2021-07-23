package context

import (
	"github.com/urfave/cli"
	"go.step.sm/cli-utils/command"
)

// init creates and registers the ca command
func init() {
	cmd := cli.Command{
		Name:      "context",
		Usage:     "manage certificate authority contexts",
		UsageText: "**step context** <subcommand> [arguments] [global-flags] [subcommand-flags]",
		Description: `**step context** command group provides facilities to manage certificate
authority contexts.

## EXAMPLES

'''
$ cat $(step path --base)/contexts.json
{
    "alpha-one": {
        "authority": "alpha-one.ca.smallstep.com",
        "profile": "alpha-one"
    },
    "alpha-two": {
        "authority": "alpha-two.ca.smallstep.com",
        "profile": "alpha-two"
    },
    "beta": {
        "authority": "beta.ca.smallstep.com",
        "profile": "beta"
    }
}
'''

Set the default certificate authority context:
'''
$ step context set alpha-one
'''

List the available certificate authority contexts:
'''
$ step context list
alpha-one
alpha-two
beta
'''`,
		Subcommands: cli.Commands{
			setCommand(),
			listCommand(),
		},
	}

	command.Register(cmd)
}
