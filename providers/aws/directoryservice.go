package aws

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

var directoryserviceAllowEmptyValues = []string{}

type DirectoryserviceGenerator struct {
	AWSService
	client *directoryservice.Client
}

func (g *DirectoryserviceGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	g.client = directoryservice.NewFromConfig(config)
	if err := g.loadDirectories(g.client); err != nil {
		return err
	}
	return nil
}

func (g *DirectoryserviceGenerator) DescribeDirectories(ctx context.Context, params *directoryservice.DescribeDirectoriesInput, opts ...func(*directoryservice.Options)) (*directoryservice.DescribeDirectoriesOutput, error) {
	return g.client.DescribeDirectories(ctx, params, opts...)
}

func (g *DirectoryserviceGenerator) loadDirectories(svc *directoryservice.Client) error {
	p := directoryservice.NewDescribeDirectoriesPaginator(g, &directoryservice.DescribeDirectoriesInput{})

	for p.HasMorePages() {
		page, err := p.NextPage(context.TODO())
		if err != nil {
			return err
		}
		for _, dir := range page.DirectoryDescriptions {
			directoryID := StringValue(dir.DirectoryId)
			directoryName := StringValue(dir.Name)
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				directoryID,
				directoryName,
				"aws_directory_service_directory",
				"aws",
				directoryserviceAllowEmptyValues))
		}
	}
	return nil
}
