package v1

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"strings"
	"time"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	"github.com/google/uuid"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Namespace struct {
	Name   string
	Labels map[string]string
}

type Secret struct {
	Name string
	Data map[string]string
}

type ConfigMap struct {
	Name string
	Data map[string]string
}

type LogEntry struct {
	Timestamp time.Time
	Content   string
}

type Metric struct {
	Name   string
	Value  float64
	Format string `json:"omitempty"`
}

type CronWorkflow struct {
	ID                         uint64
	CreatedAt                  time.Time `db:"created_at"`
	UID                        string
	Name                       string
	GenerateName               string
	Schedule                   string
	Timezone                   string
	Suspend                    bool
	ConcurrencyPolicy          string
	StartingDeadlineSeconds    *int64
	SuccessfulJobsHistoryLimit *int32
	FailedJobsHistoryLimit     *int32
	WorkflowExecution          *WorkflowExecution
}

type WorkflowTemplate struct {
	ID                   uint64
	CreatedAt            time.Time `db:"created_at"`
	UID                  string
	Name                 string
	Manifest             string
	Version              int32
	IsLatest             bool
	IsArchived           bool `db:"is_archived"`
	ArgoWorkflowTemplate *wfv1.WorkflowTemplate
	Labels               map[string]string
}

func (wt *WorkflowTemplate) GetManifestBytes() []byte {
	return []byte(wt.Manifest)
}

func (wt *WorkflowTemplate) GetParametersKeyString() (map[string]string, error) {
	mapping := make(map[interface{}]interface{})

	if err := yaml.Unmarshal(wt.GetManifestBytes(), mapping); err != nil {
		return nil, err
	}

	arguments, ok := mapping["arguments"]
	if !ok {
		return nil, nil
	}

	argumentsMap, ok := arguments.(map[interface{}]interface{})
	if !ok {
		return nil, nil
	}

	parameters, ok := argumentsMap["parameters"]
	if !ok {
		return nil, nil
	}

	parametersAsArray, ok := parameters.([]interface{})
	if !ok {
		return nil, nil
	}

	result := make(map[string]string)
	for _, parameter := range parametersAsArray {
		parameterMap, ok := parameter.(map[interface{}]interface{})
		if !ok {
			continue
		}

		key := parameterMap["name"]
		keyAsString, ok := key.(string)
		if !ok {
			continue
		}

		remainingParameters, err := yaml.Marshal(parameterMap)
		if err != nil {
			continue
		}

		result[keyAsString] = string(remainingParameters)
	}

	return result, nil
}

func (wt *WorkflowTemplate) GenerateUID() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	wt.UID = uid.String()

	return wt.UID, nil
}

func (wt *WorkflowTemplate) GetWorkflowManifestBytes() ([]byte, error) {
	if wt.ArgoWorkflowTemplate == nil {
		return []byte{}, nil
	}

	wt.ArgoWorkflowTemplate.TypeMeta.Kind = "Workflow"
	wt.ArgoWorkflowTemplate.ObjectMeta = metav1.ObjectMeta{
		GenerateName: wt.ArgoWorkflowTemplate.ObjectMeta.GenerateName,
		Labels:       wt.ArgoWorkflowTemplate.ObjectMeta.Labels,
	}

	return json.Marshal(wt.ArgoWorkflowTemplate)
}

const (
	WorfklowPending   WorkflowExecutionPhase = "Pending"
	WorfklowRunning   WorkflowExecutionPhase = "Running"
	WorfklowSucceeded WorkflowExecutionPhase = "Succeeded"
	WorfklowSkipped   WorkflowExecutionPhase = "Skipped"
	WorfklowFailed    WorkflowExecutionPhase = "Failed"
	WorfklowError     WorkflowExecutionPhase = "Error"
)

type WorkflowExecutionPhase string

type WorkflowExecution struct {
	ID               uint64
	CreatedAt        time.Time `db:"created_at"`
	UID              string
	Name             string
	GenerateName     string
	Parameters       []WorkflowExecutionParameter
	Manifest         string
	Phase            WorkflowExecutionPhase
	StartedAt        time.Time
	FinishedAt       time.Time
	WorkflowTemplate *WorkflowTemplate
	Labels           map[string]string
}

type WorkflowExecutionParameter struct {
	Name  string
	Value *string
}

type ListOptions = metav1.ListOptions

type PodGCStrategy = wfv1.PodGCStrategy

