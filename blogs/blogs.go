package blogs

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

func GetBlogPosts() ([]template.HTML, error) {
	var posts []template.HTML
	files, err := os.ReadDir("./templates/blogs")
	if err != nil {
		return nil, err
	}

	fileInfos := make([]fs.FileInfo, 0, len(files))
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}
		fileInfos = append(fileInfos, info)
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime().After(fileInfos[j].ModTime())
	})

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == ".html" {
			content, err := os.ReadFile(filepath.Join("./templates/blogs", fileInfo.Name()))
			if err != nil {
				return nil, err
			}
			posts = append(posts, template.HTML(content))
		}
	}
	return posts, nil
}
