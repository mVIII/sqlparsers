antlr_path:="/usr/local/lib/antlr-4.7.1-complete.jar"
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

#help
compile-pg:
	@echo Compiling postgres grammar
	java -Xmx500M -cp "$(antlr_path):$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go  -o $(ROOT_DIR)/postgres/parser $(ROOT_DIR)/postgres/PostgreSQLLexer.g4  -o $(ROOT_DIR)/postgres/parser $(ROOT_DIR)/postgres/PostgreSQLParser.g4 -package parser

compile-mysql:
	@echo Compiling mysql grammar
	java -Xmx500M -cp "$(antlr_path):$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go  -o $(ROOT_DIR)/mysql/parser $(ROOT_DIR)/mysql/MySqlLexer.g4  -o $(ROOT_DIR)/mysql/parser $(ROOT_DIR)/mysql/MySqlParser.g4 -package parser