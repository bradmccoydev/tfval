package model

import "time"

type InfracostBreakdown struct {
	Version  string `json:"version"`
	Currency string `json:"currency"`
	Projects []struct {
		Name     string `json:"name"`
		Metadata struct {
			Path       string `json:"path"`
			Type       string `json:"type"`
			VcsRepoURL string `json:"vcsRepoUrl"`
			VcsSubPath string `json:"vcsSubPath"`
		} `json:"metadata"`
		PastBreakdown struct {
			Resources        []interface{} `json:"resources"`
			TotalHourlyCost  string        `json:"totalHourlyCost"`
			TotalMonthlyCost string        `json:"totalMonthlyCost"`
		} `json:"pastBreakdown"`
		Breakdown struct {
			Resources []struct {
				Name     string `json:"name"`
				Metadata struct {
					Calls []struct {
						BlockName string `json:"blockName"`
						Filename  string `json:"filename"`
					} `json:"calls"`
					Filename string `json:"filename"`
				} `json:"metadata"`
				HourlyCost  interface{} `json:"hourlyCost"`
				MonthlyCost interface{} `json:"monthlyCost"`
				Tags        struct {
					Criticality     string `json:"Criticality"`
					Environment     string `json:"Environment"`
					Name            string `json:"Name"`
					Owner           string `json:"Owner"`
					SecurityPosture string `json:"SecurityPosture"`
					SourcePath      string `json:"SourcePath"`
					Terraform       string `json:"Terraform"`
				} `json:"tags,omitempty"`
				CostComponents []struct {
					Name            string `json:"name"`
					Unit            string `json:"unit"`
					HourlyQuantity  string `json:"hourlyQuantity"`
					MonthlyQuantity string `json:"monthlyQuantity"`
					Price           string `json:"price"`
					HourlyCost      string `json:"hourlyCost"`
					MonthlyCost     string `json:"monthlyCost"`
				} `json:"costComponents,omitempty"`
			} `json:"resources"`
			TotalHourlyCost  string `json:"totalHourlyCost"`
			TotalMonthlyCost string `json:"totalMonthlyCost"`
		} `json:"breakdown"`
		Diff struct {
			Resources []struct {
				Name string `json:"name"`
				Tags struct {
					Criticality     string `json:"Criticality"`
					Environment     string `json:"Environment"`
					Name            string `json:"Name"`
					Owner           string `json:"Owner"`
					SecurityPosture string `json:"SecurityPosture"`
					SourcePath      string `json:"SourcePath"`
					Terraform       string `json:"Terraform"`
				} `json:"tags"`
				Metadata struct {
				} `json:"metadata"`
				HourlyCost     string `json:"hourlyCost"`
				MonthlyCost    string `json:"monthlyCost"`
				CostComponents []struct {
					Name            string `json:"name"`
					Unit            string `json:"unit"`
					HourlyQuantity  string `json:"hourlyQuantity"`
					MonthlyQuantity string `json:"monthlyQuantity"`
					Price           string `json:"price"`
					HourlyCost      string `json:"hourlyCost"`
					MonthlyCost     string `json:"monthlyCost"`
				} `json:"costComponents"`
			} `json:"resources"`
			TotalHourlyCost  string `json:"totalHourlyCost"`
			TotalMonthlyCost string `json:"totalMonthlyCost"`
		} `json:"diff"`
		Summary struct {
			TotalDetectedResources    int `json:"totalDetectedResources"`
			TotalSupportedResources   int `json:"totalSupportedResources"`
			TotalUnsupportedResources int `json:"totalUnsupportedResources"`
			TotalUsageBasedResources  int `json:"totalUsageBasedResources"`
			TotalNoPriceResources     int `json:"totalNoPriceResources"`
			UnsupportedResourceCounts struct {
				AwsAPIGatewayIntegrationResponse int `json:"aws_api_gateway_integration_response"`
			} `json:"unsupportedResourceCounts"`
			NoPriceResourceCounts struct {
				AwsAPIGatewayDeployment     int `json:"aws_api_gateway_deployment"`
				AwsAPIGatewayIntegration    int `json:"aws_api_gateway_integration"`
				AwsAPIGatewayMethod         int `json:"aws_api_gateway_method"`
				AwsAPIGatewayMethodResponse int `json:"aws_api_gateway_method_response"`
				AwsAPIGatewayResource       int `json:"aws_api_gateway_resource"`
			} `json:"noPriceResourceCounts"`
		} `json:"summary"`
	} `json:"projects"`
	TotalHourlyCost      string    `json:"totalHourlyCost"`
	TotalMonthlyCost     string    `json:"totalMonthlyCost"`
	PastTotalHourlyCost  string    `json:"pastTotalHourlyCost"`
	PastTotalMonthlyCost string    `json:"pastTotalMonthlyCost"`
	DiffTotalHourlyCost  string    `json:"diffTotalHourlyCost"`
	DiffTotalMonthlyCost string    `json:"diffTotalMonthlyCost"`
	TimeGenerated        time.Time `json:"timeGenerated"`
	Summary              struct {
		TotalDetectedResources    int `json:"totalDetectedResources"`
		TotalSupportedResources   int `json:"totalSupportedResources"`
		TotalUnsupportedResources int `json:"totalUnsupportedResources"`
		TotalUsageBasedResources  int `json:"totalUsageBasedResources"`
		TotalNoPriceResources     int `json:"totalNoPriceResources"`
		UnsupportedResourceCounts struct {
			AwsAPIGatewayIntegrationResponse int `json:"aws_api_gateway_integration_response"`
		} `json:"unsupportedResourceCounts"`
		NoPriceResourceCounts struct {
			AwsAPIGatewayDeployment     int `json:"aws_api_gateway_deployment"`
			AwsAPIGatewayIntegration    int `json:"aws_api_gateway_integration"`
			AwsAPIGatewayMethod         int `json:"aws_api_gateway_method"`
			AwsAPIGatewayMethodResponse int `json:"aws_api_gateway_method_response"`
			AwsAPIGatewayResource       int `json:"aws_api_gateway_resource"`
		} `json:"noPriceResourceCounts"`
	} `json:"summary"`
}
