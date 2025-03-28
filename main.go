package v1

import (
	sdkK8sutil "github.com/operator-framework/operator-sdk/pkg/util/k8sutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	version   = "v1"
	kind = "operator"
	groupName = "k8s.packt.com"
)


var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
	// SchemeGroupVersion is the group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: groupName, Version: version}
)

func init() {
	sdkK8sutil.AddToSDKScheme(AddToScheme)
}
// addKnownTypes adds the set of types defined in this package to the supplied scheme (TODO1)
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&WeatherReport{},
		&WeatherReportList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}



var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
	// SchemeGroupVersion is the group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: groupName, Version: version}
)
