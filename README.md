# Go Fuse-day: Distributed MSA Soccer #

## Abstract ##
In this Go workshop we'll build and execute a distributed simulation of a soccer game.
We'll use Go builtin concurrency concepts such as goroutines and channels, [Redis](https://redis.io/) as a
message broker and [Vice](https://github.com/matryer/vice) library to connect internal 
go-channels with Redis.

This workshop is written as a workbook. Follow the instructions in the 
`TODO Challenge` comments in order to complete the challenge.


## Dependencies ## 
This workshop uses go modules. Make sure your imports are synced in the IDE and go.mod file is updated on build

## Getting Started ##

### Prerequisites ###

In order to get your hands dirty in this workshop, 
make sure you got the following prerequisites set up:
* Redis server
* GO SDK
* Go IDE
* GIT

### Installing ###

#### Redis Server ####

* **MacOS** - Assumed you have [Homebrew](https://www.howtogeek.com/211541/homebrew-for-os-x-easily-installs-desktop-apps-and-terminal-utilities/) installed, 
find installation instructions [here](https://medium.com/@petehouston/install-and-config-redis-on-mac-os-x-via-homebrew-eb8df9a4f298).
* **Windows** - Installation instructions [here](https://redislabs.com/ebook/appendix-a/a-3-installing-on-windows/a-3-2-installing-redis-on-window/).

Once installed, launch redis-server with the default settings using this command:
```$xslt
redis-server
```

#### Go SDK ####
Install the latest Go SDK (or v1.12 at least) following these [instructions](https://golang.org/doc/install).

#### GIT ####
1. Install GIT (if not installed already) following these [instructions](https://www.atlassian.com/git/tutorials/install-git).
2. clone this project:
```$xslt
git clone git@github.com:tikalk/go-distribution-fuzeday.git
```

#### GO IDE ####
In this workshop we'll use [GoLand](https://www.jetbrains.com/go/). Feel free to use your preferred IDE if any.
 
### Dependencies ### 
This workshop is dependent on a several libraries and uses [go modules](https://github.com/golang/go/wiki/Modules) 
to manage them. Make sure your project contains a `go.mod` file in any step, and that it is updated on build
or IDE sync.


### Execution ###
Whether you run the examples from the code or the provided executable, these are the supported CLI commands:
```$bash
./go-distribution-fuzeday join
./go-distribution-fuzeday throw
./go-distribution-fuzeday simulate
./go-distribution-fuzeday fuzedaydisplay
```

To get more info about global and command-specific flags, just use this command to get the help documentation on the console:
```$bash
./go-distribution-fuzeday help
```

## License ##
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details