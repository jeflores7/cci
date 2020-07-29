package main

import "errors"

type graph struct {
	nodes map[string]struct{}
	edges map[string][]string
}

func buildGraph(projects []string, deps [][]string) graph {
	g := graph{
		nodes: make(map[string]struct{}),
		edges: make(map[string][]string),
	}
	for _, project := range projects {
		g.nodes[project] = struct{}{}
	}
	for _, dep := range deps {
		g.edges[dep[1]] = append(g.edges[dep[1]], dep[0])
	}
	return g
}

func findBuildPath(projects []string, deps [][]string) ([]string, error) {
	if len(projects) == 0 {
		return nil, errors.New("no build path")
	}
	g := buildGraph(projects, deps)
	return getBuildPath(g, projects, nil)
}

func getBuildPath(g graph, projects []string, buildPath []string) ([]string, error) {
	if len(projects) == 0 {
		return buildPath, nil
	}
	for i, project := range projects {
		if checkBuildError(g, project, buildPath) {
			continue
		}
		// Didn't notice problem without IDE
		// remaining := projects[0:i] + projects[i+1:]
		remaining := append(projects[0:i], projects[i+1:]...)
		path, err := getBuildPath(g, remaining, append(buildPath, project))
		if err != nil {
			continue
		}
		// Didn't notice missing error return without IDE
		return path, nil
	}
	return nil, errors.New("no build path")
}

func checkBuildError(g graph, project string, buildPath []string) bool {
	for _, dep := range g.edges[project] {
		depMet := false
		for _, p := range buildPath {
			if p == dep {
				depMet = true
				break
			}
		}
		if !depMet {
			return true
		}
	}
	return false
}
