build:
	GOOS=linux GOARCH=arm go build -o pkg/hue-control_arm
	go build -o pkg/hue-control_amd64
