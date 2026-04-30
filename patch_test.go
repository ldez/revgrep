package revgrep

import (
	"context"
	"testing"
)

func TestGitPatch_nonGitDir(t *testing.T) {
	t.Chdir(t.TempDir())

	patch, newFiles, err := GitPatch(context.Background(), patchOption{})
	if err != nil {
		t.Errorf("error expected nil, got: %v", err)
	}

	if patch != nil {
		t.Errorf("patch expected nil, got: %v", patch)
	}

	if newFiles != nil {
		t.Errorf("newFiles expected nil, got: %v", newFiles)
	}
}
