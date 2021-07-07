# Particle
[![License](https://img.shields.io/github/license/little-angry-clouds/particle.svg)](https://github.com/little-angry-clouds/particle/blob/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/little-angry-clouds/particle)](https://goreportcard.com/report/github.com/little-angry-clouds/particle) [![Tests](https://github.com/little-angry-clouds/particle/actions/workflows/generic-tests.yml/badge.svg)](https://github.com/little-angry-clouds/particle/actions/workflows/generic-tests.yml) <a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-81%25-brightgreen.svg?longCache=true&style=flat)</a>

## About Particle
Particle is a project designed to aid in the development and testing of Helm charts and other kubernetes manifests.

It provides a consistent set of steps to execute when testing kubernetes manifests, even when using different tools to do so. Those steps would be:
- Lint your manifests
- Create a kubernetes cluster
- Deploy your manifests
- Verify that what's deployed on the cluster is in its desired state
- Destroy the cluster

It encourages an approach that results in consistently developed manifests, that are well-written, easily understood and maintained.

As you may identified by now, Particle is heavily inspired on [Molecule](https://github.com/ansible-community/molecule), which provides de same as mentioned but for Ansible roles.

## Getting started

You may want to begin with [this](https://github.com/little-angry-clouds/particle/wiki/Getting-Started).

## Get involved

If you have an idea or you want to implement an idea from the roadmap, open an issue and we can talk about it!
