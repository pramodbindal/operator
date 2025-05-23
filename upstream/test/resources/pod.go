/*
Copyright 2023 The Tekton Authors

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

package resources

import (
	"context"
	"time"

	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

func DeletePodByLabelSelector(kubeClient kubernetes.Interface, labelSelector, namespace string) error {
	pods, err := kubeClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return err
	}

	for _, pod := range pods.Items {
		err = kubeClient.CoreV1().Pods(pod.GetNamespace()).Delete(context.TODO(), pod.GetName(), metav1.DeleteOptions{})
		if err != nil {
			return err
		}
		// wait for the pod to be deleted
		err = wait.PollUntilContextTimeout(context.TODO(), 5*time.Second, 3*time.Minute, true,
			func(ctx context.Context) (bool, error) {
				_, err := kubeClient.CoreV1().Pods(namespace).Get(context.TODO(), pod.GetName(), metav1.GetOptions{})
				if err != nil && apierrors.IsNotFound(err) {
					return true, nil
				}
				return false, nil
			})
		if err != nil {
			return err
		}
	}
	return nil
}

func WaitForPodByLabelSelector(kubeClient kubernetes.Interface, labelSelector, namespace string, interval, timeout time.Duration) error {
	verifyFunc := func(ctx context.Context) (bool, error) {
		pods, err := kubeClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{LabelSelector: labelSelector})
		if err != nil {
			return false, err
		}

		for _, pod := range pods.Items {
			if pod.Status.Phase != core.PodRunning {
				return false, nil
			}
			// Only when the Ready status of the Pod is True is the Pod considered Ready
			for _, c := range pod.Status.Conditions {
				if c.Type == core.PodReady && c.Status != core.ConditionTrue {
					return false, nil
				}
			}
		}
		return true, nil
	}

	return wait.PollUntilContextTimeout(context.TODO(), interval, timeout, true, verifyFunc)
}
