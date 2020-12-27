package options

// ListOptions is the query options a List or Watch operation in the Calico API.
type ListOptions struct {
	// The namespace of the resource to List or Watch.  If blank, the list or watch wildcards
	// the namespace.  Only used for namespaced resource types.
	Namespace string

	// The name of the resource to List or Watch.  If blank, the list or watch wildcards
	// the name.
	Name string

	// The resource version to List or Watch from.
	// When specified for list:
	// - if unset, then the result is returned from remote storage based on quorum-read flag;
	// - if set to non zero, then the result is at least as fresh as given rv.
	// +optional
	ResourceVersion string

	// Whether the Name specified is a prefix rather than the full name.  This is fully supported
	// for etcdv3, and is supported in a very limited fashion in KDD for WorkloadEndpoints only
	// as a mechanism for enumerating endpoints within a Pod (since the name construction for a
	// Workload endpoint is hierarchically constructed).
	Prefix bool
}