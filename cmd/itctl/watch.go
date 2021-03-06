package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func watchHeart(c *cli.Context) error {
	heartCh, err := client.WatchHeartRate(c.Context)
	if err != nil {
		return err
	}

	for {
		select {
		case heartRate := <-heartCh:
			if c.Bool("json") {
				json.NewEncoder(os.Stdout).Encode(
					map[string]uint8{"heartRate": heartRate},
				)
			} else if c.Bool("shell") {
				fmt.Printf("HEART_RATE=%d\n", heartRate)
			} else {
				fmt.Println(heartRate, "BPM")
			}
		case <-c.Done():
			return nil
		}
	}
}

func watchBattLevel(c *cli.Context) error {
	battLevelCh, err := client.WatchBatteryLevel(c.Context)
	if err != nil {
		return err
	}

	for {
		select {
		case battLevel := <-battLevelCh:
			if c.Bool("json") {
				json.NewEncoder(os.Stdout).Encode(
					map[string]uint8{"battLevel": battLevel},
				)
			} else if c.Bool("shell") {
				fmt.Printf("BATTERY_LEVEL=%d\n", battLevel)
			} else {
				fmt.Printf("%d%%\n", battLevel)
			}
		case <-c.Done():
			return nil
		}
	}
}

func watchStepCount(c *cli.Context) error {
	stepCountCh, err := client.WatchStepCount(c.Context)
	if err != nil {
		return err
	}

	for {
		select {
		case stepCount := <-stepCountCh:
			if c.Bool("json") {
				json.NewEncoder(os.Stdout).Encode(
					map[string]uint32{"stepCount": stepCount},
				)
			} else if c.Bool("shell") {
				fmt.Printf("STEP_COUNT=%d\n", stepCount)
			} else {
				fmt.Println(stepCount, "Steps")
			}
		case <-c.Done():
			return nil
		}
	}
}

func watchMotion(c *cli.Context) error {
	motionCh, err := client.WatchMotion(c.Context)
	if err != nil {
		return err
	}

	for {
		select {
		case motionVals := <-motionCh:
			if c.Bool("json") {
				json.NewEncoder(os.Stdout).Encode(motionVals)
			} else if c.Bool("shell") {
				fmt.Printf(
					"X=%d\nY=%d\nZ=%d\n",
					motionVals.X,
					motionVals.Y,
					motionVals.Z,
				)
			} else {
				fmt.Println(motionVals)
			}
		case <-c.Done():
			return nil
		}
	}
}
