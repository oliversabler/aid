#!bin/bash/

if ! go version 2>&1 > /dev/null; then
    echo "Go is required to build this application"
    echo "exiting..."
    exit 1
fi

go build -o aid

application_dir="${HOME}/.go_cli"
if [ ! -d $application_dir ]; then
    echo "Application directory '$application_dir' did not exist"
    echo "Creating directory '$application_dir'"
    mkdir $application_dir
fi

mv ./aid $application_dir

echo "alias aid=$application_dir/aid" >> ~/.zprofile

source "${HOME}/.zprofile"
