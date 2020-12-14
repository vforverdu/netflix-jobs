package jobs

import (
	"encoding/json"
	"log"
	"time"
)

// JobResponse is the response obtained from the Netflix Jobs endpoint
type JobResponse struct {
	Records JobRecords `json:"records"`
}

// JobRecords contains a list of job postings
type JobRecords struct {
	Jobs []JobWrapper `json:"postings"`
}

// JobWrapper representes the job json structure received in a JobResponse
type JobWrapper struct {
	ID          string `json:"external_id"`
	Title       string `json:"text"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"updated_at"`
	URL         string `json:"url"`
}

// NewJobResponse crates a JobResponse from a response body
func NewJobResponse(body []byte) (*JobResponse, error) {
	var jobResponse JobResponse
	err := json.Unmarshal(body, &jobResponse)
	return &jobResponse, err
}

// GetJobs returns an slice of jobs
func (jobResponse *JobResponse) GetJobs() []*Job {
	var jobs []*Job

	for _, jobWrapper := range jobResponse.Records.Jobs {
		job := jobWrapper.ToJob()
		jobs = append(jobs, job)
	}

	return jobs
}

// ToJob converts a JobWrapper into a Job
func (jobWrapper *JobWrapper) ToJob() *Job {
	date, err := time.Parse(time.RFC3339, jobWrapper.Date)

	if err != nil {
		log.Fatal("Error parsing date")
	}

	job := &Job{
		ID:          jobWrapper.ID,
		Company:     "NETFLIX",
		Title:       jobWrapper.Title,
		Category:    "SOFTWARE",
		Description: jobWrapper.Description,
		Location:    jobWrapper.Location,
		Date:        date.Format("2006-01-02 15:04:05"),
		URL:         jobWrapper.URL,
	}

	return job
}
