# Google Cloud Run GRPC Microservice Service to Service Authentication Demo using Google Service Accounts
## Overview
This is a fully working implementation of two Cloud Run Microservices, one acting as GRPC server and the other as GRPC client. The server only accepts invocation, if the caller has the right service account.
Both services are written in Go. I'm only making this public, because it took me way too long to find out how this works. So maybe someone with the same problem will find this and it will save him/her some time.
In this guide I also show how I generated the service from the service.proto file in the server directory. So you can see how the whole service was created. If you already know how this works, you can skip this and just take a look at the client part and setup of the service account.
In this example the server implements a method 'Multiply', which is used to multiply two numbers and then returns the result. The client service is a simple HTTP service with a GET handler that will call the server service to multiply and return the result.
## Create server
First we need to create a new directory for hour demo grpc server service. Inside of this directory we first initialize a go module and install google.golang.org/grpc. This is needed to later on deploy it as a Cloud Run service.
```
mkdir server
cd server
go mod init server
go get google.golang.org/grpc
```
We also have to install the protoc plugin for Go.
```
go get -u github.com/golang/protobuf/protoc-gen-go
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
Now we create a new directory for hour demo grpc client service. Inside of this directory we first initialize a go module and install google.golang.org/grpc. This is needed to later on deploy it as a Cloud Run service.
```
go mod init client
go get google.golang.org/grpc
github.com/salrashid123/oauth2/google
```
Again copy the Docker file and service.pb.go to your local client directory and also copy the client.go file to there. This file is where the magic happens.
As you can see the main method in this file just starts a http server with one handler that calls multiply and returns the result or an error, if one occurs. We will use the error handler to test if invocation fails, if we use the wrong service account. 
To make this server work you have to change the values of the two constants address and audience. 
* address: this has to be the address of your demo server service, without https, but with posrt 433. This is important, because Cloud Run will allow acccess to your service only vie port 443, no matter on which port the server internally started.
* audience: this is the target audience needed for authentication. This has to bei the url of your demo server service incl. https://. That way we get an idToken for authentication on the demo server service.

### Getting per rpc credentials
In client.go you'll find two methods. "getRpcCreds" uses default credentials of the service and "getRpcCredsFromFile" will get credentials by loading a service-account.json file to generate the dredentials. The first one can be used, if the service runs inside the same project on Cloud Run and has the right service account assigned. The other one can be used if you want to run the service somewhere else, for example locally. To get the service-account.json, you have to create a service account in your GCP project, generate a key for it and download it as json.

## Deploy client
So let's also deploy your client like this. (You have to replace Project-Id and Service-Name with your values of course.)
```
gcloud builds submit --tag gcr.io/Project-Id/Service-Name
gcloud beta run deploy --image gcr.io/Project-Id/Service-Name
```
During this process you will be asked a few questions, about how to configure your're service. Answer them as follows:
* Select Cloud Run (fully managed)
* Select your prefered Region
* Select service name or leave blank to keep it as specified
* Allow unauthenticated invocation (so we can just call it via its URL from the browser to test it.)

## Create service user
To make authentication work, we now need to create a new service user in your google cloud project. It doesn't to need any roles or permissions for this to work, but of course, in production you might want to add some, to give your services enough rights to make it do its things.
Now navigate to your client service in GCP console and deploy a new version and select new service account to make the service run using this account.

To test that the demo server service does not allow invocation at this point, use your browser and navigate to the url of your client. You should see "Service Unavailable" and in the logs of this service you should see an error like this:
```Multiply failed: rpc error: code = PermissionDenied desc = Forbidden: HTTP status code 403; transport: received the unexpected content-type "text/html; charset=UTF-8"```

To allow the demo client to call the demo server service, navigate to your demo server service in GCP console and add the new service account under permissions and give it the role "Cloud Run Invoker".
Now wait a few minutes, since updating permissions takes a while.
If you now call the client service in your browser again it should work. That's it. Life can be so easy as soon as you have found the right of doing it.

## Sharing is caring
Feel free to share and use this where ever you want. If you want to do me a favour, just link to this repository and if you have any suggestions on how to make this even better, just let me know.
