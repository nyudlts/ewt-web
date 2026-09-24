package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Job struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	EwtRoot     string `json:"ewt_root"`
	ProjectName string `json:"project_name"`
	Status      string `json:"status"`
	Started     string `json:"started"`
	Completed   string `json:"completed"`
}

type Worker struct {
}

func WriteJobToFile(w io.Writer, job Job) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(job)
}

// Project Job routes
var projectJobs = []string{"archive"}

func GetProjectJobs(ewtRoot string, projectName string) []string {
	return projectJobs
}

func RouteProjectJob(ewtRoot string, projectName string, job string) (*gin.H, error) {
	fmt.Println("DEBUG: routing job", job)
	switch job {
	case "archive":
		if jobID, err := archiveProject(ewtRoot, projectName); err != nil {
			return nil, err

		} else {
			return &gin.H{"message": "Job accepted", "job_id": jobID}, nil
		}
	default:
		return nil, fmt.Errorf("unknown job: %s", job)
	}
}

// project jobs
func archiveProject(ewtRoot string, projectName string) (string, error) {
	fmt.Println("DEBUG: archiving project", projectName)

	job := Job{}
	id := fmt.Sprintf("%08x", rand.Intn(0x100000000))
	job.ID = id
	job.Type = "archive_project"
	job.EwtRoot = ewtRoot
	job.ProjectName = projectName
	job.Status = "queued"
	jobID := fmt.Sprintf("%s_archive_%s.json", projectName, job.ID)

	jobFile, err := os.Create(filepath.Join(ewtRoot, "jobs", "queued", jobID))
	if err != nil {
		return "", err
	}
	defer jobFile.Close()

	if err := WriteJobToFile(jobFile, job); err != nil {
		return "", err
	}

	return jobID, nil

}
