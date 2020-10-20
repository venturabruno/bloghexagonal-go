package usecase

import "github.com/venturabruno/bloghexagonal-go/domain"

type PostUseCase struct {
	postRepository domain.PostRepository
}

func NewPostUseCase(repository domain.PostRepository) *PostUseCase {
	return &PostUseCase{repository}
}

func (usecase *PostUseCase) CreatePost(post *domain.Post) (*domain.Post, error) {
	_, err := usecase.postRepository.Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (usecase *PostUseCase) GetPost(id domain.EntityID) (*domain.Post, error) {
	return usecase.postRepository.FindID(id)
}

func (usecase *PostUseCase) ListPosts() ([]*domain.Post, error) {
	return usecase.postRepository.All()
}

func (usecase *PostUseCase) PublishPost(id domain.EntityID) (*domain.Post, error) {
	post, err := usecase.postRepository.FindID(id)
	if err != nil {
		return nil, err
	}

	post.Publish()
	err = usecase.postRepository.Update(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
