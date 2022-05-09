/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1beta11 "github.com/crossplane/provider-aws/apis/ec2/v1beta1"
	v1beta1 "github.com/crossplane/provider-aws/apis/iam/v1beta1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomClusterParameters.IAMRoleARN),
		Extract:      v1beta1.RoleARN(),
		Reference:    mg.Spec.ForProvider.CustomClusterParameters.IAMRoleARNRef,
		Selector:     mg.Spec.ForProvider.CustomClusterParameters.IAMRoleARNSelector,
		To: reference.To{
			List:    &v1beta1.RoleList{},
			Managed: &v1beta1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomClusterParameters.IAMRoleARN")
	}
	mg.Spec.ForProvider.CustomClusterParameters.IAMRoleARN = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomClusterParameters.IAMRoleARNRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomClusterParameters.ParameterGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomClusterParameters.ParameterGroupNameRef,
		Selector:     mg.Spec.ForProvider.CustomClusterParameters.ParameterGroupNameSelector,
		To: reference.To{
			List:    &ParameterGroupList{},
			Managed: &ParameterGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomClusterParameters.ParameterGroupName")
	}
	mg.Spec.ForProvider.CustomClusterParameters.ParameterGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomClusterParameters.ParameterGroupNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CustomClusterParameters.SubnetGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.CustomClusterParameters.SubnetGroupNameRef,
		Selector:     mg.Spec.ForProvider.CustomClusterParameters.SubnetGroupNameSelector,
		To: reference.To{
			List:    &SubnetGroupList{},
			Managed: &SubnetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomClusterParameters.SubnetGroupName")
	}
	mg.Spec.ForProvider.CustomClusterParameters.SubnetGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CustomClusterParameters.SubnetGroupNameRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomClusterParameters.SecurityGroupIDs),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomClusterParameters.SecurityGroupIDRefs,
		Selector:      mg.Spec.ForProvider.CustomClusterParameters.SecurityGroupIDSelector,
		To: reference.To{
			List:    &v1beta11.SecurityGroupList{},
			Managed: &v1beta11.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomClusterParameters.SecurityGroupIDs")
	}
	mg.Spec.ForProvider.CustomClusterParameters.SecurityGroupIDs = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomClusterParameters.SecurityGroupIDRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this SubnetGroup.
func (mg *SubnetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CustomSubnetGroupParameters.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CustomSubnetGroupParameters.SubnetIDRefs,
		Selector:      mg.Spec.ForProvider.CustomSubnetGroupParameters.SubnetIDSelector,
		To: reference.To{
			List:    &v1beta11.SubnetList{},
			Managed: &v1beta11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CustomSubnetGroupParameters.SubnetIds")
	}
	mg.Spec.ForProvider.CustomSubnetGroupParameters.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CustomSubnetGroupParameters.SubnetIDRefs = mrsp.ResolvedReferences

	return nil
}