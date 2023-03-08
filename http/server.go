package http

import (
	"context"
	"encoding/json"
	"lazzytchik/L0/storage"
	"log"
	"net"
	"net/http"
	"strconv"
)

type OrderServer struct {
	Storage    storage.Orders
	Logger     *log.Logger
	HttpServer *http.Server
}

func New(storage storage.Orders, logger *log.Logger) *OrderServer {
	return &OrderServer{
		Storage: storage,
		Logger:  logger,
	}
}

func (s *OrderServer) getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		s.Logger.Println("wrong request method", "Want:", http.MethodGet, "Got:", r.Method)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	keys, found := r.URL.Query()["id"]
	if !found {
		s.Logger.Println("Wrong query parameter")

		notFoundResponse, err := json.Marshal(map[string]string{"reason": "wrong query parameter"})
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(notFoundResponse)
		if err != nil {
			return
		}
		return
	}

	id, convErr := strconv.Atoi(keys[0])
	if convErr != nil {
		s.Logger.Println("Wrong type of query parameter")

		wrongParamResponse, err := json.Marshal(map[string]string{"reason": "wrong type query parameter"})
		if err != nil {
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(wrongParamResponse)
		if err != nil {
			return
		}
	}
	order, err := s.Storage.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	orderResponse, _ := json.Marshal(order)
	_, err = w.Write(orderResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *OrderServer) Serve(ctx context.Context) error {
	router := http.NewServeMux()
	router.HandleFunc("/orders", s.getOrdersHandler)
	address := ":3000"

	srv := http.Server{
		Addr:        address,
		Handler:     router,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	s.HttpServer = &srv

	s.Logger.Printf("Server started on http://localhost%s.", address)

	return srv.ListenAndServe()
}

func (s *OrderServer) Stop(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}
