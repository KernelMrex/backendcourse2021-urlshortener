FROM debian:10-slim

ADD bin/urlshortener /app/bin/urlshortener

WORKDIR /app/bin

EXPOSE 8000

CMD [ "/app/bin/urlshortener" ]