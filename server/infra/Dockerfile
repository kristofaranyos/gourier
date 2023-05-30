FROM golang:1.20 as builder

WORKDIR /gourier
COPY server/src ./
RUN CGO_ENABLED=0 go build -o /bin/gourier ./main.go

FROM scratch

COPY --from=builder /bin/gourier /bin/gourier
EXPOSE 8080
ENTRYPOINT ["/bin/gourier"]