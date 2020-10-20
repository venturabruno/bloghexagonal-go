package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/venturabruno/bloghexagonal-go/application/usecase"
	"github.com/venturabruno/bloghexagonal-go/domain"
	"github.com/venturabruno/bloghexagonal-go/infrastructure/presenter"
)

func MakePostHandlers(route *mux.Router, negroniHandler negroni.Negroni, postUseCase usecase.PostUseCase) {
	route.Handle("/v1/posts", negroniHandler.With(
		negroni.Wrap(createPost(postUseCase)),
	)).Methods("POST").Name("createPost")

	route.Handle("/v1/posts/{id}/publish", negroniHandler.With(
		negroni.Wrap(publishPost(postUseCase)),
	)).Methods("POST").Name("publishPost")

	route.Handle("/v1/posts/{id}", negroniHandler.With(
		negroni.Wrap(getPost(postUseCase)),
	)).Methods("GET").Name("getPost")

	route.Handle("/v1/posts", negroniHandler.With(
		negroni.Wrap(listPosts(postUseCase)),
	)).Methods("GET").Name("listPosts")
}

func createPost(postUseCase usecase.PostUseCase) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var input struct {
			Title    string `json:"title" valid:"required~Título é obrigatório"`
			Subtitle string `json:"subtitle" valid:"required~Subtitulo é obrigatório"`
			Content  string `json:"content" valid:"required~Conteúdo é obrigatório"`
		}

		err := json.NewDecoder(request.Body).Decode(&input)
		if err != nil {
			responseInternalServer(response, err)
			return
		}

		_, err = govalidator.ValidateStruct(input)
		if err != nil {
			log.Println(err.Error())
			response.WriteHeader(http.StatusUnprocessableEntity)
			response.Write([]byte(err.Error()))
			return
		}

		post, err := domain.NewPost(input.Title, input.Subtitle, input.Content)
		if err != nil {
			responseInternalServer(response, err)
			return
		}

		_, err = postUseCase.CreatePost(post)
		if err != nil {
			responseInternalServer(response, err)
			return
		}

		response.WriteHeader(http.StatusCreated)

		jPost := &presenter.Post{
			ID:          post.ID,
			Title:       post.Title,
			Subtitle:    post.Subtitle,
			Context:     post.Content,
			Status:      post.Status,
			CreatedAt:   post.CreatedAt,
			PublishedAt: post.PublishedAt,
		}

		if err := json.NewEncoder(response).Encode(jPost); err != nil {
			responseInternalServer(response, err)
			return
		}
	})
}

func getPost(postUseCase usecase.PostUseCase) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := domain.StringToEntityID(vars["id"])
		if err != nil {
			responseNotFound(response, err)
			return
		}

		post, err := postUseCase.GetPost(id)
		if err != nil {
			responseInternalServer(response, err)
			return
		}

		if post == nil {
			responseNotFound(response, err)
			return
		}

		jPost := &presenter.Post{
			ID:          post.ID,
			Title:       post.Title,
			Subtitle:    post.Subtitle,
			Context:     post.Content,
			Status:      post.Status,
			CreatedAt:   post.CreatedAt,
			PublishedAt: post.PublishedAt,
		}

		if err := json.NewEncoder(response).Encode(jPost); err != nil {
			responseInternalServer(response, err)
			return
		}
	})
}

func listPosts(postUseCase usecase.PostUseCase) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		posts, err := postUseCase.ListPosts()
		if err != nil {
			responseInternalServer(response, err)
			return
		}

		if len(posts) == 0 {
			responseNotFound(response, err)
			return
		}

		var jPosts []*presenter.Post
		for _, post := range posts {
			jPosts = append(jPosts, &presenter.Post{
				ID:          post.ID,
				Title:       post.Title,
				Subtitle:    post.Subtitle,
				Context:     post.Content,
				Status:      post.Status,
				CreatedAt:   post.CreatedAt,
				PublishedAt: post.PublishedAt,
			})
		}

		if err := json.NewEncoder(response).Encode(jPosts); err != nil {
			responseInternalServer(response, err)
			return
		}
	})
}

func publishPost(postUseCase usecase.PostUseCase) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		id, err := domain.StringToEntityID(vars["id"])
		if err != nil {
			responseNotFound(response, err)
			return
		}

		post, err := postUseCase.PublishPost(id)
		if err != nil {
			responseInternalServer(response, err)
			return
		}

		if post == nil {
			responseNotFound(response, err)
			return
		}

		jPost := &presenter.Post{
			ID:          post.ID,
			Title:       post.Title,
			Subtitle:    post.Subtitle,
			Context:     post.Content,
			Status:      post.Status,
			CreatedAt:   post.CreatedAt,
			PublishedAt: post.PublishedAt,
		}

		if err := json.NewEncoder(response).Encode(jPost); err != nil {
			responseInternalServer(response, err)
			return
		}
	})
}

func responseInternalServer(response http.ResponseWriter, err error) {
	log.Println(err.Error())
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte("Internal Server Error"))
}

func responseNotFound(response http.ResponseWriter, err error) {
	log.Println(err.Error())
	response.WriteHeader(http.StatusNotFound)
	response.Write([]byte("Post not found"))
}
