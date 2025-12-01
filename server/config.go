package server

import (
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func Config() fiber.Config {
	readTimeOutSecondsCount, err := strconv.Atoi(os.Getenv("READ_TIMEOUT_SECONDS"))
	if err != nil {
		readTimeOutSecondsCount = 10
	}

	return fiber.Config{
		ReadTimeout: time.Duration(readTimeOutSecondsCount) * time.Second,
		Prefork:     os.Getenv("PREFORK") == "true",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	}
}
