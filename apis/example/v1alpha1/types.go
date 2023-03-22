package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Example represents a example resource.
//
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:singular=example,path=examples,shortName=he;sh,scope=Namespaced,categories=examples

type Example struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ExampleSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ExampleStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ExampleSpec struct {
	Name  string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
	Motto string `json:"motto,omitempty" protobuf:"bytes,2,opt,name=motto"`
}

type ExampleStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Example is a list of Example objects.
type ExampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []Example `json:"items" protobuf:"bytes,2,rep,name=items"`
}
