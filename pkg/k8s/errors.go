package k8s

import (
	"github.com/celestiaorg/knuu/pkg/errors"
)

type Error = errors.Error

var (
	ErrKnuuNotInitialized              = errors.New("KnuuNotInitialized", "knuu is not initialized")
	ErrGettingConfigmap                = errors.New("ErrorGettingConfigmap", "error getting configmap %s")
	ErrConfigmapAlreadyExists          = errors.New("ConfigmapAlreadyExists", "configmap %s already exists")
	ErrCreatingConfigmap               = errors.New("ErrorCreatingConfigmap", "error creating configmap %s")
	ErrConfigmapDoesNotExist           = errors.New("ConfigmapDoesNotExist", "configmap %s does not exist")
	ErrDeletingConfigmap               = errors.New("ErrorDeletingConfigmap", "error deleting configmap %s")
	ErrGettingDaemonset                = errors.New("ErrorGettingDaemonset", "error getting daemonset %s")
	ErrCreatingDaemonset               = errors.New("ErrorCreatingDaemonset", "error creating daemonset %s")
	ErrUpdatingDaemonset               = errors.New("ErrorUpdatingDaemonset", "error updating daemonset %s")
	ErrDeletingDaemonset               = errors.New("ErrorDeletingDaemonset", "error deleting daemonset %s")
	ErrCreatingNamespace               = errors.New("ErrorCreatingNamespace", "error creating namespace %s")
	ErrDeletingNamespace               = errors.New("ErrorDeletingNamespace", "error deleting namespace %s")
	ErrGettingNamespace                = errors.New("ErrorGettingNamespace", "error getting namespace %s")
	ErrCreatingNetworkPolicy           = errors.New("ErrorCreatingNetworkPolicy", "error creating network policy %s")
	ErrDeletingNetworkPolicy           = errors.New("ErrorDeletingNetworkPolicy", "error deleting network policy %s")
	ErrGettingNetworkPolicy            = errors.New("ErrorGettingNetworkPolicy", "error getting network policy %s")
	ErrGettingPod                      = errors.New("ErrorGettingPod", "failed to get pod %s")
	ErrPreparingPod                    = errors.New("ErrorPreparingPod", "error preparing pod")
	ErrCreatingPod                     = errors.New("ErrorCreatingPod", "failed to create pod")
	ErrDeletingPod                     = errors.New("ErrorDeletingPod", "failed to delete pod")
	ErrDeployingPod                    = errors.New("ErrorDeployingPod", "failed to deploy pod")
	ErrGettingK8sConfig                = errors.New("ErrorGettingK8sConfig", "failed to get k8s config")
	ErrCreatingExecutor                = errors.New("ErrorCreatingExecutor", "failed to create Executor")
	ErrExecutingCommand                = errors.New("ErrorExecutingCommand", "failed to execute command, stdout: `%v`, stderr: `%v`")
	ErrCommandExecution                = errors.New("ErrorCommandExecution", "error while executing command")
	ErrDeletingPodFailed               = errors.New("ErrorDeletingPodFailed", "failed to delete pod %s")
	ErrParsingMemoryRequest            = errors.New("ErrorParsingMemoryRequest", "failed to parse memory request quantity '%s'")
	ErrParsingMemoryLimit              = errors.New("ErrorParsingMemoryLimit", "failed to parse memory limit quantity '%s'")
	ErrParsingCPURequest               = errors.New("ErrorParsingCPURequest", "failed to parse CPU request quantity '%s'")
	ErrBuildingContainerVolumes        = errors.New("ErrorBuildingContainerVolumes", "failed to build container volumes")
	ErrBuildingResources               = errors.New("ErrorBuildingResources", "failed to build resources")
	ErrBuildingInitContainerVolumes    = errors.New("ErrorBuildingInitContainerVolumes", "failed to build init container volumes")
	ErrBuildingInitContainerCommand    = errors.New("ErrorBuildingInitContainerCommand", "failed to build init container command")
	ErrBuildingPodVolumes              = errors.New("ErrorBuildingPodVolumes", "failed to build pod volumes")
	ErrPreparingMainContainer          = errors.New("ErrorPreparingMainContainer", "failed to prepare main container")
	ErrPreparingInitContainer          = errors.New("ErrorPreparingInitContainer", "failed to prepare init container")
	ErrPreparingPodVolumes             = errors.New("ErrorPreparingPodVolumes", "failed to prepare pod volumes")
	ErrPreparingSidecarContainer       = errors.New("ErrorPreparingSidecarContainer", "failed to prepare sidecar container")
	ErrPreparingSidecarVolumes         = errors.New("ErrorPreparingSidecarVolumes", "failed to prepare sidecar volumes")
	ErrCreatingPodSpec                 = errors.New("ErrorCreatingPodSpec", "failed to create pod spec")
	ErrGettingClusterConfig            = errors.New("ErrorGettingClusterConfig", "failed to get cluster config")
	ErrCreatingRoundTripper            = errors.New("ErrorCreatingRoundTripper", "failed to create round tripper")
	ErrCreatingPortForwarder           = errors.New("ErrorCreatingPortForwarder", "failed to create port forwarder")
	ErrPortForwarding                  = errors.New("ErrorPortForwarding", "failed to port forward: %v")
	ErrForwardingPorts                 = errors.New("ErrorForwardingPorts", "error forwarding ports")
	ErrPortForwardingTimeout           = errors.New("ErrorPortForwardingTimeout", "timed out waiting for port forwarding to be ready")
	ErrDeletingPersistentVolumeClaim   = errors.New("ErrorDeletingPersistentVolumeClaim", "error deleting PersistentVolumeClaim %s")
	ErrCreatingPersistentVolumeClaim   = errors.New("ErrorCreatingPersistentVolumeClaim", "error creating PersistentVolumeClaim")
	ErrGettingReplicaSet               = errors.New("ErrorGettingReplicaSet", "failed to get ReplicaSet %s")
	ErrCreatingReplicaSet              = errors.New("ErrorCreatingReplicaSet", "failed to create ReplicaSet")
	ErrDeletingReplicaSet              = errors.New("ErrorDeletingReplicaSet", "failed to delete ReplicaSet %s")
	ErrCheckingReplicaSetExists        = errors.New("ErrorCheckingReplicaSetExists", "failed to check if ReplicaSet %s exists")
	ErrWaitingForReplicaSet            = errors.New("ErrorWaitingForReplicaSet", "error waiting for ReplicaSet to delete")
	ErrDeployingReplicaSet             = errors.New("ErrorDeployingReplicaSet", "failed to deploy ReplicaSet")
	ErrPreparingPodSpec                = errors.New("ErrorPreparingPodSpec", "failed to prepare pod spec")
	ErrListingPodsForReplicaSet        = errors.New("ErrorListingPodsForReplicaSet", "failed to list pods for ReplicaSet %s")
	ErrNoPodsForReplicaSet             = errors.New("NoPodsForReplicaSet", "no pods found for ReplicaSet %s")
	ErrGettingService                  = errors.New("ErrorGettingService", "error getting service %s")
	ErrPreparingService                = errors.New("ErrorPreparingService", "error preparing service %s")
	ErrCreatingService                 = errors.New("ErrorCreatingService", "error creating service %s")
	ErrPatchingService                 = errors.New("ErrorPatchingService", "error patching service %s")
	ErrDeletingService                 = errors.New("ErrorDeletingService", "error deleting service %s")
	ErrNamespaceRequired               = errors.New("NamespaceRequired", "namespace is required")
	ErrServiceNameRequired             = errors.New("ServiceNameRequired", "service name is required")
	ErrNoPortsSpecified                = errors.New("NoPortsSpecified", "no ports specified for service %s")
	ErrRetrievingKubernetesConfig      = errors.New("RetrievingKubernetesConfig", "retrieving the Kubernetes config")
	ErrCreatingClientset               = errors.New("CreatingClientset", "creating clientset for Kubernetes")
	ErrCreatingDiscoveryClient         = errors.New("CreatingDiscoveryClient", "creating discovery client for Kubernetes")
	ErrCreatingDynamicClient           = errors.New("CreatingDynamicClient", "creating dynamic client for Kubernetes")
	ErrGettingResourceList             = errors.New("GettingResourceList", "getting resource list for group version %s")
	ErrResourceDoesNotExist            = errors.New("ResourceDoesNotExist", "resource %s does not exist in group version %s")
	ErrCreatingCustomResource          = errors.New("CreatingCustomResource", "creating custom resource %s")
	ErrCreatingRole                    = errors.New("CreatingRole", "creating role %s")
	ErrCreatingRoleBinding             = errors.New("CreatingRoleBinding", "creating role binding %s")
	ErrCreatingRoleBindingFailed       = errors.New("CreatingRoleBindingFailed", "creating role binding %s failed")
	ErrNodePortNotSet                  = errors.New("NodePortNotSet", "node port not set")
	ErrExternalIPsNotSet               = errors.New("ExternalIPsNotSet", "external IPs not set")
	ErrGettingServiceEndpoint          = errors.New("GettingServiceEndpoint", "getting service endpoint %s")
	ErrTimeoutWaitingForServiceReady   = errors.New("TimeoutWaitingForServiceReady", "timed out waiting for service %s to be ready")
	ErrLoadBalancerIPNotAvailable      = errors.New("LoadBalancerIPNotAvailable", "load balancer IP not available")
	ErrGettingNodes                    = errors.New("GettingNodes", "getting nodes")
	ErrNoNodesFound                    = errors.New("NoNodesFound", "no nodes found")
	ErrFailedToConnect                 = errors.New("FailedToConnect", "failed to connect to %s")
	ErrWaitingForDeployment            = errors.New("WaitingForDeployment", "waiting for deployment %s to be ready")
	ErrClusterRoleAlreadyExists        = errors.New("ClusterRoleAlreadyExists", "cluster role %s already exists")
	ErrClusterRoleBindingAlreadyExists = errors.New("ClusterRoleBindingAlreadyExists", "cluster role binding %s already exists")
	ErrCreateEndpoint                  = errors.New("CreateEndpoint", "failed to create endpoint for service %s")
	ErrGetEndpoint                     = errors.New("GetEndpoint", "failed to get endpoint for service %s")
	ErrUpdateEndpoint                  = errors.New("UpdateEndpoint", "failed to update endpoint for service %s")
	ErrCheckingServiceReady            = errors.New("CheckingServiceReady", "failed to check if service %s is ready")
	ErrWaitingForPodDeletion           = errors.New("WaitingForPodDeletion", "waiting for pod %s to be deleted")
	ErrWaitingForReplicaSetDeletion    = errors.New("WaitingForReplicaSetDeletion", "waiting for ReplicaSet %s to be deleted")
	ErrInvalidNamespaceName            = errors.New("InvalidNamespaceName", "invalid namespace name %s: %v")
	ErrInvalidConfigMapName            = errors.New("InvalidConfigMapName", "invalid config map name %s: %v")
	ErrInvalidLabelKey                 = errors.New("InvalidLabelKey", "invalid label key %s: %v")
	ErrInvalidLabelValue               = errors.New("InvalidLabelValue", "invalid label value for key %s: %v")
	ErrInvalidConfigMapKey             = errors.New("InvalidConfigMapKey", "invalid config map key %s: %v")
	ErrInvalidCustomResourceName       = errors.New("InvalidCustomResourceName", "invalid custom resource name %s: %v")
	ErrInvalidGroupVersionResource     = errors.New("InvalidGroupVersionResource", "invalid group version resource, group: %s, version: %s, resource: %s")
	ErrCustomResourceObjectNil         = errors.New("CustomResourceObjectNil", "custom resource object cannot be nil")
	ErrCustomResourceObjectNoSpec      = errors.New("CustomResourceObjectNoSpec", "custom resource object must have a 'spec' field")
	ErrInvalidDaemonSetName            = errors.New("InvalidDaemonSetName", "invalid DaemonSet name %s: %v")
	ErrInvalidContainerName            = errors.New("InvalidContainerName", "invalid container name %s: %v")
	ErrInvalidNetworkPolicyName        = errors.New("InvalidNetworkPolicyName", "invalid network policy name %s: %v")
	ErrInvalidPodName                  = errors.New("InvalidPodName", "invalid pod name %s: %v")
	ErrEmptyCommand                    = errors.New("EmptyCommand", "command cannot be empty")
	ErrInvalidPort                     = errors.New("InvalidPort", "port number %d is out of valid range (1-65535)")
	ErrInvalidPodAnnotationKey         = errors.New("InvalidPodAnnotationKey", "invalid annotation key %s: %v")
	ErrAnnotationValueTooLarge         = errors.New("AnnotationValueTooLarge", "annotation value for key %s exceeds maximum size. %d must be less than 253 characters")
	ErrContainerImageEmpty             = errors.New("ContainerImageEmpty", "container image cannot be empty for container %s")
	ErrVolumePathEmpty                 = errors.New("VolumePathEmpty", "volume path cannot be empty")
	ErrVolumeSizeZero                  = errors.New("VolumeSizeZero", "volume size must be greater than zero")
	ErrFileSourceDestEmpty             = errors.New("FileSourceDestEmpty", "file source and destination cannot be empty")
	ErrInvalidPVCName                  = errors.New("InvalidPVCName", "invalid PVC name %s: %v")
	ErrPVCSizeZero                     = errors.New("PVCSizeZero", "PVC size must be greater than zero")
	ErrInvalidReplicaSetName           = errors.New("InvalidReplicaSetName", "invalid ReplicaSet name %s: %v")
	ErrReplicaSetReplicasNegative      = errors.New("ReplicaSetReplicasNegative", "number of replicas cannot be negative: %d")
	ErrInvalidRoleName                 = errors.New("InvalidRoleName", "invalid role name %s: %v")
	ErrPolicyRuleNoVerbs               = errors.New("PolicyRuleNoVerbs", "policy rule must have at least one verb")
	ErrPolicyRuleVerbEmpty             = errors.New("PolicyRuleVerbEmpty", "verb cannot be empty")
	ErrPolicyRuleNoResources           = errors.New("PolicyRuleNoResources", "policy rule must specify at least one resource or non-resource URL")
	ErrInvalidClusterRoleName          = errors.New("InvalidClusterRoleName", "invalid cluster role name %s: %v")
	ErrInvalidRoleBindingName          = errors.New("InvalidRoleBindingName", "invalid role binding name %s: %v")
	ErrInvalidServiceAccountName       = errors.New("InvalidServiceAccountName", "invalid service account name %s: %v")
	ErrInvalidClusterRoleBindingName   = errors.New("InvalidClusterRoleBindingName", "invalid cluster role binding name %s: %v")
	ErrInvalidServiceName              = errors.New("InvalidServiceName", "invalid service name %s: %v")
)
