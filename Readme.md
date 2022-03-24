# IBFD Standard Go Project Layout

## Overview
This is a basic layout for IBFD Go application projects. It's not an official standard defined by the core Go dev team.

With Go 1.14 [Go Modules][1] are finally ready for production. Use [Go Modules][1] unless you have a specific reason not to use them and if you do then you don’t need to worry about **$GOPATH** and where you put your project.

[1]: https://github.com/golang/go/wiki/Modules


## How to use
- Copy all files/directories under the **project-layout** folder to your **project directory**
- Replace the text _"ibfd.org/app"_ with _"ibfd.org/[your-project-name]"_ throughout the project
- Fill empty field values in _app.json_ <br>
e.g. edit _database credentials_, _server (host, port) info_, etc.
- Rename _"app.go"_ and _"app.json"_ by your project name. <br>
e.g. project name = _tarantula_ <br>
Rename - <br>
    _app.go_ as _tarantula.go_ <br>
    _app.json_ as _tarantula.json_
- Fix all required **TODO** operations in different files
- Execute commands
    1. run **go mod init ibfd.org/[your-project-name]**
    2. run **go get** or **go mod download** command
    3. run **go build**
    4. run **./[your-project-name]**
-  Update **Readme.md**

