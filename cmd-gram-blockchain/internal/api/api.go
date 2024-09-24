package api

import (
	"cmd-gram-blockchain/internal/models"
	"cmd-gram-blockchain/pkg/blockchain"
	"cmd-gram-blockchain/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const (
	bindAddr = "0.0.0.0:10000"
	logLevel = "debug"
)

type api struct {
	//config 	TODO IF there are will be more vars
	logger *logrus.Logger
	r      *mux.Router
	bc     *blockchain.Blockchain
	//db     	TODO
}

func New(router *mux.Router) *api {
	return &api{
		logger: logrus.New(),
		r:      router,
	}
}

func (api *api) Start() error {
	if err := api.configureLogger(); err != nil {
		return err
	}

	api.Handle()

	if err := api.configureBlockChain(); err != nil {
		return err
	}

	api.logger.Info("starting api server")

	return http.ListenAndServe(bindAddr, api.r)
}

func (api *api) configureLogger() error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(level)

	return nil
}

func (api *api) configureBlockChain() error {
	bc, err := blockchain.New()
	if err != nil {
		return err
	}

	api.bc = bc
	return nil
}

func (api *api) Handle() {
	api.r.HandleFunc("/api/new-block", api.handleNewBlock).Methods(http.MethodPost)
}

func (api *api) handleNewBlock(w http.ResponseWriter, r *http.Request) {
	msg := models.MessageDTO{}
	utils.ParseBody(r, &msg)
	err := api.bc.AddBlock(msg)
	if err != nil {
		res := models.Error{Text: fmt.Sprintf("cannot create block due to error: %v", err)}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		out, err := json.Marshal(res)
		if err != nil {
			api.logger.Debug(err)
			return
		}
		w.Write(out)
		return
	}
	api.bc.PrintChain()
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
