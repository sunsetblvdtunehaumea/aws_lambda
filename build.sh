clear
ls -lrt
rm myFunction.zip bootstrap

GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go

zip myFunction.zip bootstrap


ls -lrt
