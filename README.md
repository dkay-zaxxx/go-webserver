# My simple GO webserver

This is a simple webserver that can be hosted on a machine that has golang installed. It also provides access logging similar to apache and nginx. The purpose of this is eliminate the need for installing apache or nginx on a machine for the purpose of just viewing web content without needing to go through the setup of apache or nginx.

# How to get started - Linux
First start by cloning the repository onto your machine

`git clone https://github.com/dkay-zaxxx/go-webserver`

# Installing Golang
You can install golang from snapcraft

`snap install go --classic`

# Starting the webserver
Navigate to the main directory and then run

`go run main.go`

# Your GO webserver is now running!
You can access the webserver via the IP of your machine on port 8888

`http://localhost:8888`
