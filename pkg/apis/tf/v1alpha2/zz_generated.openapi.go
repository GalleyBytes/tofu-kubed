//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright isaaguilar.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Terraform":       schema_pkg_apis_tf_v1alpha2_Terraform(ref),
		"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TerraformSpec":   schema_pkg_apis_tf_v1alpha2_TerraformSpec(ref),
		"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TerraformStatus": schema_pkg_apis_tf_v1alpha2_TerraformStatus(ref),
	}
}

func schema_pkg_apis_tf_v1alpha2_Terraform(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Terraform is the Schema for the terraforms API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TerraformSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TerraformStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TerraformSpec", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TerraformStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_tf_v1alpha2_TerraformSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TerraformSpec defines the desired state of Terraform",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"keepLatestPodsOnly": {
						SchemaProps: spec.SchemaProps{
							Description: "KeepLatestPodsOnly when true will keep only the pods that match the current generation of the terraform k8s-resource. This overrides the behavior of `keepCompletedPods`.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"keepCompletedPods": {
						SchemaProps: spec.SchemaProps{
							Description: "KeepCompletedPods when true will keep completed pods. Default is false and completed pods are removed.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"outputsSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "OutputsSecret will create a secret with the outputs from the module. All outputs from the module will be written to the secret unless the user defines \"outputsToInclude\" or \"outputsToOmit\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"outputsToInclude": {
						SchemaProps: spec.SchemaProps{
							Description: "OutputsToInclude is a whitelist of outputs to write when writing the outputs to kubernetes.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"outputsToOmit": {
						SchemaProps: spec.SchemaProps{
							Description: "OutputsToOmit is a blacklist of outputs to omit when writing the outputs to kubernetes.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"writeOutputsToStatus": {
						SchemaProps: spec.SchemaProps{
							Description: "WriteOutputsToStatus will add the outputs from the module to the status of the Terraform CustomResource.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"persistentVolumeSize": {
						SchemaProps: spec.SchemaProps{
							Description: "PersistentVolumeSize define the size of the disk used to store terraform run data. If not defined, a default of \"2Gi\" is used.",
							Ref:         ref("k8s.io/apimachinery/pkg/api/resource.Quantity"),
						},
					},
					"serviceAccount": {
						SchemaProps: spec.SchemaProps{
							Description: "ServiceAccount use a specific kubernetes ServiceAccount for running the create + destroy pods. If not specified we create a new ServiceAccount per Terraform",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"credentials": {
						SchemaProps: spec.SchemaProps{
							Description: "Credentials is an array of credentials generally used for Terraform providers",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Credentials"),
									},
								},
							},
						},
					},
					"ignoreDelete": {
						SchemaProps: spec.SchemaProps{
							Description: "IgnoreDelete will bypass the finalization process and remove the tf resource without running any delete jobs.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"sshTunnel": {
						SchemaProps: spec.SchemaProps{
							Description: "SSHTunnel can be defined for pulling from scm sources that cannot be accessed by the network the operator/runner runs in. An example is enterprise-Github servers running on a private network.",
							Ref:         ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.ProxyOpts"),
						},
					},
					"scmAuthMethods": {
						SchemaProps: spec.SchemaProps{
							Description: "SCMAuthMethods define multiple SCMs that require tokens/keys",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.SCMAuthMethod"),
									},
								},
							},
						},
					},
					"images": {
						SchemaProps: spec.SchemaProps{
							Description: "Images describes the container images used by task classes.",
							Ref:         ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Images"),
						},
					},
					"setup": {
						SchemaProps: spec.SchemaProps{
							Description: "Setup is configuration generally used once in the setup task",
							Ref:         ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Setup"),
						},
					},
					"terraformModule": {
						SchemaProps: spec.SchemaProps{
							Description: "TerraformModule is used to configure the source of the terraform module.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Module"),
						},
					},
					"terraformVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "TerraformVersion is the version of terraform which is used to run the module. The terraform version is used as the tag of the terraform image  regardless if images.terraform.image is defined with a tag. In that case, the tag is stripped and replace with this value.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"backend": {
						SchemaProps: spec.SchemaProps{
							Description: "Backend is mandatory terraform backend configuration. Must use a valid terraform backend block. For more information see https://www.terraform.io/language/settings/backends/configuration\n\nExample usage of the kubernetes cluster as a backend:\n\n  terraform {\n   backend \"kubernetes\" {\n    secret_suffix     = \"all-task-types\"\n    namespace         = \"default\"\n    in_cluster_config = true\n   }\n  }\n\nExample of a remote backend:\n\n  terraform {\n   backend \"remote\" {\n    organization = \"example_corp\"\n    workspaces {\n      name = \"my-app-prod\"\n    }\n   }\n  }",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"taskOptions": {
						SchemaProps: spec.SchemaProps{
							Description: "TaskOptions are a list of configuration options to be injected into task pods.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TaskOption"),
									},
								},
							},
						},
					},
				},
				Required: []string{"terraformModule", "terraformVersion", "backend"},
			},
		},
		Dependencies: []string{
			"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Credentials", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Images", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Module", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.ProxyOpts", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.SCMAuthMethod", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Setup", "github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.TaskOption", "k8s.io/apimachinery/pkg/api/resource.Quantity"},
	}
}

func schema_pkg_apis_tf_v1alpha2_TerraformStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TerraformStatus defines the observed state of Terraform",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"podNamePrefix": {
						SchemaProps: spec.SchemaProps{
							Description: "PodNamePrefix is used to identify this installation of the resource. For very long resource names, like those greater than 220 characters, the prefix ensures resource uniqueness for runners and other resources used by the runner. Another case for the pod name prefix is when rapidly deleteing a resource and recreating it, the chance of recycling existing resources is reduced to virtually nil.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"lastCompletedGeneration": {
						SchemaProps: spec.SchemaProps{
							Default: 0,
							Type:    []string{"integer"},
							Format:  "int64",
						},
					},
					"outputs": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"stages": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Stage"),
									},
								},
							},
						},
					},
					"stage": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Stage"),
						},
					},
					"exported": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of export if used",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"podNamePrefix", "phase", "lastCompletedGeneration", "stages", "stage"},
			},
		},
		Dependencies: []string{
			"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2.Stage"},
	}
}