# sftptest

This is a program for testing SFTP connection. I created this program to monitor SFTP connectivity with [Monit](https://mmonit.com/monit/).

## Getting Started

Make sure port number is specified with hostname.

```shell
git clone git@github.com:hongster/sftptest.git
cd sftptest
go mod tidy
go build
./sftptest ftp.example.com:22 userFoo passwordBar
```