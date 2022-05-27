# sftptest

This is a program for testing SFTP connection. I created this program to monitor SFTP connectivity with [Monit](https://mmonit.com/monit/).

## Usage

Make sure port number is specified with hostname.

```bash
go build
./sftptest ftp.example.com:22 userFoo passwordBar
```