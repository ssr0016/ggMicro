package main

import "log"

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLogginService(svc)

	apiServer := NewApiServer(svc)
	log.Fatal(apiServer.start(":3000"))

}
