FROM ubuntu:latest
LABEL authors="mrcla"

ENTRYPOINT ["top", "-b"]
