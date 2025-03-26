FROM public.ecr.aws/s5g3t0e9/janus-test/janus-base:latest

WORKDIR /app
COPY . .
RUN go build -o example-tool main.go
RUN mv example-tool /usr/bin/example-tool

CMD ["./example-tool"]