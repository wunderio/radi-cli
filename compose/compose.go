package compose
/**
 * Wrapper for libCompose
 */

import (
	"github.com/james-nesbitt/wundertools-go/log"
	"github.com/james-nesbitt/wundertools-go/config"

 	"github.com/docker/libcompose"
)

type Compose struct {

	*libcompose.Project

}