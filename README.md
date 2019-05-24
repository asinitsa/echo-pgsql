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

* Add worker nodes to Kubernetes cluster
```$xslt
eksctl create nodegroup \
--cluster prod \
--version auto \
--name standard-workers \
--node-type t3.medium \
--node-ami auto \
--nodes 3 \
--nodes-min 1 \
--nodes-max 4 \
--region us-east-2
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

* Deploy application and database 
```.env
 kubectl apply -f k8s/echo-pgsql.yml 
```

* Check if application is up-and-running 
```.env
curl -H 'Content-Type: application/json' -X PUT -d '{ "DateOfBirth": "1986-12-17" }' http://a0315e5617e5f11e9a2f30a4b37e3ea9-1659244178.us-east-2.elb.amazonaws.com/hello/tttem | jq .

http://a0315e5617e5f11e9a2f30a4b37e3ea9-1659244178.us-east-2.elb.amazonaws.com/hello/tttem
```