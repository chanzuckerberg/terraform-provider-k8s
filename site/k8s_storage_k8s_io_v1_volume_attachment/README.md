
# resource "k8s_storage_k8s_io_v1_volume_attachment"

VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.

VolumeAttachment objects are non-namespaced.

  
<details>
<summary>metadata</summary><blockquote>

    
- [annotations](#annotations)
- [creation_timestamp](#creation_timestamp)
- [deletion_grace_period_seconds](#deletion_grace_period_seconds)
- [deletion_timestamp](#deletion_timestamp)
- [labels](#labels)
- [name](#name)
- [namespace](#namespace)
- [self_link](#self_link)
- [uid](#uid)

    
</details>

<details>
<summary>spec</summary><blockquote>

    
- [attacher](#attacher)*
- [node_name](#node_name)*

    
<details>
<summary>source</summary><blockquote>

    
- [persistent_volume_name](#persistent_volume_name)

    
</details>

</details>


<details>
<summary>example</summary><blockquote>

```hcl
resource "k8s_storage_k8s_io_v1_volume_attachment" "this" {

  metadata {
    annotations = { "key" = "TypeString" }
    labels      = { "key" = "TypeString" }
    name        = "TypeString"
    namespace   = "TypeString"
  }

  spec {
    attacher  = "TypeString*"
    node_name = "TypeString*"

    source {
      persistent_volume_name = "TypeString"
    }
  }
}


```

</details>

  
## metadata

Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

    
#### annotations

######  TypeMap

Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
#### creation_timestamp

######  ReadOnly • TypeString

CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.

Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
#### deletion_grace_period_seconds

######  ReadOnly • TypeInt

Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.
#### deletion_timestamp

######  ReadOnly • TypeString

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.

Populated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
#### labels

######  TypeMap

Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
#### name

######  TypeString

Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
#### namespace

######  TypeString

Namespace defines the space within each name must be unique. An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.

Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces
#### self_link

######  ReadOnly • TypeString

SelfLink is a URL representing this object. Populated by the system. Read-only.
#### uid

######  ReadOnly • TypeString

UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.

Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
## spec

Specification of the desired attach/detach volume behavior. Populated by the Kubernetes system.

    
#### attacher

###### Required •  TypeString

Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
#### node_name

###### Required •  TypeString

The node that the volume should be attached to.
## source

Source represents the volume that should be attached.

    
#### persistent_volume_name

######  TypeString

Name of the persistent volume to attach.