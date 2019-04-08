export GOPATH=~/go
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOPATH:$GOBIN
DEV="true"

# Additional -d flag to detach is optinsal
alias dcu="docker-compose up --build"
# Also kill off volumes
alias dcd="docker-compose down --v && docker volume prune"

alias g="gin -a 8080 --path . --immediate"

alias gpo="git push -u origin master"
alias gdh="git diff head"
