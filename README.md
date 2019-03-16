### Known issues

#### REDIRECT_URI
The Golang client expects a REDIRECT_URI, even though the
one-time Server auth flow involves a JS client retrieving auth, forwarding it
back to server, then the server simply does an exchange for access token
(and hence no REDIRECT_URI needed). The client credentials may not even need
to have such a REDIRECT_URI defined, but to make the Go client happy,
specify REDIRECT_URI as the Authorized JS origin defined in the client
credentials in GCP.

### Tools

#### Gin

https://github.com/codegangsta/gin

Gin watches for changes to .go in the source directory and rebuilds.
It is a proxy server which redirects traffic to the app's port.

```
go get github.com/codegangsta/gin
gin -a 8080 --path . --immediate
```

### Git

```
$ git remote add origin https://github.com/alephnan/cetea.git
$ git push -u origin master
```

```
$ git diff HEAD
$ git diff origin/master
```

```
$ git add <file> # stage file
$ git reset <file> # unstage file
$ git commit -m "message"
$ git commit --amend
```

### Golang
```
$ cd $GOPATH/src/github.com/alephnan/cetea
$ go build -o cetea
$ ./cetea
```

### Docker

List images
```
$ docker images
```

Remove image
```
$ docker rmi -f <image_name>
```

List the containers
```
$ docker containers ls
```

Build the docker image
```
$ docker build -t cetea .
```

Run the docker image
```
$ docker run -p 8081:8080 -ti cetea
```

### Unix

```
$ vim ~/bashrc
$ source ~/.bashrc
$ export $GOPATH=~/go_workspace
$ export $GOBIN=$GOPATH/bin
```

```
$ ps ax | grep java
$ top
$ netstat -an | grep 80
```
### Useful links

*  [Deploying Go web app on GCP](https://medium.com/martinomburajr/building-a-go-web-app-from-scratch-to-deploying-on-google-cloud-part-0-intro-a6bf26972ce5)
*  [Error handling in Go](https://blog.golang.org/error-handling-and-go)
*  [vs io.WriteString vs responseWriter.Write vs fmt.Fprintf](https://stackoverflow.com/questions/37863374/whats-the-difference-between-responsewriter-write-and-io-writestring)
*  [How to write go code](https://golang.org/doc/code.html)
*  [Git reset vs checkout vs revert](https://www.atlassian.com/git/tutorials/resetting-checking-out-and-reverting)
*  [Passing arguments with struct](https://stackoverflow.com/questions/26211954/how-do-i-pass-arguments-to-my-handler)