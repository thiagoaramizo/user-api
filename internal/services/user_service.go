package service

import (
    "gerenciamento-servicos/internal/models"
    "gerenciamento-servicos/internal/repository"
)

func GetUser(id uint) (models.User, error) {
    return repository.GetUserByID(id)
}