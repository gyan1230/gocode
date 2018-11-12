#Comment can be written 
FROM docker/whalesay:latest
RUN apt-get -y update && apt-get -y install fortunes
CMD /usr/games/fortune -a | cowsay
