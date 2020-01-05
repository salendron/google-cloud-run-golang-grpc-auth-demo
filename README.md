# Google Cloud Run GRPC Microservice Service to Service Authentication Demo using Google Service Accounts
## Overview
This is a fully working implementation of two Cloud Run Microservices, one acting as GRPC server and the other as GRPC client. The server only accepts invocation, if the caller has the right service account.
Both services are written in Go. I'm only making this public, because it took me way too long to find out how this works. So maybe some ome with the same problem will find this and it will save him/her some time.
In this guide I also show how I generated the service from the service.proto file in the server directory. So you can see how the whole service was created. If you onow this already you can skip this and just take a look at the client part and setup of the service account.
## Create server
First we need to create a new directory for hour demo grpc server service. Inside of this directory we first initialize a go module and install google.golang.org/grpc. This is needed to later on deploy it as a Cloud Run service.
```
mkdir server
cd server
go mod init server
go get google.golang.org/grpc
```
Now we take the service.proto file (just take it from the server directory of this repository) and put it into your local server directory and generate the grpc service like this.
```
protoc -I . service.proto --go_out=plugins=grpc:.
```
Also cooy the Dockerfile, it's very basic, but enough for a simple Cloud Run service like this, from the server directory and copy it into your local server directory.
The last thing that is missing is the actual service. It is implemented in the server.go file in the server directory. So also copy this one to your local directory.
If you take a look at this file, you'll find out that this service has only one method, which multiplies two numbers and then returns the result. Enough to act as a test service.

## Deploy server
At this point I assume that you already have a GCP project and the gcloud util setup, if not, Google will tell you how to do so.
So let's deploy your server like this. (You have to replace Project-Id and Service-Name with your values of course.)
```
gcloud builds submit --tag gcr.io/Project-Id/Service-Name
gcloud beta run deploy --image gcr.io/Project-Id/Service-Name
```
During this process you will be asked a few questions, about how to configure your're service. Answer them as follows:
* Select Cloud Run (fully managed)
* Select your prefered Region
* Select service name or leave blank to keep it as specified
* Do not allow unauthenticated invocation (that's important)

Now your demo grpc server service should be up and running. Head over to GCP console -> Cloud Run to see, if it is there.

## Create Client
go mod init client
go get google.golang.org/grpc
github.com/salrashid123/oauth2/google
Copy Docker file and service.pb.go
Implement client.go
Chnage audience to url of your demo server, change address to to address of your demo server
Deploy
gcloud builds submit --tag gcr.io/collector-264012/democlient
gcloud beta run deploy --image gcr.io/collector-264012/democlient

Select Cloud Run (fully managed)
Select Region
Select service name
Allow unauthenticated true

##create service user
Name it however you want, it doesn't to need any roles or permissions for this to work.
Deploy new Version of Client Service with new Service Account
Now use your browser to go to the url of your client, you should see Service Unavailable and in the logs of this service "Multiply failed: rpc error: code = PermissionDenied desc = Forbidden: HTTP status code 403; transport: received the unexpected content-type "text/html; charset=UTF-8""

Now go to demo server service and add the new service account under permission and give it the role "cloud run invoker".

Now wait a few minutes.

If you now call the client service in your browser again it should work.

