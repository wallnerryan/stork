// +build integrationtest

package integrationtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testBasicCloudMigartion(t *testing.T) {
	t.Run("sanityCloudMigrationTest", sanityCloudMigrationTest)
}

func sanityCloudMigrationTest(t *testing.T) {
	// TODO: get px token information from torpedo' get storage token call
	// once it's merged in topedo , govendor and upate here
	specReq := ClusterPairRequest{
		PairName:       remotePairName,
		ConfigMapName:  remoteConfig,
		SpecDirPath:    ".",
		PxIP:           "test-ip",
		PxClusterToken: token,
		PxPort:         "9001",
	}

	err := CreateClusterPairSpec(specReq)
	require.NoError(t, err, "Error creating cluster Spec")

	// appply cluster pair spec and check status
	// apply cloudmigration spec and check status
}
