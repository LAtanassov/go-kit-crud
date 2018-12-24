package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	command := newCommand()
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type cli struct{}

func newCommand() *cobra.Command {

	c := cli{}

	var rootCmd = cobra.Command{
		Use:   "usercli",
		Short: "User CLI",
		Long:  `User CLI is grpc client to perform action to User service`,
	}

	var createCmd = &cobra.Command{
		Use:   "create <username> <givenname> <familyname>",
		Short: "Creates a new user",
		Long:  `Creates a new user`,
		Run:   c.createRun,
	}

	var readCmd = &cobra.Command{
		Use:   "read <username>",
		Short: "Reads a existing user",
		Long:  `Reads a existing user`,
		Run:   c.readRun,
	}

	var updateCmd = &cobra.Command{
		Use:   "update <username> <givenname> <familyname>",
		Short: "Updates an existing user",
		Long:  `Updates an existing user`,
		Run:   c.updateRun,
	}

	var deleteCmd = &cobra.Command{
		Use:   "delete <username>",
		Short: "Deletes a existing user",
		Long:  `Deletes a existing user`,
		Run:   c.deleteRun,
	}

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)

	return &rootCmd
}

func (c *cli) createRun(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		cmd.Help()
		os.Exit(1)
	}
}

func (c *cli) readRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		os.Exit(1)
	}
}

func (c *cli) updateRun(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		cmd.Help()
		os.Exit(1)
	}
}

func (c *cli) deleteRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		os.Exit(1)
	}
}
