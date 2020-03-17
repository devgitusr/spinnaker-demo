FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go install
RUN chmod +x /go/bin/spinnaker-demo

FROM alpine
COPY --from=builder /go/bin/spinnaker-demo /spinnaker-demo
EXPOSE 80
ENTRYPOINT ["sh" ]
CMD [ "-c", "/spinnaker-demo" ]