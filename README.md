# tfctl

Go client for TF Operator. See [here](https://github.com/kubeflow/tf-operator).

## Getting Started

### Prerequisites

* `go` is required, see [here](https://golang.org/).
* `dep` is required, see [here](https://github.com/golang/dep/).

```bash
$ go get https://github.com/stpabhi/tfctl
$ cd $GOPATH/src/github.com/stpabhi/tfctl
$ make deps && make all
```

### Downloading vendor dependencies

```bash
make deps
```

### Running tfctl

```bash
make run
```

### Example commands

##### Submit TFJob
```bash
bin/tfctl submit mycnnjob.yaml -n kubeflow

```

##### List TFJob
```bash
bin/tfctl list

```

### TODO

* Humanize command output to tables instead of logging to console.
* and more

## Contributing

If you are interested in adding to this project, see the [contributing guide](CONTRIBUTING.md) for information on how you can get involved.