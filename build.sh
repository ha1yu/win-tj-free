CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags "-s -w -H windowsgui" main.go

# go build -a -ldflags "-s -w -H windowsgui" main.go