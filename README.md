# udpsample

Sample UDP client/echo server commands implementation

* [echo server](/cmd/x/udpserver/main.go)

* [client](/cmd/x/udpclient/main.go)



## Build the project

Build server sample command :
```
go build -buildmode=exe github.com/hellgate75/udpsample/cmd/x/udpserver
```

Build client sample command :
```
go build -buildmode=exe github.com/hellgate75/udpsample/cmd/x/udpclient
```


## Install the project

Install server sample command :
```
go install -buildmode=exe github.com/hellgate75/udpsample/cmd/x/udpserver
```

Install client sample command :
```
go install -buildmode=exe github.com/hellgate75/udpsample/cmd/x/udpclient
```


## Commands usage

Start the UDP Echo server :
```
udpserver <host|ip mask|ip address|empty>:<port>
```

Execute the UDP client :
```
udpclient <host|ip address>:<port> <message>
```

Note: ```exit``` command will cause the server stop and exit



Enjoy the experience.


## License

The library is licensed with [LGPL v. 3.0](/LICENSE) clauses, with prior authorization of author before any production or commercial use. Use of this library or any extension is prohibited due to high risk of damages due to improper use. No warranty is provided for improper or unauthorized use of this library or any implementation.

Any request can be prompted to the author [Fabrizio Torelli](https://www.linkedin.com/in/fabriziotorelli) at the following email address:

[hellgate75@gmail.com](mailto:hellgate75@gmail.com)
