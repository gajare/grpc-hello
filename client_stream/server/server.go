package main

import (
	streaming "client-streaming/proto/generated"
)

type MyServer struct {
	streaming.UnimplementedClientStreamServer
}

func main() {

}
