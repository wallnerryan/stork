// +build integrationtest

package integrationtest

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func testBasicCloudMigration(t *testing.T) {
	t.Run("sanityCloudMigrationTest", sanityCloudMigrationTest)
}

func sanityCloudMigrationTest(t *testing.T) {
	// TODO: get px token information from torpedo' get storage token call
	// once it's merged in topedo , govendor and upate here
	logrus.Info("Basic sanity tests for cloud migration")
	err := dumpRemoteKubeConfig(remoteConfig)
	logrus.Info("file writing to /opt", err)
	token, err := volumeDriver.GetStorageInfo()
	logrus.Info("token received:", token, err)
	specReq := ClusterPairRequest{
		PairName:       remotePairName,
		ConfigMapName:  remoteConfig,
		SpecDirPath:    "./",
		PxIP:           "test-ip",
		PxClusterToken: "53f94470b8205820bd3bdaf5a8f09f9b18a948058c56f8344b6429f5b3e58c6d60b0909e55ba51cc9ab2d6c88579f27f85dddf5998c4ab3169c75ec18ee14d64",
		PxPort:         "9001",
	}
	logrus.Info("writing to spec file", specReq)
	err = CreateClusterPairSpec(specReq)
	require.NoError(t, err, "Error creating cluster Spec")
	// appply cluster pair spec and check status
	_, err = schedulerDriver.CreateCRDObjects("./" + pairFileName)
	// apply cloudmigration spec and check status
}
