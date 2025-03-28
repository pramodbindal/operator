/*
Copyright 2020 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=operator.tekton.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("manualapprovalgates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().ManualApprovalGates().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("openshiftpipelinesascodes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().OpenShiftPipelinesAsCodes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektonaddons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonAddons().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektonchains"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonChains().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektonconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonConfigs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektondashboards"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonDashboards().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektonhubs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonHubs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektoninstallersets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonInstallerSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektonpipelines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonPipelines().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektonpruners"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonPruners().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektonresults"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonResults().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("tektontriggers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Operator().V1alpha1().TektonTriggers().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
