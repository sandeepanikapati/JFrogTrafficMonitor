package main

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/commands"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "JFrogTrafficMonitor"
	app.Description = "Monitors all traffic of the repositories "
	app.Version = "v0.0.1"
	app.Commands = getCommands()
	return app
}
func getCommands() []components.Command {
        return []components.Command{
                commands.TopNDownloads(),
                commands.GetTotalSum(),  
                commands.TopNUploads(),           
}
}
