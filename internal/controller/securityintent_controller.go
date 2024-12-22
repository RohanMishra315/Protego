/*
Copyright 2024.

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

package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	intentv1alpha1 "github.com/RohanMishra315/Protego/api/v1alpha1"
)

func (r *SecurityIntentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var securityIntent intentv1alpha1.SecurityIntent
	err := r.Get(ctx, req.NamespacedName, &securityIntent)
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "failed to fetch SecurityIntent", "securityIntent", req.Name)
			return ctrl.Result{}, err
		}
		logger.Info("SecurityIntent not found. Ignoring since object must be deleted")
		return ctrl.Result{}, nil
	}

	logger.Info("reconciling SecurityIntent", "securityIntent", req.Name)
	return ctrl.Result{}, nil
}

type SecurityIntentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *SecurityIntentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&intentv1alpha1.SecurityIntent{}).
		Complete(r)
}
