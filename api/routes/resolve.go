package routes

import (
	"fmt"
	"url_shortner/database"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	rdb := database.CreateClient(0)
	defer rdb.Close()

	value, err := rdb.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		fmt.Println("URL not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "URL Not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can not connect to DB"})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	// return c.Redirect(value, 301) // shown in video but not working
	c.Redirect(value, 301)
	return nil
}
