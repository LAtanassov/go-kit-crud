package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LAtanassov/go-kit-crud/pkg/pb"

	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

type cli struct {
	timeout time.Duration
	conn    *grpc.ClientConn
	client  pb.UserClient
}

func main() {

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	command := newCommand(&cli{
		timeout: 15 * time.Second,
		conn:    conn,
		client:  client,
	})
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newCommand(c *cli) *cobra.Command {

	var rootCmd = cobra.Command{
		Use:   "usercli",
		Short: "User CLI",
		Long:  `User CLI is grpc client to perform action to User service`,
	}

	var createCmd = cobra.Command{
		Use:   "create <username> <givenname> <familyname>",
		Short: "Creates a new user",
		Long:  `Creates a new user`,
		Run:   c.createRun,
	}

	var readCmd = cobra.Command{
		Use:   "read <username>",
		Short: "Reads a existing user",
		Long:  `Reads a existing user`,
		Run:   c.readRun,
	}

	var updateCmd = cobra.Command{
		Use:   "update <username> <givenname> <familyname>",
		Short: "Updates an existing user",
		Long:  `Updates an existing user`,
		Run:   c.updateRun,
	}

	var deleteCmd = cobra.Command{
		Use:   "delete <username>",
		Short: "Deletes a existing user",
		Long:  `Deletes a existing user`,
		Run:   c.deleteRun,
	}

	rootCmd.AddCommand(&createCmd)
	rootCmd.AddCommand(&readCmd)
	rootCmd.AddCommand(&updateCmd)
	rootCmd.AddCommand(&deleteCmd)

	return &rootCmd
}

func (c *cli) createRun(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		cmd.Help()
		os.Exit(1)
	}

	req := pb.CreateRequest{Username: args[0], Givenname: args[1], Familyname: args[2]}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.client.Create(ctx, &req)
	if err != nil {
		log.Fatalf("fail to create new user with err: %v", err)
	}
}

func (c *cli) readRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		os.Exit(1)
	}

	req := pb.ReadRequest{Username: args[0]}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.client.Read(ctx, &req)
	if err != nil {
		log.Fatalf("fail to read user with err: %v", err)
	}
}

func (c *cli) updateRun(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		cmd.Help()
		os.Exit(1)
	}

	req := pb.UpdateRequest{Username: args[0], Givenname: args[1], Familyname: args[2]}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.client.Update(ctx, &req)
	if err != nil {
		log.Fatalf("fail to upate user with err: %v", err)
	}
}

func (c *cli) deleteRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Help()
		os.Exit(1)
	}

	req := pb.DeleteRequest{Username: args[0]}

	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	_, err := c.client.Delete(ctx, &req)
	if err != nil {
		log.Fatalf("fail to delete user with err: %v", err)
	}
}
