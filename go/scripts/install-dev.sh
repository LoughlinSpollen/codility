#! /usr/bin/env zsh

if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "installing privatus dev env on MacOS"
else
    echo "command install-dev only supports MacOS"
    exit
fi   

eval xcodebuild -version 2> /dev/null 2>&1
if [ ! $? -eq 0 ] ; then
    echo "Please Install XCode first - bye"
    exit
fi

echo "Installing go lang"
brew install go

echo "" >> ~/.zshrc
GOPATH=$(go env GOPATH)
echo "wrting 'GOPATH=${GOPATH}/go' to ~/.zshrc"
echo "export GOPATH=${GOPATH}/go" >> ~/.zshrc
echo 'export PATH="$PATH:$GOPATH"' >> ~/.zshrc
unset GOPATH

GOROOT=$(go env GOROOT)
echo "wrting 'GOROOT=${GOROOT}' to ~/.zshrc"
echo "export GOROOT=${GOROOT}" >> ~/.zshrc
unset GOROOT

echo "wrting 'GOBIN=$GOPATH/bin' to ~/.zshrc"
echo "export GOBIN=$GOPATH/bin" >> ~/.zshrc
echo 'export PATH="$PATH:$GOROOT/bin"' >> ~/.zshrc
echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.zshrc

# reload shell
exec $SHELL -l

LINT_VERSION="v1.30.0"
eval golangci-lint 2> /dev/null 2>&1
if [ $? -eq 0 ] ; then
    echo "Linter installed"
else
    echo "Installing Linter"
    eval brew install golangci-lint
fi

echo
eval go mod init
eval go mod tidy
# eval go mod vendor
eval go mod download
eval go mod verify
