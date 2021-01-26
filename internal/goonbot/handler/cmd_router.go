package handler

import (
	"strings"
)

// CmdHandler is a function which handles a command line style command.
//
// The result consists of the result of your command which will be sent to the user.
// This can be nil if you don't want to respond to the user who called your command.
type CmdHandler func(args []string) interface{}

type cmdPath struct {
	Path    string
	Handler CmdHandler
}

// CmdRouter routes incoming string invocations to registered handlers based on the invocation's prefixes.
// For example:
//		r := NewCmdRouter("!bigboy", defaultHandler)
//		r.HandlePath("biscuit", handlerA)
// 		r.HandlePath("cookie", handlerB)
//
//		r.Route("!bigboy biscuit arg1 arg2") // returns the result of handlerA(flags, []{"arg1", "arg2"})
// 		r.Route("!bigboy whatttt") // return the result of defaultHandler(flags, []{"whatttt"})
//      r.Route("not a bigboy command") // ignores the message and returns nil & false
type CmdRouter interface {

	// HandlePath registers a given path to a given handler on this CmdRouter
	//
	// If a message matches multiple paths, the one you registered first will be used. So be sure to register your
	// paths in the correct order.
	HandlePath(path string, handler CmdHandler)

	// Route directs a string msg to a registered handler based on the message's prefix.
	//
	// Route returns the result of the handler whose path is prefixed on the input message.
	// It will return the result of the CmdRouter's default handler if there is no match.
	//
	// If a message does not begin with the CmdRouter's root path, it will be ignored and nil will be returned.
	// If a message's prefix matches multiple handlers' paths, the handler you registered first will be used.
	Route(msg string) interface{}
}

// NewCmdRouter is used for constructing new instances of CmdRouter
func NewCmdRouter(rootCommand string, defaultHandler CmdHandler) CmdRouter {
	return &cmdRouter{
		rootPath:       rootCommand,
		defaultHandler: defaultHandler,
	}
}

type cmdRouter struct {
	rootPath       string
	defaultHandler CmdHandler
	paths          []cmdPath
}

func (c *cmdRouter) HandlePath(path string, handler CmdHandler) {
	c.paths = append(c.paths, cmdPath{
		Path:    path,
		Handler: handler,
	})
}

func (c *cmdRouter) Route(msg string) interface{} {
	if !matchesPrefix(msg, c.rootPath) {
		return nil
	}

	for _, p := range c.paths {
		cmdPrefix := c.rootPath + " " + p.Path

		if matchesPrefix(msg, cmdPrefix) {
			return p.Handler(removePrefixAndSplit(msg, cmdPrefix))
		}
	}

	return c.defaultHandler(removePrefixAndSplit(msg, c.rootPath))
}

func matchesPrefix(msg, prefix string) bool {
	return len(msg) >= len(prefix) && msg[:len(prefix)] == prefix
}

func removePrefixAndSplit(msg, prefix string) []string {
	if len(msg) == len(prefix) {
		return nil
	}

	withoutPrefix := strings.Trim(msg[len(prefix):], " ")

	return strings.Split(withoutPrefix, " ")
}
