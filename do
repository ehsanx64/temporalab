#!/bin/bash

cmd::temporal-start() {
	temporal server start-dev
}

cmd::temporal-start-wui-8080() {
	temporal server start-dev --ui-port 8080
}

cmd::temporal-start-docker() {
	# UI would be accessible from host at http://localhost:8233/
	docker run --rm -p 7233:7233 -p 8233:8233 temporalio/temporal server start-dev --ip 0.0.0.0
}

cmd::temporal-mac-install() {
	brew install temporal
}

cmd::help() {
    echo "cmd::help"
}


set +x
cmd=$1

# If no command given run the help
if [ "$cmd" = "" ]; then
    cmd::help
    exit 1
fi

if declare -F "cmd::$cmd"; then
    { set +x; } 2> /dev/null
    eval "cmd::$cmd $@"
    { set -x; } 2> /dev/null
else
    echo "No such command: $cmd"
    exit 1
fi

