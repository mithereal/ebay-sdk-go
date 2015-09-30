Go Ebay  API
===
ebay api written in go.

===

the examples directory conatains ready to go code snippets, you can make certain calls like fetch orders and getmessages to ebay servers, if you wish to contribute please follow the same schema as the examples.

usage: 

 Linux: 

Copy ebay credentials to ~/.ebayapi or set the following env varibles 

EBAY_APP_ID

EBAY_DEV_ID

EBAY_CERT_ID

EBAY_TOKEN

Copy example getorders to $gopath/src/ebaygetorders

go install ebaygetorders

cd to $gopath/bin

./ebayfetchorders -s 2015-07-16T20:34:44 -e 2015-08-14T20:34:44

Copy ebaygetorders to bin

Follow the same for other examples
 
