package service

import (
	"context"
	"errors"
	"math/rand"

	"strangerbot/repository"
	"strangerbot/repository/model"
)

var (
	ErrUserNotFillAllQuestion = errors.New("user not fill all question")
)

func ServiceMatch(ctx context.Context, chatId int64) (*model.User, error) {

	repo := repository.GetRepository()

	// find all question
	questions, err := repo.GetAllQuestion(ctx)
	if err != nil {
		return nil, err
	}

	// find all question option
	options, err := repo.GetAllOption(ctx)
	if err != nil {
		return nil, err
	}

	// find user all user question data
	userQuestionData, err := repo.GetUserQuestionDataByUser(ctx, chatId)
	if err != nil {
		return nil, err
	}

	// check user fill full?
	//fillFull, _ := model.Questions(questions).CheckUserFillFull(userQuestionData)
	//if !fillFull {
	//	return nil, ErrUserNotFillAllQuestion
	//}

	// find matching options and user matching options value
	chatIds, err := repo.GetChatByMatching(ctx, chatId, questions, options, userQuestionData)
	if err != nil {
		return nil, err
	}

	// shuffle chat id
	matchChatId := shuffleChatId(chatIds)

	// find user
	user, err := repo.GetUserByChatId(ctx, matchChatId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func shuffleChatId(a []int64) int64 {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a[0]
}
