## Messaging queue code sample

### Getting started
1. run `$ dep ensure` to get all required dependencies. 
2. run `$ go build` .
3. run `$ ./userq httpsrv` to run http server.
4. open a new terminal and run `$ ./userq msgqsrv` to run http server.

No you are ready to open postman and post to http://localhost:3000

```
{
    "firstname": "ayman",
    "lastname": "hassan",
    "address": "enkhalifapro@mail.com",
    "gender":"malex",
    "timestamp": 1537599615
}
```