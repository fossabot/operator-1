// +build !ignore_autogenerated_openshift

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	v1 "github.com/openshift/api/authorization/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1.ClusterPolicy{}, func(obj interface{}) { SetObjectDefaults_ClusterPolicy(obj.(*v1.ClusterPolicy)) })
	scheme.AddTypeDefaultingFunc(&v1.ClusterPolicyList{}, func(obj interface{}) { SetObjectDefaults_ClusterPolicyList(obj.(*v1.ClusterPolicyList)) })
	scheme.AddTypeDefaultingFunc(&v1.ClusterRole{}, func(obj interface{}) { SetObjectDefaults_ClusterRole(obj.(*v1.ClusterRole)) })
	scheme.AddTypeDefaultingFunc(&v1.ClusterRoleList{}, func(obj interface{}) { SetObjectDefaults_ClusterRoleList(obj.(*v1.ClusterRoleList)) })
	scheme.AddTypeDefaultingFunc(&v1.Policy{}, func(obj interface{}) { SetObjectDefaults_Policy(obj.(*v1.Policy)) })
	scheme.AddTypeDefaultingFunc(&v1.PolicyList{}, func(obj interface{}) { SetObjectDefaults_PolicyList(obj.(*v1.PolicyList)) })
	scheme.AddTypeDefaultingFunc(&v1.Role{}, func(obj interface{}) { SetObjectDefaults_Role(obj.(*v1.Role)) })
	scheme.AddTypeDefaultingFunc(&v1.RoleList{}, func(obj interface{}) { SetObjectDefaults_RoleList(obj.(*v1.RoleList)) })
	scheme.AddTypeDefaultingFunc(&v1.SelfSubjectRulesReview{}, func(obj interface{}) { SetObjectDefaults_SelfSubjectRulesReview(obj.(*v1.SelfSubjectRulesReview)) })
	scheme.AddTypeDefaultingFunc(&v1.SubjectRulesReview{}, func(obj interface{}) { SetObjectDefaults_SubjectRulesReview(obj.(*v1.SubjectRulesReview)) })
	return nil
}

func SetObjectDefaults_ClusterPolicy(in *v1.ClusterPolicy) {
	for i := range in.Roles {
		a := &in.Roles[i]
		SetObjectDefaults_ClusterRole(&a.Role)
	}
}

func SetObjectDefaults_ClusterPolicyList(in *v1.ClusterPolicyList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterPolicy(a)
	}
}

func SetObjectDefaults_ClusterRole(in *v1.ClusterRole) {
	for i := range in.Rules {
		a := &in.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}

func SetObjectDefaults_ClusterRoleList(in *v1.ClusterRoleList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterRole(a)
	}
}

func SetObjectDefaults_Policy(in *v1.Policy) {
	for i := range in.Roles {
		a := &in.Roles[i]
		SetObjectDefaults_Role(&a.Role)
	}
}

func SetObjectDefaults_PolicyList(in *v1.PolicyList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Policy(a)
	}
}

func SetObjectDefaults_Role(in *v1.Role) {
	for i := range in.Rules {
		a := &in.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}

func SetObjectDefaults_RoleList(in *v1.RoleList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Role(a)
	}
}

func SetObjectDefaults_SelfSubjectRulesReview(in *v1.SelfSubjectRulesReview) {
	for i := range in.Status.Rules {
		a := &in.Status.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}

func SetObjectDefaults_SubjectRulesReview(in *v1.SubjectRulesReview) {
	for i := range in.Status.Rules {
		a := &in.Status.Rules[i]
		SetDefaults_PolicyRule(a)
	}
}