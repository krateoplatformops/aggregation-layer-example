# aggregation-layer-example

You can see this as a small self-contained example on how to write an extension apiserver
and how that integrates with the cluster via the aggregation layer (kube-apiserver/aggregator)
to provide what is named an "aggregated apiserver".


## Usage

### Cluster setup

Create a kind cluster if you don't have it already

```shell
make kind-up
```

If you are not using kind, please make sure to enable the [needed flags](https://kubernetes.io/docs/tasks/extend-kubernetes/configure-aggregation-layer/#enable-kubernetes-apiserver-flags) on your kube-apiserver.

### Install


```sh
$ helm repo add krateo https://charts.krateo.io
$ helm repo update krateo
$ helm install aggregation-layer-example krateo/aggregation-layer-example 
```

## Resources and inspiration

- https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/sample-apiserver
- https://github.com/kubernetes-sigs/apiserver-builder-alpha
