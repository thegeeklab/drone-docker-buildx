package plugin

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

// helper function to create the docker login command.
func commandLogin(login Login) *exec.Cmd {
	if login.Email != "" {
		return commandLoginEmail(login)
	}
	return exec.Command(
		dockerExe, "login",
		"-u", login.Username,
		"-p", login.Password,
		login.Registry,
	)
}

// helper to check if args match "docker pull <image>"
func isCommandPull(args []string) bool {
	return len(args) > 2 && args[1] == "pull"
}

func commandPull(repo string) *exec.Cmd {
	return exec.Command(dockerExe, "pull", repo)
}

func commandLoginEmail(login Login) *exec.Cmd {
	return exec.Command(
		dockerExe, "login",
		"-u", login.Username,
		"-p", login.Password,
		"-e", login.Email,
		login.Registry,
	)
}

// helper function to create the docker info command.
func commandVersion() *exec.Cmd {
	return exec.Command(dockerExe, "version")
}

// helper function to create the docker info command.
func commandInfo() *exec.Cmd {
	return exec.Command(dockerExe, "info")
}

func commandBuilder(daemon Daemon) *exec.Cmd {
	args := []string{
		"buildx", 
		"create", 
		"--use",
	}

	if(daemon.Config != "") {
		args = append(args, "--config", daemon.Config)
	}

	return exec.Command(dockerExe, args...)
}

func commandBuildx() *exec.Cmd {
	return exec.Command(dockerExe, "buildx", "ls")
}

// helper function to create the docker build command.
func commandBuild(build Build, dryrun bool) *exec.Cmd {
	args := []string{
		"buildx",
		"build",
		"--rm=true",
		"-f", build.Dockerfile,
	}

	defaultBuildArgs := []string{
		fmt.Sprintf("DOCKER_IMAGE_CREATED=%s", time.Now().Format(time.RFC3339)),
	}

	args = append(args, build.Context)
	if ! dryrun {
		args = append(args, "--push")
	}
	if build.Squash {
		args = append(args, "--squash")
	}
	if build.Compress {
		args = append(args, "--compress")
	}
	if build.Pull {
		args = append(args, "--pull=true")
	}
	if build.NoCache {
		args = append(args, "--no-cache")
	}
	for _, arg := range build.CacheFrom.Value() {
		args = append(args, "--cache-from", arg)
	}
	for _, arg := range build.ArgsEnv.Value() {
		addProxyValue(&build, arg)
	}
	for _, arg := range append(defaultBuildArgs, build.Args.Value()...) {
		args = append(args, "--build-arg", arg)
	}
	for _, host := range build.AddHost.Value() {
		args = append(args, "--add-host", host)
	}
	if build.Target != "" {
		args = append(args, "--target", build.Target)
	}
	if build.Quiet {
		args = append(args, "--quiet")
	}

	if len(build.Platforms.Value()) > 0 {
		args = append(args, "--platform", strings.Join(build.Platforms.Value()[:], ","))
	}

	for _, arg := range build.Tags.Value() {
		args = append(args, "-t", fmt.Sprintf("%s:%s", build.Repo, arg))
	}	

	return exec.Command(dockerExe, args...)
}

// helper function to add proxy values from the environment
func addProxyBuildArgs(build *Build) {
	addProxyValue(build, "http_proxy")
	addProxyValue(build, "https_proxy")
	addProxyValue(build, "no_proxy")
}

// helper function to add the upper and lower case version of a proxy value.
func addProxyValue(build *Build, key string) {
	value := getProxyValue(key)

	if len(value) > 0 && !hasProxyBuildArg(build, key) {
		build.Args = *cli.NewStringSlice(append(build.Args.Value(), fmt.Sprintf("%s=%s", key, value))...)
		build.Args = *cli.NewStringSlice(append(build.Args.Value(), fmt.Sprintf("%s=%s", strings.ToUpper(key), value))...)
	}
}

// helper function to get a proxy value from the environment.
//
// assumes that the upper and lower case versions of are the same.
func getProxyValue(key string) string {
	value := os.Getenv(key)

	if len(value) > 0 {
		return value
	}

	return os.Getenv(strings.ToUpper(key))
}

// helper function that looks to see if a proxy value was set in the build args.
func hasProxyBuildArg(build *Build, key string) bool {
	keyUpper := strings.ToUpper(key)

	for _, s := range build.Args.Value() {
		if strings.HasPrefix(s, key) || strings.HasPrefix(s, keyUpper) {
			return true
		}
	}

	return false
}

// helper function to create the docker daemon command.
func commandDaemon(daemon Daemon) *exec.Cmd {
	args := []string{
		"--data-root", daemon.StoragePath,
		"--host=unix:///var/run/docker.sock",
	}

	if daemon.StorageDriver != "" {
		args = append(args, "-s", daemon.StorageDriver)
	}
	if daemon.Insecure && daemon.Registry != "" {
		args = append(args, "--insecure-registry", daemon.Registry)
	}
	if daemon.IPv6 {
		args = append(args, "--ipv6")
	}
	if len(daemon.Mirror) != 0 {
		args = append(args, "--registry-mirror", daemon.Mirror)
	}
	if len(daemon.Bip) != 0 {
		args = append(args, "--bip", daemon.Bip)
	}
	for _, dns := range daemon.DNS.Value() {
		args = append(args, "--dns", dns)
	}
	for _, dnsSearch := range daemon.DNSSearch.Value() {
		args = append(args, "--dns-search", dnsSearch)
	}
	if len(daemon.MTU) != 0 {
		args = append(args, "--mtu", daemon.MTU)
	}
	if daemon.Experimental {
		args = append(args, "--experimental")
	}
	return exec.Command(dockerdExe, args...)
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}
