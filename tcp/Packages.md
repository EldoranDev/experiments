# Packages

## Header

| Bytes | Name    | Description            |
|-------|---------|------------------------|
| 1     | Version | Version of the Package |
| 1     | Length  | Length of the Header   |
| 1     | Message | Message Type           |

## Commands

| Bytes | Name    | Description                                    |
|-------|---------|------------------------------------------------|
| 1     | Command | Which Command                                  |
| 8     | Length  | Length of the Payload                          |
|       | Data    | Payload (length defined by Length field above) |
