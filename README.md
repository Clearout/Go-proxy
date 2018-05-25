# Go-proxy

Proxy to add CORS headers to any request, and also change request type to GET for any incoming requests.

To build app
``` bash
cd src\proxy
go install
```
or just run the .exe if on windows.

Runs on localhost:3000 and connects to DSB's elasticsearch server http://ryzen2.utv.lokal:9200
This connection only works if you are on DSB's network or VPN.
