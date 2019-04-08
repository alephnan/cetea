# PC GCP

A plain and curated UI for GCP

# Mission

Organize Google Cloud Platform information and make it accessible, useful, and fun.

# Opinions

- Global search should be a cornerstone feature and powerful
- Product listings should be made by an intelligent recommendation system, or atleast data-driven, rather than
  politicaly
- THIS application should run on Google Cloud
- Experiments should run fast and not pre-requisite a PhD
- Dark mode.
- Keyboard accessible
- Mobile friendly
- "Below the fold" lazy loading of AJAX requests
- Will not aim for feature parity

# Execution

- Avoid Angular like the plague
- Avoid Observables / ReactiveJS
- API served with Golang
- CSS should tally less than 5kB
- Avoid fetching Google and custom fonts
- Avoid corporate blue

## Bets

- TypeScript
- Vue.JS

### TODO

- prevent CORs to API

### Known issues

#### REDIRECT_URI

The Golang client expects a REDIRECT_URI, even though the
one-time Server auth flow involves a JS client retrieving auth, forwarding it
back to server, then the server simply does an exchange for access token
(and hence no REDIRECT_URI needed). The client credentials may not even need
to have such a REDIRECT_URI defined, but to make the Go client happy,
specify REDIRECT_URI as the Authorized JS origin defined in the client
credentials in GCP.

### Dev guide

#### Mac

1. Add to ~/.bash_profile

   https://stackoverflow.com/questions/7780030/how-to-fix-terminal-not-loading-bashrc-on-os-x-lion

   ```
   [[ -s ~/.bashrc ]] && source ~/.bashrc
   ```

2. Install docker and docker-compose

3. Brew

   ```
   /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
   ```

4. NGINX

   ```
   brew install nginx
   ```

5. Install watch-rebuild tool

   https://github.com/codegangsta/gin

   Gin watches for changes to .go in the source directory and rebuilds.
   It is a proxy server which redirects traffic to the app's port.

   ```
   go get github.com/codegangsta/gin
   ```

#### GCP

1. Create a 'Web Application' client credential in Google Cloud Console.
   Whilist http://localhost:8080 and http://localhost:8081 as valid JS
   origins.

2. Download client secret json

3. Store client config in ./backend/config/client_secret.json

4. Define \$VUE_APP_GAPI_CLIENT_ID in ./frontend/.env.local

#### Run

#### Dev mode

1. Move frontend build to static directory

```
$ cp -rf frontend/dist backend/static
```

2. Flag Go server to serve statc

   ```
   $ export DEV=true
   ```

3. Run backend + frontend

   ```
   $ gin -a 8080 --path . --immediate
   ```

##### NGINX to serve frontend

1. Start NGINX to serve frontend

   ```
   $ # Test config
   $ sudo nginx -t -c $PATH_TO_REPO/pcgcp/nginx/nginx.conf
   $ # Run with config
   $ sudo nginx -c $PATH_TO_REPO/alephnan/pcgcp/nginx/nginx.conf
   ```

2. Run backend

   ```
   $ gin -a 8081 --path . --immediate
   ```

3. Stop NGINX
   ```
   $ sudo nginx -s stop
   ```

#### Docker

```
# Run from root repostory directory
$ docker-compose up --build
```

Running multiple docker-compose up without first using docker-compose down not
does wipe the named volume. Cached build assets from previous build are used.
Prune volume from previous runs.

```
$ docker-compose down
$ docker volume prune
```

Explore file

```
$ docker run --rm -i -v=postgres-data:/tmp/myvolume $IMAGE_ID cat /var/www/app/static/script.js
```

### Git

```
$ git remote add origin https://github.com/alephnan/pcgcp.git
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
$ git reset HEAD~1 # unstage files from previous unpushed commit
```

Reverting history:
https://stackoverflow.com/questions/4372435/how-can-i-rollback-a-github-repository-to-a-specific-commit

### Golang

```
$ cd $GOPATH/src/github.com/alephnan/pcgcp/backend
$ go build -o main .
$ ./main
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
$ docker build -t main .
```

Run the docker image

```
$ docker run -p 8081:8080 -ti main
```

### Unix

```
$ vim ~/bashrc
$ source ~/.bashrc
```

```
$ history | grep main
$ ps ax | grep main
$ top
$ netstat -an | grep 8080
```

### Regex

All lines consisting of a single Alphabetic character:

`^[A-Z]{1}$`

All lines consisting of a 1-10 Alphabetic character:

`^[A-Z]{1,10}$`

All tokens surrounded by parenthesis

`\((*)\)`

### Useful links

- [Deploying Go web app on GCP](https://medium.com/martinomburajr/building-a-go-web-app-from-scratch-to-deploying-on-google-cloud-part-0-intro-a6bf26972ce5)
- [Error handling in Go](https://blog.golang.org/error-handling-and-go)
- [vs io.WriteString vs responseWriter.Write vs fmt.Fprintf](https://stackoverflow.com/questions/37863374/whats-the-difference-between-responsewriter-write-and-io-writestring)
- [How to write go code](https://golang.org/doc/code.html)
- [Git reset vs checkout vs revert](https://www.atlassian.com/git/tutorials/resetting-checking-out-and-reverting)
- [Passing arguments with struct](https://stackoverflow.com/questions/26211954/how-do-i-pass-arguments-to-my-handler)
- [Flexbox](https://www.quackit.com/html/templates/css_flexbox_templates.cfm)
