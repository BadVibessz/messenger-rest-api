# Домашнее задание 6

---

Вам необходимо продолжить разработку веб сервера чата из домашнего задания №5  

На этот раз нужно дополнить существующую реализацию новыми функциональными и нефункциональными требованиями

---

### Функциональные требования

- Реализованы все функциональные требования из ДЗ №5
- Добавить авторизацию и аутентификацию через JWT 
  - Должна быть возможность переключаться между различными типами авторизации и аутентификации через конфиг приложения

### Нефункциональные требования

- Реализованы все нефункциональные требования из ДЗ №5
- Добавлена кастомная **_logger middleware_** (вам нужно написать ее самими).
  - Поддерживает различные уровни логирования 
  - Выводить время
  - Вывод уровня ошибки
  - Вывод ошибки
  - Вывод сообщения
  - Вывод http статус кода
- Добавлена recovery middleware с логированием 
- Добавлен конфиг приложения
- Добавлена реализация repository с поддержкой PostgresSQL
  - Есть миграции (up/down)
  - Есть инструкция в Makefile для того, чтобы поднимать и опускать миграции
  - Есть возможность переключаться между in_memory дб и postgres_sql через конфиг приложения
- Тестовое покрытие unit тестами не менее 85%  


---

### Уточнения

- Вам может потребоваться изменение сущностей и переработка интерфейсов для того, чтобы можно было безболезненно 
переключаться между in_memory и postgres db

---

### Критерии приемки дз

- Скриншоты схемы БД
- В коде все должно быть реализованно в соответствии с нефункциональными требованиями

---

### Доп ссылки

* https://habr.com/ru/companies/ruvds/articles/566198/ (про мидлвары)
* https://habr.com/ru/companies/oleg-bunin/articles/461935/ (работа с бд)
* https://habr.com/ru/companies/avito/articles/716516/ (работа с бд)
* https://habr.com/ru/articles/780280/ (работа с миграциями)
* Unit тесты
  * https://www.youtube.com/watch?v=fMUNBJPhP6Y&list=PLbTTxxr-hMmxZMXsvaE-PozXxktdJ5zLR&index=3&pp=iAQB (Жашкевич 1 часть)
  * https://www.youtube.com/watch?v=Mvw5fbHGJFw&list=PLbTTxxr-hMmxZMXsvaE-PozXxktdJ5zLR&index=4&pp=iAQB (Жашкевич 2 часть)
  * https://www.youtube.com/watch?v=QJq3PZ1V-5Y&list=PLbTTxxr-hMmxZMXsvaE-PozXxktdJ5zLR&index=5&pp=iAQB (Жашкевич 3 часть)
  * https://habr.com/ru/companies/otus/articles/739468/ (про тестовое покрытие)
  * https://habr.com/ru/companies/avito/articles/658907/ (про моки)
  * 
* Структура проекта
  * https://habr.com/ru/companies/inDrive/articles/690088/ (про то за что отвечает каждый пакет)
* Конфигурация приложения
  * https://habr.com/ru/articles/479882/ (статья о том как конфигурировать и что это такое)
  * https://github.com/spf13/viper (библиотека для конфига приложений)
