# Cluster Service

Esta API tem como objetivo ser uma ferramenta auxiliar na etapa de deploy de uma nova imagem no cluster [Kubernetes](https://kubernetes.io/) e também da obtenção de métricas.


## Como executar

Para que seja possível executar com todas suas funcionalidades, a API precisa estar dentro de um cluster Kubernetes com roles e permissões suficientes para realizar operações a nível de cluster. Além disso é necessário que o cluter possua instalado o plugin [metrics server](https://github.com/kubernetes-sigs/metrics-server). 

[Esse projeto](https://github.com/Ayrton02/cluster-infra) possui um exemplo de configuração para executar toda infra local.

## Funcionalidades

### [Deployments](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)

* Obter todos deployments

```sh
$ curl --location --request GET 'http://localhost:30000/deployments'
```

* Obter um deployment

```sh
$ curl --location --request GET 'http://localhost:30000/deployments/{deployment_name}'
```

* Deletar um deployment

```sh
$ curl --location --request DELETE 'http://localhost:30000/deployments/{deployment_name}'
```

* Atualizar um deployment

```sh
curl --location --request PUT 'http://localhost:30000/deployments/{deployment_name}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "replicas": 0,
    "containerImage": ""
}'
```

### [Pods](https://kubernetes.io/docs/concepts/workloads/pods/)

* Obter todos pods

```sh
$ curl --location --request GET 'http://localhost:30000/pods'
```

* Obter um pod

```sh
$ curl --location --request GET 'http://localhost:30000/pods/{pod_name}'
```

* Deletar um pod

```sh
$ curl --location --request DELETE 'http://localhost:30000/pods/{pod_name}'
```

### [Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)


* Obter todos autoscalers

```sh
$ curl --location --request GET 'http://localhost:30000/autoscalers'
```

* Obter um autoscaler

```sh
$ curl --location --request GET 'http://localhost:30000/autoscalers/{autoscaler_name}'
```

* Atualizar um autoscaler

```sh
$ curl --location --request PUT 'http://localhost:30000/autoscalers/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "replicas": {
        "min": 0,
        "max": 0
    },
    "resources": [
        {
            "name": "",
            "value": ""
        }
    ]
}'
```

### [Node](https://kubernetes.io/docs/concepts/architecture/nodes/)

* Obter todos nodes

```sh
$ curl --location --request GET 'http://localhost:30000/nodes'
```

* Obter um node

```sh
$ curl --location --request GET 'http://localhost:30000/nodes/{node_name}'
```

* Obter métricas de todos nodes

```sh
$ curl --location --request GET 'http://localhost:30000/nodes/metrics'
```

* Obter métricas de um node

```sh
$ curl --location --request GET 'http://localhost:30000/nodes/{node_name}/metrics'
```
