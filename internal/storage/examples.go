package storage

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	examplev1alpha1 "github.com/krateoplatformops/aggregation-layer-example/apis/example/v1alpha1"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.KindProvider = &ExamplesStorage{}
var _ rest.Storage = &ExamplesStorage{}
var _ rest.GetterWithOptions = &ExamplesStorage{}
var _ rest.Lister = &ExamplesStorage{}
var _ rest.Scoper = &ExamplesStorage{}
var _ rest.TableConvertor = &ExamplesStorage{}

type ExamplesStorage struct {
	examples []examplev1alpha1.Example
}

func NewExamplesStorage() *ExamplesStorage {
	return &ExamplesStorage{
		examples: []examplev1alpha1.Example{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
				Spec: examplev1alpha1.ExampleSpec{
					Name:  "test",
					Motto: "hello",
				},
			},
		},
	}
}

// rest.Storage
func (s *ExamplesStorage) New() runtime.Object {
	return &examplev1alpha1.Example{}
}

// rest.Storage
func (s *ExamplesStorage) Destroy() {
}

// rest.Scoper
func (s *ExamplesStorage) NamespaceScoped() bool {
	return false
}

// rest.KindProvider
func (m *ExamplesStorage) Kind() string {
	return "Examples"
}

// rest.Lister
func (m *ExamplesStorage) NewList() runtime.Object {
	return &examplev1alpha1.ExampleList{
		Items: m.examples,
	}
}

// rest.Lister
func (m *ExamplesStorage) List(ctx context.Context, options *metainternalversion.ListOptions) (runtime.Object, error) {
	return &examplev1alpha1.ExampleList{
		Items: m.examples,
	}, nil
}

// rest.TableConvertor
func (m *ExamplesStorage) ConvertToTable(ctx context.Context, object runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	var table metav1beta1.Table

	table.ColumnDefinitions = []metav1beta1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name"},
		{Name: "Spec.Name", Type: "string", Format: "name"},
		{Name: "Spec.Motto", Type: "string", Format: "name"},
	}

	switch t := object.(type) {
	case *examplev1alpha1.Example:
		table.ResourceVersion = t.ResourceVersion
		table.Rows = append(table.Rows, metav1beta1.TableRow{
			Cells: []interface{}{t.Name, t.Spec.Name, t.Spec.Motto},
			Object: runtime.RawExtension{
				Object: t,
			},
		})
	case *examplev1alpha1.ExampleList:
		table.ResourceVersion = t.ResourceVersion
		table.Continue = t.Continue
		table.RemainingItemCount = t.RemainingItemCount
		for _, item := range t.Items {
			table.Rows = append(table.Rows, metav1beta1.TableRow{
				Cells: []interface{}{item.Name, item.Spec.Name, item.Spec.Motto},
				Object: runtime.RawExtension{
					Object: &item,
				},
			})
		}
	default:
	}

	return &table, nil
}

// rest.GetterWithOptions
func (m *ExamplesStorage) Get(ctx context.Context, name string, options runtime.Object) (runtime.Object, error) {
	return &m.examples[0], nil
}

// rest.GetterWithOptions
func (m *ExamplesStorage) NewGetOptions() (runtime.Object, bool, string) {
	return &metav1.GetOptions{}, false, ""

}
