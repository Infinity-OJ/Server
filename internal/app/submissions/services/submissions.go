package services

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/infinity-oj/server/internal/app/submissions/repositories"
	"github.com/infinity-oj/server/internal/pkg/models"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var specialKey = "imf1nlTy0j"

type SubmissionsService interface {
	Create(submitterID uint64, problemID string, userSpace string) (s *models.Submission, err error)
	Judge(submission *models.Submission) error
}

type DefaultSubmissionService struct {
	logger           *zap.Logger
	ProblemService   ProblemsService
	JudgementService JudgementsService
	FileService      FilesService
	Repository       repositories.SubmissionRepository
}

func (d DefaultSubmissionService) Create(submitterID uint64, problemId string, userSpace string) (s *models.Submission, err error) {
	s, err = d.Repository.Create(submitterID, problemId, userSpace)
	if err != nil {
		return nil, err
	}

	err = d.Judge(s)
	return
}

type Meta struct {
	TestCases []string `yaml:"testcases"`
}

func (d DefaultSubmissionService) Judge(submission *models.Submission) error {
	problem, err := d.ProblemService.Fetch(submission.ProblemID)
	if err != nil {
		return err
	}

	meta, err := d.FileService.FetchMetaFile(problem.PrivateSpace)
	if err != nil {
		return errors.Wrap(err, "judge error: fetch meta file error")
	}

	m := Meta{}
	err = yaml.Unmarshal(meta, &m)
	if err != nil {
		d.logger.Error("error: %v", zap.Error(err))
		return errors.Wrap(err, "judge error: parse meta file error")
	}

	fmt.Println(problem.PublicSpace)
	fmt.Println(problem.PrivateSpace)
	for k, v := range m.TestCases {
		fmt.Println(k, v)
		//if err := d.JudgementService.Create(submission.ID, problem.PublicSpace, problem.PrivateSpace, submission.UserSpace, v); err != nil {
		//	return errors.Wrap(err, "judge error: submit judgement error")
		//}
	}
	return nil
}

func (d DefaultSubmissionService) DispatchJudgement(submissionId string) error {

	submission, err := d.Repository.FetchSubmissionBySubmissionId(submissionId)
	if err != nil {
		return err
	}

	submissionElement := d.Repository.CreateSubmissionInQueue(submission)

	upstreams := submissionElement.FindUpstreams()

	for _, upstream := range upstreams {

		upstreamType := upstream.Type

		upstreamProperty, err := json.Marshal(upstream.Properties)
		if err != nil {
			continue
		}
		err = d.JudgementService.Create(
			upstreamType,
			string(upstreamProperty),
			upstream.Inputs,
		)
		if err != nil {

		}
	}

	return nil

}

func NewSubmissionService(
	logger *zap.Logger,
	ProblemService ProblemsService,
	Repository repositories.SubmissionRepository,
	FileService FilesService,
	JudgementService JudgementsService,
) SubmissionsService {
	go func() {

	}()
	return &DefaultSubmissionService{
		logger:           logger.With(zap.String("type", "DefaultSubmissionService")),
		ProblemService:   ProblemService,
		FileService:      FileService,
		JudgementService: JudgementService,
		Repository:       Repository,
	}
}

/*

	judgement, task := m.findJudgementTask(judgementId, taskId)
	if task == nil {
		return errors.New("unknown task: " + taskId)
	}

	task.Status = "done"

	//fmt.Println(task.block.Output)
	//fmt.Println(results)

	if len(task.block.Output) != len(results) {
		return errors.New("output slots mismatch")
	}

	blockId := task.block.Id
	for index, result := range results {
		fmt.Println(blockId, index)
		links := judgement.Graph.FindLinkBySourcePort(blockId, index)
		fmt.Println(links, result)
		for _, link := range links {
			fmt.Println(link.Id)
			judgement.Result[link.Id] = result
			fmt.Println("set link", link.Id, "to", result)
		}
	}
	task.block.Done()

*/
