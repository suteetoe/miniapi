# Mini Api

Log File Store in : `/var/log/vfapi/`

Log File pattern : `app-yyyy-MM-dd.log`

Example  : `/var/log/vfapi/app-2024-03-05.log`

## Run App

run app with docker

```yaml
version: '3.9'

services:
  miniapipure:
    image: ghcr.io/suteetoe/miniapi-pureapp:main
    ports:
     - 8080:8080
```


*run api*

```
go run main.go
```


*send request*

```wget
### List Demo Data
GET http://localhost:8080/demo


### Get Demo Code 001
GET http://localhost:8080/demo/002


### Get XDemo should return 404

GET http://localhost:8080/xdemo

```

