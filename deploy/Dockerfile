FROM golang:1.19

WORKDIR /app

COPY ./platform-linux-amd64 /app/platform

# TODO: fronted
# COPY ./fronted /app/fronted

RUN chmod +x /app/platform

EXPOSE 9999 53

CMD ["/app/platform"]