type WorkflowExecutionOptions struct {
	Name           string
	GenerateName   string
	Entrypoint     string
	Parameters     []WorkflowExecutionParameter
	ServiceAccount string
	Labels         *map[string]string
	ListOptions    *ListOptions
	PodGCStrategy  *PodGCStrategy
}

type File struct {
	Path         string
	Name         string
	Size         int64
	Extension    string
	ContentType  string
	LastModified time.Time
	Directory    bool
}

func FilePathToName(path string) string {
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	lastSlashIndex := strings.LastIndex(path, "/")
	if lastSlashIndex < 0 {
		return path
	}

	return path[lastSlashIndex+1:]
}

// Given a path, returns the parent path, asssuming a '/' delimitor
// Result does not have a trailing slash.
// -> a/b/c/d would return a/b/c
// -> a/b/c/d/ would return a/b/c
// If path is empty string, it is returned.
// If path is '/' (root) it is returned as is.
// If there is no '/', '/' is returned.
func FilePathToParentPath(path string) string {
	separator := "/"
	if path == "" || path == separator {
		return path
	}

	if strings.HasSuffix(path, "/") {
		path = path[0 : len(path)-1]
	}

	lastIndexOfForwardSlash := strings.LastIndex(path, separator)
	if lastIndexOfForwardSlash <= 0 {
		return separator
	}

	return path[0:lastIndexOfForwardSlash]
}

func FilePathToExtension(path string) string {
	dotIndex := strings.LastIndex(path, ".")

	if dotIndex == -1 {
		return ""
	}

	if dotIndex == (len(path) - 1) {
		return ""
	}

	return path[dotIndex+1:]
}

func WrapSpecInK8s(data []byte) ([]byte, error) {
	mapping := make(map[interface{}]interface{})
	if err := yaml.Unmarshal(data, mapping); err != nil {
		return nil, err
	}

	contentMap := map[interface{}]interface{}{
		"metadata": make(map[interface{}]interface{}),
		"spec":     mapping,
	}

	finalBytes, err := yaml.Marshal(contentMap)
	if err != nil {
		return nil, nil
	}

	return finalBytes, nil
}

func RemoveAllButSpec(manifest []byte) (map[interface{}]interface{}, error) {
	mapping := make(map[interface{}]interface{})

	if err := yaml.Unmarshal(manifest, mapping); err != nil {
		return nil, err
	}

	deleteEmptyValuesMapping(mapping)

	spec, ok := mapping["spec"].(map[interface{}]interface{})
	if !ok {
		return nil, nil
	}

	return spec, nil
}

func AddWorkflowTemplateParametersFromAnnotations(spec map[interface{}]interface{}, annotations map[string]string) {
	// TODO put in parameters here, decoding as you go.
	// @todo don't forget to clean up the code and centralize it somewhere.
	// maybe we should have a manifest builder or something.

	spec["arguments"] = make(map[interface{}]interface{})
	arguments := spec["arguments"].(map[interface{}]interface{})
	arguments["parameters"] = make([]interface{}, 0)
	parameters := arguments["parameters"].([]interface{})

	for _, value := range annotations {
		data := make(map[interface{}]interface{})
		err := yaml.Unmarshal([]byte(value), data)
		if err != nil {
			// todo log
			continue
		}

		parameters = append(parameters, data)
	}

	arguments["parameters"] = parameters
}

// Returns the number of keys in the map
func deleteEmptyValuesMapping(mapping map[interface{}]interface{}) int {
	keys := 0
	for key, value := range mapping {
		keys++
		valueAsMapping, ok := value.(map[interface{}]interface{})
		if ok {
			if deleteEmptyValuesMapping(valueAsMapping) == 0 {
				delete(mapping, key)
			}
		}

		valueAsArray, ok := value.([]interface{})
		if ok {
			deleteEmptyValuesArray(valueAsArray)
		}

		valueAsString, ok := value.(string)
		if ok && valueAsString == "" {
			delete(mapping, key)
		}
	}

	return keys
}

// Returns the number of items in the array.
func deleteEmptyValuesArray(values []interface{}) int {
	count := 0
	for _, value := range values {
		count++

		valueAsMapping, ok := value.(map[interface{}]interface{})
		if ok {
			deleteEmptyValuesMapping(valueAsMapping)
		}

		valueAsArray, ok := value.([]interface{})
		if ok {
			deleteEmptyValuesArray(valueAsArray)
		}
	}

	return count
}
