#!/bin/sh
cd /app
/go/bin/dlv exec /bin/discord-clone --headless --log -l 0.0.0.0:2345 --api-version=2
