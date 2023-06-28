package plugin

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
	"golang.org/x/sys/execabs"
)

// helper function to create the docker login command.
func commandLogin(login Login) *execabs.Cmd {
	if login.Email != "" {
		return commandLoginEmail(login)
	}

	args := []string{
		"login",
		"-u", login.Username,
		"-p", login.Password,
		login.Registry,
	}

	return execabs.Command(
		dockerBin, args...,
	)
}

func commandLoginEmail(login Login) *execabs.Cmd {
	args := []string{
		"login",
		"-u", login.Username,
		"-p", login.Password,
		"-e", login.Email,
		login.Registry,
	}

	return execabs.Command(
		dockerBin, args...,
	)
}

// helper function to create the docker info command.
func commandVersion() *execabs.Cmd {
	return execabs.Command(dockerBin, "version")
}

// helper function to create the docker info command.
func commandInfo() *execabs.Cmd {
	return execabs.Command(dockerBin, "info")
}

func commandBuilder(daemon Daemon) *execabs.Cmd {
	args := []string{
		"buildx",
		"create",
		"--use",
	}

	if daemon.BuildkitConfig != "" {
		args = append(args, "--config", buildkitConfig)
	}

	return execabs.Command(dockerBin, args...)
}

func commandBuildx() *execabs.Cmd {
	return execabs.Command(dockerBin, "buildx", "ls")
}

// helper function to create the docker build command.
func commandBuild(build Build, dryrun bool) *execabs.Cmd {
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
	if !dryrun && build.Output == "" && len(build.Tags.Value()) > 0 {
		args = append(args, "--push")
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

	for _, arg := range build.CacheFrom {
		args = append(args, "--cache-from", arg)
	}

	if build.CacheTo != "" {
		args = append(args, "--cache-to", build.CacheTo)
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

	if build.Output != "" {
		args = append(args, "--output", build.Output)
	}

	for _, arg := range build.NamedContext.Value() {
		args = append(args, "--build-context", arg)
	}

	if len(build.Platforms.Value()) > 0 {
		args = append(args, "--platform", strings.Join(build.Platforms.Value(), ","))
	}

	for _, arg := range build.Tags.Value() {
		args = append(args, "-t", fmt.Sprintf("%s:%s", build.Repo, arg))
	}

	for _, arg := range build.ExtraTags.Value() {
		args = append(args, "-t", arg)
	}

	for _, arg := range build.Labels.Value() {
		args = append(args, "--label", arg)
	}

	if build.Provenance != "" {
		args = append(args, "--provenance", build.Provenance)
	}

	if build.SBOM != "" {
		args = append(args, "--sbom", build.SBOM)
	}

	if build.Secret != "" {
		args = append(args, "--secret", build.Secret)
	}

	for _, secret := range build.SecretEnvs.Value() {
		if arg, err := getSecretStringCmdArg(secret); err == nil {
			args = append(args, "--secret", arg)
		}
	}

	for _, secret := range build.SecretFiles.Value() {
		if arg, err := getSecretFileCmdArg(secret); err == nil {
			args = append(args, "--secret", arg)
		}
	}

	return execabs.Command(dockerBin, args...)
}

// helper function to parse string secret key-pair
func getSecretStringCmdArg(kvp string) (string, error) {
	return getSecretCmdArg(kvp, false)
}

// helper function to parse file secret key-pair
func getSecretFileCmdArg(kvp string) (string, error) {
	return getSecretCmdArg(kvp, true)
}

// helper function to parse secret key-pair
func getSecretCmdArg(kvp string, file bool) (string, error) {
	delimIndex := strings.IndexByte(kvp, '=')
	if delimIndex == -1 {
		return "", fmt.Errorf("%s is not a valid secret", kvp)
	}

	key := kvp[:delimIndex]
	value := kvp[delimIndex+1:]

	if key == "" || value == "" {
		return "", fmt.Errorf("%s is not a valid secret", kvp)
	}

	if file {
		return fmt.Sprintf("id=%s,src=%s", key, value), nil
	}

	return fmt.Sprintf("id=%s,env=%s", key, value), nil
}

// helper function to add proxy values from the environment.
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
func commandDaemon(daemon Daemon) *execabs.Cmd {
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

	return execabs.Command(dockerdBin, args...)
}

// trace writes each command to stdout with the command wrapped in an xml
// tag so that it can be extracted and displayed in the logs.
func trace(cmd *execabs.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}
