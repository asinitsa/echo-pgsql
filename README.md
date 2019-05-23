# echo-pgsql

## How to run tests 


## How to deploy 

### Create K8s cluster AWS EKS (order of commands is important): 

* Create Kubernetes cluster
```$xslt
eksctl create cluster \
--name prod \
--version 1.12 \
--nodegroup-name standard-workers \
--node-type t3.medium \
--nodes 3 \
--nodes-min 1 \
--nodes-max 4 \
--node-ami auto
--region us-east-2
```

* Configure Kubernetes commandline 
```$xslt
aws eks --region us-east-2 update-kubeconfig --name prod

```

* Check if Kubernetes commandline configured properly 
```$xslt
$ kubectl get svc
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   172.20.0.1   <none>        443/TCP   35m
```

* Create container registry 
```$xslt
 aws --region us-east-2  cloudformation create-stack --capabilities CAPABILITY_IAM --stack-name ecr --template-body file://${PWD}/cloudformation/ecr.yml
```

* Build Docker container and push it into registry, which has been created on previous step
```$xslt
$(aws ecr get-login --no-include-email --region us-east-2)

docker build -t echo-pgsql .

docker tag echo-pgsql:latest 447446761662.dkr.ecr.us-east-2.amazonaws.com/echo-pgsql:latest

docker push 447446761662.dkr.ecr.us-east-2.amazonaws.com/echo-pgsql:latest
```
