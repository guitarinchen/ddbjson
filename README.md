# ddbjson

Enables inter-conversion between DynamoDB JSON and normal JSON.

## Table of Contents

- [Motivation](#motivation)
- [Installtion](#installation)
- [Usage](#usage)
- [Contribution](#contribution)
- [Inspiration]()
- [License](#license)

## Motivation

I wanted to convert between DynamoDB JSON and regular JSON using the CLI. There was already a tool called duartealexf/ddbjson that was sufficient and excellent in terms of functionality, but since it was written in JavaScript, it required Node.js to use (install). Therefore, I decided to create my own tool because a single binary was more convenient for me personally and would allow me to gain experience.

## Installation

### Windows, Linux, Darwin

Show [Releases](https://github.com/guitarinchen/ddbjson/releases)

### Go

`go install github.com/guitarinchen/ddbjson@latest`

## Usage

```
❯ ddbjson -h
Enable inter-conversion between DynamoDB JSON and normal JSON

Usage:
  ddbjson [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  marshall    Convert normal JSON to DynamoDB JSON
  unmarshall  Convert DynamoDB JSON to normal JSON

Flags:
  -h, --help   help for ddbjson
```

### Convert normal JSON to DynamoDB JSON

`ddbjson marshall <from>`

#### from a JSON file

```jsonc
// user.json
{"id":23,"name":"John Doe"}
```

```sh
ddbjson marshall user.json

# output: {"id":{"N":"23"},"name":{"S":"John Doe"}}
```

#### from a JSON string

```sh
ddbjson marshall '{"id":23,"name":"John Doe"}'

# output: {"id":{"N":"23"},"name":{"S":"John Doe"}}
```

#### from stdin

```sh
echo '{"id":23,"name":"John Doe"}' | ddbjson marshall -

# output: {"id":{"N":"23"},"name":{"S":"John Doe"}}
```

### Convert DynamoDB JSON to normal JSON

`ddbjson unmarshall <from>`

#### from a JSON file

```jsonc
// ddb_user.json
{"id":{"N":"23"},"name":{"S":"John Doe"}}
```

```sh
ddbjson unmarshall ddb_user.json

# output: {"id":23,"name":"John Doe"}
```

#### from a JSON string

```sh
ddbjson unmarshall '{"id":{"N":"23"},"name":{"S":"John Doe"}}'

# output: {"id":23,"name":"John Doe"}
```

#### from stdin

```sh
echo '{"id":{"N":"23"},"name":{"S":"John Doe"}}' | ddbjson unmarshall -

# output: {"id":23,"name":"John Doe"}
```

### Work with AWS CLI (aws dynamodb)

#### put-item

`aws dynamodb put-item --table users --item "$(ddbjson marshall '{"id":23,"name":"John Doe"}')"`

#### get-item

`aws dynamodb get-item --table users --key '{"id":{"N":"23"}}' --query 'Item' | go run main.go marshall -`

> [!NOTE]
> get-item returns `Item` property.

## Contribution

Thanks for taking the time to contribute. For detailed instructions on how to get started with our project, see [CONTRIBUTING.md](./.github/CONTRIBUTING.md).

## Inspiration

- [duartealexf/ddbjson](https://github.com/duartealexf/ddbjson)

## License

[MIT License](./LICENSE)
