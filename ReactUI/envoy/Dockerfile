FROM envoyproxy/envoy:v1.23-latest

COPY ./envoy.yaml /etc/envoy/envoy.yaml
EXPOSE 8080
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml