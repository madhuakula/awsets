package lister

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"

	"github.com/trek10inc/awsets/context"

	"github.com/trek10inc/awsets/resource"
)

type AWSElasticBeanstalkApplication struct {
}

func init() {
	i := AWSElasticBeanstalkApplication{}
	listers = append(listers, i)
}

func (l AWSElasticBeanstalkApplication) Types() []resource.ResourceType {
	return []resource.ResourceType{resource.ElasticBeanstalkApplication}
}

func (l AWSElasticBeanstalkApplication) List(ctx context.AWSetsCtx) (*resource.Group, error) {

	svc := elasticbeanstalk.New(ctx.AWSCfg)

	rg := resource.NewGroup()

	apps, err := svc.DescribeApplicationsRequest(&elasticbeanstalk.DescribeApplicationsInput{}).Send(ctx.Context)
	if err != nil {
		return rg, fmt.Errorf("failed to list elastic beanstalk applications: %w", err)
	}
	for _, v := range apps.Applications {
		r := resource.New(ctx, resource.ElasticBeanstalkApplication, v.ApplicationName, v.ApplicationName, v)
		rg.AddResource(r)
	}

	return rg, nil
}
