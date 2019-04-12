package command

import "github.com/unjello/commit-activity-log/config"

func Execute() {
	config.Load()
	List()
}
