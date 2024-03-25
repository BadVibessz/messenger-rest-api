// nolint
package in_memory

import (
	"context"
	"errors"
	"github.com/ew0s/ewos-to-go-hw/chat-server/internal/repository"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/ew0s/ewos-to-go-hw/chat-server/internal/domain/entity"

	inmemory "github.com/ew0s/ewos-to-go-hw/chat-server/pkg/db/in-memory"
)

type PublicMessageRepo struct {
	DB    inmemory.InMemoryDB
	mutex sync.RWMutex
}

func NewPublicMessageRepo(db inmemory.InMemoryDB) *PublicMessageRepo {
	repo := PublicMessageRepo{
		DB:    db,
		mutex: sync.RWMutex{},
	}

	_, err := repo.DB.GetTable(PublicMessageTableName)
	if errors.Is(err, inmemory.ErrNotExistedTable) {
		repo.DB.CreateTable(PublicMessageTableName)
	}

	return &repo
}

func (pr *PublicMessageRepo) AddPublicMessage(_ context.Context, msg entity.PublicMessage) (*entity.PublicMessage, error) {
	pr.mutex.Lock()
	defer pr.mutex.Unlock()

	idOffset, err := pr.DB.GetTableCounter(PublicMessageTableName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	msg.ID = idOffset + 1
	msg.SentAt = now
	msg.EditedAt = now

	if err = pr.DB.AddRow(PublicMessageTableName, strconv.Itoa(msg.ID), msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func (pr *PublicMessageRepo) getAllPublicMessages(_ context.Context, offset, limit int) []*entity.PublicMessage {
	rows, err := pr.DB.GetAllRows(PublicMessageTableName, offset, limit)
	if err != nil {
		return nil
	}

	res := make([]*entity.PublicMessage, 0, len(rows))

	for _, row := range rows {
		msg, ok := row.(entity.PublicMessage)
		if ok {
			res = append(res, &msg)
		}
	}

	sort.Slice(res, func(i, j int) bool { return res[i].SentAt.Before(res[j].SentAt) })

	return res
}

func (pr *PublicMessageRepo) GetAllPublicMessages(ctx context.Context, offset, limit int) []*entity.PublicMessage {
	pr.mutex.RLock()
	defer pr.mutex.RUnlock()

	return pr.getAllPublicMessages(ctx, offset, limit)
}

func (pr *PublicMessageRepo) getPublicMessage(_ context.Context, id int) (*entity.PublicMessage, error) {
	row, err := pr.DB.GetRow(PublicMessageTableName, strconv.Itoa(id))
	if err != nil {
		return nil, repository.ErrNoSuchPublicMessage
	}

	msg, ok := row.(entity.PublicMessage)
	if !ok {
		return nil, repository.ErrNoSuchPublicMessage
	}

	return &msg, nil
}

func (pr *PublicMessageRepo) GetPublicMessage(ctx context.Context, id int) (*entity.PublicMessage, error) {
	pr.mutex.RLock()
	defer pr.mutex.RUnlock()

	return pr.getPublicMessage(ctx, id)
}
