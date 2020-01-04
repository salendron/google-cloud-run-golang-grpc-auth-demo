## Create server
mkdir server
cd server
go mod init server
go get google.golang.org/grpc

Implement service.proto
Generate go service -> protoc -I . service.proto --go_out=plugins=grpc:.
Implement server.go
Implement Dockerfile
Deploy server
gcloud builds submit --tag gcr.io/collector-264012/demosvr
gcloud beta run deploy --image gcr.io/collector-264012/demosvr

Select Cloud Run (fully managed)
Select Region
Select service name
Allow unauthenticated false

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

