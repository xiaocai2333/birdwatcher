package repair

import (
	"context"
	"fmt"
	etcdpbv2 "github.com/milvus-io/birdwatcher/proto/v2.2/etcdpb"
	"github.com/milvus-io/birdwatcher/states/etcd/common"
	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// PartitionCommand return repair partition command.
func PartitionCommand(cli clientv3.KV, basePath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "partition",
		Aliases: []string{"partitions"},
		Short:   "do segment & index meta check and try to repair",
		Run: func(cmd *cobra.Command, args []string) {
			collID, err := cmd.Flags().GetInt64("collectionID")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			partitionID, err := cmd.Flags().GetInt64("partitionID")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			partitionName, err := cmd.Flags().GetString("partitionName")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			run, err := cmd.Flags().GetBool("run")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			ctx := context.Background()
			partitions, err := common.ListCollectionPartitions(ctx, cli, basePath, collID)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for _, partition := range partitions {
				if partition.ID == partitionID {
					fmt.Printf("partition: %d is exist, no need to repair\n", partitionID)
					return
				}
			}

			collections, err := common.ListCollectionsV2(cli, basePath, func(info *etcdpbv2.CollectionInfo) bool {
				return info.ID == collID
			})
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if len(collections) == 0 {
				fmt.Println("no collection found")
				return
			}

			collection := collections[0]
			partition := &etcdpbv2.PartitionInfo{
				PartitionID:               partitionID,
				PartitionName:             partitionName,
				PartitionCreatedTimestamp: collection.CreateTime,
				CollectionId:              collID,
				State:                     etcdpbv2.PartitionState_PartitionCreated,
			}

			fmt.Println("new partition:")
			fmt.Printf("collectionID: %d, partitionID: %d, partitionName: %s, createdTime： %d, state: %s\n",
				partition.CollectionId, partition.PartitionID, partition.PartitionName,
				partition.PartitionCreatedTimestamp, partition.State.String())
			if !run {
				return
			}

			if err := common.WritePartition(ctx, cli, basePath, partition); err != nil {
				fmt.Println(err.Error())
				return
			}
		},
	}
	cmd.Flags().Int64("collectionID", 0, "collection id to filter with")
	cmd.Flags().Int64("partitionID", 0, "partition id to filter with")
	cmd.Flags().String("partitionName", "_default", "partition id to filter with")
	cmd.Flags().Bool("run", false, "actual do repair")
	return cmd
}
