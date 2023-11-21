package main

import (
	"log"
	"os"
	"path/filepath"

	arg "github.com/alexflint/go-arg"

	"github.com/31333337/bmrng/go/0kn/pkg/utils"
)

// from cmd/server/server.go
func LaunchServer(args ArgsServer) {
	serversFile := args.ServerFile
	groupsFile := args.GroupFile
	addr := args.Addr
	errors.Addr = addr

	logger := utils.InitLogging()
	sugar := logger.Sugar()
	defer sugar.Sync()
	sugar.Infow(
		"Launching server",
		"address %s", addr,
	)

	servers, err := config.UnmarshalServersFromFile(serversFile)
	if err != nil {
		sugar.Fatalf("Could not read servers file %s", serversFile)
	}

	groups, err := config.UnmarshalGroupsFromFile(groupsFile)
	if err != nil {
		sugar.Fatalf("Could not read group file %s", groupsFile)
	}

	// will start in blocked state
	h := server.NewHandler()
	server := server.NewServer(&config.Servers{Servers: servers}, &config.Groups{Groups: groups}, h, addr)

	server.TcpConnections.LaunchAccepts()
	network.RunServer(h, server, servers, addr)
	config.Flush()
}

// from cmd/client/client.go
func LaunchClient(args ArgsClient) {
	serversFile := args.ServerFile
	groupsFile := args.GroupFile
	clientsFile := args.ClientFile
	addr := args.Addr
	errors.Addr = addr

	logger := utils.InitLogging()
	sugar := logger.Sugar()
	defer sugar.Sync()
	sugar.Infow(
		"Launching client",
		"address %s", addr,
	)

	servers, err := config.UnmarshalServersFromFile(serversFile)
	if err != nil {
		log.Fatalf("Could not read servers file %s", serversFile)
  }

	groups, err := config.UnmarshalGroupsFromFile(groupsFile)
	if err != nil {
		log.Fatalf("Could not read group file %s", groupsFile)
	}

	clients, err := config.UnmarshalServersFromFile(clientsFile)
	if err != nil {
		log.Fatalf("Could not read clients file %s", clientsFile)
	}

	clientRunner := client.NewClientRunner(servers, groups)

	// This fails @ network/rpc_call.go callee.HandleSignedMessageStream
	/*
		err = clientRunner.Connect()
		if err != nil {
			log.Fatalf("Could not make clients %v", err)
		}
	*/

	network.RunServer(nil, clientRunner, clients, addr)
}

// set the working directory from env var and change to the directory
func setWorkingDirectory() {
	// get working directory from env var, ensure set for children processes
  workingDir := getWorkingDirectory()
	if workingDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic("Error getting user's home directory: " + err.Error())
		}
		defaultDir := filepath.Join(homeDir, ".0KN")

		log.Println("Env _0KN_WORKDIR not set, using default", defaultDir)
		workingDir = defaultDir
		os.Setenv("_0KN_WORKDIR", workingDir)
	}

	// create working directory if it does not exist
	subDirs := []string{
		"certificates",
	}
	for _, d := range subDirs {
		dir := workingDir + "/" + d
		if err := os.MkdirAll(dir, os.FileMode(0700)); err != nil {
			panic("Failed to create directory: " + dir + "; " + err.Error())
		}
	}

	// change to working directory
	if err := os.Chdir(workingDir); err != nil {
		panic("Failed to change the working directory: " + err.Error())
	}
}

func getWorkingDirectory() string {
	return os.Getenv("_0KN_WORKDIR")
}

func main() {
	var args Args
	argParser := arg.MustParse(&args)

	utils.SetDebugLogEnabled(args.Debug)
	utils.SetDebugLogCallerEnabled(args.DebugCaller)

	logger := utils.InitLogging()

	////////////////////////////////////////////////////////////////////////
	// change current working directory for trellis hardcoded paths TODO
	////////////////////////////////////////////////////////////////////////
	// TODO: handle this cleanly with a config file(?)
	trellisCoordinatorDirectory := "../../../trellis/cmd/coordinator/"
	if err := os.Chdir(trellisCoordinatorDirectory); err != nil {
		sugar := logger.Sugar()
		sugar.Fatalw(
			"Failed to change the working directory",
			"error: %v", err,
		)
		sugar.Sync()
		return
	}

	setWorkingDirectory()


	switch {
	case args.Coordinator != nil:
		LaunchCoordinator(*args.Coordinator, argParser, logger)

	case args.Client != nil:
		LaunchClient(*args.Client)

	case args.Server != nil:
		LaunchServer(*args.Server)

	default:
		argParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}
