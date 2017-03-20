### Docker 101

* Uses Linux features like cgroups, chroot and namespaces
* cgroups - a Linux kernel feature that isolates the resource usage (CPU, memory, disk, I/O, network etc)
* namespaces - every container lives in its own namespace (mount, network interfaces, hostname)
* Docker engine runs as a daemon, comes with a CLI and has a REST API.
* [Docker commands list](https://github.com/wsargent/docker-cheat-sheet)

![Docker](https://github.com/navkar/user_api/blob/master/images/Docker-vs-Virtualization.jpg)

![Docker Architecture](https://github.com/navkar/user_api/blob/master/images/docker-architecture.png)

![Cheat-sheet](https://github.com/navkar/user_api/blob/master/images/docker_cheat_sheet.png)

### a docker hub

* a cloud for sharing container images and automating workflows

#### Building the docker image

* Run the postgres container

```bash

navkar$ docker start userapi_db_1
```