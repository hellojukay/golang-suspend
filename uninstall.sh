#!/usr/bin/env bash

systemctl disable  web-suspend
systemctl stop  web-suspend
rm -f  /usr/lib/systemd/system/web-suspend.service
rm /usr/bin/web-suspend