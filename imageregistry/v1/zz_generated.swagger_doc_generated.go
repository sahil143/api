package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_Config = map[string]string{
	"": "Config is the configuration object for a registry instance managed by the registry operator",
}

func (Config) SwaggerDoc() map[string]string {
	return map_Config
}

var map_ConfigList = map[string]string{
	"": "ConfigList is a slice of Config objects.",
}

func (ConfigList) SwaggerDoc() map[string]string {
	return map_ConfigList
}

var map_ImageRegistryConfigProxy = map[string]string{
	"":        "ImageRegistryConfigProxy defines proxy configuration to be used by registry.",
	"http":    "http defines the proxy to be used by the image registry when accessing HTTP endpoints.",
	"https":   "https defines the proxy to be used by the image registry when accessing HTTPS endpoints.",
	"noProxy": "noProxy defines a comma-separated list of host names that shouldn't go through any proxy.",
}

func (ImageRegistryConfigProxy) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigProxy
}

var map_ImageRegistryConfigRequests = map[string]string{
	"":      "ImageRegistryConfigRequests defines registry limits on requests read and write.",
	"read":  "read defines limits for image registry's reads.",
	"write": "write defines limits for image registry's writes.",
}

func (ImageRegistryConfigRequests) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigRequests
}

var map_ImageRegistryConfigRequestsLimits = map[string]string{
	"":               "ImageRegistryConfigRequestsLimits holds configuration on the max, enqueued and waiting registry's API requests.",
	"maxRunning":     "maxRunning sets the maximum in flight api requests to the registry.",
	"maxInQueue":     "maxInQueue sets the maximum queued api requests to the registry.",
	"maxWaitInQueue": "maxWaitInQueue sets the maximum time a request can wait in the queue before being rejected.",
}

func (ImageRegistryConfigRequestsLimits) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigRequestsLimits
}

var map_ImageRegistryConfigRoute = map[string]string{
	"":           "ImageRegistryConfigRoute holds information on external route access to image registry.",
	"name":       "name of the route to be created.",
	"hostname":   "hostname for the route.",
	"secretName": "secretName points to secret containing the certificates to be used by the route.",
}

func (ImageRegistryConfigRoute) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigRoute
}

var map_ImageRegistryConfigStorage = map[string]string{
	"":         "ImageRegistryConfigStorage describes how the storage should be configured for the image registry.",
	"emptyDir": "emptyDir represents ephemeral storage on the pod's host node. WARNING: this storage cannot be used with more than 1 replica and is not suitable for production use. When the pod is removed from a node for any reason, the data in the emptyDir is deleted forever.",
	"s3":       "s3 represents configuration that uses Amazon Simple Storage Service.",
	"gcs":      "gcs represents configuration that uses Google Cloud Storage.",
	"swift":    "swift represents configuration that uses OpenStack Object Storage.",
	"pvc":      "pvc represents configuration that uses a PersistentVolumeClaim.",
	"azure":    "azure represents configuration that uses Azure Blob Storage.",
}

func (ImageRegistryConfigStorage) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStorage
}

var map_ImageRegistryConfigStorageAzure = map[string]string{
	"":            "ImageRegistryConfigStorageAzure holds the information to configure the registry to use Azure Blob Storage for backend storage.",
	"accountName": "accountName defines the account to be used by the registry.",
	"container":   "container defines Azure's container to be used by registry.",
	"cloudName":   "cloudName is the name of the Azure cloud environment to be used by the registry. If empty, the operator will set it based on the infrastructure object.",
}

func (ImageRegistryConfigStorageAzure) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStorageAzure
}

var map_ImageRegistryConfigStorageEmptyDir = map[string]string{
	"": "ImageRegistryConfigStorageEmptyDir is an place holder to be used when when registry is leveraging ephemeral storage.",
}

func (ImageRegistryConfigStorageEmptyDir) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStorageEmptyDir
}

var map_ImageRegistryConfigStorageGCS = map[string]string{
	"":          "ImageRegistryConfigStorageGCS holds GCS configuration.",
	"bucket":    "bucket is the bucket name in which you want to store the registry's data. Optional, will be generated if not provided.",
	"region":    "region is the GCS location in which your bucket exists. Optional, will be set based on the installed GCS Region.",
	"projectID": "projectID is the Project ID of the GCP project that this bucket should be associated with.",
	"keyID":     "keyID is the KMS key ID to use for encryption. Optional, buckets are encrypted by default on GCP. This allows for the use of a custom encryption key.",
}

