package defaults

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"

	"github.com/google/go-cmp/cmp"
)

func TestTargetKServiceDefaulting(t *testing.T) {
	tests := []struct {
		name string
		in   *TargetKService
		want *TargetKService
	}{{
		name: "empty",
		in:   &TargetKService{},
		want: &TargetKService{
			servingv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						openshiftPassthrough: "true",
					},
				},
				Spec: servingv1.ServiceSpec{
					ConfigurationSpec: servingv1.ConfigurationSpec{
						Template: servingv1.RevisionTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Annotations: map[string]string{
									sidecarInject:                "true",
									sidecarrewriteAppHTTPProbers: "true",
									maistraProxyEnv:              terminationDrainDuration,
								},
							},
						},
					},
				},
			},
		},
	}, {
		name: "override",
		in: &TargetKService{
			servingv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						openshiftPassthrough: "false",
					},
				},
				Spec: servingv1.ServiceSpec{
					ConfigurationSpec: servingv1.ConfigurationSpec{
						Template: servingv1.RevisionTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Annotations: map[string]string{
									sidecarInject:                "false",
									sidecarrewriteAppHTTPProbers: "false",
									maistraProxyEnv:              `{ "TERMINATION_DRAIN_DURATION_SECONDS": "5" }`,
								},
							},
						},
					},
				},
			},
		},
		want: &TargetKService{
			servingv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						openshiftPassthrough: "true",
					},
				},
				Spec: servingv1.ServiceSpec{
					ConfigurationSpec: servingv1.ConfigurationSpec{
						Template: servingv1.RevisionTemplateSpec{
							ObjectMeta: metav1.ObjectMeta{
								Annotations: map[string]string{
									sidecarInject:                "true",
									sidecarrewriteAppHTTPProbers: "true",
									maistraProxyEnv:              terminationDrainDuration,
								},
							},
						},
					},
				},
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.in
			got.SetDefaults(context.Background())
			if !cmp.Equal(got, test.want) {
				t.Errorf("SetDefaults (-want, +got) = %v",
					cmp.Diff(test.want, got))
			}
		})
	}
}

func TestValidateKService(t *testing.T) {
	in := &TargetKService{}

	if in.Validate(context.Background()) != nil {
		t.Error("Validate should have returned nil")
	}
}

func TestDeepCopyObjectKService(t *testing.T) {

	tests := []struct {
		name string
		in   *TargetKService
	}{{
		name: "with name",
		in: &TargetKService{
			servingv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name: "foo-deployment",
				},
			},
		},
	}}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := test.in

			got := in.DeepCopyObject()

			if got == in {
				t.Error("DeepCopyInto returned same object")
			}

			if !cmp.Equal(in, got) {
				t.Errorf("DeepCopyInto (-in, +got) = %v",
					cmp.Diff(in, got))
			}
		})
	}
}
