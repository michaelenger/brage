package cmd

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"brage/site"
	"brage/utils"
	"github.com/spf13/cobra"
)

// Port to serve the site on
var port int32

// Server handler based on a Site
type siteHandler struct {
	sitePath string
	logger   *log.Logger
}

func runServeCommand(cmd *cobra.Command, args []string) {
	logger := log.Default()

	var sourcePath string
	if len(args) > 0 {
		sourcePath = args[0]
	} else {
		sourcePath = "."
	}
	sourcePath = utils.AbsolutePath(sourcePath)

	logger.Printf("Loading site from: %v", sourcePath)

	// Load the site to ensure that everything we need is there
	_, err := site.Load(sourcePath)
	if err != nil {
		logger.Fatalf("ERROR! Unable to load site: %v", err)
	}

	handler := siteHandler{
		sourcePath,
		logger,
	}

	logger.Printf("Running server on: http://localhost:%v", port)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%v", port),
		Handler:        &handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Fatal(server.ListenAndServe())
}

func (handler *siteHandler) serveFile(assetFile string, w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(assetFile); err != nil {
		handler.logger.Print("404 Not Found")
		http.NotFound(w, r)
		return
	}

	fileBytes, err := os.ReadFile(assetFile)
	if err != nil {
		handler.logger.Print("500 Server Error")
		errorText := fmt.Sprintf("Unable to read asset file: %v", err)
		handler.logger.Print(errorText)
		http.Error(w, errorText, 500)
		return
	}

	mimeType := mime.TypeByExtension(filepath.Ext(assetFile))
	handler.logger.Printf("200 OK %v", mimeType)
	w.Header().Set("Content-Type", mimeType)
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)
}

// Handle an HTTP request on the server
func (handler *siteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path
	handler.logger.Printf("Request: %v %v", r.Method, requestPath)

	site, err := site.Load(handler.sitePath)
	if err != nil {
		handler.logger.Fatalf("ERROR! Unable to load site: %v", err)
	}

	if len(requestPath) >= 7 && requestPath[:7] == "/assets" {
		assetPath := path.Join(site.SourceDirectory, requestPath)
		handler.serveFile(assetPath, w, r)
		return
	}

	for _, page := range site.Pages {
		if page.Path == requestPath {
			content, err := page.Render(site)
			if err != nil {
				handler.logger.Print("500 Server Error")
				errorText := fmt.Sprintf("Unable to render page file: %v", err)
				handler.logger.Print(errorText)
				http.Error(w, errorText, 500)
				return
			}

			handler.logger.Print("200 OK")
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, content)
			return
		}
	}

	handler.logger.Print("404 Not Found")
	http.NotFound(w, r)
}

var serveCommand = &cobra.Command{
	Use:   "serve [PATH]",
	Short: "Run a local server preview of the site",
	Long:  "Run a local server preview of the site",
	Args:  cobra.MaximumNArgs(1),
	Run:   runServeCommand,
}

func init() {
	serveCommand.Flags().Int32VarP(&port, "port", "p", 8080, "Port to serve the site on")

	rootCmd.AddCommand(serveCommand)
}
