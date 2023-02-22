package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils/execext"
)

func archiveTargets(name string, together bool) error {
	releaseTarget := fmt.Sprintf("%s/%s", RELEASE_DIR, name)
	command := `
		rm -rf ${RELEASE}* && mkdir -p ${RELEASE} && \
		mv ${TARGET}/${NAME}-* ${RELEASE}/ && \
		cd ${RELEASE} && \
		shasum -a 256 ${NAME}* > checksum256 && \
		if [ -n "$TOGETHER" ]; then \
			tar -czf ../${NAME}.tar.gz ./* ; \
			cd .. && rm -rf ${NAME} ; \
		else \
			for f in $(ls ${NAME}*); do \
				tar -czf ${f}.tar.gz $f; \
				rm ${f}; \
			done; \
			shasum -a 256 ${NAME}* >> checksum256
		fi
		`
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var env = append(os.Environ(),
		"RELEASE="+releaseTarget,
		"TARGET="+TARGET_DIR,
		"NAME="+name,
	)
	if together {
		env = append(env, "TOGETHER=1")
	}
	opts := &execext.RunCommandOptions{
		Env:     env,
		Command: command,
		Dir:     ".",
		Stdout:  &stdout,
		Stderr:  &stderr,
	}
	if err := execext.RunCommand(context.Background(), opts); err != nil {
		logger.Debug(command)
		logger.Error(err)
		logger.Debug(stdout.String())
		logger.Error(stderr.String())
		return err
	}
	logger.SuccessF("Archived %s successfully", name)
	return nil
}
