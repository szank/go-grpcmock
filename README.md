# go-grpcmock
My take on mocking gRPC services

This is an experimental tool to generate mock grpc services from the proto files. The goal of this project is to 
write a gRPC server that would expose all the methods defined in the proto files and in the MVP iteration: 
* start a mocking gRPC server
* start another gRPC service that for each "standard" grpc call exposes 
* a method to set the next N return values for the method being mocked 
* capture all the request messages for the grpc method 
* allows inspection of the arguments sent to the rpc being mocked


MVP non-goals 
* Make this tool work with other languages than go 

## Why am I writing this tool? 
I could not find any tool that would allow me to inspect the data sent to the rpc methods, nor verify if the method has
been actually called or not. 

## Why not extend gripmock?
Want to try doing it by myself