
export SFTPHost=127.0.0.1:5000
export SFTPPath=/wmsshipment


export SFTP_HOST=10.7.57.162
export SFTP_PORT=22
export SFTP_USERNAME=mock
export SFTP_PASSWORD=mock
export SFTP_BASEPATH=ttb/dev



run-main:
	PORT=3000 go run cmd/main.go

run-test:
	golangci-lint run
	go test -race -cover ./... -count=1 -failfast