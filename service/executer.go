package service

import "os/exec"

type executer struct{}

type Executer interface {
	Exec(cmd string, useBash bool) ([]byte, error)
}

func NewExecuter() Executer {
	return &executer{}
}

func (e *executer) Exec(cmd string, useBash bool) ([]byte, error) {
	if useBash {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			return out, err
		}
		return out, nil
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			return out, err
		}
		return out, nil
	}
}
