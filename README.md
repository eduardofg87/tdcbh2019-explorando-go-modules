# TDC BH 2019 - Explorando Go Modules
Palestra Trilha GO - TDC BH 2019 - Explorando Go Modules
Palestrante: Eduardo Figueiredo Gonçalves - @eduardofg87

## Outside Docker

### Export GO111MODULE
`export GO111MODULE=on`

### Cleaning all the trash
`rm go.mod go.sum`

### Initiate go mod
`go mod init`

### Creating vendor folder
`go mod vendor`

### Build the app
`go build`

### Run the app
`./tdcbh2019-explorando-go-modules`


# Using Docker

## Build
`sudo docker build -t tdcbh2019-explorando-go-modules .`

## RUN
`sudo docker run -p 8080:8080 tdcbh2019-explorando-go-modules`