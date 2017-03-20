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

### AWS docker hub

* Create a repository in Amazon ECS container service
* Use aws-cli to upload images using the repository URI

```bash

navkar$ aws ecr get-login --region us-west-2
navkar$ docker login -u AWS -p AQECAHj6lc4... -e none 707944905670.dkr.ecr.us-west-2.amazonaws.com
navkar$ docker push 707944905670.dkr.ecr.us-west-2.amazonaws.com/go_api:latest
The push refers to a repository [707944905670.dkr.ecr.us-west-2.amazonaws.com/go_api]
01bbfa17fb8b: Pushed 
06e07303de7e: Pushed 
7d9c10f89b29: Pushed 
b8882fed7c5a: Pushed 
3270d804862e: Pushed 
5426f7c4a6c8: Pushed 
0d079ee8dc4b: Pushed 
12c8823c9d5e: Pushed 
56c5d34cc6a5: Pushed 
b7590724968a: Pushed 
38f9a02a0390: Pushed 
72917b86e825: Pushed 
100396c46221: Pushed 
7b4b54c74241: Pushed 
d17d48b2382a: Pushed 
latest: digest: sha256:56c4597339014df82a31b79e9e1854f81629c7c60f50ca314d0b80103d3471a7 size: 3459
```

#### remove all containers based on a pattern

```bash

$ docker ps -a | grep "cfd11f2a0955" | awk '{print $1}' | xargs docker rm
````

#### References

[remove-docker-images-containers-and-volumes]
(https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes)

[authorization-token-has-expired] (http://stackoverflow.com/questions/41379808/authorization-token-has-expired-issue-aws-cli-on-macos-sierra)