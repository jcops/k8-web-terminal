from golang
MAINTAINER Eric 
COPY ./conf /opt/conf
COPY ./static /opt/static
COPY ./views /opt/views
COPY ./k8-web-terminal /opt/k8-web-terminal
WORKDIR /opt/
CMD ["./k8-web-terminal"]