package common

import (
	"context"
	"fmt"
	"path"
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/samber/lo"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/milvus-io/birdwatcher/models"
	etcdpbv2 "github.com/milvus-io/birdwatcher/proto/v2.2/etcdpb"
)

const (
	PartitionPrefix = `root-coord/partitions/`
)

// ListCollectionPartitions returns partition list of collection.
func ListCollectionPartitions(ctx context.Context, cli clientv3.KV, basePath string, collectionID int64) ([]*models.Partition, error) {
	prefix := path.Join(basePath, PartitionPrefix, fmt.Sprintf("%d", collectionID))

	infos, keys, err := ListProtoObjects[etcdpbv2.PartitionInfo](ctx, cli, prefix)
	if err != nil {
		return nil, err
	}

	return lo.Map(infos, func(info etcdpbv2.PartitionInfo, idx int) *models.Partition {
		return models.NewPartition(&info, keys[idx])
	}), nil
}

func WritePartition(ctx context.Context, cli clientv3.KV, basePath string, partition *etcdpbv2.PartitionInfo) error {
	prefix := []string{
		path.Join(basePath, PartitionPrefix, strconv.FormatInt(partition.CollectionId, 10), strconv.FormatInt(partition.PartitionID, 10)),
		path.Join(basePath, SnapshotPrefix, PartitionPrefix, strconv.FormatInt(partition.CollectionId, 10),
			fmt.Sprintf("%d_ts%d", partition.PartitionID, partition.PartitionCreatedTimestamp)),
	}

	bs, err := proto.Marshal(partition)
	if err != nil {
		return err
	}
	for _, key := range prefix {
		_, err = cli.Put(ctx, key, string(bs))
		if err != nil {
			return err
		}
	}

	return nil
}
