gRPC is roughly 7 times faster than REST when receiving data & roughly 10 times faster than REST when sending data for this specific payload


https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go


https://grpc.io/docs/languages/go/quickstart/


https://grpc.io/docs/protoc-installation/

https://github.com/casey/just

Here are some notes I have when scanning your project:
- Use spf13/cobra to create your cli.
https://cobra.dev/
- Don't use println for logging using a dedicated library zap, logrus...
- https://github.com/uber-go/zap
- Unit tests are missing
- Project structure can be improved
- File names why do you use multiple main?
- You can improve your readme file
some useful links
https://go.dev/doc/effective_go
https://github.com/evrone/go-clean-template


https://medium.com/tcp-ip/kerberos-authentication-fc41aa320a01


https://github.com/jcmturner/gokrb5
https://www.youtube.com/watch?v=5N242XcKAsM
https://github.com/dosvath/kerberos-containers
https://pkg.go.dev/gopkg.in/jcmturner/gokrb5.v8#section-readme