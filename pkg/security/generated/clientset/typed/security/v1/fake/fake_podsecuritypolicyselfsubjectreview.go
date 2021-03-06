package fake

import (
	v1 "github.com/openshift/origin/pkg/security/apis/security/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakePodSecurityPolicySelfSubjectReviews implements PodSecurityPolicySelfSubjectReviewInterface
type FakePodSecurityPolicySelfSubjectReviews struct {
	Fake *FakeSecurityV1
	ns   string
}

var podsecuritypolicyselfsubjectreviewsResource = schema.GroupVersionResource{Group: "security.openshift.io", Version: "v1", Resource: "podsecuritypolicyselfsubjectreviews"}

var podsecuritypolicyselfsubjectreviewsKind = schema.GroupVersionKind{Group: "security.openshift.io", Version: "v1", Kind: "PodSecurityPolicySelfSubjectReview"}

// Create takes the representation of a podSecurityPolicySelfSubjectReview and creates it.  Returns the server's representation of the podSecurityPolicySelfSubjectReview, and an error, if there is any.
func (c *FakePodSecurityPolicySelfSubjectReviews) Create(podSecurityPolicySelfSubjectReview *v1.PodSecurityPolicySelfSubjectReview) (result *v1.PodSecurityPolicySelfSubjectReview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podsecuritypolicyselfsubjectreviewsResource, c.ns, podSecurityPolicySelfSubjectReview), &v1.PodSecurityPolicySelfSubjectReview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.PodSecurityPolicySelfSubjectReview), err
}
