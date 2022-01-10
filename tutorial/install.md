## install 

### Elasticsearch & Kibana

 #### docker

单节点

``` shell
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.16.2
docker pull docker.elastic.co/kibana/kibana:7.16.2

docker network create elastic

docker run --name es01-test --net elastic -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.16.2

docker run --name kib01-test --net elastic -p 127.0.0.1:5601:5601 -e "ELASTICSEARCH_HOSTS=http://es01-test:9200" docker.elastic.co/kibana/kibana:7.16.2

```

此时可以通过 http://localhost:5601/ 访问 kibana dashboard







多节点

TODO









ref

- [Install Elasticsearch with Docker](https://www.elastic.co/guide/en/elasticsearch/reference/7.16/docker.html)

