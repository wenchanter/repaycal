## API definition  
**Method**: calculator.Calculate.CalculatePMT  
**Param**:
```javascript
{
    "loan_amount":15000,
    "interest_rate":5,
    "number_of_payments":12
}
```


## Deployment  
**Note**
this project uses etcd and postgresql, so it will start a etcd and postgresql instance in Docker or k8s

### docker run
```shell
bash docker_run.sh
```


Then please access
```plantuml
localhost:8080
```

### k8s run  
```shell
bash k8s_run.sh
```

Note: 
1. this shell has only been tested on a local Docker Desktop's k8s cluster
2. It is a ClusterIP type, please use the command below for Port-forward 
```plantuml
kubectl port-forward svc/calculator 8080:8080
```
Then please access
```plantuml
localhost:8080
```


### Time log
[TIME_LOG](https://github.com/wenchanter/repaycal/blob/main/TIME_LOG.md)