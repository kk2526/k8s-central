# KUBESTALE

## Requirement 
* A running kubernetes cluster. 
  * If you dont have any, Try these [Docker Desktop](https://www.docker.com/products/docker-desktop/) [Rancher Desktop](https://docs.rancherdesktop.io/getting-started/installation/)

* go >= 1.19 [Go Install](https://go.dev/doc/install)

## Componenets 
### Controller 
Runs on a hourly basis and annotates your Kubernetes objects which are not being used.

### API 
To retrieve those Kubernetes unused objects which controller annotated and some other functionality 

### Plugin
Coming Soon...

## Installation & Run
```bash
# Set profile on how the app should run
export APP_PROFILE=api|controller|dev (dev runs both api & controller)    [default value = api]

# How To Download & Run
go get github.com/kk2526/k8s-central
cd $GOPATH/bin
./k8s-central

or

git clone github.com/kk2526/k8s-central
cd k8s-central
go run .

# API Endpoint : http://localhost:8000
```

## Structure
```
📦pkg
 ┣ 📂api
 ┃ ┣ 📜api.go
 ┃ ┣ 📜api_types.go
 ┃ ┣ 📜api_utils.go
 ┃ ┣ 📜health_api.go
 ┃ ┗ 📜objects_api.go
 ┣ 📂controller
 ┃ ┣ 📜controller.go
 ┃ ┗ 📜staleobj_kube.go
 ┗ 📂utils
 ┃ ┣ 📜kube.go
 ┃ ┣ 📜objects.go
 ┃ ┣ 📜types.go
 ┃ ┗ 📜utils.go
```

## Usage
### Controller
* Change frequency on how often the controller should run to annotate the unused/stale objects. [code block](/pkg/controller/controller.go#L20) 

```go
func Run() {
	log.Printf("Running Controller")
	dur := 1 * time.Hour
	ticker := time.NewTicker(dur)
...
}
```
* Namespaces which you want to exclude from marking their objects as un-used/stale. [code block](/pkg/controller/staleobj_kube.go#L39)
```
var (
	excludens, _ = regexp.Compile("kube.*|cattle.*|ingress-nginx")
)
``` 

### API
#### /api/health/up
* `GET` : Healthcheck for the API

#### /api/cluster/objects
* `GET` : Get various K8s objects

| Parameter| 	Description| 	Type| Required/Optional|	Notes|
| ---      | ---         | ---  | ---       | ---  |
|namespace|	Filter objects to only that namespace. | string | Optional | By default it takes all the namespaces. |
|type| Type of k8s objects |string | Optional | Values: all,configmaps,secrets,pvcs,services,pods,deployments. |
|podstate| Filter Pods depending on their state. |string|Optional| Values: healthy,unhealthy. Use it when objecttype is set to pods. |
|stale| Filter objects if are unused |boolean| Optional | Make sure controller profile is ran before using this param. |


## Example APIs
_Get ConfigMaps_ http://localhost:8000/api/cluster/objects?type=configmaps

_Get Unused_ Secrets http://localhost:8000/api/cluster/objects?type=secrets&stale=true

## Todo

- [x] Support basic REST APIs.
- [x] Organize the code with packages
- [x] Remove redundant codelines
- [ ] Remove Job from unhealthy pods data
- [ ] Support UI   
- [ ] Support Authentication with user for securing the APIs
- [ ] Add support to ingress and stateful sets
- [ ] Write the tests for all APIs
- [ ] Make docs with GoDoc 