### Docker 101

* Uses Linux features like cgroups and namespaces
* cgroups - This is a Linux kernel feature that isolates the resource usage (CPU, memory, disk, I/O, network etc)
* namespaces - every container lives in its own namespace (mount, network interfaces, hostname)
* Docker engine runs as a daemon, comes with a CLI and has a REST API.

(docker-cheat-sheet) [https://github.com/wsargent/docker-cheat-sheet]

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