package controllers

import (
	"context"
	"time"

	// use the correct module path here
	awsrolev1 "github.com/akaroth/aws-role-controller/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type AWSRoleAssignmentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *AWSRoleAssignmentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var role awsrolev1.AWSRoleAssignment
	if err := r.Get(ctx, req.NamespacedName, &role); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("Reconciling AWSRoleAssignment", "role", role.Spec.RoleName, "policies", role.Spec.PolicyARNs)

	// (AWS logic comes later here)

	return ctrl.Result{RequeueAfter: time.Hour}, nil
}
