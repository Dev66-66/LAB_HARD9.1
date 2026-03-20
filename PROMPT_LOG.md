# ОБЩИЙ PROMPT LOG

## [High] Лабораторная работа №1: Go Compute + Python Client
- **Request:** Создать микросервис на Go для тяжелых вычислений и Python-клиент (HTTP).
- **Response:** Реализован Go-бэкенд с воркер-пулом (горутины/каналы), написан Python-клиент, добавлены Unit-тесты.

## [High] Лабораторная работа №2: Go Orchestrator + Rust Crypto + Python
- **Request:** Использовать Go для оркестрации и Rust-библиотеку (PyO3) для криптографии.
- **Response:** Реализован Go-оркестратор задач, Rust-библиотека с XOR-шифрованием, Python-клиент для связи компонентов.

## [Average] Лабораторная работа №3: Simple Go HTTP Server
- **Request:** Написать простой HTTP-сервер на Go с одним эндпоинтом.
- **Response:** Реализован сервер с эндпоинтом `/hello` и соответствующим тестом.

## [Average] Лабораторная работа №4: Go Binary + Python (Subprocess)
- **Request:** Скомпилировать Go-программу в бинарный файл и вызвать из Python через subprocess.
- **Response:** Создан Go CLI инструмент, скомпилирован в `.exe`, реализован вызов через `subprocess.run` в Python.

## [Average] Задание №6: Rust Library
- **Request:** Установить Rust и создать библиотеку с одной функцией.
- **Response:** Создана Rust-библиотека `simple_lib` с функцией `add` и встроенными тестами.

## Реструктуризация проекта
- **Request:** Разделить проекты на папки High и Average, объединить логи и провести Code Review.
- **Response:** Проведен анализ кода, проекты перемещены, создан единый PROMPT_LOG.md.
