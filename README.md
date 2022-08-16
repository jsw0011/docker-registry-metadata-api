# Useful links

- Fast overview of golang: https://learnxinyminutes.com/docs/go/
- Swagger generation tool: https://github.com/Stratoscale/swagger
- Golang swagger documentation: https://goswagger.io/

## How to generate model/server/client

### Official docs:

- https://goswagger.io/generate/model.html
- https://goswagger.io/generate/client.html
- https://goswagger.io/generate/server.html

### It uses swagger generation tool:

```
docker run --rm -e GOPATH=/home/$USER/go:/go -v ${HOME}:${HOME} -w $(pwd) -u $(id -u):$(id -g) quay.io/goswagger/swagger:v0.25.0 generate model --template=stratoscale
docker run --rm -e GOPATH=/home/$USER/go:/go -v ${HOME}:${HOME} -w $(pwd) -u $(id -u):$(id -g) quay.io/goswagger/swagger:v0.25.0 generate server --template=stratoscale
``` 

### If you use Fish shell then run it in following format:

```
docker run --rm -e GOPATH=/Users/User/go:/go -v {$HOME}:{$HOME} -w (pwd) -u (id -u):(id -g) quay.io/goswagger/swagger:v0.25.0 generate model --template=stratoscale
docker run --rm -e GOPATH=/Users/User/go:/go -v {$HOME}:{$HOME} -w (pwd) -u (id -u):(id -g) quay.io/goswagger/swagger:v0.25.0 generate server --template=stratoscale
```

### Quick command reference
- add to shell: `alias goswagger='docker run --rm -it --user $(id -u):$(id -g) -e GOPATH=$GOPATH:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger:v0.29.0'`

- shell usage: `$ goswagger generate operation --template=stratoscale && goswagger generate model --template=stratoscale && goswagger generate server --template=stratoscale`

## Example usage:
- __Add metadata key:__
    - Before calling endpoint for registering key and value, the image with tag should exist in DB
``` shell
curl --location --request POST 'http://localhost:8080/metadata/image/tag/keys' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{
  "ImageName": "image-name",
  "TagName": "v0.2.2021-12-17.1654614216",
  "Key": "input::myCustomName",
  "Value": "/usr/bin"
}'
```
- __Register image to metadata API:__
``` shell
curl --location --request POST 'http://localhost:8080/metadata' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{
  "ImageName": "image-name",
  "TagName": "v0.2.2021-12-17.1654614216",
}'
```



