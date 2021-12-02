package commands
import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
        "fmt"
        "os/exec"
	"strings"
)

func TopNUploads() components.Command {
	return components.Command{
		Name:        "top20u",
		Description: "This command results in the Top 20 uploads",
		Aliases:     []string{"Top20U"},
		Arguments:   getHelloArguments2(),
		Flags:       getHelloFlags2(),
		EnvVars:     getHelloEnvVar2(),
		Action: func(c *components.Context) error {
			return topnuploadsmethod(c)
		},
	}
}

func getHelloArguments2() []components.Argument {
	return []components.Argument{
		{
			Name:        "addressee",
			Description: "The name of the person you would like to greet.",
		},
	}
}

func getHelloFlags2() []components.Flag {
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

func getHelloEnvVar2() []components.EnvVar {
	return []components.EnvVar{
		{
			Name:        "HELLO_FROG_GREET_PREFIX",
			Default:     "A new greet from your plugin template: ",
			Description: "Adds a prefix to every greet.",
		},
	}
}

type helloConfiguration2 struct {
	addressee string
	shout     bool
	repeat    int
	prefix    string
}

func doGreet2(c *helloConfiguration2) string {
	greet := c.prefix + "Hello " + c.addressee + "!\n"

	if c.shout {
		greet = strings.ToUpper(greet)
	}

	return strings.TrimSpace(strings.Repeat(greet, c.repeat))
}

func topnuploadsmethod(c *components.Context) error {
	cmd, err := exec.Command("/bin/sh", "top20uploads.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	fmt.Println("Noofdownloads repository path Sizeoftheartifact")
	fmt.Println(string(cmd))
	return nil
}
