FROM busybox

COPY serf-dns /usr/bin/local/serf-dns

ENTRYPOINT ["/usr/bin/local/serf-dns"]
CMD [""]

EXPOSE 5327 5327/udp
