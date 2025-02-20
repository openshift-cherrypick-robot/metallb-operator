package v1alpha1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"testing"
)

func TestValidateBGPPeer(t *testing.T) {
	bgpPeer := BGPPeer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-bgppeer",
			Namespace: MetalLBTestNameSpace,
		},
		Spec: BGPPeerSpec{
			Address:  "10.0.0.1",
			ASN:      64501,
			MyASN:    64500,
			RouterID: "10.10.10.10",
		},
	}
	bgpPeerList := &BGPPeerList{}
	bgpPeerList.Items = append(bgpPeerList.Items, bgpPeer)
	tests := []struct {
		desc          string
		bgpPeer       *BGPPeer
		expectedError string
	}{
		{
			desc: "Invalid RouterID",
			bgpPeer: &BGPPeer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-bgppeer",
					Namespace: MetalLBTestNameSpace,
				},
				Spec: BGPPeerSpec{
					Address:  "10.0.0.1",
					ASN:      64501,
					MyASN:    64500,
					RouterID: "11.11.11.300",
				},
			},
			expectedError: "Invalid RouterID",
		},
		{
			desc: "Invalid BGP Peer IP address",
			bgpPeer: &BGPPeer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-bgppeer",
					Namespace: MetalLBTestNameSpace,
				},
				Spec: BGPPeerSpec{
					Address:  "10.0.1",
					ASN:      64501,
					MyASN:    64500,
					RouterID: "10.10.10.10",
				},
			},
			expectedError: "Invalid BGPPeer address",
		},
		{
			desc: "Duplicate BGP Peer",
			bgpPeer: &BGPPeer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-bgppeer",
					Namespace: MetalLBTestNameSpace,
				},
				Spec: BGPPeerSpec{
					Address:  "10.0.0.1",
					ASN:      64501,
					MyASN:    64500,
					RouterID: "10.10.10.10",
				},
			},
			expectedError: "Duplicate BGPPeer",
		},
	}

	for _, test := range tests {
		err := test.bgpPeer.validateBGPPeer(bgpPeerList)
		if err == nil {
			t.Errorf("%s: ValidateBGPPeer failed, no error found while expected: \"%s\"", test.desc, test.expectedError)
		} else {
			if !strings.Contains(fmt.Sprint(err), test.expectedError) {
				t.Errorf("%s: ValidateBGPPeer failed, expected error: \"%s\" to contain: \"%s\"", test.desc, err, test.expectedError)
			}
		}
	}
}
