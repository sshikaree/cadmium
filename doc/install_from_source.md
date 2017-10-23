# Installing Cadmium from source

Firstly, you have to get Golang compiler. Refer to https://golang.org
for information about that.

Next big dependency is a [therecipe's bindings to Qt5](https://github.com/therecipe/qt),
which should be installed manually.

After that clone this repo:

```
git clone https://github.com/pztrn/cadmium && cd cadmium
```

Install Go's dep tool:

```
go get -u github.com/golang/dep/cmd/dep
```

Install other dependencies:

```
dep ensure
```

*This documentation is incomplete*