package plugin

import (
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

// DefaultTagSuffix returns a set of default suggested tags
// based on the commit ref with an attached suffix.
func DefaultTagSuffix(ref, suffix string) ([]string, error) {
	tags, err := DefaultTags(ref)
	if err != nil {
		return nil, err
	}

	if len(suffix) == 0 {
		return tags, nil
	}

	for i, tag := range tags {
		if tag == "latest" {
			tags[i] = suffix
		} else {
			tags[i] = fmt.Sprintf("%s-%s", tag, suffix)
		}
	}

	return tags, nil
}

func splitOff(input, delim string) string {
	const splits = 2
	parts := strings.SplitN(input, delim, splits)

	if len(parts) == splits {
		return parts[0]
	}

	return input
}

// DefaultTags returns a set of default suggested tags based on
// the commit ref.
func DefaultTags(ref string) ([]string, error) {
	if !strings.HasPrefix(ref, "refs/tags/") {
		return []string{"latest"}, nil
	}

	rawVersion := stripTagPrefix(ref)

	version, err := semver.NewVersion(rawVersion)
	if err != nil {
		return []string{"latest"}, err
	}

	if version.PreRelease != "" || version.Metadata != "" {
		return []string{
			version.String(),
		}, nil
	}

	rawVersion = stripTagPrefix(ref)
	rawVersion = splitOff(splitOff(rawVersion, "+"), "-")
	//nolint:gomnd
	dotParts := strings.SplitN(rawVersion, ".", 3)

	if version.Major == 0 {
		return []string{
			fmt.Sprintf("%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor),
			fmt.Sprintf(
				"%0*d.%0*d.%0*d",
				len(dotParts[0]),
				version.Major,
				len(dotParts[1]),
				version.Minor,
				len(dotParts[2]),
				version.Patch,
			),
		}, nil
	}

	return []string{
		fmt.Sprintf("%0*d", len(dotParts[0]), version.Major),
		fmt.Sprintf("%0*d.%0*d", len(dotParts[0]), version.Major, len(dotParts[1]), version.Minor),
		fmt.Sprintf(
			"%0*d.%0*d.%0*d",
			len(dotParts[0]),
			version.Major,
			len(dotParts[1]),
			version.Minor,
			len(dotParts[2]),
			version.Patch,
		),
	}, nil
}

// UseDefaultTag to keep only default branch for latest tag.
func UseDefaultTag(ref, defaultBranch string) bool {
	if strings.HasPrefix(ref, "refs/tags/") {
		return true
	}

	if stripHeadPrefix(ref) == defaultBranch {
		return true
	}

	return false
}

func stripHeadPrefix(ref string) string {
	return strings.TrimPrefix(ref, "refs/heads/")
}

func stripTagPrefix(ref string) string {
	ref = strings.TrimPrefix(ref, "refs/tags/")
	ref = strings.TrimPrefix(ref, "v")

	return ref
}
