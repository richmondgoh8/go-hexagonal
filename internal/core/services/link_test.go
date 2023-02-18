package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/richmondgoh8/boilerplate/internal/core/domain"
	mock_ports "github.com/richmondgoh8/boilerplate/internal/mocks/core/ports"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetURLData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	tests := []struct {
		name       string
		doMockRepo func(*mock_ports.MockLinkRepository)
		err        error
	}{
		{
			name: "Test Case Positive",
			doMockRepo: func(repository *mock_ports.MockLinkRepository) {
				repository.EXPECT().GetURL(gomock.Any(), gomock.Any()).Return(domain.Link{
					ID:   "1",
					Url:  "www.google.com.sg",
					Name: "Google",
				}, nil)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockLinkRepo := mock_ports.NewMockLinkRepository(mockCtrl)
			tc.doMockRepo(mockLinkRepo)

			linkSvc := NewLinkSvc(mockLinkRepo)
			_, err := linkSvc.GetURLData(context.Background(), "")
			assert.Equal(t, tc.err, err)
		})
	}
}
