FROM scratch

COPY ./main /

EXPOSE 17171

ENTRYPOINT ["./main"]