#!/bin/bash

# 🔄 Скрипт для обновления proto-файлов

# Цвета для вывода
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}🔄 Обновление proto-файлов...${NC}"

# Создаем директории, если их нет
mkdir -p api-gateway/gen/go/

# Копируем proto-файлы из user-service в api-gateway
echo -e "${YELLOW}📋 Копирование proto-файлов из user-service...${NC}"
cp -r user-service/gen/go/* api-gateway/gen/go/

echo -e "${GREEN}✅ Proto-файлы успешно обновлены!${NC}"

# Проверяем, есть ли изменения в git
if git diff --quiet api-gateway/gen/go/; then
    echo -e "${YELLOW}ℹ️ Нет изменений в proto-файлах.${NC}"
else
    echo -e "${GREEN}🔄 Есть изменения в proto-файлах, которые будут добавлены в git.${NC}"
    git add api-gateway/gen/go/
    git status | grep api-gateway/gen/go/
fi

echo -e "${CYAN}🎯 Готово!${NC}" 