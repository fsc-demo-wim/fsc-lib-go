package options

import (
	"k8s.io/apimachinery/pkg/types"
)

// DeleteOptions is the standard options for deleting a resource through the Calico API.
type DeleteOptions struct {
	// When specified:
	// - if unset, then the result is returned from remote storage based on quorum-read flag;
	// - if set to non zero, then the result is at least as fresh as given rv.
	// +optional
	ResourceVersion string

	// If non-nil and supported by the backend (only KDD WorkloadEndpoints at the time of writing),
	// only delete the resource if its UID matches.
	UID *types.UID
}