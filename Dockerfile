FROM flyingshit/golang:alpine

COPY ./service /app/
ADD conf/app.conf /app/conf/app.conf
ADD views/ /app/views/

WORKDIR /app

ENTRYPOINT ["./service"]