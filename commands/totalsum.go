package commands
import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
        "fmt"
        "os/exec"
	"strings"
)

func GetTotalSum() components.Command {
	return components.Command{
		Name:        "totalsum",
		Description: "This command results in the Top 20 downloads",
		Aliases:     []string{"TotalSum"},
		Arguments:   getHelloArguments(),
		Flags:       getHelloFlags(),
		EnvVars:     getHelloEnvVar(),
		Action: func(c *components.Context) error {
			return sumofdatausage(c)
		},
	}
}

func getHelloArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "addressee",
			Description: "The name of the person you would like to greet.",
		},
	}
}

func getHelloFlags() []components.Flag {
	return []components.Flag{
		components.BoolFlag{
			Name:         "shout",
			Description:  "Makes output uppercase.",
			DefaultValue: false,
		},
		components.StringFlag{
			Name:         "repeat",
			Description:  "Greets multiple times.",
			DefaultValue: "1",
		},
	}
}

func getHelloEnvVar() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

type helloConfiguration struct {
	addressee string
	shout     bool
	repeat    int
	prefix    string
}

func doGreet(c *helloConfiguration) string {
	greet := c.prefix + "Hello " + c.addressee + "!\n"

	if c.shout {
		greet = strings.ToUpper(greet)
	}

	return strings.TrimSpace(strings.Repeat(greet, c.repeat))
}

func sumofdatausage(c *components.Context) error {
	cmd, err := exec.Command("/bin/sh", "sumofdatausage.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println("total size of artifacts in bytes")
	fmt.Println(string(cmd))
	return nil
}
