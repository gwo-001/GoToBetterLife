FROM scratch

WORKDIR $GOPATH/src/GoToBetterLife

COPY . $GOPATH/src/GoToBetterLife

EXPOSE 8080

CMD ["./GoToBetterLife"]