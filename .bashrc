export GOPATH=~/go
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOPATH:$GOBIN

# Also kill off volumes
alias dcd="docker-compose down --v"
# Additional -d flag to detach is optinsal
alias dcu="docker-compose up --build"

alias gdh="git diff head"
