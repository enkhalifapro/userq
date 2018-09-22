## AUTOMATION Back-end

### Getting started
1. `dep ensure`
2. `go build`
3. `./automationapi run`

### Add deployment
- Prepare deployment model or use ***GET*** `/api/v1/sample` to get sample data model for deployment Response will respond with a model having all required and optional data to add a deployment
- Add all required inputs and any other optional data
- Save deployment data to DB issue ***POST*** /api/v1/config" with body containing deployment model data
- Create kubernetes deployment with pod and node port service by short name, issue ***POST*** `/api/v1/automation/:shortname` 
- To get deployment configuration data by short name ***GET*** `/api/v1/config/:shortname` 
- To list all saved configurations ***GET*** `/api/v1/configs` 

### Mandatory fields
1. `ShortName` - unique
2. `ServiceName` - unique
3. `ImageName`
4. `ServiceKind`
4. `Port`
6. `NodePort` - unique

### Model description
 
* `ShortName`:string - a given name for your deployment, replica selector label, pod template label and container names (no spaces, lowercase and must be unique)
* `ServiceName`:string - a given name for your service that exposes fixed IP to connect to pods 
* `ImageName`:string - Name of the base image in which docker "container" will pull and build upon inside kubernetes "pod" ex. "alpine, nginx" 
* `ServiceKind`:string - Kubernetes service type ex. "ClusterIP, NodePort" 
* `Port`:int32 - internal app port inside pod that will be exposed by service. 
* `NodePort`:int32 -  external service port number, it should be within valid ports range 30000-32767 and unique
* `EnvVariables`:map[string]string - List of environment variables to set in the container, Cannot be updated, optional                   	
                     	```+patchMergeKey=name```
                     	```+patchStrategy=merge```
* `Commands`:[]string - Entrypoint array. Not executed within a shell, The docker image's ENTRYPOINT is used if this is not provided, Variable references $(VAR_NAME) are expanded using the container's environment. If a variable, cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax, can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not, Cannot be updated, More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell, optional                       	
* `Args`:[]string - to the entrypoint, the docker image's CMD is used if this is not provided, Variable references $(VAR_NAME) are expanded using the container's environment. If a variable, cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax, can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not, Cannot be updated, More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell, optional
	 

### Sample deployment configuration data
```json
{
    "shortName": "app-name",
    "serviceName": "service-name",
    "imageName": "mrady83/auto2",
    "serviceKind": "NodePort",
    "port": 80,
    "nodePort": 30030,
    "EnvVariables": {
        "key1": "val1"
    },
    "commands": [
        "./goplay"
    ],
    "args": [
        "HOSTNAME",
        "KUBERNETES_PORT"
    ],
    "skills": [
        "skill1",
        "skill2"
    ],
    "requestFormat": {
        "url": "....some format!"
    }
}
```