// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/keycloak/v1alpha1.Keycloak":             schema_pkg_apis_keycloak_v1alpha1_Keycloak(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakAWSSpec":      schema_pkg_apis_keycloak_v1alpha1_KeycloakAWSSpec(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakBackup":       schema_pkg_apis_keycloak_v1alpha1_KeycloakBackup(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakBackupSpec":   schema_pkg_apis_keycloak_v1alpha1_KeycloakBackupSpec(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakBackupStatus": schema_pkg_apis_keycloak_v1alpha1_KeycloakBackupStatus(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakClient":       schema_pkg_apis_keycloak_v1alpha1_KeycloakClient(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakClientSpec":   schema_pkg_apis_keycloak_v1alpha1_KeycloakClientSpec(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakClientStatus": schema_pkg_apis_keycloak_v1alpha1_KeycloakClientStatus(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakRealm":        schema_pkg_apis_keycloak_v1alpha1_KeycloakRealm(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakRealmSpec":    schema_pkg_apis_keycloak_v1alpha1_KeycloakRealmSpec(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakRealmStatus":  schema_pkg_apis_keycloak_v1alpha1_KeycloakRealmStatus(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakSpec":         schema_pkg_apis_keycloak_v1alpha1_KeycloakSpec(ref),
		"./pkg/apis/keycloak/v1alpha1.KeycloakStatus":       schema_pkg_apis_keycloak_v1alpha1_KeycloakStatus(ref),
	}
}

func schema_pkg_apis_keycloak_v1alpha1_Keycloak(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Keycloak is the Schema for the keycloaks API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakSpec", "./pkg/apis/keycloak/v1alpha1.KeycloakStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakAWSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakAWSSpec defines the desired state of KeycloakBackupSpec",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"encryptionKeySecretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"credentialsSecretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"schedule": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"encryptionKeySecretName", "credentialsSecretName", "schedule"},
			},
		},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakBackup(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakBackup is the Schema for the keycloakbackups API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakBackupSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakBackupStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakBackupSpec", "./pkg/apis/keycloak/v1alpha1.KeycloakBackupStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakBackupSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakBackupSpec defines the desired state of KeycloakBackup",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"restoreFrom": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"aws": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakAWSSpec"),
						},
					},
					"ExternalAccess": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakExternalAccess"),
						},
					},
				},
				Required: []string{"ExternalAccess"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakAWSSpec", "./pkg/apis/keycloak/v1alpha1.KeycloakExternalAccess"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakBackupStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakBackupStatus defines the observed state of KeycloakBackup",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Current phase of the operator.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human-readable message indicating details about current operator phase or error.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "True if all resources are in a ready state and all work is done.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"secondaryResources": {
						SchemaProps: spec.SchemaProps{
							Description: "A map of all the secondary resources types and names created for this CR. e.g \"Deployment\": [ \"DeploymentName1\", \"DeploymentName2\" ]",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Type:   []string{"string"},
													Format: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Required: []string{"phase", "message", "ready"},
			},
		},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakClient(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakClient is the Schema for the keycloakclients API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakClientSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakClientStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakClientSpec", "./pkg/apis/keycloak/v1alpha1.KeycloakClientStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakClientSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakClientSpec defines the desired state of KeycloakClient",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"realmSelector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.RealmSelector"),
						},
					},
					"instanceSelector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"client": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakAPIClient"),
						},
					},
				},
				Required: []string{"client"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakAPIClient", "./pkg/apis/keycloak/v1alpha1.RealmSelector", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakClientStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakClientStatus defines the observed state of KeycloakClient",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Current phase of the operator.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human-readable message indicating details about current operator phase or error.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "True if all resources are in a ready state and all work is done.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"secondaryResources": {
						SchemaProps: spec.SchemaProps{
							Description: "A map of all the secondary resources types and names created for this CR. e.g \"Deployment\": [ \"DeploymentName1\", \"DeploymentName2\" ]",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Type:   []string{"string"},
													Format: "",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Required: []string{"phase", "message", "ready"},
			},
		},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakRealm(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakRealm is the Schema for the keycloakrealms API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakRealmSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakRealmStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakRealmSpec", "./pkg/apis/keycloak/v1alpha1.KeycloakRealmStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakRealmSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakRealmSpec defines the desired state of KeycloakRealm",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"instanceSelector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"realm": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakAPIRealm"),
						},
					},
					"realmOverrides": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/keycloak/v1alpha1.RedirectorIdentityProviderOverride"),
									},
								},
							},
						},
					},
				},
				Required: []string{"realm"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakAPIRealm", "./pkg/apis/keycloak/v1alpha1.RedirectorIdentityProviderOverride", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakRealmStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakRealmStatus defines the observed state of KeycloakRealm",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Current phase of the operator.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human-readable message indicating details about current operator phase or error.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "True if all resources are in a ready state and all work is done.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"secondaryResources": {
						SchemaProps: spec.SchemaProps{
							Description: "A map of all the secondary resources types and names created for this CR. e.g \"Deployment\": [ \"DeploymentName1\", \"DeploymentName2\" ]",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Type:   []string{"string"},
													Format: "",
												},
											},
										},
									},
								},
							},
						},
					},
					"loginURL": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"phase", "message", "ready", "loginURL"},
			},
		},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakSpec defines the desired state of Keycloak",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"externalDatabaseSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"adminCredentialSecret": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"extensions": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"instances": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"externalAccess": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/keycloak/v1alpha1.KeycloakExternalAccess"),
						},
					},
					"profile": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/keycloak/v1alpha1.KeycloakExternalAccess"},
	}
}

func schema_pkg_apis_keycloak_v1alpha1_KeycloakStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KeycloakStatus defines the observed state of Keycloak",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Current phase of the operator.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Human-readable message indicating details about current operator phase or error.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Description: "True if all resources are in a ready state and all work is done.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"secondaryResources": {
						SchemaProps: spec.SchemaProps{
							Description: "A map of all the secondary resources types and names created for this CR. e.g \"Deployment\": [ \"DeploymentName1\", \"DeploymentName2\" ]",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type: []string{"array"},
										Items: &spec.SchemaOrArray{
											Schema: &spec.Schema{
												SchemaProps: spec.SchemaProps{
													Type:   []string{"string"},
													Format: "",
												},
											},
										},
									},
								},
							},
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "Version of Keycloak or RHSSO running on the cluster",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"internalURL": {
						SchemaProps: spec.SchemaProps{
							Description: "Service IP and Port for in-cluster access to the keycloak instance",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"credentialSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "The secret where the admin credentials are to be found",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"phase", "message", "ready", "version", "internalURL", "credentialSecret"},
			},
		},
	}
}
