* Install go as per https://golang.org/doc/install
* Copy this directory into your  $GOPATH /src dir
* Setup your keys as either ENV vars or $HOME/.ebayapi
* ENV vars example

    export GOPATH="/home/username/Projects/go/"
    export EBAY_TOKEN=""
    export EBAY_APP_ID=""
    export EBAY_DEV_ID=""
    export EBAY_CERT_ID=""
	

* Run go install ebaygetmessages
* Go to $GOPATH /bin
* Run ex. ./ebaygetmessages -s 2015-07-16T20:34:44 -e 2015-08-14T20:34:44