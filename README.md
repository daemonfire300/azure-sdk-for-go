# Fork of the Microsoft Azure SDK for Go


# Installation

    go get -d github.com/daemonfire/azure-sdk-for-go/storage

# Purpose of this Fork

The purpose of this fork is to extend the BlobStorage capacities of this package. Most of the ground work is done by `client.go` and `blob.go` but is not exposed to the user as interface. I will try to gradually add more functions and easier ways for controlloing Blob objects, such as specifying the `Content-Type:` when uploading a block.

# Usage (of the original package)

Read Godoc of the repository at: http://godoc.org/github.com/Azure/azure-sdk-for-go/

The client currently supports authentication to the Service Management
API with certificates or Azure `.publishSettings` file. You can 
download the `.publishSettings` file for your subscriptions
[here](https://manage.windowsazure.com/publishsettings).


# License

This project is published under [Apache 2.0 License](LICENSE).
