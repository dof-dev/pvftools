package api

import "github.com/wailsapp/wails/v2/pkg/runtime"

func (a *App) SelectDirectory(defaultDir string) (string, error) {
	return runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		DefaultDirectory:     defaultDir,
		CanCreateDirectories: true,
	})
}
