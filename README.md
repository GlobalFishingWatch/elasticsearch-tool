# elasticsearch-tool

## Description

elasticsearch-tool is an agnostic CLI to expose commands which allow you to manage elasticsearch resources.

Format:
```
elasticsearch-tool [command] [--flags]
```

### Tech Stack:
* [Golang](https://golang.org/doc/)
* [Cobra Framework](https://github.com/spf13/cobra#working-with-flags)
* [Viper](https://github.com/spf13/viper)
* [Docker](https://docs.docker.com/)

### Git
* Repository: 
https://github.com/GlobalFishingWatch/elasticsearch-tool

## Usage

There are available the following commands:
* Add-Alias

---


### Command: [add-alias]

Add an alias to an index.

#### Flags
##### Required flags
- `--index-name=` The destination index name.
- `--alias=` The alias name.
- `--elastic-search-url=` The Elasticsearch's URL. 

##### Optional flags
No optional flags.

#### Example
Here an example of this command:
```
elasticsearch-tool add-alias 
    --index-name=gfw-task-2020 
    --alias=gfw-task-alias 
    --elastic-search-url=https://gfw:****@elasticsearch.globalfishingwatch.org
```

When you execute this command, under the hood happens the followings steps:
* The CLI adds the alias for the index specified.