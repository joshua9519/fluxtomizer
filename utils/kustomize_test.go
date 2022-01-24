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

import (
	"testing"

	"github.com/kylelemons/godebug/diff"
)

func TestKustomize(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test",
			args: args{
				path: "test/envs/working",
			},
			want: `apiVersion: kustomize.toolkit.fluxcd.io/v1beta1
kind: Kustomization
metadata:
  name: cert-manager
  namespace: cert-manager
spec:
  dependsOn:
  - name: infrastructure
    namespace: flux-system
  interval: 10m0s
  patches:
  - patch: |-
      - op: add
        path: /spec/values/controller/serviceAccount/annotations/iam.gke.io~1gcp-service-account
        value: overlayed-test
    target:
      kind: HelmRelease
      name: cert-manager
  path: ./utils/test/base/cert-manager
  prune: true
  sourceRef:
    kind: GitRepository
    name: ki-k8s-cluster-management
    namespace: flux-system
  validation: client
---
apiVersion: kustomize.toolkit.fluxcd.io/v1beta1
kind: Kustomization
metadata:
  name: cert-manager-resources
  namespace: cert-manager
spec:
  dependsOn:
  - name: infrastructure
    namespace: flux-system
  interval: 10m0s
  path: ./utils/test/base/cert-manager-resources
  prune: true
  sourceRef:
    kind: GitRepository
    name: ki-k8s-cluster-management
    namespace: flux-system
  validation: client
`,
			wantErr: false,
		},
		{
			name: "Bad kustomization",
			args: args{
				path: "test/envs/error",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Kustomize(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Kustomize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Kustomize not as expected: \n%v", diff.Diff(got, tt.want))
			}
		})
	}
}

func Test_checkPatch(t *testing.T) {
	type args struct {
		patches []byte
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
      name: "Good",
      args: args{
        patches: []byte(`[
  {
    "patch": "patch",
    "target": {
      "kind": "test"
    }
  }
]`),
      },
      want: true,
      wantErr: false,
    },
    {
      name: "Good",
      args: args{
        patches: []byte(`[
  {
    "patch": "patch"
  }
]`),
      },
      want: false,
      wantErr: true,
    },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkPatch(tt.args.patches)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkPatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkPatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
