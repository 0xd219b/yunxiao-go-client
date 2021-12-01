package config

type Config struct {
	AK       string        `yaml:"ak"`
	SK       string        `yaml:"sk"`
	Endpoint string        `yaml:"endpoint"`
	Project  ProjectConfig `yaml:"project"`
	Task     TaskConfig    `yaml:"task"`
	Action   string        `yaml:"action"`
}

type ProjectConfig struct {
	ProjectID string `yaml:"project_id"`
	OrgID     string `yaml:"org_id"`
}

type TaskConfig struct {
	TaskID string `yaml:"task_id"`
	Status string `yaml:"status"`
}

var Cfg Config

func GetAK() string {
	return Cfg.AK
}

func GetSK() string {
	return Cfg.SK
}

func GetEndpoint() string {
	return Cfg.Endpoint
}

func GetProjectID() string {
	return Cfg.Project.ProjectID
}

func GetOrgID() string {
	return Cfg.Project.OrgID
}

func GetAction() string {
	return Cfg.Action
}
