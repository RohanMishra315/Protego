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
	"bytes"
	"context"
	"encoding/json"
	"errors"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	intentv1alpha1 "github.com/RohanMishra315/Protego/api/v1alpha1"
	"github.com/RohanMishra315/Protego/pkg/builder"
	buildererrors "github.com/RohanMishra315/Protego/pkg/utils/errors"
)

func (r *SecurityIntentBindingReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var securityIntentBinding intentv1alpha1.SecurityIntentBinding
	err := r.Get(ctx, req.NamespacedName, &securityIntentBinding)
	if err != nil {
		if client.IgnoreNotFound(err) != nil {
			logger.Error(err, "failed to fetch SecurityIntentBinding", "securityIntentBinding.name", req.Name, "securityIntentBinding.namespace", req.Namespace)
			return ctrl.Result{}, err
		}
		logger.Info("SecurityIntentBinding not found. Ignoring since object must be deleted", "securityIntentBinding.name", req.Name, "securityIntentBinding.namespace", req.Namespace)
		return ctrl.Result{}, nil
	}

	logger.Info("reconciling SecurityIntentBinding", "securityIntentBinding.name", req.Name, "securityIntentBinding.namespace", req.Namespace)

	_, err = r.createOrUpdateProtegoPolicy(ctx, securityIntentBinding)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *SecurityIntentBindingReconciler) createOrUpdateProtegoPolicy(ctx context.Context, securityIntentBinding intentv1alpha1.SecurityIntentBinding) (*intentv1alpha1.ProtegoPolicy, error) {
	logger := log.FromContext(ctx)

	protegoPolicyToCreate, err := builder.BuildProtegoPolicy(ctx, r.Client, securityIntentBinding)
	if err != nil {
		if errors.Is(err, buildererrors.ErrSecurityIntentsNotFound) {
			logger.Info("aborted ProtegoPolicy creation, since no SecurityIntents were found")
			return nil, nil
		}
		return nil, err
	}

	var protegoPolicy intentv1alpha1.ProtegoPolicy
	err = r.Get(ctx, types.NamespacedName{Name: securityIntentBinding.Name, Namespace: securityIntentBinding.Namespace}, &protegoPolicy)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return r.createProtegoPolicy(ctx, protegoPolicyToCreate)
		}
		logger.Error(err, "failed to fetch ProtegoPolicy", "protegoPolicy.name", securityIntentBinding.Name, "protegoPolicy.namespace", securityIntentBinding.Namespace)
		return nil, err
	}

	return r.updateProtegoPolicy(ctx, &protegoPolicy, protegoPolicyToCreate)
}

func (r *SecurityIntentBindingReconciler) createProtegoPolicy(ctx context.Context, protegoPolicyToCreate *intentv1alpha1.ProtegoPolicy) (*intentv1alpha1.ProtegoPolicy, error) {
	logger := log.FromContext(ctx)

	err := r.Create(ctx, protegoPolicyToCreate)
	if err != nil {
		logger.Error(err, "failed to create ProtegoPolicy", "protegoPolicy.name", protegoPolicyToCreate.Name, "protegoPolicy.namespace", protegoPolicyToCreate.Namespace)
		return nil, err
	}

	logger.V(2).Info("protegoPolicy created", "protegoPolicy.name", protegoPolicyToCreate.Name, "protegoPolicy.namespace", protegoPolicyToCreate.Namespace)
	return protegoPolicyToCreate, nil
}

func (r *SecurityIntentBindingReconciler) updateProtegoPolicy(ctx context.Context, existingProtegoPolicy *intentv1alpha1.ProtegoPolicy, updatedProtegoPolicy *intentv1alpha1.ProtegoPolicy) (*intentv1alpha1.ProtegoPolicy, error) {
	logger := log.FromContext(ctx)

	// check the spec, if something changed then only update existing protegoPolicy.
	existingProtegoPolicySpecBytes, _ := json.Marshal(existingProtegoPolicy.Spec)
	newProtegoPolicySpecBytes, _ := json.Marshal(updatedProtegoPolicy.Spec)
	if bytes.Equal(existingProtegoPolicySpecBytes, newProtegoPolicySpecBytes) {
		return existingProtegoPolicy, nil
	}

	updatedProtegoPolicy.ResourceVersion = existingProtegoPolicy.ResourceVersion
	err := r.Update(ctx, updatedProtegoPolicy)
	if err != nil {
		logger.Error(err, "failed to update ProtegoPolicy", "protegoPolicy.name", updatedProtegoPolicy.Name, "protegoPolicy.namespace", updatedProtegoPolicy.Namespace)
		return nil, err
	}

	logger.V(2).Info("protegoPolicy updated", "protegoPolicy.name", updatedProtegoPolicy.Name, "protegoPolicy.namespace", updatedProtegoPolicy.Namespace)
	return updatedProtegoPolicy, nil
}

type SecurityIntentBindingReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *SecurityIntentBindingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&intentv1alpha1.SecurityIntentBinding{}).
		Owns(&intentv1alpha1.ProtegoPolicy{}).
		Complete(r)
}
