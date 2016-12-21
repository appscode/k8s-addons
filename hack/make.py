#!/usr/bin/env python

import subprocess
import sys
from os.path import expandvars, expanduser

REPO_ROOT = expanduser('~') + '/go/src/github.com/appscode/k8s-addons'


def call(cmd, stdin=None, cwd=REPO_ROOT):
    print(cmd)
    return subprocess.call([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)


def die(status):
    if status:
        sys.exit(status)


def fmt():
    die(call('goimports -w api client pkg'))
    call('gofmt -s -w api client pkg')


def vet():
    call('go vet ./api/... ./client/... ./pkg/...')


def lint():
    call('golint ./api/...')
    call('golint ./client/...')
    call('golint ./pkg/...')


def default():
    fmt()
    die(call('GO15VENDOREXPERIMENT=1 go install ./...'))


if __name__ == "__main__":
    default()
