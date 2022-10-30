package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()

  m.Get("/", func() string {
    return "Hello world!"
  })

  // Martini listens on 3000 by default if the PORT environment variable does not exist or is not defined
  m.Run()
}