# Smoothly
Configure and control your containerization and IaC in a simple way.  If you are a developer and you don't like to touch topics like pipelines, docker, deployments, etc... Just install smoothly and forget about it.

Current stack supported:
- Angular app running in nginx

# Dev configuration
Requirements:
- Npm
- Angular-cli
- Go
- Docker

Steps:
- Run ng new test-app from root folder
- Make sure you have docker running
- Include or run the export PATH env variable with the path where you clone the repo
- Run go install
- Run smoothly with the flag you wanna test
