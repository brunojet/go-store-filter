

# go-store-filter

## Proposta do Projeto

Este projeto será usado para uma prova de conceito de filtros dinâmicos para aplicativos, com foco em arquitetura limpa, manutenção e escalabilidade.

O objetivo é criar uma API RESTful em Go para gerenciamento de filtros e tipos de filtros, com integração ao MySQL, suporte a hierarquia e identificadores públicos.

## Arquitetura
- **Clean Architecture**: Separação clara entre domínio, casos de uso, repositórios e handlers.
- **Camadas:**
	- domain: entidades de negócio
	- repository: interfaces e implementações de acesso a dados
	- usecase: regras de negócio e orquestração
	- handler: endpoints HTTP
	- dto: objetos de transferência de dados (request/response)
- **Banco de Dados**: MySQL, charset utf8mb4, collation utf8mb4_0900_ai_ci
- **ORM**: GORM

## Entidades
- **tipo_filtro**: representa tipos de filtros, com suporte a hierarquia (parent/children) e identificador público (cod_tip_flo_cto)
- **filtro**: representa filtros, também com suporte a hierarquia e identificador público (cod_flo_cto)

## Status Atual
- Estrutura de pastas criada seguindo Clean Architecture
- Models e migrations para tipo_filtro e filtro implementados
- Checklist de etapas para expor endpoints CRUD documentado em `docs/api_crud_checklist.md`
- Integração com repositório GitHub e versionamento inicial realizado

## Próximos Passos
1. Implementar camada de repositórios (interfaces e GORM)
2. Implementar camada de casos de uso/serviços
3. Criar DTOs e validação
4. Implementar handlers HTTP e rotas
5. Testes unitários e documentação dos endpoints

---

Este projeto serve como base para sistemas que necessitam de filtros dinâmicos, hierárquicos e facilmente integráveis a outros domínios de negócio.
