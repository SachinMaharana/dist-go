#dist-go

docker build . -t hello-man

docker run hello-man

docker run -e NAME=Maharana hello-man

go test -v ./{variadic_test.go,variadic.go}
