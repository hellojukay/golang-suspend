#!/usr/bin/env bash
go build -o web-suspend  main.go

install web-suspend /usr/bin/

install web-suspend.service /usr/lib/systemd/system/web-suspend.service
systemctl enable  web-suspend
systemctl start  web-suspend