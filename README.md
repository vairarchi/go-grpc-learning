# go-grpc-learning
The companion repo for the GRPC course I am doing on Udemy


Commands used:

Generate proto code:

protoc -I greet/proto --go_out=. --go_opt=module=github.com/vairarchi/go-grpc-learning  --go-grpc_out=. --go-grpc_opt=module=github.com/vairarchi/go-grpc-learning  greet/proto/dummy.proto 

To avoid using such a long command we will make use of script at - https://raw.githubusercontent.com/Clement-Jean/grpc-go-course/master/Makefile

So after hsving the makefile we can simple type the command:
    make greet
    make help // to see what all commands are associated with make


