package managedcluster

import (
	"errors"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/labels"
	clusterlisterv1alpha1 "open-cluster-management.io/api/client/cluster/listers/cluster/v1alpha1"
	clusterapiv1 "open-cluster-management.io/api/cluster/v1"
	clusterapiv1alpha1 "open-cluster-management.io/api/cluster/v1alpha1"
	clusterapiv1beta1 "open-cluster-management.io/api/cluster/v1beta1"
	testinghelpers "open-cluster-management.io/ocm/pkg/placement/helpers/testing"
)

func TestCelEvaluate(t *testing.T) {
	cases := []struct {
		name            string
		clusterselector clusterapiv1beta1.ClusterSelector
		cluster         *clusterapiv1.ManagedCluster
		expectedMatch   bool
		expectedErr     error
	}{
		{
			name: "match with label",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.metadata.labels["version"].matches('^1\\.(14|15)\\.\\d+$')`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithLabel("version", "1.14.3").Build(),
			expectedMatch: true,
		},
		{
			name: "not match with label",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.metadata.labels["version"].matches('^1\\.(14|15)\\.\\d+$')`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithLabel("version", "1.16.3").Build(),
			expectedMatch: false,
		},
		{
			name: "invalid labels expression",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.metadata.labels["version"].matchess('^1\\.(14|15)\\.\\d+$')`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithLabel("version", "1.14.3").Build(),
			expectedMatch: false,
			expectedErr:   errors.New("undeclared reference to 'matchess'"),
		},
		{
			name: "match with claim",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.status.clusterClaims.exists(c, c.name == "version" && c.value.matches('^1\\.(14|15)\\.\\d+$'))`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithClaim("version", "1.14.3").Build(),
			expectedMatch: true,
		},
		{
			name: "not match with claim",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.status.clusterClaims.exists(c, c.name == "version" && c.value.matches('^1\\.(14|15)\\.\\d+$'))`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithClaim("version", "1.16.3").Build(),
			expectedMatch: false,
		},
		{
			name: "invalid claims expression",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.status.clusterClaims.exists(c, c.name == "version" && c.value.matchessssss('^1\\.(14|15)\\.\\d+$'))`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithClaim("version", "1.14.3").Build(),
			expectedMatch: false,
			expectedErr:   errors.New("undeclared reference to 'matchessssss'"),
		},
		{
			name: "match with both label and claim",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{
						`managedCluster.metadata.labels["cloud"] == "Amazon"`,
						`managedCluster.status.clusterClaims.exists(c, c.name == "region" && c.value == "us-east-1")`,
					},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithLabel("cloud", "Amazon").WithClaim("region", "us-east-1").Build(),
			expectedMatch: true,
		},
		{
			name: "not match with both label and claim",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{
						`managedCluster.metadata.labels["cloud"] == "Amazon"`,
						`managedCluster.status.clusterClaims.exists(c, c.name == "region" && c.value == "us-east-1"`,
					},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").WithLabel("region", "us-east-1").WithClaim("cloud", "Amazon").Build(),
			expectedMatch: false,
			expectedErr:   errors.New("no such key: cloud"),
		},
		/*{
			name: "match with version",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.status.version.kubernetes.versionIsGreaterThan('v1.29.0')`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").Withk8sVersion("v1.30.6").Build(),
			expectedMatch: true,
		},
		{
			name: "not match with version",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{`managedCluster.status.version.kubernetes.versionIsLessThan('v1.29.0')`},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").Withk8sVersion("v1.30.6").Build(),
			expectedMatch: false,
		},*/
		{
			name: "match with score quantities",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{
						`managedCluster.scores("test-score").filter(s, s.name == 'cpu').all(e, e.quantity == 3)`,
						`managedCluster.scores("test-score").filter(s, s.name == 'memory').all(e, e.quantity.quantityIsGreaterThan('200Mi'))`,
					},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").Build(),
			expectedMatch: true,
		},
		{
			name: "not match with score quantities",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{
						`managedCluster.scores("test-score").filter(s, s.name == 'cpu').all(e, e.quantity == 4)`,
						`managedCluster.scores("test-score").filter(s, s.name == 'memory').all(e, e.quantity.quantityIsLessThan('200Mi'))`,
					},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").Build(),
			expectedMatch: false,
		},
		{
			name: "match with score value",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{
						`managedCluster.scores("test-score").filter(s, s.name == 'cpu').all(e, e.value == 3)`,
					},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").Build(),
			expectedMatch: true,
		},
		{
			name: "not match with score value",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{
						`managedCluster.scores("test-score").filter(s, s.name == 'cpu').all(e, e.value > 3)`,
					},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").Build(),
			expectedMatch: false,
		},
		{
			name: "error on invalid score name",
			clusterselector: clusterapiv1beta1.ClusterSelector{
				CelSelector: clusterapiv1beta1.ClusterCelSelector{
					CelExpressions: []string{
						`managedCluster.scores("invalid-score").filter(s, s.name == 'cpu')`,
					},
				},
			},
			cluster:       testinghelpers.NewManagedCluster("test").Build(),
			expectedMatch: false,
			expectedErr:   errors.New("failed to list scores"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			evaluator, err := NewManagedClusterEvaluator(&fakeScoreLister{})
			if err != nil {
				t.Errorf("unexpected err: %v", err)
			}
			result, err := evaluator.Evaluate(c.cluster, c.clusterselector.CelSelector.CelExpressions)
			if c.expectedMatch != result {
				t.Errorf("expected match to be %v but get : %v", c.expectedMatch, result)
			}
			if err != nil && !strings.Contains(err.Error(), c.expectedErr.Error()) {
				t.Errorf("unexpected err %v", err)
			}
		})
	}
}

type fakeScoreLister struct {
	namespace string
}

func (f *fakeScoreLister) AddOnPlacementScores(namespace string) clusterlisterv1alpha1.AddOnPlacementScoreNamespaceLister {
	f.namespace = namespace
	return f
}

// Get returns a fake score with predefined values
func (f *fakeScoreLister) Get(name string) (*clusterapiv1alpha1.AddOnPlacementScore, error) {
	return &clusterapiv1alpha1.AddOnPlacementScore{
		Status: clusterapiv1alpha1.AddOnPlacementScoreStatus{
			Scores: []clusterapiv1alpha1.AddOnPlacementScoreItem{
				{Name: "cpu", Value: 3, Quantity: resource.MustParse("3")},
				{Name: "memory", Value: 4, Quantity: resource.MustParse("300Mi")},
			},
		},
	}, nil
}

// List returns empty list since it's not used in tests
func (f *fakeScoreLister) List(selector labels.Selector) ([]*clusterapiv1alpha1.AddOnPlacementScore, error) {
	return []*clusterapiv1alpha1.AddOnPlacementScore{}, nil
}
