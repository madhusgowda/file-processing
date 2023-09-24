package handler

import (
	"encoding/json"
	domains "fileProcessing/internal/core/domain"
	"fileProcessing/internal/core/services"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/raft"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	raftNode    *raft.Raft
	fileService *services.FileService
}

func NewHandler(raftNode *raft.Raft, fileService *services.FileService) *Handler {
	return &Handler{raftNode: raftNode, fileService: fileService}
}

func (h *Handler) UploadFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("filename")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		fileName := fileHeader.Filename

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileData, err := json.Marshal(domains.File{
			FileName: fileName,
			Size:     int64(len(fileBytes)),
		})

		raftNodeErr := h.raftNode.Apply(fileData, 0)
		if raftNodeErr.Error() != nil {
			//Error Handling
			return
		}
		response := struct {
			Message string `json:"message"`
		}{
			Message: "File uploaded and processed successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (h *Handler) GetFileSizeHandler(fileService *services.FileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fileName := vars["filename"]

		fileSize, err := fileService.GetFileSize(fileName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get file size %v", err), http.StatusNotFound)
			return
		}

		response := struct {
			FileName string `json:"filename"`
			FileSize int64  `json:"file_size"`
		}{
			FileName: fileName,
			FileSize: fileSize,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
