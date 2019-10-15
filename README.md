# go-slack-bot-workshop
## Install Golang 1.13
###Macos
```shell script
curl -o golang.pkg https://dl.google.com/go/go1.13.1.darwin-amd64.pkg
sudo open golang.pkg
export GOROOT=/usr/local/go
export GOPATH=$HOME/Projects/Proj1
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go version
```
###Ubuntu
```shell script
wget https://dl.google.com/go/go1.13.1.linux-amd64.tar.gz
sudo tar -xvf go1.13.1.linux-amd64.tar.gz
sudo mv go /usr/local
export GOROOT=/usr/local/go
export GOPATH=$HOME/Projects/Proj1
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
go version
```
###Windows
```shell script
go to https://golang.org/dl/
download and install go1.13.1.windows-386.msi
go version
```


