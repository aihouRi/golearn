package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}


//以下是发现问题后修改的，更新github后本地会删除它（2024年10月23日）
func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// 获取所有页面模板文件
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// 获取所有布局模板文件，只需要查找一次
	layoutFiles, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return myCache, err
	}

	// 遍历每个页面模板
	for _, page := range pages {
		name := filepath.Base(page)

		// 解析页面模板
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// 如果有布局模板，解析布局模板
		if len(layoutFiles) > 0 {
			ts, err = ts.ParseFiles(layoutFiles...)
			if err != nil {
				return myCache, err
			}
		}

		// 将解析后的模板存入缓存
		myCache[name] = ts
	}

	return myCache, nil
}

