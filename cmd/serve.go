package cmd

import (
	"log"
	"net/http"

	"github.com/danesparza/Dashboard-service/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	serverInterface   string
	serverPort        int
	serverUIDirectory string
	allowedOrigins    string
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `The serve command starts hosting the dashboard`,
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	//	If we have a config file, report it:
	if viper.ConfigFileUsed() != "" {
		log.Println("[INFO] Using config file:", viper.ConfigFileUsed())
	}

	//	Create a router and setup our REST endpoints...
	var router = mux.NewRouter()

	//	Setup our routes
	router.HandleFunc("/", api.ShowUI)
	router.HandleFunc("/config", nil).Methods("GET")
	router.HandleFunc("/config", nil).Methods("POST")

	//	If we don't have a UI directory specified...
	if viper.GetString("server.ui-dir") == "" {
		//	Use the static assets file generated with
		//	https://github.com/elazarl/go-bindata-assetfs using the dashboard UI from
		//	https://github.com/danesparza/Dashboard.
		//
		//	To generate this file, place the 'ui'
		//	directory under the main dashboard-service directory and run the commands:
		//	go-bindata-assetfs.exe -pkg cmd ./ui/...
		//	mv bindata_assetfs.go cmd
		//	go install ./...

		router.PathPrefix("/ui").Handler(http.StripPrefix("/ui", http.FileServer(assetFS())))
	} else {
		//	Use the supplied directory:
		log.Printf("[INFO] Using UI directory: %s\n", viper.GetString("server.ui-dir"))
		router.PathPrefix("/ui").Handler(http.StripPrefix("/ui", http.FileServer(http.Dir(viper.GetString("server.ui-dir")))))
	}

	// Setup CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	//	Format the bound interface:
	formattedInterface := viper.GetString("server.bind")
	if formattedInterface == "" {
		formattedInterface = "127.0.0.1"
	}

	//	If we have an SSL cert specified, use it:
	if viper.GetString("server.sslcert") != "" {
		log.Printf("[INFO] Using SSL cert: %s\n", viper.GetString("server.sslcert"))
		log.Printf("[INFO] Using SSL key: %s\n", viper.GetString("server.sslkey"))
		log.Printf("[INFO] Starting HTTPS server: https://%s:%s\n", formattedInterface, viper.GetString("server.port"))

		//	Start the service with SSL
		log.Printf("[ERROR] %v\n", http.ListenAndServeTLS(viper.GetString("server.bind")+":"+viper.GetString("server.port"), viper.GetString("server.sslcert"), viper.GetString("server.sslkey"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	} else {
		log.Printf("[INFO] Starting HTTP server: http://%s:%s\n", formattedInterface, viper.GetString("server.port"))

		//	Start the service with HTTP
		log.Printf("[ERROR] %v\n", http.ListenAndServe(viper.GetString("server.bind")+":"+viper.GetString("server.port"), handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	}

}

func init() {
	rootCmd.AddCommand(serveCmd)

	//	Setup our flags
	serveCmd.Flags().IntVarP(&serverPort, "port", "p", 1313, "port on which the server will listen")
	serveCmd.Flags().StringVarP(&serverInterface, "bind", "i", "", "interface to which the server will bind")
	serveCmd.Flags().StringVarP(&serverUIDirectory, "ui-dir", "u", "", "directory for the UI")
	serveCmd.Flags().StringVarP(&allowedOrigins, "allowed-origins", "o", "", "comma seperated list of allowed CORS origins")

	//	Bind config flags for optional config file override:
	viper.BindPFlag("server.port", serveCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.bind", serveCmd.Flags().Lookup("bind"))
	viper.BindPFlag("server.ui-dir", serveCmd.Flags().Lookup("ui-dir"))
	viper.BindPFlag("server.allowed-origins", serveCmd.Flags().Lookup("allowed-origins"))
}
