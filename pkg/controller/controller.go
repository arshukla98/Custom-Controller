package controller

import (
	"context"
	"fmt"
	upgradev1 "github.com/arshukla98/sample-controller/pkg/apis/upgrade/v1"
)

func ProcessKubeUpgrade(ctx context.Context, uk *upgradev1.UpgradeKube) (*upgradev1.UpgradeKube, error) {
	defer fmt.Println("Exiting ProcessKubeUpgrade")
	fmt.Println("Entering ProcessKubeUpgrade")
	fmt.Println("Payload:", uk)
	return nil, nil
}
