package servicetemplate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/mesg-foundation/core/x/xgit"
	"github.com/mesg-foundation/core/x/xos"
)

const templatesURL = "https://raw.githubusercontent.com/mesg-foundation/awesome/master/templates.json"

type Template struct {
	Name string
	URL  string
}

func (s *Template) String() string {
	return fmt.Sprintf("%s (%s)", s.Name, s.URL)
}

type ConfigOption struct {
	Name        string
	Description string
}

func List() ([]*Template, error) {
	resp, err := http.Get(templatesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var list []*Template
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func Download(t *Template, dst string) error {
	path, err := ioutil.TempDir("", "mesg-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(path)

	if err = xgit.Clone(t.URL, path); err != nil {
		return err
	}

	return xos.CopyDir(filepath.Join(path, "template"), dst)
}

func ExecuteTemplateConfig(dst string, option ConfigOption) error {
	var (
		file = filepath.Join(dst, "mesg.yml")
		buf  *bytes.Buffer
	)

	t, err := template.ParseFiles(file)
	if err != nil {
		return err
	}

	si, err := os.Stat(file)
	if err != nil {
		return err
	}

	if err := t.Execute(buf, option); err != nil {
		return err
	}

	return ioutil.WriteFile(file, buf.Bytes(), si.Mode())
}
