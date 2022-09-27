package main

import (
	"context"
	"github.com/edaniels/golog"
	"go.viam.com/rdk/robot/client"
	"go.viam.com/rdk/utils"
	"go.viam.com/utils/rpc"
)

func connect() (*client.RobotClient, golog.Logger) {
	logger := golog.NewDevelopmentLogger("client")
	robot, err := client.New(
		context.Background(),
		"starter-bot-main.k4xl69bmso.viam.cloud",
		logger,
		client.WithDialOptions(rpc.WithCredentials(rpc.Credentials{
			Type:    utils.CredentialsTypeRobotLocationSecret,
			Payload: "e7ztx7s67d4pnnnor54qkn2wlhf58cd7erukcldonha0cc62",
		})),
	)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("resources : ", robot.ResourceNames())
	return robot, logger
}

func close(logger golog.Logger, robot *client.RobotClient) {
	if err := robot.Close(context.Background()); err != nil {
		logger.Fatal(err)
	}
}

func main() {
	robot, logger := connect()
	defer close(logger, robot)
}
