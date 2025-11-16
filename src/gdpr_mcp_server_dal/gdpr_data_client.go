package gdpr_mcp_server_dal

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/settings"
	"go.uber.org/zap"
)

type GdprDataClient struct {
	logger       *zap.Logger
	dataSettings *settings.DataSettings

	recitalsSet map[string]*models.Recital
	chaptersSet map[string]*models.Chapter

	// articlesSet and articleParagraphsSet use the same key (article ID)
	articlesSet          map[string]*models.Article
	articleParagraphsSet map[string][]*models.ArticleParagraph

	mu sync.RWMutex
}

func NewGdprDataClient(dataSettings *settings.DataSettings, logger *zap.Logger) (*GdprDataClient, error) {
	c := &GdprDataClient{
		logger:               logger,
		dataSettings:         dataSettings,
		recitalsSet:          make(map[string]*models.Recital),
		chaptersSet:          make(map[string]*models.Chapter),
		articlesSet:          make(map[string]*models.Article),
		articleParagraphsSet: make(map[string][]*models.ArticleParagraph),
	}

	if err := c.loadData(); err != nil {
		return nil, err
	}

	return c, nil
}

func decodeJSONFile[T any](path string, out *T) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(out)
}

func listDirEntries(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("read dir error (%s): %v", dir, err)
		return nil
	}

	return entries
}

func (c *GdprDataClient) loadData() error {
	var wg sync.WaitGroup
	wg.Add(4)

	var errsMu sync.Mutex
	var errs []error

	runLoader := func(f func() error) {
		defer wg.Done()
		if err := f(); err != nil {
			errsMu.Lock()
			errs = append(errs, err)
			errsMu.Unlock()
		}
	}

	go runLoader(c.loadRecitals)
	go runLoader(c.loadChapters)
	go runLoader(c.loadArticles)
	go runLoader(c.loadArticleParagraphs)

	wg.Wait()

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (c *GdprDataClient) loadRecitals() error {
	dir := c.dataSettings.RecitalsDataFilePath
	for _, e := range listDirEntries(dir) {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		path := filepath.Join(dir, e.Name())
		var r models.Recital
		if err := decodeJSONFile(path, &r); err != nil {
			log.Printf("recital decode error (%s): %v", path, err)
			continue
		}
		if r.ID == "" {
			return fmt.Errorf("recital missing ID (path=%s)", path)
		}
		c.recitalsSet[r.ID] = &r
	}

	return nil
}

func (c *GdprDataClient) loadChapters() error {
	dir := c.dataSettings.ChaptersDataFilePath
	for _, e := range listDirEntries(dir) {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		path := filepath.Join(dir, e.Name())
		var ch models.Chapter
		if err := decodeJSONFile(path, &ch); err != nil {
			log.Printf("chapter decode error (%s): %v", path, err)
			continue
		}
		if ch.ID == "" {
			return fmt.Errorf("chapter missing ID (path=%s)", path)
		}
		c.chaptersSet[ch.ID] = &ch
	}

	return nil
}

func (c *GdprDataClient) loadArticles() error {
	dir := c.dataSettings.ArticlesDataFilePath
	for _, d := range listDirEntries(dir) {
		if !d.IsDir() {
			continue
		}
		artPath := filepath.Join(dir, d.Name(), "art.json")
		var a models.Article
		if err := decodeJSONFile(artPath, &a); err != nil {
			// Some directories may not yet have art.json; skip silently.
			continue
		}
		if a.ID == "" {
			return fmt.Errorf("article missing ID (path=%s)", artPath)
		}
		c.articlesSet[a.ID] = &a
	}

	return nil
}

func (c *GdprDataClient) loadArticleParagraphs() error {
	root := c.dataSettings.ArticlesDataFilePath
	for _, d := range listDirEntries(root) {
		if !d.IsDir() {
			continue
		}
		subdir := filepath.Join(root, d.Name())
		for _, fEnt := range listDirEntries(subdir) {
			if fEnt.IsDir() {
				continue
			}
			name := fEnt.Name()
			if !strings.HasPrefix(name, "para-") || filepath.Ext(name) != ".json" {
				continue
			}
			path := filepath.Join(subdir, name)
			var p models.ArticleParagraph
			if err := decodeJSONFile(path, &p); err != nil {
				log.Printf("paragraph decode error (%s): %v", path, err)
				continue
			}
			if p.ArticleId == "" {
				return fmt.Errorf("paragraph missing ArticleId (path=%s)", path)
			}
			c.articleParagraphsSet[p.ArticleId] = append(c.articleParagraphsSet[p.ArticleId], &p)
		}
	}

	return nil
}

func (c *GdprDataClient) RecitalsSetSnapshot() map[string]*models.Recital {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make(map[string]*models.Recital, len(c.recitalsSet))
	for id, r := range c.recitalsSet {
		copyVal := *r
		out[id] = &copyVal
	}
	return out
}

func (c *GdprDataClient) ChaptersSetSnapshot() map[string]*models.Chapter {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make(map[string]*models.Chapter, len(c.chaptersSet))
	for id, ch := range c.chaptersSet {
		copyVal := *ch
		out[id] = &copyVal
	}
	return out
}

func (c *GdprDataClient) ArticlesSetSnapshot() map[string]*models.Article {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make(map[string]*models.Article, len(c.articlesSet))
	for id, a := range c.articlesSet {
		copyVal := *a
		out[id] = &copyVal
	}
	return out
}

func (c *GdprDataClient) ArticleParagraphsSetSnapshot() map[string][]*models.ArticleParagraph {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make(map[string][]*models.ArticleParagraph, len(c.articleParagraphsSet))
	for id, ps := range c.articleParagraphsSet {
		cp := make([]*models.ArticleParagraph, len(ps))
		copy(cp, ps)
		out[id] = cp
	}
	return out
}
