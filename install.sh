#!/bin/bash

APP_TARGET_DIR="${HOME}/.go_cli"
PROFILE="${HOME}/.zprofile"

if ! go version > /dev/null 2>&1; then
    echo "Go is required to build this application"
    exit 1
fi

go build -o aid

if [ ! -d "$APP_TARGET_DIR" ]; then
    mkdir "$APP_TARGET_DIR"
fi

mv ./aid "$APP_TARGET_DIR"

if ! grep -q "alias aid=" "$PROFILE" 2>/dev/null; then
    echo "alias aid=$APP_TARGET_DIR/aid" >> "$PROFILE"
fi

echo "Installed to $APP_TARGET_DIR/aid"
echo "Run 'source ~/.zprofile' to use the 'aid' command"
