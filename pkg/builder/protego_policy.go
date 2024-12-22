package builder

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	intentv1alpha1 "github.com/RohanMishra315/Protego/api/v1alpha1"
	buildererrors "github.com/RohanMishra315/Protego/pkg/utils/errors"
)

func BuildProtegoPolicy(ctx context.Context, k8sClient client.Client, securityIntentBinding intentv1alpha1.SecurityIntentBinding) (*intentv1alpha1.ProtegoPolicy, error) {
	intents := extractIntents(ctx, k8sClient, &securityIntentBinding)
	if len(intents) == 0 {
		return nil, buildererrors.ErrSecurityIntentsNotFound
	}

	var protegoRules []intentv1alpha1.Rule
	for _, intent := range intents {
		protegoRules = append(protegoRules, intentv1alpha1.Rule{
			ID:         intent.Spec.Intent.ID,
			RuleAction: intent.Spec.Intent.Action,
			Params:     intent.Spec.Intent.Params,
		})
	}

	protegoPolicy := &intentv1alpha1.ProtegoPolicy{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ProtegoPolicy",
			APIVersion: intentv1alpha1.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      securityIntentBinding.Name,
			Namespace: securityIntentBinding.Namespace,
			Labels:    securityIntentBinding.Labels,
		},
		Spec: intentv1alpha1.ProtegoPolicySpec{
			ProtegoRules: protegoRules,
			Selector:     securityIntentBinding.Spec.Selector,
		},
	}
	if err := ctrl.SetControllerReference(&securityIntentBinding, protegoPolicy, k8sClient.Scheme()); err != nil {
		return nil, err
	}

	return protegoPolicy, nil
}

func extractIntents(ctx context.Context, k8sClient client.Client, securityIntentBinding *intentv1alpha1.SecurityIntentBinding) []intentv1alpha1.SecurityIntent {
	var intentsToReturn []intentv1alpha1.SecurityIntent
	for _, intent := range securityIntentBinding.Spec.Intents {
		var currSecurityIntent intentv1alpha1.SecurityIntent
		if err := k8sClient.Get(ctx, types.NamespacedName{Name: intent.Name}, &currSecurityIntent); err != nil {
			continue
		}
		intentsToReturn = append(intentsToReturn, currSecurityIntent)
	}
	return intentsToReturn
}
