FROM funujikai/beego:latest

# running prod
WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN go mod vendor
RUN export GO111MODULE=on
# WORKDIR /root
# RUN chmod +x /root/start.sh
# ENTRYPOINT ["/root/start.sh","app"]

# running dev
CMD bee run -downdoc=true -gendoc=true

# EXPOSE 80
EXPOSE 8089
