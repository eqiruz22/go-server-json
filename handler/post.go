package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/eqiruz22/go-server-json/utils"
)


func PostHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		getPost(w,r)
	case http.MethodPost:
		addPost(w,r)
	default:
		utils.JSONResponse(w,r,"not ok","method not allowed",nil,http.StatusMethodNotAllowed)
		return
	}

}

func HandleIdWithPath(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet :
		getIdPost(w,r)
	case http.MethodPatch :
		updatePost(w,r)
	case http.MethodDelete :
		deletePost(w,r)
	default:
		utils.JSONResponse(w,r,"not ok","method not allowed",nil,http.StatusMethodNotAllowed)
		return
	}
}


func getPost(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w,r,"OK","Retrieved Data",utils.Db.Posts,http.StatusOK)
}

func addPost(w http.ResponseWriter, r *http.Request){
	var newPost utils.Post
	err := json.NewDecoder(r.Body).Decode(&newPost)
	if err != nil {
		utils.JSONResponse(w, r, "OK", "Invalid request body", nil, http.StatusBadRequest)
		return
	}

	newID := 1
    if len(utils.Db.Posts) > 0 {
		// check last id from database
        lastID := utils.Db.Posts[len(utils.Db.Posts)-1]
        newID = lastID.ID + 1
    }

	// id increment from last id
    newPost.ID = newID

    utils.Db.Posts = append(utils.Db.Posts, newPost)

	err = utils.SaveDB()
	if err != nil {
		utils.JSONResponse(w,r,"not ok",err.Error(),nil,http.StatusInternalServerError)
		return
	}
	utils.JSONResponse(w,r,"OK","Post Created",newPost,http.StatusCreated)
}

func getIdPost(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path,"/posts/")
	// parse string id to int
	id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.JSONResponse(w, r, "not ok", "Invalid ID format", nil, http.StatusBadRequest)
        return
    }
	for _, post := range utils.Db.Posts {
		if post.ID == id {
			utils.JSONResponse(w,r,"OK","Get post by id",post,http.StatusOK)
			return
		} 
	}
	utils.JSONResponse(w,r,"OK","post not found",nil,http.StatusNotFound)
}

func updatePost(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path,"/posts/")
	id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.JSONResponse(w, r, "not ok", "Invalid ID format", nil, http.StatusBadRequest)
        return
    }
	var updated utils.Post
	errBody := json.NewDecoder(r.Body).Decode(&updated)
	if errBody != nil {
		utils.JSONResponse(w,r,"OK","invalid request payload",nil,http.StatusBadRequest)
		return
	}

	updated.ID = id

	for i, post := range utils.Db.Posts {
		if post.ID == id {
			utils.Db.Posts[i] = updated

			err = utils.SaveDB()

			if err != nil {
				utils.JSONResponse(w,r,"not ok","error while saved to database",nil,http.StatusInternalServerError)
				return
			}
			utils.JSONResponse(w,r,"OK","post update success",updated,http.StatusOK)
			return
		}
	}
	utils.JSONResponse(w,r,"OK","post not found",nil,http.StatusNotFound)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path,"/posts/")
	id, err := strconv.Atoi(idStr)
    if err != nil {
        utils.JSONResponse(w, r, "not ok", "Invalid ID format", nil, http.StatusBadRequest)
        return
    }
	for i, post := range utils.Db.Posts {
		if post.ID == id {
			utils.Db.Posts = append(utils.Db.Posts[:i], utils.Db.Posts[i+1:]...)

			err := utils.SaveDB()

			if err != nil {
				utils.JSONResponse(w,r,"not ok","error while saved to database",nil,http.StatusInternalServerError)
				return
			}
			utils.JSONResponse(w,r,"OK","post deleted",nil,http.StatusNoContent)
			return
		}
	}
	utils.JSONResponse(w,r,"OK","post not found",nil,http.StatusNotFound)
}




