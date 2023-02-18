package custommiddleware

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/richmondgoh8/boilerplate/pkg/uuid"
	"net/http"
)

type contextKey struct {
	name string
}

var (
	// TrackingCtxKey is the context.Context key to store the tracking id for a request.
	trackingCtxKey = &contextKey{"trackingID"}
)

func InjectTrackingID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		randomTrackingID := uuid.GenerateTrackingID()

		ctx = context.WithValue(ctx, trackingCtxKey.name, randomTrackingID)
		ctx = context.WithValue(ctx, middleware.RequestIDKey, randomTrackingID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
