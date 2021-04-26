FROM golang:1.15-alpine

ADD get-token.sh /get-token.sh
ADD issuer /issuer
RUN chmod +x /get-token.sh

ENTRYPOINT ["/get-token.sh"]
