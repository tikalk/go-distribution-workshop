# Go Workshop: Distributed MSA Soccer #

## Abstract ##
In this Go workshop we'll build and execute a distributed simulation of a soccer game.
We'll use Go builtin concurrency concepts such as goroutines and channels, [Redis](https://redis.io/) as a
message broker and [Vice](https://github.com/matryer/vice) library to connect internal 
go-channels with Redis.

This workshop is divided to 5 development phases, each one provides an initial state
as separated branch, with a challenge to develop. At the end of this workshop we'll
sit back and enjoy the soccer game in which every participant plays the role of a single
player of field.

## Phases ##
* `phase_1-concurrency_setup` - have a single player kick the ball locally using internal go channels
* `phase_2-distribution_setup` - prepare channels for distribution
* `phase_3-vice_integration` - running simulation locally using [Redis](https://redis.io/) and [Vice](https://github.com/matryer/vice)
* `phase_4-display_service` - viewing game on local UI with display service
* `phase_5-lets_rock` - connecting to a shared Redis server and playing together!

## Dependencies ## 


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
1. Install the latest Go SDK following these [instructions](https://golang.org/doc/install).
2. If not already created, create a folder named `go` under your user folder (`~/go` on MAC and Linux, `Users/[user]/go` on Windows)
3. Set both `GOROOT` and `GOPATH` environment variables to the folder above.
4. Inside the folder specified in `GOPATH` create a `src` folder and clone this git project there

#### GIT ####
1. Install GIT (if not installed already) following these [instructions](https://www.atlassian.com/git/tutorials/install-git).
2. cd to the `src` folder in your `GOPATH`
3. clone this project by this line
```$xslt
git clone git@github.com:tikalk/go-distribution-workshop.git
```

#### GO IDE ####
In this workshop we'll use [GoLand](https://www.jetbrains.com/go/).Feel free to use your preferred IDE if any.
 
### Dependencies ### 
This workshop is dependent on a several libraries. In order to install them all run this line:

MacOS:
```$bash
chmod +x dep.sh
./dep.sh
```

Windows:
```$bash
dep.bat
```

### Execution ###
Whether you run the examples from the code or the provided executable, these are the supported CLI commands:
```$bash
./go-distribution-workshop join
./go-distribution-workshop throw
./go-distribution-workshop simulate
./go-distribution-workshop display
```

To get more info about global and command-specific flags, just use this command to get the help documentation on the console:
```$bash
./go-distribution-workshop help
```

## License ##
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details