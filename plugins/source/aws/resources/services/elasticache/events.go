package elasticache

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func Events() *schema.Table {
	tableName := "aws_elasticache_events"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Event.html`,
		Resolver:    fetchElasticacheEvents,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticache"),
		Transform:   transformers.TransformWithStruct(&types.Event{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "_event_hash",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveObjectHash,
				PrimaryKey: true,
			},
		},
	}
}

func fetchElasticacheEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input elasticache.DescribeEventsInput
	c := meta.(*client.Client)
	paginator := elasticache.NewDescribeEventsPaginator(meta.(*client.Client).Services().Elasticache, &input)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(options *elasticache.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- v.Events
	}
	return nil
}
