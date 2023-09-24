package handler

import (
	"fileProcessing/internal/core/services"
	"github.com/gorilla/mux"
	"github.com/hashicorp/raft"
)

func NewRoutes(raftNode *raft.Raft, fileService *services.FileService) *mux.Router {
	r := mux.NewRouter()
	webHandler := NewHandler(raftNode, fileService)

	r.HandleFunc("/upload", webHandler.UploadFileHandler()).Methods("POST")
	r.HandleFunc("/file/{filename}", webHandler.GetFileSizeHandler(fileService)).Methods("GET")
	return r
}
