# DistriAI-Index-Solana

## System Requirements
- Linux-amd64

## Build
Requires Go1.20 or higher.
```
GOOS=linux GOARCH=amd64 go build
```
If all goes well, you will get a program called `distriai-index-solana`.

## Run
### Step 1: Prepare configuration file
- Copy `/config` folder to where `distriai-index-solana` program locate.
- Edit Database configuration in `./config/config.yml`.
```
Server:
  Mode: release
  Port: 8800

Database:
  Host:
  Port: 3306
  UserName:
  Password:
  Database: distriai-index-solana

Mailbox:
  Host:
  Port: 25
  Username:
  Password:

Chain:
  Rpc:
  ProgramId:
```

### Step 2: Start the distriai-index-solana service
- New a screen window
```
screen -S distriai-index-solana
```
- start service
```
./distriai-index-solana
```
When the service is started, the `machines` and `orders` tables in database will be cleared, the latest data on the chain will be pulled.

## Stop
- Attach the screen window
```
screen -r distriai-index-solana
```
- stop service

`CTRL +  C`
