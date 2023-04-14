package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func GetLinks(ctx *fiber.Ctx) error {
	folderPath := "./documents/"

	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {

		fmt.Printf("An error occured while reading the folder: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "false", "err": err})
	}
	body := "<html>"
	body += "<head><title>Masti</title></head>"
	body += "<body>"
	body += "<h1>Links to happiness</h1>"
	body += "<ul>"

	for _, file := range files {
		body += fmt.Sprintf("<li><a href='%s'>%s</a></li>", file, file)
	}

	body += "</ul>"
	body += "</body>"
	body += "</html>"

	fmt.Println("Files in the folder:")
	// for _, file := range files {
	// 	fmt.Println(file)
	// }
	// err = databaseClient.RedisClient.Set(context.Background(), token, newUser.Id.String(), 0).Err()
	// if err != nil {
	// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "false", "err": "Could not update refreshToken"})
	// }

	return ctx.Type("html").SendString(body)
}
