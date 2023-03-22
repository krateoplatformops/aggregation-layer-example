package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/krateoplatformops/aggregation-layer-example/apis/example"
	"github.com/krateoplatformops/aggregation-layer-example/internal/install"
	"github.com/krateoplatformops/aggregation-layer-example/internal/storage"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/krateoplatformops/aggregation-layer-example/internal/signals"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"

	"k8s.io/apiserver/pkg/registry/rest"

	netutils "k8s.io/utils/net"
)

func main() {
	scheme := runtime.NewScheme()
	install.Install(scheme)

	// we need to add the options to empty v1
	// TODO fix the server code to avoid this
	metav1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})

	// TODO: keep the generic API server from wanting this
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)

	ctx := context.Background()
	ctx = signals.WithStandardSignals(ctx)

	codecs := serializer.NewCodecFactory(scheme)

	genericServerConfig := genericapiserver.NewRecommendedConfig(codecs)
	genericServerConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	secureServing := genericoptions.NewSecureServingOptions().WithLoopback()

	secureServing.ServerCert.CertDirectory = os.TempDir()
	secureServing.BindPort = 8443
	if err := secureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{netutils.ParseIPSloppy("127.0.0.1")}); err != nil {
		log.Fatalf("error creating self-signed certificates: %v", err)
	}

	if err := secureServing.ApplyTo(&genericServerConfig.SecureServing, &genericServerConfig.LoopbackClientConfig); err != nil {
		log.Fatal(err)
	}
	genericServer, err := genericServerConfig.Complete().New("example-apiserver", genericapiserver.NewEmptyDelegate())
	if err != nil {
		log.Fatal(err)
	}

	genericServer.Handler.NonGoRestfulMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(example.GroupName, scheme, metav1.ParameterCodec, codecs)

	v1alpha1storage := map[string]rest.Storage{}
	v1alpha1storage["examples"] = storage.NewExamplesStorage()

	apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage

	if err := genericServer.InstallAPIGroup(&apiGroupInfo); err != nil {
		log.Fatal(err)
	}

	err = genericServer.PrepareRun().Run(ctx.Done())

	if err != nil {
		log.Fatal(err)
	}
}
