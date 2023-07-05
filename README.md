# rclone-go
Rclone HTTP Api Client

## Usage

```go
package main

import (
    "github.com/hoglandets-it/gorclone"
)

client := gorclone.NewClient("http://rclone-server.com:1234", "username", "password")

// List Remotes
remotes, err := client.Run(ConfigListremotes{})

// Get Remote
remote, err := client.Run(ConfigGet{Name: "remote-name"})

// Create Remote
result, err := client.Run(ConfigCreate{
    Name: "remote-name",
    Type: "sftp",
    Parameters: map[string]string{
        "host": "ftp.server.com",
        "user": "sftp_user",
        "pass": "password",
    },
})

// Create Remote with JsonData
result, err := client.Run(ConfigCreate{
    JsonData: `{"name":"remote-name","type":"sftp","parameters":{"host":"ftp.server.com","user":"sftp_user","pass":"password"}}`,
})

// Run asynchronous operations returns the Job ID
result, err := client.Run(ConfigCreate{
    Name: "remote-name",
    Type: "sftp",
    Parameters: map[string]string{
        "host": "ftp.server.com",
        "user": "sftp_user",
        "pass": "password",
    },
    Async: true,
})
```

## API
All commands are implemented as structs with a name based on the path of the command, for example backend/commanbd becomes BackendCommand and Config/dumpconfig becomes ConfigDumpconfig.
If you pass data to the "JsonData" field of a struct, that data will be used instead of the manual fields. Some structs will only have the JsonData field specified, this data will always be passed into the body of the command even when the specific command doesn't have any parameters (e.g. ConfigListremotes)

You can find a complete documentation of the available commands and parameters at https://rclone.org/rc/#rc-commands
