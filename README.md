
# GIN (GO Framework)


## Project setup

Add Compile Daemon (https://github.com/githubnemo/CompileDaemon):
```bash
  go get github.com/githubnemo/CompileDaemon
```
then install it with :

```bash
  go install github.com/githubnemo/CompileDaemon
```


Add Go Env (https://github.com/joho/godotenv):
```bash
  go get github.com/joho/godotenv
```

Add GIN Web Framework(https://gin-gonic.com/docs/quickstart/):
```bash
  go get -u github.com/gin-gonic/gin
```


Add GORM(https://gorm.io/index.html):
```bash
  go get -u gorm.io/gorm 
```
and after that add driver
```bash
  go get -u gorm.io/driver/DATABASE
```




To run this project use tap in cmd:
```bash
  CompileDaemon -command="./go-crud"
```






#### Test API with postman

you have example for http request for testing your end-point :

add file GO-crud.postman_collection.json to your postman locally 



## Authors

- [@MiBoutaj](https://github.com/MiBoutaj)
