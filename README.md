Mass-Telnet command 
=============

This is a program to connect to multiple telnet hosts and execute a set of commands on them.

The format of the credential file is
IP:username:password

Modify line 61 to change the command executed on telnet servers.

Compilation
-------------

Compiling it easy!

Grab the depdendency: 
```bash
$ go get github.com/ziutek/telnet
```

### Lastly ###

```bash
$ go build main.go
$ ./main < file name >
```