package v7

import (
	"strings"

	"gopkg.in/yaml.v2"
)

const version = "apps/v1"

// NewDeployment constructer
func NewDeployment(appName string, registryURL string) (*Deployment, *Container) {
	d := Deployment{}
	d.APIVersion = version
	d.Kind = "Deployment"
	d.Metadata.Name = appName + "-deployment"
	d.Metadata.DeploymentLabels.DeploymentApp = appName
	d.DeploymentSpec.Replicas = 1
	d.DeploymentSpec.Selector.MatchLabels.MatchApp = appName
	d.DeploymentSpec.Template.Metadata.Labels.App = appName
	d.DeploymentSpec.Template.Spec.Containers = []Container{*newContainer(appName, registryURL)}

	return &d, &d.DeploymentSpec.Template.Spec.Containers[0]
}

func newContainer(appName string, registryURL string) *Container {
	c := Container{}
	c.Name = appName
	c.Image = registryURL + "/" + appName + ":v1"
	return &c
}

/*
 * The gopkg.in/yaml.v2 Marshaller always wraps double
 * quotes in single quotes makeing things like:
 *
 * 		image: '"dockerimage:v1"'
 *
 * I have resorted to placing double quote tokens around strings
 * that should be double quoted prior to marshaling, and then
 * replacing the tokens with actual double quotes after
 * marshling into a string.
 *
 * The final output looks like:
 *
 *			image: "dockerimage:v1"
 */
const dblQuoteToken = "&#dbl_quote#&"

// Deployment structure
type Deployment struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name             string `yaml:"name"`
		DeploymentLabels struct {
			DeploymentApp string `yaml:"app"`
		} `yaml:"labels"`
	} `yaml:"metadata"`
	DeploymentSpec struct {
		Replicas int
		Selector struct {
			MatchLabels `yaml:"matchLabels"`
		} `yaml:"selector"`
		Template `yaml:"template"`
	} `yaml:"spec"`
}

// Marshal Deployment Struct to string
func (d Deployment) Marshal() string {
	d.DeploymentSpec.Template.Spec.doubleQuote()

	byteArray, err := yaml.Marshal(&d)
	check(err)

	return doubleQuote(byteArray)
}

// Unmarshal string to Struct Deployment
func Unmarshal(s string) Deployment {
	deployment := Deployment{}

	err := yaml.Unmarshal([]byte(s), &deployment)
	check(err)

	return deployment
}

// MatchLabels struct {
type MatchLabels struct {
	MatchApp string `yaml:"app"`
}

// Template structure
type Template struct {
	Metadata `yaml:"metadata"`
	Spec     `yaml:"spec"`
}

// Metadata structure
type Metadata struct {
	Labels `yaml:"labels"`
}

// Labels structure
type Labels struct {
	App string `yaml:"app"`
}

// Spec structure
type Spec struct {
	Containers []Container `yaml:"containers"`
}

func (s *Spec) doubleQuote() {
	for i, c := range s.Containers {
		s.Containers[i] = c.doubleQuote()
	}
}

// Container structure
type Container struct {
	Name      string `yaml:"name"`
	Image     string `yaml:"image"`
	Resources `yaml:"resources,omitempty"`
	Envs      []Env    `yaml:"env,omitempty"`
	Ports     []Port   `yaml:"ports,omitempty"`
	Command   []string `yaml:"command,flow,omitempty"`
	Args      []string `yaml:"args,flow,omitempty"`
}

// AddEnv and an Env
func (c *Container) AddEnv(n string, v interface{}) {
	if c.Envs == nil {
		c.Envs = []Env{}
	}
	e := Env{}
	e.setEnv(n, v)
	c.Envs = append(c.Envs, e)
}

func (c *Container) doubleQuote() Container {
	c.Image = dblQuoteToken + c.Image + dblQuoteToken
	for i, env := range c.Envs {
		c.Envs[i] = env.doubleQuote()
	}
	for i, cmd := range c.Command {
		c.Command[i] = dblQuoteToken + cmd + dblQuoteToken
	}
	for i, arg := range c.Args {
		c.Args[i] = dblQuoteToken + arg + dblQuoteToken
	}
	return *c
}

// Resources structure
type Resources struct {
	Limits   `yaml:"limits,omitempty"`
	Requests `yaml:"requests,omitempty"`
}

// Limits structure
type Limits struct {
	CPU    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory"`
}

// Requests structure
type Requests struct {
	CPU    string `yaml:"cpu,omitempty"`
	Memory string `yaml:"memory"`
}

// Env Structure
type Env struct {
	Name  string      `yaml:"name"`
	Value interface{} `yaml:"value"`
}

func (e *Env) doubleQuote() Env {
	s, ok := e.Value.(string)
	if ok {
		e.Value = dblQuoteToken + s + dblQuoteToken
	}
	return *e
}

func (e *Env) setEnv(n string, v interface{}) Env {
	e.Name = n
	e.Value = v
	return *e
}

// Port structure
type Port struct {
	ContainerPort int `yaml:"containerPort"`
}

func doubleQuote(b []byte) string {
	s := strings.ReplaceAll(string(b), "'"+dblQuoteToken, "\"")
	return strings.ReplaceAll(s, dblQuoteToken+"'", "\"")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
