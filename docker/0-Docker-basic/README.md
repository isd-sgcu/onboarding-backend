# What is Docker ?

- Virtualization software
- Makes **developing** and **deploying applications** much easier
- Packages application with all the necessary dependencies, configuration, system tools and runtime
- Portable artifact, easily shared and distributed

> **Container**: a standardized unit, that has everything the application needs to run

### What problems solved ?

#### The development process before containers?

- Each developer needs to **install and configure** all services **directly on their OS** on their local machine

> - **Installation process different** for each OS environment
> - **Many steps**, where something can go wrong

for example, If your app uses 10 services, each developer needs to install these 10 services

#### Development process with ==Container== ?

with container you don't have to install any of these services directly

- Own **isolated environment**
- Postgres packaged with all dependencies and configs
- Start service as a Docker container using a **1 Docker command**
- **Command same for all services and all OS**

It **Standardizes process** of running any service on any local dev environment

- **Easy to run different versions** of same app without any conflicts
- **No configuration on the server**

---

# Docker vs Virtual Machines

- Why is Docker so widely used?
- What **advantages** does it have over virtual machines?
- What is the **differences** ?

  - Docker: virtualizes the applications layer
  - VM: has the applications layer and its own kernel
  - Docker images much smaller

- Docker Desktop: uses a **Hypervisor layer** with a lightweight Linux distro

---

# Images vs Containers

package or artifact that we produce with Docker is called a **Docker image**
Docker image

- is an executable application artifact
- includes app source code, but also **complete environment configuration**
- Immutable **template** that defines how a container will be realized

Docker container

- Actually **starts the application**
- A **running instance** of an image
  - You can run **multiple containers from 1 image**

Command
`docker ps`: List running containers

---

# Public and Private Registries

> How to we get docker images ?
> from Docker registries

- Docker registries: A **storage** and distribution system for Docker images
- Docker hosts **one of the biggest Docker Registry** called **Docker Hub**
- Docker Official Images
  - A dedicated **team responsible for reviewing and publishing all content** in the Docker Official Images repositories
  - Works in **collaboration with software maintainers, security experts**

---

# Port Binding

- Application inside container runs in an **isolated Docker network**
- This allows us to run the same app running on the **same port multiple times**
- We need to **expose** the container port **to the host**

![[port-binding]]

`-p {HOST_PORT}:{CONTAINER_PORT}`

```sh
docker run -d -p 9000:80 nginx:1.23
```

Standard to use the same port on your host as container is using

---

# Start and Stop container

- `docker stop <container-id>` = stop running container
- `docker start <container-id>` = start one or more stopped containers

---

# Public and Private Docker Registries

**Private**

- You need to **authenticate** before accessing the registry
- All big cloud provider offer private registries: Amazon ECR, Google Container Registry, etc
- Docker hub

---

# Registry vs Repository

**Docker Registry**

- A **service** providing storage
- Can be **host by a third party** like AWS
  **Docker Repository**
- **Collection of related images** with same name but different versions

---

# Building Own Docker images

- we need to create a "definition" of how to build an **image** from our application
- `Dockerfile` is a text document that contains commands to assemble an image
- Docker can them build an image by reading those instructions

### Structure of Dockerfile

- Dockerfiles start from a parent image or ==base image==
- It's a Docker image that your image is based on

**FROM**

- Dockerfile **must begin** with a `FROM` instruction
- Build this image from the specified image

**COPY**

- **Copies** files or directories from `<src>` and adds them to the filesystem of the container at the path `<dest>`
- While `RUN` is execute in te container, `COPY` is executed on the host

**WORKDIR**

- **Sets the working directory** for all following commands
- Like changing into a directory: `cd .. `

**RUN**

- Will execute any command in a shell **inside** the container environment

**CMD**

- The instruction that is to be executed when a Docker container starts
- There can **only** be **one "CMD" instruction** in a Dockerfile

Example

```Dockerfile
FROM node:19-alpine

# COPY <src> on our machine <dest> in the container
COPY package.json /app/
COPY src /app/

WORKDIR /app

RUN npm install

CMD ["node", "server.js"]
```

---

# Build docker image

```sh
# docker build -t <tag> <path-of-dockerfile>
docker build -t node-app:1.0 .
```

run

```sh
docker run -d -p 3000:3000 node-app:1.0
```

---

# Dockerize your own application

1. Write Dockerfile
2. Build Docker Image
3. Run as Docker container

---
