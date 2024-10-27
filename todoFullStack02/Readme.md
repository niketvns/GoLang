# Backend with `Go`

## Go `fiber`

An Express-inspired web framework written in Go.

> Fiber is a Go web `framework built on top of Fasthttp`, the fastest HTTP engine for Go. It's designed to ease things up for fast development with zero memory allocation and performance in mind.

```go
// main.go file
package main

import (
    "log" //
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Create app instance
    app := fiber.New()


    // create a get endpoint
    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // listen on 3000 port
    // log.Fatal is equivalent to Print followed by a call to os.Exit(1).
    log.Fatal(app.Listen(":3000"))
}
```

## Live reload with `air`

```go
go install github.com/air-verse/air@latest
```

> Setup `air.toml` file for live reload

```toml
root = "."
tmp_dir = "tmp"

[build]
    bin = "main"
    cmd = "go build -o {{.Output}} {{.Input}}"
    exclude = ["tmp/*", "client/*"]
    include = ["**/*.go"]
    ignore = ["tmp/*"]
```

After air.toml setup, run the `air` command in the terminal to run the `main.go` file. It will create `temp` folder and `main.exe` file inside it.  
`temp -> main.exe`

> 💡 Now we are good to go.

# Let's learn with a simple `todo` BE App.

## Structure of Todo

```go
type Todo struct {
	Id        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"isCompleted"`
}
```

## Create a `GET` API endpoint to get all todos

```go
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"isCompleted"`
}

func main() {
	app := fiber.New()

	todos := []Todo{}

    // Here (c *fiber.Ctx) -> context pointer
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Server is up!"})
	})

	// Get all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Success!", "todos": todos})
	})

	log.Fatal(app.Listen(":4000"))
}

```

## Create a `POST` API Endpoint to create a todo

```go
// Create a todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}
        //  store req body in todo pointer
		c.BodyParser(todo)
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required!"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})
```

## Create a `PATCH` endpoint to update todo

```go
app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.Id) == id {
			todos[i].Completed = true
			return c.Status(200).JSON(fiber.Map{
					"todo": todos[i],
				})
			}
	}

	return c.Status(404).JSON(fiber.Map{
			"msg": "Todo not found!",
	})
})
```

## Create a `DELETE` endpoint to delete a todo

```go
app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
	id := c.Params("id")

	for index, todo := range todos {
		if fmt.Sprint(todo.Id) == id {
			todos = append(todos[:index], todos[:index+1]...)
			return c.Status(200).JSON(fiber.Map{
					"msg": "Todo with ID: " + id + " deleted successfull",
			})
		}
	}

	return c.Status(404).JSON(fiber.Map{
			"msg": "Todo not exist!",
		})
})
```

> Note: Delete an item from Array/Slice
> `append`(todos[:index], todos[:index+1]...)

## How to use `.env` in `Go`

Install `godotenv` package

```bash
go get github.com/joho/godotenv
```

Create a `.env` file inside root folder of the project.

```env
PORT=500
```

```go
package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	// How to use .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}
	PORT := os.Getenv("PORT")

	log.Fatal(app.Listen(":" + PORT))
}
```

## Connect with `MongoDB`
