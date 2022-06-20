package workspace

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tliron/kutil/logging"
)

const TestSourceDir = "./testdata"

func TestNew(t *testing.T) {
	type TestCase struct {
		name     string
		rootPath string
		output   *Workspace
	}

	testsCases := []TestCase{
		{
			name:     "new workspace",
			rootPath: "/foo/bar",
			output: &Workspace{
				path:  "/foo/bar",
				plans: nil,
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			ws := New(tt.rootPath, logging.GetLogger("test"))

			assert.NotNil(t, ws)
			assert.Equal(t, tt.output.path, ws.path)
			assert.Equal(t, tt.output.plans, ws.plans)
		})
	}
}

func TestWorkspace_AddPlan(t *testing.T) {
	type TestCase struct {
		name      string
		rootPath  string
		planToAdd string
		planName  string
	}

	testsCases := []TestCase{
		{
			name:      "add simple plan",
			rootPath:  TestSourceDir,
			planToAdd: filepath.Join(TestSourceDir, "main.cue"),
			planName:  "./main.cue",
		},
		{
			name:      "plan in nested source dir",
			rootPath:  filepath.Join(TestSourceDir, "dir-multi-files"),
			planToAdd: filepath.Join(TestSourceDir, "dir-multi-files", "multi.cue"),
			planName:  "./multi.cue",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			ws := New(tt.rootPath, logging.GetLogger("test"))
			assert.NotNil(t, ws)

			err := ws.AddPlan(tt.planToAdd)
			assert.Nil(t, err)

			p := ws.GetPlan(tt.planToAdd)
			assert.NotNil(t, p)
			assert.Equal(t, tt.planName, p.RootFilePath)
		})
	}
}

func TestWorkspace_AddPlan_AlreadyExit(t *testing.T) {
	ws := New(filepath.Join(TestSourceDir, "dir-multi-files"), logging.GetLogger("test"))

	type TestCase struct {
		name       string
		plansToAdd []string
		planPath   string
		nbFiles    int
	}

	testsCases := []TestCase{
		{
			name: "plan in nested source dir",
			plansToAdd: []string{
				filepath.Join(TestSourceDir, "dir-multi-files", "multi.cue"),
				filepath.Join(TestSourceDir, "dir-multi-files", "action.cue"),
				filepath.Join(TestSourceDir, "dir-multi-files", "plan.cue"),
			},
			planPath: filepath.Join(TestSourceDir, "dir-multi-files", "plan.cue"),
			nbFiles:  3,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, ws)

			for _, n := range tt.plansToAdd {
				err := ws.AddPlan(n)
				assert.Nil(t, err)
			}

			p := ws.GetPlan(tt.planPath)
			assert.NotNil(t, p)
			assert.Equal(t, tt.nbFiles, len(p.Files()))
		})
	}
}
