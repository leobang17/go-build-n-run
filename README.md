# Build and Run shortcut script for Go

A simple shortcut script to sequeantially **build and execute standalone go program**.
You should run the command on the root directory of the go project, _which must contains a `main.go` file_.

## Requirements

- Go programming language must be installed on your system.

## Get staterd

1. Clone project & move to project directory

```
git clone https://github.com/leobang17/go-build-n-run.git
```

```
cd ./go-build-n-run
```

2. Grant permission to the shell script file to be executed

```sh
chmod +x build_and_move.sh
```

3. Run shell script to build 'go-build-n-run' application and set `gobr` as a global usable command.

```sh
./build_and_move.sh
```

- this command contains `sudo` keyword to move go binary file to `/usr/local/bin`.

4. Use `gobr` command on every standalone go project you want.

```
gobr [<optional> build file destination]
```
