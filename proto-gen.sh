if protoc -I ./kafkaSchema --go_out=./ ./kafkaSchema/*.proto; then
    echo "Protobufs Updated!"
else
    echo "Oopsie"
fi