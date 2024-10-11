package services

import (
	"errors"
	"github.com/DoktorGhost/golibrary/internal/core/library/subdomain_rental/repositories/postgres/dao"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRentalService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRentalsRepository(ctrl)
	userService := NewRentalService(mockRepo)

	// 1. CreateRentalInfo успешный сценарий
	t.Run("CreateRentalInfo success", func(t *testing.T) {
		userID := 1
		bookID := 10

		mockRepo.EXPECT().CreateRentalsInfo(userID, bookID).Return(1, nil)

		rentalID, err := userService.CreateRentalInfo(userID, bookID)

		assert.NoError(t, err)
		assert.Equal(t, 1, rentalID)
	})
	// 2. CreateRentalInfo fail
	t.Run("CreateRentalInfo fail", func(t *testing.T) {
		userID := 1
		bookID := 10

		mockRepo.EXPECT().CreateRentalsInfo(userID, bookID).Return(0, errors.New("ERROR"))

		rentalID, err := userService.CreateRentalInfo(userID, bookID)

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})
	// 3. GetRentalsInfoByID успешный сценарий
	t.Run("GetRentalsInfoByID success", func(t *testing.T) {
		rentalID := 1

		mockRepo.EXPECT().GetRentalsInfoByID(rentalID).Return(dao.RentalsTable{}, nil)

		rental, err := userService.GetRentalsInfoByID(rentalID)

		assert.NoError(t, err)
		assert.Equal(t, dao.RentalsTable{}, rental)
	})
	// 4. GetRentalsInfoByID fail
	t.Run("GetRentalsInfoByID fail", func(t *testing.T) {
		rentalID := 1

		mockRepo.EXPECT().GetRentalsInfoByID(rentalID).Return(dao.RentalsTable{}, errors.New("ERROR"))

		rental, err := userService.GetRentalsInfoByID(rentalID)

		assert.Error(t, err)
		assert.Equal(t, dao.RentalsTable{}, rental)
	})
	// 5. UpdateRentalsInfo успешный сценарий
	t.Run("UpdateRentalsInfo success", func(t *testing.T) {
		rental := dao.RentalsTable{}

		mockRepo.EXPECT().UpdateRentalsInfo(rental).Return(nil)

		err := userService.UpdateRentalsInfo(rental)

		assert.NoError(t, err)

	})
	// 6. UpdateRentalsInfo fail
	t.Run("UpdateRentalsInfo fail", func(t *testing.T) {
		rental := dao.RentalsTable{}

		mockRepo.EXPECT().UpdateRentalsInfo(rental).Return(errors.New("ERROR"))

		err := userService.UpdateRentalsInfo(rental)

		assert.Error(t, err)

	})
	// 7. DeleteRentalsInfo успешный сценарий
	t.Run("DeleteRentalsInfo success", func(t *testing.T) {
		rentalID := 1

		mockRepo.EXPECT().DeleteRentalsInfo(rentalID).Return(nil)

		err := userService.DeleteRentalsInfo(rentalID)

		assert.NoError(t, err)

	})
	// 8. DeleteRentalsInfo fail
	t.Run("DeleteRentalsInfo fail", func(t *testing.T) {
		rentalID := 1

		mockRepo.EXPECT().DeleteRentalsInfo(rentalID).Return(errors.New("ERROR"))

		err := userService.DeleteRentalsInfo(rentalID)

		assert.Error(t, err)

	})
	// 9. CreateRentals успешный сценарий
	t.Run("CreateRentals success", func(t *testing.T) {
		bookID := 1

		mockRepo.EXPECT().CreateRentals(bookID).Return(nil)

		err := userService.CreateRentals(bookID)

		assert.NoError(t, err)

	})
	// 10. CreateRentals fail
	t.Run("CreateRentals fail", func(t *testing.T) {
		bookID := 1

		mockRepo.EXPECT().CreateRentals(bookID).Return(errors.New("ERROR"))

		err := userService.CreateRentals(bookID)

		assert.Error(t, err)

	})
	// 11. GetRentalsByID успешный сценарий
	t.Run("GetRentalsByID success", func(t *testing.T) {
		bookID := 1

		mockRepo.EXPECT().GetRentalsByID(bookID).Return(1, nil)

		rentalID, err := userService.GetRentalsByID(bookID)

		assert.NoError(t, err)
		assert.Equal(t, 1, rentalID)
	})
	// 12. GetRentalsByID fail
	t.Run("GetRentalsByID fail", func(t *testing.T) {
		bookID := 1

		mockRepo.EXPECT().GetRentalsByID(bookID).Return(0, errors.New("ERROR"))

		rentalID, err := userService.GetRentalsByID(bookID)

		assert.Error(t, err)
		assert.Equal(t, 0, rentalID)
	})
	// 13. UpdateRentals успешный сценарий
	t.Run("UpdateRentals success", func(t *testing.T) {
		id := 1
		rentalID := 1

		mockRepo.EXPECT().UpdateRentals(id, rentalID).Return(nil)

		err := userService.UpdateRentals(id, rentalID)

		assert.NoError(t, err)

	})
	// 14. UpdateRentals fail
	t.Run("UpdateRentals fail", func(t *testing.T) {
		id := 1
		rentalID := 1

		mockRepo.EXPECT().UpdateRentals(id, rentalID).Return(errors.New("ERROR"))

		err := userService.UpdateRentals(id, rentalID)

		assert.Error(t, err)

	})
	// 15. DeleteRentals успешный сценарий
	t.Run("DeleteRentals success", func(t *testing.T) {
		id := 1

		mockRepo.EXPECT().DeleteRentals(id).Return(nil)

		err := userService.DeleteRentals(id)

		assert.NoError(t, err)

	})
	// 16. DeleteRentals fail
	t.Run("DeleteRentals fail", func(t *testing.T) {
		id := 1

		mockRepo.EXPECT().DeleteRentals(id).Return(errors.New("ERROR"))

		err := userService.DeleteRentals(id)

		assert.Error(t, err)

	})
	// 17. GetActiveRentals успешный сценарий
	t.Run("GetActiveRentals success", func(t *testing.T) {
		maps := make(map[int][]int)

		mockRepo.EXPECT().GetActiveRentals().Return(maps, nil)

		m, err := userService.GetActiveRentals()

		assert.NoError(t, err)
		assert.Equal(t, maps, m)
	})
	// 18. GetActiveRentals fail
	t.Run("GetActiveRentals fail", func(t *testing.T) {

		var maps map[int][]int
		mockRepo.EXPECT().GetActiveRentals().Return(nil, errors.New("ERROR"))

		m, err := userService.GetActiveRentals()

		assert.Error(t, err)
		assert.Equal(t, maps, m)

	})
	// 19. GetTopAuthorsByPeriod успешный сценарий
	t.Run("GetTopAuthorsByPeriod success", func(t *testing.T) {
		period := 2
		limit := 2

		var topAuthors []dao.TopAuthor

		mockRepo.EXPECT().GetTopAuthors(period, limit).Return(topAuthors, nil)

		m, err := userService.GetTopAuthorsByPeriod(period, limit)

		assert.NoError(t, err)
		assert.Equal(t, topAuthors, m)
	})
	// 20. GetTopAuthorsByPeriod fail
	t.Run("GetTopAuthorsByPeriod fail", func(t *testing.T) {
		period := 2
		limit := 2
		var topAuthors []dao.TopAuthor

		mockRepo.EXPECT().GetTopAuthors(period, limit).Return(topAuthors, errors.New("ERROR"))

		m, err := userService.GetTopAuthorsByPeriod(period, limit)

		assert.Error(t, err)
		assert.Equal(t, topAuthors, m)

	})
}

/*
func (s *RentalService) GetTopAuthorsByPeriod(period, limit int) ([]dao.TopAuthor, error) {
	authors, err := s.repo.GetTopAuthors(period, limit)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения топ авторов за период %d: %v", period, err)
	}
	return authors, nil
}
*/
