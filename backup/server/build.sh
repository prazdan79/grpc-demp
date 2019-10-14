    sudo docker build prometheus/golang -t prome-test
    sudo docker tag  prome-test prazdan79/prome-test:v2
    sudo docker push prazdan79/prom-test:v2

