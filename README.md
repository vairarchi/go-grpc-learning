# go-grpc-learning
The companion repo for the GRPC course I am doing on Udemy


Commands used:

Generate proto code:

protoc -I greet/proto --go_out=. --go_opt=module=github.com/vairarchi/go-grpc-learning  --go-grpc_out=. --go-grpc_opt=module=github.com/vairarchi/go-grpc-learning  greet/proto/dummy.proto 

To avoid using such a long command we will make use of script at - https://raw.githubusercontent.com/Clement-Jean/grpc-go-course/master/Makefile

So after hsving the makefile we can simple type the command:
    make greet
    make help // to see what all commands are associated with make

    make greet
    (after the executables are made)
    ./bin/greet/<client/server>


>Note: To start writing the server code
 - you have to make use of the generated .pb.go file 
 - search for <name>serviceServer interface in the file
 - in that interface you will get the function definition of the function that needs to be implemented in the server file
 - For ex: if you are adding `GreetWithDeadline` rpc, after generating the greet_grpc.pb.go file, search for `GreetServiceServer` and copy the function definition `GreetWithDeadline(context.Context, *GreetRequest) (*GreetResponse, error)` which needs to be implemented on the server side


>For SSL:
- get the file ssl.sh from https://raw.githubusercontent.com/Clement-Jean/grpc-go-course/master/ssl/ssl.sh and place it in the folder named `ssl`
- Commands:
    - cd ssl
    - chmod +x ssl.sh
    - ./ssl.sh 
- Since we will now have all the required files we can use them is `server` and `client` of `greet`

