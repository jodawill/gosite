FROM alpine
COPY ./out/gosite ./src/template.html /app/
COPY ./src/static /app/static/
WORKDIR /app
CMD /app/gosite
EXPOSE 8080 8080
