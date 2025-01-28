#!bin/bash/
# optional flag "addAlias" to add alias to your .zprofile
# . ./install.sh [addAlias]

APP_TARGET_DIR="${HOME}/.go_cli"
PROFILE="${HOME}/.zprofile"

if ! go version 2>&1 > /dev/null; then
    echo "Go is required to build this application"
    echo "exiting..."
    exit 1
fi

go build -o aid

if [ ! -d $APP_TARGET_DIR ]; then
    echo "Application directory '$APP_TARGET_DIR' did not exist"
    echo "Creating directory '$APP_TARGET_DIR'"
    mkdir $APP_TARGET_DIR
fi

mv ./aid $APP_TARGET_DIR

if [ "$1" == "addAlias" ]; then
    echo "alias aid=$APP_TARGET_DIR/aid" >> $PROFILE
    source $PROFILE
fi

