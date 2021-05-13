package main

import sockets "snippets/sockets/socket"

func main() {
	go sockets.Server()
	sockets.Client()
}
