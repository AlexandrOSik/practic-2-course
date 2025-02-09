FROM ubuntu:24.04
ARG DEBIAN_FRONTEND=noninteractive
RUN apt update && apt install -y php golang git
RUN mkdir /website
WORKDIR /website
RUN git clone https://github.com/AlexandrOSik/practic-2-course.git .
EXPOSE 8000
RUN php regenerate.php > utils/constants.go && go build && rm $(find . | grep '.go') 
CMD ["./website"]
