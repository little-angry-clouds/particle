package error

import "fmt"

type ClusterNotExists struct{}

func (e *ClusterNotExists) Error() string {
	return "there's no cluster detected"
}

type HelmRepoExists struct{
	Name string
}

func (e *HelmRepoExists) Error() string {
	return fmt.Sprintf("the helm repository '%s' is already added", e.Name)
}

type HelmChartExists struct{
	Name string
}

func (e *HelmChartExists) Error() string {
	return fmt.Sprintf("the chart '%s' already exists", e.Name)
}

type HelmChartMissingArgument struct{}

func (e *HelmChartMissingArgument) Error() string {
	return "missing argument 'CHART_NAME'"
}

type HelmDependencyType struct{}

func (e *HelmDependencyType) Error() string {
	return "charts has incorrect type, should be a list"
}

type HelmValuesType struct{}

func (e *HelmValuesType) Error() string {
	return "values has incorrect type, should be a dictionary of strings"
}

type HelmPrepareType struct{}

func (e *HelmPrepareType) Error() string {
	return "values has incorrect type, should be a dictionary of strings"
}

type ChartNotInstalled struct{
	Name string
}

func (e *ChartNotInstalled) Error() string {
	return fmt.Sprintf("chart '%s' is not installed", e.Name)
}

type ChartCantInstall struct{
	Name string
}

func (e *ChartCantInstall) Error() string {
	return fmt.Sprintf("%s, so '%s' can't be installed", &ClusterNotExists{}, e.Name)
}

type ChartCantDelete struct{
	Name string
}

func (e *ChartCantDelete) Error() string {
	return fmt.Sprintf("%s, so '%s' can't be deleted", &ClusterNotExists{}, e.Name)
}
