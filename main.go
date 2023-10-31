package main

import (
	login "project_yd/login"
	server "project_yd/server"
	"sync"
	//itemcreateevent "github.com/heroiclabs/nakama/v3/nurhyme_common/ItemCreateEvent"
)

func RegistRpc() {
	login.RegistLoginRpc()
}

func main() {
	var waitGroup sync.WaitGroup

	//-- GoRoutine Count +1
	waitGroup.Add(1)
	go func() {
		//-- GoRoutine Count -1
		defer waitGroup.Done()
		//-- Start Grpc server
		server.StartGrpcServer()
	}()

	RegistRpc()
	server.StartDBConnection()

	waitGroup.Wait()
	select {}
}
