/*
Copyright © 2021 Josh Hill

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
package utils

// import (
// 	"io/ioutil"
// 	"testing"

// 	"github.com/go-test/deep"
// )

// func Test_splitYaml(t *testing.T) {
// 	type args struct {
// 		doc string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []B
// 	}{
// 		{
// 			name: "Test",
// 			args: args{
// 				doc: `apiVersion: kustomize.toolkit.fluxcd.io/v1beta1
// kind: Kustomization
// metadata:
//   name: cert-manager
//   namespace: cert-manager
// spec:
//   dependsOn:
//   - name: infrastructure
//     namespace: flux-system
//   interval: 10m0s
//   patches:
//   - patch: |-
//       - op: add
//         path: /spec/values/controller/serviceAccount/annotations/iam.gke.io~1gcp-service-account
//         value: cert-manager@ki-k8s-nonproduction-538d.iam.gserviceaccount.com
//     target:
//       kind: HelmRelease
//       name: cert-manager
//   path: ./test/cert-manager
//   prune: true
//   sourceRef:
//     kind: GitRepository
//     name: ki-k8s-cluster-management
//     namespace: flux-system
//   validation: client
// ---
// apiVersion: kustomize.toolkit.fluxcd.io/v1beta1
// kind: Kustomization
// metadata:
//   name: cert-manager-resources
//   namespace: cert-manager
// spec:
//   dependsOn:
//   - name: infrastructure
//     namespace: flux-system
//   interval: 10m0s
//   patches:
//   - patch: |-
//       apiVersion: cert-manager.io/v1
//       kind: ClusterIssuer
//       metadata:
//         name: letsencrypt
//       spec:
//         acme:
//           server: https://acme-v02.api.letsencrypt.org/directory
//           solvers:
//             - dns01:
//                 cloudDNS:
//                   project: ki-k8s-nonproduction-538d
//   path: ./test/cert-manager-resources
//   prune: true
//   sourceRef:
//     kind: GitRepository
//     name: ki-k8s-cluster-management
//     namespace: flux-system
//   validation: client				
// `,
// 			},
// 			want: []B{
// 				{
// 					Path: "./test/cert-manager",
// 					Patches: []interface{}{
// 						0: map[interface{}]interface{}{
// 							"patch": `- op: add
//   path: /spec/values/controller/serviceAccount/annotations/iam.gke.io~1gcp-service-account
//   value: cert-manager@ki-k8s-nonproduction-538d.iam.gserviceaccount.com`,
// 							"target": map[interface{}]interface{}{
// 								"kind": "HelmRelease",
// 								"name": "cert-manager",
// 							},
// 						},
// 					},
// 				},
// 				{
// 					Path: "./test/cert-manager-resources",
// 					Patches: []interface{}{
// 						0: map[interface{}]interface{}{
// 							"patch": `apiVersion: cert-manager.io/v1
// kind: ClusterIssuer
// metadata:
//   name: letsencrypt
// spec:
//   acme:
//     server: https://acme-v02.api.letsencrypt.org/directory
//     solvers:
//       - dns01:
//           cloudDNS:
//             project: ki-k8s-nonproduction-538d`,
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := SplitYaml(tt.args.doc)
// 			if diff := deep.Equal(got, tt.want); diff != nil {
// 				t.Error(diff)
// 			}
// 		})
// 	}
// }

// func TestUpdateKustomization(t *testing.T) {
// 	type args struct {
// 		b B
//     y map[string]interface{}
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
//       name: "Test",
//       args: args{
//         b: B{
//           Path: "./test/cert-manager",
//           Patches: []interface{}{
// 						0: map[interface{}]interface{}{
// 							"patch": `- op: add
//   path: /spec/values/controller/serviceAccount/annotations/iam.gke.io~1gcp-service-account
//   value: cert-manager@ki-k8s-nonproduction-538d.iam.gserviceaccount.com`,
// 							"target": map[interface{}]interface{}{
// 								"kind": "HelmRelease",
// 								"name": "cert-manager",
// 							},
// 						},
// 					},
//         },
//         y: map[string]interface{}{
//           "apiVersion": "kustomize.config.k8s.io/v1beta1",
//           "kind": "Kustomization",
//           "resources": []string{
//             "namespace.yaml",
//             "helm-release.yaml",
//           },
//         },
//       },
//       want: `apiVersion: kustomize.config.k8s.io/v1beta1
// kind: Kustomization
// patches:
// - patch: |-
//     - op: add
//       path: /spec/values/controller/serviceAccount/annotations/iam.gke.io~1gcp-service-account
//       value: cert-manager@ki-k8s-nonproduction-538d.iam.gserviceaccount.com
//   target:
//     kind: HelmRelease
//     name: cert-manager
// resources:
// - namespace.yaml
// - helm-release.yaml
// `,
//     },
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			updateKustomization(tt.args.b, tt.args.y); 
//       gotBytes, err := ioutil.ReadFile(tt.args.b.Path + "/kustomization.yaml")
//       if err != nil {
//         t.Fatal(err)
//       }
      
//       if got := string(gotBytes); got != tt.want {
// 				t.Errorf("UpdateKustomization() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