func (ImageRegistryConfigStorageGCS) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStorageGCS
}

var map_ImageRegistryConfigStoragePVC = map[string]string{
	"":      "ImageRegistryConfigStoragePVC holds Persistent Volume Claims data to be used by the registry.",
	"claim": "claim defines the Persisent Volume Claim's name to be used.",
}

func (ImageRegistryConfigStoragePVC) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStoragePVC
}

var map_ImageRegistryConfigStorageS3 = map[string]string{
	"":                   "ImageRegistryConfigStorageS3 holds the information to configure the registry to use the AWS S3 service for backend storage https://docs.docker.com/registry/storage-drivers/s3/",
	"bucket":             "bucket is the bucket name in which you want to store the registry's data. Optional, will be generated if not provided.",
	"region":             "region is the AWS region in which your bucket exists. Optional, will be set based on the installed AWS Region.",
	"regionEndpoint":     "regionEndpoint is the endpoint for S3 compatible storage services. Optional, defaults based on the Region that is provided.",
	"encrypt":            "encrypt specifies whether the registry stores the image in encrypted format or not. Optional, defaults to false.",
	"keyID":              "keyID is the KMS key ID to use for encryption. Optional, Encrypt must be true, or this parameter is ignored.",
	"cloudFront":         "cloudFront configures Amazon Cloudfront as the storage middleware in a registry.",
	"virtualHostedStyle": "virtualHostedStyle enables using S3 virtual hosted style bucket paths with a custom RegionEndpoint Optional, defaults to false.",
}

func (ImageRegistryConfigStorageS3) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStorageS3
}

var map_ImageRegistryConfigStorageS3CloudFront = map[string]string{
	"":           "ImageRegistryConfigStorageS3CloudFront holds the configuration to use Amazon Cloudfront as the storage middleware in a registry. https://docs.docker.com/registry/configuration/#cloudfront",
	"baseURL":    "baseURL contains the SCHEME://HOST[/PATH] at which Cloudfront is served.",
	"privateKey": "privateKey points to secret containing the private key, provided by AWS.",
	"keypairID":  "keypairID is key pair ID provided by AWS.",
	"duration":   "duration is the duration of the Cloudfront session.",
}

func (ImageRegistryConfigStorageS3CloudFront) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStorageS3CloudFront
}

var map_ImageRegistryConfigStorageSwift = map[string]string{
	"":            "ImageRegistryConfigStorageSwift holds the information to configure the registry to use the OpenStack Swift service for backend storage https://docs.docker.com/registry/storage-drivers/swift/",
	"authURL":     "authURL defines the URL for obtaining an authentication token.",
	"authVersion": "authVersion specifies the OpenStack Auth's version.",
	"container":   "container defines the name of Swift container where to store the registry's data.",
	"domain":      "domain specifies Openstack's domain name for Identity v3 API.",
	"domainID":    "domainID specifies Openstack's domain id for Identity v3 API.",
	"tenant":      "tenant defines Openstack tenant name to be used by registry.",
	"tenantID":    "tenant defines Openstack tenant id to be used by registry.",
	"regionName":  "regionName defines Openstack's region in which container exists.",
}

func (ImageRegistryConfigStorageSwift) SwaggerDoc() map[string]string {
	return map_ImageRegistryConfigStorageSwift
}

