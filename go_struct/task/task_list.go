package main

import (
	"fmt"
)

type task struct {
	title string
	body  string
}

func newTask(title, body string) task {
	return task{
		title: title,
		body:  body,
	}
}

func (c *task) setTitle(title string) {
	c.title = title
}

func (c *task) setBody(body string) {
	c.body = body
}

func (c *task) getTask() string {
	return fmt.Sprintf("%s\n%s", c.title, c.body)
}

func main() {
	a := newTask("N th Prime Number", "You must find n(1 <= n <= 10^9) th prime number in 1 sek")
	fmt.Println(a.getTask())
}
