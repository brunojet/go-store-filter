# Checklist para Expor Endpoints CRUD (Filter/FilterType)

## 1. Repository Layer
- [x] Definir interfaces de repositório para Filter e FilterType
- [x] Implementar repositórios usando GORM

## 2. Usecase/Service Layer
- [x] Criar serviços/casos de uso para Filter e FilterType (CRUD)
- [x] Implementar regras de negócio básicas

## 3. DTOs e Validação
- [ ] Definir structs DTO para requests/responses
- [ ] Adicionar validação de entrada (ex: validator)

## 4. Handler/Controller Layer
- [ ] Criar handlers HTTP para Filter e FilterType
- [ ] Mapear rotas REST (GET, POST, PUT, DELETE)

## 5. Injeção de Dependências
- [ ] Montar a aplicação conectando repository, service e handler
- [ ] Configurar router (ex: gorilla/mux)

## 6. Testes
- [ ] Escrever testes unitários para repositórios
- [ ] Escrever testes unitários para serviços
- [ ] Escrever testes para handlers (opcional)

## 7. Documentação
- [ ] Documentar endpoints e exemplos de uso

---

Siga essa ordem para garantir boas práticas, separação de responsabilidades e fácil manutenção do código.
