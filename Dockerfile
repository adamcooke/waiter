FROM scratch
ENTRYPOINT ["/wait-for"]
COPY wait-for /
