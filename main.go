package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Command struct {
	Name string
	Args string
}

type Function struct {
	Name       string
	Body       []Command
	Executed   bool
	ExecutedBy []string
}

type Package struct {
	Name      string
	Functions []Function
}

func ParseEsh(filePath string) (*Package, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", filePath)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	packageRegex := regexp.MustCompile(`(?i)^\s*package\s+(\w+)`)
	functionRegex := regexp.MustCompile(`(?i)func\s+(\w+)\s*{([^}]*)}`)

	packageMatch := packageRegex.FindStringSubmatch(string(content))
	if len(packageMatch) != 2 {
		return nil, fmt.Errorf("invalid package declaration")
	}
	pkgName := packageMatch[1]

	functionMatches := functionRegex.FindAllStringSubmatch(string(content), -1)

	var functions []Function
	for _, funcMatch := range functionMatches {
		if len(funcMatch) < 3 {
			return nil, fmt.Errorf("invalid function declaration")
		}
		funcName := funcMatch[1]
		funcBody := strings.TrimSpace(funcMatch[2])

		var funcCommands []Command
		commandLines := strings.Split(funcBody, "\n")
		for _, cmdLine := range commandLines {
			cmdLine = strings.TrimSpace(cmdLine)
			if cmdLine == "" {
				continue
			}

			cmdParts := strings.Fields(cmdLine)
			cmdName := cmdParts[0]
			cmdArgs := strings.Join(cmdParts[1:], " ")
			funcCommands = append(funcCommands, Command{Name: cmdName, Args: cmdArgs})
		}

		functions = append(functions, Function{
			Name:       funcName,
			Body:       funcCommands,
			Executed:   false,
			ExecutedBy: nil,
		})
	}

	return &Package{Name: pkgName, Functions: functions}, nil
}

func ExecuteCommand(cmd Command, outputFile *os.File, abortChan chan struct{}) error {
	log.Infof("Executing command: %s %s", cmd.Name, cmd.Args)

	var command *exec.Cmd
	if runtime.GOOS == "windows" {
		psCommand := fmt.Sprintf("%s %s", cmd.Name, cmd.Args)
		command = exec.Command("powershell", "-Command", psCommand)
	} else {
		shellCommand := fmt.Sprintf("%s %s", cmd.Name, cmd.Args)
		command = exec.Command("sh", "-c", shellCommand)
	}

	command.Stdout = outputFile
	command.Stderr = outputFile

	if err := command.Run(); err != nil {
		return fmt.Errorf("command %s failed: %w", cmd.Name, err)
	}
	return nil
}

func ExecuteFunction(fn *Function, outputFile *os.File, abortChan chan struct{}) error {
	log.Infof("Executing function: %s", fn.Name)
	for _, cmd := range fn.Body {
		select {
		case <-abortChan:
			log.Warnf("Function %s aborted", fn.Name)
			return fmt.Errorf("function %s aborted", fn.Name)
		default:
		}
		if err := ExecuteCommand(cmd, outputFile, abortChan); err != nil {
			log.Errorf("Function %s failed: %s", fn.Name, err)
			return fmt.Errorf("function %s failed: %w", fn.Name, err)
		}
	}
	fn.Executed = true
	log.Infof("Function %s completed successfully", fn.Name)
	return nil
}

func ExecutePackage(pkg *Package, outputFile *os.File) {
	log.Infof("Executing package: %s", pkg.Name)

	abortChan := make(chan struct{})
	defer close(abortChan)

	for i := range pkg.Functions {
		fn := &pkg.Functions[i]
		if err := ExecuteFunction(fn, outputFile, abortChan); err != nil {
			log.Errorf("Execution of function %s failed: %v", fn.Name, err)
			return
		}
	}

	for _, fn := range pkg.Functions {
		if !fn.Executed {
			log.Errorf("Function %s was not executed.", fn.Name)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Error("Usage: esh <file.esh>")
		return
	}

	filePath := os.Args[1]
	if filepath.Ext(filePath) != ".esh" {
		log.Error("Invalid file extension. Expected .esh")
		return
	}

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	outputFile, err := os.Create("output.log")
	if err != nil {
		log.Errorf("Error creating output file: %v", err)
		return
	}
	defer outputFile.Close()

	pkg, err := ParseEsh(filePath)
	if err != nil {
		log.Errorf("Error parsing esh: %v", err)
		return
	}

	ExecutePackage(pkg, outputFile)

	log.Info("Execution completed.")
}