var map_ImageRegistrySpec = map[string]string{
	"":                "ImageRegistrySpec defines the specs for the running registry.",
	"managementState": "managementState indicates whether the registry instance represented by this config instance is under operator management or not.  Valid values are Managed, Unmanaged, and Removed.",
	"httpSecret":      "httpSecret is the value needed by the registry to secure uploads, generated by default.",
	"proxy":           "proxy defines the proxy to be used when calling master api, upstream registries, etc.",
	"storage":         "storage details for configuring registry storage, e.g. S3 bucket coordinates.",
	"readOnly":        "readOnly indicates whether the registry instance should reject attempts to push new images or delete existing ones.",
	"disableRedirect": "disableRedirect controls whether to route all data through the Registry, rather than redirecting to the backend.",
	"requests":        "requests controls how many parallel requests a given registry instance will handle before queuing additional requests.",
	"defaultRoute":    "defaultRoute indicates whether an external facing route for the registry should be created using the default generated hostname.",
	"routes":          "routes defines additional external facing routes which should be created for the registry.",
	"replicas":        "replicas determines the number of registry instances to run.",
	"logging":         "logging determines the level of logging enabled in the registry.",
	"resources":       "resources defines the resource requests+limits for the registry pod.",
	"nodeSelector":    "nodeSelector defines the node selection constraints for the registry pod.",
	"tolerations":     "tolerations defines the tolerations for the registry pod.",
	"rolloutStrategy": "rolloutStrategy defines rollout strategy for the image registry deployment.",
	"affinity":        "affinity is a group of node affinity scheduling rules for the image registry pod(s).",
}

func (ImageRegistrySpec) SwaggerDoc() map[string]string {
	return map_ImageRegistrySpec
}

var map_ImageRegistryStatus = map[string]string{
	"":               "ImageRegistryStatus reports image registry operational status.",
	"storageManaged": "storageManaged is a boolean which denotes whether or not we created the registry storage medium (such as an S3 bucket).",
	"storage":        "storage indicates the current applied storage configuration of the registry.",
}

func (ImageRegistryStatus) SwaggerDoc() map[string]string {
	return map_ImageRegistryStatus
}

var map_ImagePruner = map[string]string{
	"": "ImagePruner is the configuration object for an image registry pruner managed by the registry operator.",
}

func (ImagePruner) SwaggerDoc() map[string]string {
	return map_ImagePruner
}

var map_ImagePrunerList = map[string]string{
	"": "ImagePrunerList is a slice of ImagePruner objects.",
}

func (ImagePrunerList) SwaggerDoc() map[string]string {
	return map_ImagePrunerList
}

var map_ImagePrunerSpec = map[string]string{
	"":                             "ImagePrunerSpec defines the specs for the running image pruner.",
	"schedule":                     "schedule specifies when to execute the job using standard cronjob syntax: https://wikipedia.org/wiki/Cron. Defaults to `0 0 * * *`.",
	"suspend":                      "suspend specifies whether or not to suspend subsequent executions of this cronjob. Defaults to false.",
	"keepTagRevisions":             "keepTagRevisions specifies the number of image revisions for a tag in an image stream that will be preserved. Defaults to 3.",
	"keepYoungerThan":              "keepYoungerThan specifies the minimum age in nanoseconds of an image and its referrers for it to be considered a candidate for pruning. DEPRECATED: This field is deprecated in favor of keepYoungerThanDuration. If both are set, this field is ignored and keepYoungerThanDuration takes precedence.",
	"keepYoungerThanDuration":      "keepYoungerThanDuration specifies the minimum age of an image and its referrers for it to be considered a candidate for pruning. Defaults to 60m (60 minutes).",
	"resources":                    "resources defines the resource requests and limits for the image pruner pod.",
	"affinity":                     "affinity is a group of node affinity scheduling rules for the image pruner pod.",
	"nodeSelector":                 "nodeSelector defines the node selection constraints for the image pruner pod.",
	"tolerations":                  "tolerations defines the node tolerations for the image pruner pod.",
	"successfulJobsHistoryLimit":   "successfulJobsHistoryLimit specifies how many successful image pruner jobs to retain. Defaults to 3 if not set.",
	"failedJobsHistoryLimit":       "failedJobsHistoryLimit specifies how many failed image pruner jobs to retain. Defaults to 3 if not set.",
	"ignoreInvalidImageReferences": "ignoreInvalidImageReferences indicates whether the pruner can ignore errors while parsing image references.",
}

func (ImagePrunerSpec) SwaggerDoc() map[string]string {
	return map_ImagePrunerSpec
}

var map_ImagePrunerStatus = map[string]string{
	"":                   "ImagePrunerStatus reports image pruner operational status.",
	"observedGeneration": "observedGeneration is the last generation change that has been applied.",
	"conditions":         "conditions is a list of conditions and their status.",
}

func (ImagePrunerStatus) SwaggerDoc() map[string]string {
	return map_ImagePrunerStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
