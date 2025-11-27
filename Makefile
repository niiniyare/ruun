# ============================================================================
# 🧱 Environment & Defaults
# ============================================================================

SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables --no-builtin-rules --no-print-directory
.DEFAULT_GOAL := help

# ============================================================================
# 📋 Variables & Configuration
# ============================================================================

#V Database Configuration
# DB_NAME ?= ledger
# DB_USER ?= admin
# DB_PSSWD ?= admin
# DB_HOST ?= localhost
# DB_PORT ?= 5432
# DB_URL ?= postgresql://$(DB_USER):$(DB_PSSWD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Paths
MIGRATION_PATH := db/migration
SQLC_OUT := db/sqlc
DOCS_PATH := docs
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

# Ports
DOC_PORT ?= 8081
GRPC_PORT ?= 9090

# Tools
REQUIRED_TOOLS := go node templ pnpm




# Colors for output
ESC := $(shell printf '\033')
RED := $(ESC)[0;31m
GREEN := $(ESC)[0;32m
YELLOW := $(ESC)[0;33m
BLUE := $(ESC)[0;34m
PURPLE := $(ESC)[0;35m
CYAN := $(ESC)[0;36m
NC := $(ESC)[0m

# ============================================================================
# 🆘 Help & Utilities
# ============================================================================

.PHONY: help
help: ## 📚 Show this help message
	@echo "$(CYAN)Available targets:$(NC)"
	@awk 'BEGIN {FS = ":.*##"; printf "\n"} \
	/^[a-zA-Z0-9_.-]+:.*?##/ { \
		gsub(/^[ \t]+/, "", $$2); \
		printf "  $(GREEN)%-20s$(NC) %s\n", $$1, $$2 \
	}' $(MAKEFILE_LIST)
	@echo ""

.PHONY: check-tools
check-tools: ## 🔧 Check if required tools are installed
	@echo "$(BLUE)Checking required tools...$(NC)"
	@$(foreach tool,$(REQUIRED_TOOLS),\
		command -v $(tool) >/dev/null 2>&1 || { \
			echo "$(RED)❌ Missing: $(tool)$(NC)"; exit 1; \
		};)
	@echo "$(GREEN)✅ All required tools are installed$(NC)"


.PHONY: status
status: ## 📊 Show project status and configuration
	@echo "$(CYAN)Project Configuration:$(NC)"
	@echo "  Database: $(DB_USER)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)"
	@echo "  Migration Path: $(MIGRATION_PATH)"
	@echo "  Documentation Port: $(DOC_PORT)"
	@echo "  Test Timeout: $(TEST_TIMEOUT)"
	@echo "  Parallel Jobs: $(TEST_PARALLEL_JOBS)"

# ============================================================================
# 🏗️ Build & Generation
# ============================================================================


.PHONY: build
build: ## 🗄️ Build and minify Static files 
	@echo "$(BLUE)Build and minify Static files ...$(NC)"
	@cd static && pnpm  build && cd .. 
	@echo "$(GREEN)✅ Building is complete$(NC)"
