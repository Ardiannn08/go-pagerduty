package main

import (
	"fmt"
	"strings"

	"github.com/Ardiannn08/go-pagerduty"
	"github.com/mitchellh/cli"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type UserList struct {
	Meta
}

func UserListCommand() (cli.Command, error) {
	return &UserList{}, nil
}

func (c *UserList) Help() string {
	helpText := `
	pd user list List users

	Options:

		 -query     Filter users with certain name
		 -team      Filter users by team id(s)
		 -include   Additional details to include

	`
	return strings.TrimSpace(helpText)
}

func (c *UserList) Synopsis() string {
	return "List users of your PagerDuty account, optionally filtered by a search query"
}

func (c *UserList) Run(args []string) int {
	var query string
	var teamIDs []string
	var includes []string

	flags := c.Meta.FlagSet("user list")
	flags.Usage = func() { fmt.Println(c.Help()) }
	flags.StringVar(&query, "query", "", "Show users whose names contain the query")
	flags.Var((*ArrayFlags)(&teamIDs), "team", "Filter users by team ids (can be specified multiple times)")
	flags.Var((*ArrayFlags)(&includes), "include", "Additional details to include (can be specified multiple times)")

	if err := flags.Parse(args); err != nil {
		log.Error(err)
		return -1
	}
	if err := c.Meta.Setup(); err != nil {
		log.Error(err)
		return -1
	}
	client := c.Meta.Client()
	opts := pagerduty.ListUsersOptions{
		Query:    query,
		TeamIDs:  teamIDs,
		Includes: includes,
	}
	if resp, err := client.ListUsers(opts); err != nil {
		log.Error(err)
		return -1
	} else {
		for i, p := range resp.Users {
			fmt.Println("Entry: ", i)
			data, err := yaml.Marshal(p)
			if err != nil {
				log.Error(err)
				return -1
			}
			fmt.Println(string(data))
		}
	}
	return 0
}
