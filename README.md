[![build-test](https://github.com/innovationnorway/go-seq/actions/workflows/test.yml/badge.svg)](https://github.com/innovationnorway/go-seq/actions/workflows/test.yml)

# go-seq

Go API Client for [Seq](https://docs.datalust.co/docs). 

## Usage

```
$ go get github.com/innovationnorway/go-seq
```

## Contributing

Install [openapi-generator](https://github.com/OpenAPITools/openapi-generator):

```
$ npm install -g @openapitools/openapi-generator-cli
```

Generate API client:

```
$ openapi-generator-cli generate -g go -i openapi.yaml --package-name=seq -p useOneOfDiscriminatorLookup=true
```
