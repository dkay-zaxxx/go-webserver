# My simple GO webserver

This is a simple webserver that can be hosted on a machine that has golang installed. It also provides access logging similar to apache and nginx. The purpose of this is eliminate the need for installing apache or nginx on a machine that you just want to see web content without needing to go through the setup of an apache or nginx.

# How to get started - Linux
First start by cloning the repository onto your machine
`git clone https://github.com/dkay-zaxxx/go-webserver`

Once the repository has been cloned successfully onto your machine, you may navigate into the directory and create a folder logs and web content
I would suggest create these two folders which are referenced by main.go `logs` and ` public_html` 

# Starting the webserver

`go run main.go`

# Your GO webserver is now running!
You can access the webserver via the IP of your machine on port 8888
`http://localhost:8888`
