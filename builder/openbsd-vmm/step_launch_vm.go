package openbsdvmm

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/helper/multistep"
)

type stepLaunchVM struct {
	outputPath string
	image  string
	name   string
	mem    string
	kernel string
}

func (step *stepLaunchVM) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	var usedoas bool = true;
	ui := state.Get("ui").(packer.Ui)
	path := filepath.Join(step.outputPath, step.image)

	command := []string{
		"start",
		step.name,
		"-L -i 1",
		"-m " + step.mem,
		"-d " + path,
		"-b " + step.kernel,
	}
	ui.Say("Bring up VM...")
	if err := driver.VmctlCmd(usedoas, command...); err != nil {
		err := fmt.Errorf("Error bringing VM up: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	state.Put("boot_image", path)

	return multistep.ActionContinue
}

func (step *stepLaunchVM) Cleanup(state multistep.StateBag) {
}
