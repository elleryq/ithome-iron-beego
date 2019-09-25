# ithome-iron-beego

## Build

After clone repository, just enter directory then type `go build`

## Docker

Build
```
docker build -t elleryq/ithome-iron-beego:0.1.0
```

Run
```
docker run --name beego-mysql -e MYSQL_DATABASE=hellodb -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql:5.7
docker run -p 8080:8080 -e ORM_DRIVER=mysql -e "ORM_SOURCE=root:my-secret-pw@tcp(beego-mysql:3306)/hellodb?charset=utf8" --link beego-mysql elleryq/ithome-iron-beego:0.1.0
```
