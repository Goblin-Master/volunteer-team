package test

import (
	"fmt"
	"testing"
	"volunteer-team/backend/internal/infrastructure/configs"
)

func TestConfigs(t *testing.T) {
	configs.LoadConfig()
	fmt.Println(configs.Conf)
}
