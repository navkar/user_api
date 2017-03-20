### Docker 101

* Uses Linux features like cgroups, chroot and namespaces
* cgroups - This is a Linux kernel feature that isolates the resource usage (CPU, memory, disk, I/O, network etc)
* namespaces - every container lives in its own namespace (mount, network interfaces, hostname)
* Docker engine runs as a daemon, comes with a CLI and has a REST API.
* Docker commands list [docker-cheat-sheet]! (https://github.com/wsargent/docker-cheat-sheet)

![Docker](https://github.com/navkar/user_api/blob/master/images/Docker-vs-Virtualization.jpg)

![Docker Architecture](https://github.com/navkar/user_api/blob/master/images/docker-architecture.png)

![Cheat-sheet](https://github.com/navkar/user_api/blob/master/images/docker_cheat_sheet.png)

### How to setup the golang api with postgres as backend

### What is a docker hub?

* a cloud for sharing container images and automating workflows

#### Building the docker image

* Run the postgres container

```bash

navkar$ docker start userapi_db_1
```

#### Building golang api container

* Run the following command

```bash

navkar$ docker build -t go_api .
...
Successfully built 5e1ca53de33c

navkar$ docker images
REPOSITORY                                               TAG                 IMAGE ID            CREATED             SIZE
go_api                                                   latest              5e1ca53de33c        8 minutes ago       739 MB
postgresql                                               latest              b5940e4e2b8e        32 hours ago        398 MB
golang                                                   onbuild             5e66373f9a5d        2 weeks ago         697 MB

navkar$ docker ps -a
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                    PORTS                    NAMES
c8efaf4a74e0        postgresql          "/usr/lib/postgres..."   49 minutes ago      Up 40 minutes             5432/tcp                 userapi_db_1

navkar$ docker run --env-file ./env.list --publish 8080:8080 --name go_api --link userapi_db_1:postgresql go_api
+ exec app
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/users             --> main.PostUser (4 handlers)
[GIN-debug] GET    /api/v1/users             --> main.GetUsers (4 handlers)
[GIN-debug] GET    /api/v1/users/:id         --> main.GetUser (4 handlers)
[GIN-debug] PUT    /api/v1/users/:id         --> main.UpdateUser (4 handlers)
[GIN-debug] DELETE /api/v1/users/:id         --> main.DeleteUser (4 handlers)
[GIN-debug] Listening and serving HTTP on :8080

navkar$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
3726245ea441        go_api              "go-wrapper run"         3 minutes ago       Up 3 minutes        0.0.0.0:8080->8080/tcp   go_api
c8efaf4a74e0        postgresql          "/usr/lib/postgres..."   39 minutes ago      Up 31 minutes       5432/tcp                 userapi_db_1
```

#### Container monitoring

```bash

navkar$ docker stats <id>
```

### Docker compose file

* The docker compose file consists of create and running the go_api and the database container
* The command docker-compose up executes the file docker-compose.yml file

```bash

navkar$ docker-compose up
Creating network "userapi_default" with the default driver
Creating user_api_db
Creating user_api
Attaching to user_api_db, user_api
user_api_db | 2017-03-20 03:00:28.915 UTC [6] LOG:  database system was interrupted; last known up at 2017-03-14 15:55:27 UTC
user_api_db | 2017-03-20 03:00:29.548 UTC [6] LOG:  database system was not properly shut down; automatic recovery in progress
user_api_db | 2017-03-20 03:00:29.560 UTC [6] LOG:  redo starts at 0/17843C8
user_api_db | 2017-03-20 03:00:29.560 UTC [6] LOG:  record with zero length at 0/1784408
user_api_db | 2017-03-20 03:00:29.560 UTC [6] LOG:  redo done at 0/17843C8
user_api_db | 2017-03-20 03:00:29.560 UTC [6] LOG:  last completed transaction was at log time 2017-03-14 15:55:27.964861+00
user_api_db | 2017-03-20 03:00:29.647 UTC [6] LOG:  MultiXact member wraparound protections are now enabled
user_api_db | 2017-03-20 03:00:29.650 UTC [10] LOG:  autovacuum launcher started
user_api_db | 2017-03-20 03:00:29.651 UTC [1] LOG:  database system is ready to accept connections
user_api | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
user_api |  - using env:	export GIN_MODE=release
user_api |  - using code:	gin.SetMode(gin.ReleaseMode)
user_api | 
user_api | [GIN-debug] POST   /api/v1/users             --> main.PostUser (4 handlers)
user_api | [GIN-debug] GET    /api/v1/users             --> main.GetUsers (4 handlers)
user_api | [GIN-debug] GET    /api/v1/users/:id         --> main.GetUser (4 handlers)
user_api | [GIN-debug] PUT    /api/v1/users/:id         --> main.UpdateUser (4 handlers)
user_api | [GIN-debug] DELETE /api/v1/users/:id         --> main.DeleteUser (4 handlers)
user_api | [GIN-debug] Listening and serving HTTP on :8080
user_api | connection string: host=user_api_db user=docker dbname=postgres sslmode=disable password=docker
user_api | 
user_api | (/go/src/app/main.go:39) 
...
```