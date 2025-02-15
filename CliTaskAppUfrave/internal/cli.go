package internal

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v3"
)

func Cli() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add task to json file",
				Action: func(ctx context.Context, c *cli.Command) error {
					desciption := c.Args().First()
					task := NewTask(desciption)
					fmt.Println(task)
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete task",
				Action: func(ctx context.Context, c *cli.Command) error {
					idString := c.Args().First()
					id, err := strconv.Atoi(idString)
					if err != nil {
						return err
					}
					DeleteTask(id)
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "update task by id",
				Action: func(ctx context.Context, c *cli.Command) error {
					Update(c.Args().Slice())
					return nil
				},
			},
			{
				Name:  "mark-in-progress",
				Usage: "in progress",
				Action: func(ctx context.Context, c *cli.Command) error {
					idString := c.Args().First()
					id, err := strconv.Atoi(idString)
					if err != nil {
						return err
					}
					InProgress(id)
					return nil
				},
			},
			{
				Name:  "mark-done",
				Usage: "done",
				Action: func(ctx context.Context, c *cli.Command) error {
					idString := c.Args().First()
					id, err := strconv.Atoi(idString)
					if err != nil {
						return err
					}
					Done(id)
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "list issues",
				Action: func(ctx context.Context, c *cli.Command) error {
					ListTasks(c.Args().Slice())
					return nil
				},
			},
		},
	}

	cmd.Run(context.Background(), os.Args)
}
