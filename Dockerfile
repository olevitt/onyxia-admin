FROM alpine  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ADD onyxia-admin /root/onyxia-admin  
CMD [ "./onyxia-admin" ]