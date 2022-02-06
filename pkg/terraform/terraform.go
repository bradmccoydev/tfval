package terraform

import (
	"context"
	"fmt"
	"log"

	config "github.com/bradmccoydev/terraform-plan-validator/util"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

func ProduceTerraformPlan(cfg config.Config) string {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.1.4")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	}

	workingDir := "/Users/bmccoy/Development/bradmccoydev/aks/src"
	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	// err = tf.Init(context.Background(), tfexec.Upgrade(true))
	// if err != nil {
	// 	log.Fatalf("error running Init: %s", err)
	// }

	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	fmt.Println(state.FormatVersion) // "0.1"

	return "hi"
}
