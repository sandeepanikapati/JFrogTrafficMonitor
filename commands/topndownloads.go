package commands
import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
        "fmt"
        "os/exec"
	"strings"
)

func TopNDownloads() components.Command {
	return components.Command{
		Name:        "top20d",
		Description: "This command results in the Top 20 downloads",
		Aliases:     []string{"Top20D"},
		Arguments:   getHelloArguments1(),
		Flags:       getHelloFlags1(),
		EnvVars:     getHelloEnvVar1(),
		Action: func(c *components.Context) error {
			return top20downloads(c)
		},
	}
}

func getHelloArguments1() []components.Argument {
	return []components.Argument{
		{
			Name:        "addressee",
			Description: "The name of the person you would like to greet.",
		},
	}
}

func getHelloFlags1() []components.Flag {
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

func getHelloEnvVar1() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

type helloConfiguration1 struct {
	addressee string
	shout     bool
	repeat    int
	prefix    string
}

func doGreet1(c *helloConfiguration1) string {
	greet := c.prefix + "Hello " + c.addressee + "!\n"

	if c.shout {
		greet = strings.ToUpper(greet)
	}

	return strings.TrimSpace(strings.Repeat(greet, c.repeat))
}

func top20downloads(c *components.Context) error {
	cmd, err := exec.Command("/bin/sh", "top20downloads.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println("Noofdownloads repository path Sizeoftheartifact")
	fmt.Println(string(cmd))
	return nil
}
