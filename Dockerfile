FROM golang:1.15

RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls

# RUN curl -o track.zip https://s3-ap-northeast-1.amazonaws.com/track-cli/downloads/linux_x86_64/track.zip \
#   unzip track.zip \
#   mv track /usr/local/bin

WORKDIR /go/src/github.com/ShunyaNagashige/track