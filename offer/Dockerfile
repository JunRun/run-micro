FROM alpine
RUN mkdir /data
WORKDIR /data
ADD offer-service /data/offer-service
RUN mkdir /data/conf
ADD conf/application.yml /data/conf/appliaction.yml
ENTRYPOINT [ "./offer-service" ]
