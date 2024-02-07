package blogs

import (
	"html/template"
	"os"
	"path/filepath"
)

func GetBlogPosts() ([]template.HTML, error) {
	var posts []template.HTML
	files, err := os.ReadDir("./templates/blogs")
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".html" {
			content, err := os.ReadFile(filepath.Join("./templates/blogs", file.Name()))
			if err != nil {
				return nil, err
			}
			posts = append(posts, template.HTML(content))
		}
	}
	return posts, nil
}
