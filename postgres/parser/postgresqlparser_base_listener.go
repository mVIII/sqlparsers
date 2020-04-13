// Code generated from /home/kovacs/Projects/github.com/mVIII/sqlparsers/postgres/PostgreSQLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PostgreSQLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePostgreSQLParserListener is a complete listener for a parse tree produced by PostgreSQLParser.
type BasePostgreSQLParserListener struct{}

var _ PostgreSQLParserListener = &BasePostgreSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePostgreSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePostgreSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePostgreSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePostgreSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BasePostgreSQLParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasePostgreSQLParserListener) ExitRoot(ctx *RootContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasePostgreSQLParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasePostgreSQLParserListener) ExitStmt(ctx *StmtContext) {}

// EnterAbort_stmt is called when production abort_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAbort_stmt(ctx *Abort_stmtContext) {}

// ExitAbort_stmt is called when production abort_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAbort_stmt(ctx *Abort_stmtContext) {}

// EnterAlter_stmt is called when production alter_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_stmt(ctx *Alter_stmtContext) {}

// ExitAlter_stmt is called when production alter_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_stmt(ctx *Alter_stmtContext) {}

// EnterAlter_aggregate_stmt is called when production alter_aggregate_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_aggregate_stmt(ctx *Alter_aggregate_stmtContext) {}

// ExitAlter_aggregate_stmt is called when production alter_aggregate_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_aggregate_stmt(ctx *Alter_aggregate_stmtContext) {}

// EnterAlter_collation_stmt is called when production alter_collation_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_collation_stmt(ctx *Alter_collation_stmtContext) {}

// ExitAlter_collation_stmt is called when production alter_collation_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_collation_stmt(ctx *Alter_collation_stmtContext) {}

// EnterAlter_conversion_stmt is called when production alter_conversion_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_conversion_stmt(ctx *Alter_conversion_stmtContext) {
}

// ExitAlter_conversion_stmt is called when production alter_conversion_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_conversion_stmt(ctx *Alter_conversion_stmtContext) {}

// EnterAlter_database_stmt is called when production alter_database_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_database_stmt(ctx *Alter_database_stmtContext) {}

// ExitAlter_database_stmt is called when production alter_database_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_database_stmt(ctx *Alter_database_stmtContext) {}

// EnterAlter_default_privileges_stmt is called when production alter_default_privileges_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_default_privileges_stmt(ctx *Alter_default_privileges_stmtContext) {
}

// ExitAlter_default_privileges_stmt is called when production alter_default_privileges_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_default_privileges_stmt(ctx *Alter_default_privileges_stmtContext) {
}

// EnterAlter_domain_stmt is called when production alter_domain_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_domain_stmt(ctx *Alter_domain_stmtContext) {}

// ExitAlter_domain_stmt is called when production alter_domain_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_domain_stmt(ctx *Alter_domain_stmtContext) {}

// EnterAlter_event_trigger_stmt is called when production alter_event_trigger_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_event_trigger_stmt(ctx *Alter_event_trigger_stmtContext) {
}

// ExitAlter_event_trigger_stmt is called when production alter_event_trigger_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_event_trigger_stmt(ctx *Alter_event_trigger_stmtContext) {
}

// EnterAlter_extension_stmt is called when production alter_extension_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_extension_stmt(ctx *Alter_extension_stmtContext) {}

// ExitAlter_extension_stmt is called when production alter_extension_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_extension_stmt(ctx *Alter_extension_stmtContext) {}

// EnterAlter_foreign_data_wrapper_stmt is called when production alter_foreign_data_wrapper_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_foreign_data_wrapper_stmt(ctx *Alter_foreign_data_wrapper_stmtContext) {
}

// ExitAlter_foreign_data_wrapper_stmt is called when production alter_foreign_data_wrapper_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_foreign_data_wrapper_stmt(ctx *Alter_foreign_data_wrapper_stmtContext) {
}

// EnterAlter_foreign_table_action is called when production alter_foreign_table_action is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_foreign_table_action(ctx *Alter_foreign_table_actionContext) {
}

// ExitAlter_foreign_table_action is called when production alter_foreign_table_action is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_foreign_table_action(ctx *Alter_foreign_table_actionContext) {
}

// EnterAlter_foreign_table_action_list is called when production alter_foreign_table_action_list is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_foreign_table_action_list(ctx *Alter_foreign_table_action_listContext) {
}

// ExitAlter_foreign_table_action_list is called when production alter_foreign_table_action_list is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_foreign_table_action_list(ctx *Alter_foreign_table_action_listContext) {
}

// EnterAlter_foreign_table_stmt is called when production alter_foreign_table_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_foreign_table_stmt(ctx *Alter_foreign_table_stmtContext) {
}

// ExitAlter_foreign_table_stmt is called when production alter_foreign_table_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_foreign_table_stmt(ctx *Alter_foreign_table_stmtContext) {
}

// EnterAlter_function_stmt is called when production alter_function_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_function_stmt(ctx *Alter_function_stmtContext) {}

// ExitAlter_function_stmt is called when production alter_function_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_function_stmt(ctx *Alter_function_stmtContext) {}

// EnterAlter_group_stmt is called when production alter_group_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_group_stmt(ctx *Alter_group_stmtContext) {}

// ExitAlter_group_stmt is called when production alter_group_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_group_stmt(ctx *Alter_group_stmtContext) {}

// EnterAlter_index_stmt is called when production alter_index_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_index_stmt(ctx *Alter_index_stmtContext) {}

// ExitAlter_index_stmt is called when production alter_index_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_index_stmt(ctx *Alter_index_stmtContext) {}

// EnterAlter_language_stmt is called when production alter_language_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_language_stmt(ctx *Alter_language_stmtContext) {}

// ExitAlter_language_stmt is called when production alter_language_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_language_stmt(ctx *Alter_language_stmtContext) {}

// EnterAlter_large_object_stmt is called when production alter_large_object_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_large_object_stmt(ctx *Alter_large_object_stmtContext) {
}

// ExitAlter_large_object_stmt is called when production alter_large_object_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_large_object_stmt(ctx *Alter_large_object_stmtContext) {
}

// EnterAlter_materialize_view_stmt is called when production alter_materialize_view_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_materialize_view_stmt(ctx *Alter_materialize_view_stmtContext) {
}

// ExitAlter_materialize_view_stmt is called when production alter_materialize_view_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_materialize_view_stmt(ctx *Alter_materialize_view_stmtContext) {
}

// EnterAlter_operator_stmt is called when production alter_operator_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_operator_stmt(ctx *Alter_operator_stmtContext) {}

// ExitAlter_operator_stmt is called when production alter_operator_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_operator_stmt(ctx *Alter_operator_stmtContext) {}

// EnterAlter_operator_class_stmt is called when production alter_operator_class_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_operator_class_stmt(ctx *Alter_operator_class_stmtContext) {
}

// ExitAlter_operator_class_stmt is called when production alter_operator_class_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_operator_class_stmt(ctx *Alter_operator_class_stmtContext) {
}

// EnterAlter_operator_family_stmt is called when production alter_operator_family_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_operator_family_stmt(ctx *Alter_operator_family_stmtContext) {
}

// ExitAlter_operator_family_stmt is called when production alter_operator_family_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_operator_family_stmt(ctx *Alter_operator_family_stmtContext) {
}

// EnterAlter_policy_stmt is called when production alter_policy_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_policy_stmt(ctx *Alter_policy_stmtContext) {}

// ExitAlter_policy_stmt is called when production alter_policy_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_policy_stmt(ctx *Alter_policy_stmtContext) {}

// EnterAlter_publication_stmt is called when production alter_publication_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_publication_stmt(ctx *Alter_publication_stmtContext) {
}

// ExitAlter_publication_stmt is called when production alter_publication_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_publication_stmt(ctx *Alter_publication_stmtContext) {
}

// EnterAlter_role_options is called when production alter_role_options is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_role_options(ctx *Alter_role_optionsContext) {}

// ExitAlter_role_options is called when production alter_role_options is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_role_options(ctx *Alter_role_optionsContext) {}

// EnterAlter_role_stmt is called when production alter_role_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_role_stmt(ctx *Alter_role_stmtContext) {}

// ExitAlter_role_stmt is called when production alter_role_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_role_stmt(ctx *Alter_role_stmtContext) {}

// EnterAlter_rule_stmt is called when production alter_rule_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_rule_stmt(ctx *Alter_rule_stmtContext) {}

// ExitAlter_rule_stmt is called when production alter_rule_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_rule_stmt(ctx *Alter_rule_stmtContext) {}

// EnterAlter_schema_stmt is called when production alter_schema_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_schema_stmt(ctx *Alter_schema_stmtContext) {}

// ExitAlter_schema_stmt is called when production alter_schema_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_schema_stmt(ctx *Alter_schema_stmtContext) {}

// EnterAlter_sequence_stmt is called when production alter_sequence_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_sequence_stmt(ctx *Alter_sequence_stmtContext) {}

// ExitAlter_sequence_stmt is called when production alter_sequence_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_sequence_stmt(ctx *Alter_sequence_stmtContext) {}

// EnterAlter_server_options_list is called when production alter_server_options_list is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_server_options_list(ctx *Alter_server_options_listContext) {
}

// ExitAlter_server_options_list is called when production alter_server_options_list is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_server_options_list(ctx *Alter_server_options_listContext) {
}

// EnterAlter_server_stmt is called when production alter_server_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_server_stmt(ctx *Alter_server_stmtContext) {}

// ExitAlter_server_stmt is called when production alter_server_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_server_stmt(ctx *Alter_server_stmtContext) {}

// EnterAlter_statistics_stmt is called when production alter_statistics_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_statistics_stmt(ctx *Alter_statistics_stmtContext) {
}

// ExitAlter_statistics_stmt is called when production alter_statistics_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_statistics_stmt(ctx *Alter_statistics_stmtContext) {}

// EnterAlter_subscription_stmt is called when production alter_subscription_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_subscription_stmt(ctx *Alter_subscription_stmtContext) {
}

// ExitAlter_subscription_stmt is called when production alter_subscription_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_subscription_stmt(ctx *Alter_subscription_stmtContext) {
}

// EnterAlter_system_stmt is called when production alter_system_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_system_stmt(ctx *Alter_system_stmtContext) {}

// ExitAlter_system_stmt is called when production alter_system_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_system_stmt(ctx *Alter_system_stmtContext) {}

// EnterAlter_table_stmt is called when production alter_table_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// ExitAlter_table_stmt is called when production alter_table_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// EnterAlter_tablespace_stmt is called when production alter_tablespace_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_tablespace_stmt(ctx *Alter_tablespace_stmtContext) {
}

// ExitAlter_tablespace_stmt is called when production alter_tablespace_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_tablespace_stmt(ctx *Alter_tablespace_stmtContext) {}

// EnterAlter_text_search_config_stmt is called when production alter_text_search_config_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_text_search_config_stmt(ctx *Alter_text_search_config_stmtContext) {
}

// ExitAlter_text_search_config_stmt is called when production alter_text_search_config_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_text_search_config_stmt(ctx *Alter_text_search_config_stmtContext) {
}

// EnterAlter_text_search_dict_stmt is called when production alter_text_search_dict_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_text_search_dict_stmt(ctx *Alter_text_search_dict_stmtContext) {
}

// ExitAlter_text_search_dict_stmt is called when production alter_text_search_dict_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_text_search_dict_stmt(ctx *Alter_text_search_dict_stmtContext) {
}

// EnterAlter_text_search_parser_stmt is called when production alter_text_search_parser_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_text_search_parser_stmt(ctx *Alter_text_search_parser_stmtContext) {
}

// ExitAlter_text_search_parser_stmt is called when production alter_text_search_parser_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_text_search_parser_stmt(ctx *Alter_text_search_parser_stmtContext) {
}

// EnterAlter_text_search_template_stmt is called when production alter_text_search_template_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_text_search_template_stmt(ctx *Alter_text_search_template_stmtContext) {
}

// ExitAlter_text_search_template_stmt is called when production alter_text_search_template_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_text_search_template_stmt(ctx *Alter_text_search_template_stmtContext) {
}

// EnterAlter_trigger_stmt is called when production alter_trigger_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_trigger_stmt(ctx *Alter_trigger_stmtContext) {}

// ExitAlter_trigger_stmt is called when production alter_trigger_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_trigger_stmt(ctx *Alter_trigger_stmtContext) {}

// EnterAlter_type_stmt is called when production alter_type_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_type_stmt(ctx *Alter_type_stmtContext) {}

// ExitAlter_type_stmt is called when production alter_type_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_type_stmt(ctx *Alter_type_stmtContext) {}

// EnterAlter_user_stmt is called when production alter_user_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_user_stmt(ctx *Alter_user_stmtContext) {}

// ExitAlter_user_stmt is called when production alter_user_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_user_stmt(ctx *Alter_user_stmtContext) {}

// EnterAlter_user_mapping_stmt is called when production alter_user_mapping_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_user_mapping_stmt(ctx *Alter_user_mapping_stmtContext) {
}

// ExitAlter_user_mapping_stmt is called when production alter_user_mapping_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_user_mapping_stmt(ctx *Alter_user_mapping_stmtContext) {
}

// EnterAlter_view_stmt is called when production alter_view_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAlter_view_stmt(ctx *Alter_view_stmtContext) {}

// ExitAlter_view_stmt is called when production alter_view_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAlter_view_stmt(ctx *Alter_view_stmtContext) {}

// EnterAnalyze_stmt is called when production analyze_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterAnalyze_stmt(ctx *Analyze_stmtContext) {}

// ExitAnalyze_stmt is called when production analyze_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitAnalyze_stmt(ctx *Analyze_stmtContext) {}

// EnterClose_stmt is called when production close_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterClose_stmt(ctx *Close_stmtContext) {}

// ExitClose_stmt is called when production close_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitClose_stmt(ctx *Close_stmtContext) {}

// EnterCluster_stmt is called when production cluster_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCluster_stmt(ctx *Cluster_stmtContext) {}

// ExitCluster_stmt is called when production cluster_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCluster_stmt(ctx *Cluster_stmtContext) {}

// EnterComment_stmt is called when production comment_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterComment_stmt(ctx *Comment_stmtContext) {}

// ExitComment_stmt is called when production comment_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitComment_stmt(ctx *Comment_stmtContext) {}

// EnterCommit_stmt is called when production commit_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCommit_stmt(ctx *Commit_stmtContext) {}

// ExitCommit_stmt is called when production commit_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCommit_stmt(ctx *Commit_stmtContext) {}

// EnterCommit_prepared_stmt is called when production commit_prepared_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCommit_prepared_stmt(ctx *Commit_prepared_stmtContext) {}

// ExitCommit_prepared_stmt is called when production commit_prepared_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCommit_prepared_stmt(ctx *Commit_prepared_stmtContext) {}

// EnterCopy_stmt is called when production copy_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCopy_stmt(ctx *Copy_stmtContext) {}

// ExitCopy_stmt is called when production copy_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCopy_stmt(ctx *Copy_stmtContext) {}

// EnterCreate_stmt is called when production create_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_stmt(ctx *Create_stmtContext) {}

// ExitCreate_stmt is called when production create_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_stmt(ctx *Create_stmtContext) {}

// EnterCreate_access_method_stmt is called when production create_access_method_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_access_method_stmt(ctx *Create_access_method_stmtContext) {
}

// ExitCreate_access_method_stmt is called when production create_access_method_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_access_method_stmt(ctx *Create_access_method_stmtContext) {
}

// EnterCreate_aggregate_stmt is called when production create_aggregate_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_aggregate_stmt(ctx *Create_aggregate_stmtContext) {
}

// ExitCreate_aggregate_stmt is called when production create_aggregate_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_aggregate_stmt(ctx *Create_aggregate_stmtContext) {}

// EnterCreate_cast_stmt is called when production create_cast_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_cast_stmt(ctx *Create_cast_stmtContext) {}

// ExitCreate_cast_stmt is called when production create_cast_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_cast_stmt(ctx *Create_cast_stmtContext) {}

// EnterCreate_collation_opt is called when production create_collation_opt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_collation_opt(ctx *Create_collation_optContext) {}

// ExitCreate_collation_opt is called when production create_collation_opt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_collation_opt(ctx *Create_collation_optContext) {}

// EnterCreate_collation_opt_list is called when production create_collation_opt_list is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_collation_opt_list(ctx *Create_collation_opt_listContext) {
}

// ExitCreate_collation_opt_list is called when production create_collation_opt_list is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_collation_opt_list(ctx *Create_collation_opt_listContext) {
}

// EnterCreate_collation_stmt is called when production create_collation_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_collation_stmt(ctx *Create_collation_stmtContext) {
}

// ExitCreate_collation_stmt is called when production create_collation_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_collation_stmt(ctx *Create_collation_stmtContext) {}

// EnterCreate_conversion_stmt is called when production create_conversion_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_conversion_stmt(ctx *Create_conversion_stmtContext) {
}

// ExitCreate_conversion_stmt is called when production create_conversion_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_conversion_stmt(ctx *Create_conversion_stmtContext) {
}

// EnterCreate_database_stmt is called when production create_database_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_database_stmt(ctx *Create_database_stmtContext) {}

// ExitCreate_database_stmt is called when production create_database_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_database_stmt(ctx *Create_database_stmtContext) {}

// EnterDomain_constraint is called when production domain_constraint is entered.
func (s *BasePostgreSQLParserListener) EnterDomain_constraint(ctx *Domain_constraintContext) {}

// ExitDomain_constraint is called when production domain_constraint is exited.
func (s *BasePostgreSQLParserListener) ExitDomain_constraint(ctx *Domain_constraintContext) {}

// EnterCreate_domain_stmt is called when production create_domain_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_domain_stmt(ctx *Create_domain_stmtContext) {}

// ExitCreate_domain_stmt is called when production create_domain_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_domain_stmt(ctx *Create_domain_stmtContext) {}

// EnterCreate_event_trigger_cond is called when production create_event_trigger_cond is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_event_trigger_cond(ctx *Create_event_trigger_condContext) {
}

// ExitCreate_event_trigger_cond is called when production create_event_trigger_cond is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_event_trigger_cond(ctx *Create_event_trigger_condContext) {
}

// EnterCreate_event_trigger_stmt is called when production create_event_trigger_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_event_trigger_stmt(ctx *Create_event_trigger_stmtContext) {
}

// ExitCreate_event_trigger_stmt is called when production create_event_trigger_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_event_trigger_stmt(ctx *Create_event_trigger_stmtContext) {
}

// EnterCreate_foreign_data_options is called when production create_foreign_data_options is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_foreign_data_options(ctx *Create_foreign_data_optionsContext) {
}

// ExitCreate_foreign_data_options is called when production create_foreign_data_options is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_foreign_data_options(ctx *Create_foreign_data_optionsContext) {
}

// EnterCreate_foreign_data_stmt is called when production create_foreign_data_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_foreign_data_stmt(ctx *Create_foreign_data_stmtContext) {
}

// ExitCreate_foreign_data_stmt is called when production create_foreign_data_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_foreign_data_stmt(ctx *Create_foreign_data_stmtContext) {
}

// EnterCreate_foreign_table_stmt is called when production create_foreign_table_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_foreign_table_stmt(ctx *Create_foreign_table_stmtContext) {
}

// ExitCreate_foreign_table_stmt is called when production create_foreign_table_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_foreign_table_stmt(ctx *Create_foreign_table_stmtContext) {
}

// EnterCreate_function_stmt is called when production create_function_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_function_stmt(ctx *Create_function_stmtContext) {}

// ExitCreate_function_stmt is called when production create_function_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_function_stmt(ctx *Create_function_stmtContext) {}

// EnterCreate_group_stmt is called when production create_group_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_group_stmt(ctx *Create_group_stmtContext) {}

// ExitCreate_group_stmt is called when production create_group_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_group_stmt(ctx *Create_group_stmtContext) {}

// EnterCreate_index_stmt is called when production create_index_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_index_stmt(ctx *Create_index_stmtContext) {}

// ExitCreate_index_stmt is called when production create_index_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_index_stmt(ctx *Create_index_stmtContext) {}

// EnterCreate_language_stmt is called when production create_language_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_language_stmt(ctx *Create_language_stmtContext) {}

// ExitCreate_language_stmt is called when production create_language_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_language_stmt(ctx *Create_language_stmtContext) {}

// EnterCreate_materialized_view_stmt is called when production create_materialized_view_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_materialized_view_stmt(ctx *Create_materialized_view_stmtContext) {
}

// ExitCreate_materialized_view_stmt is called when production create_materialized_view_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_materialized_view_stmt(ctx *Create_materialized_view_stmtContext) {
}

// EnterCreate_operator_stmt is called when production create_operator_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_operator_stmt(ctx *Create_operator_stmtContext) {}

// ExitCreate_operator_stmt is called when production create_operator_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_operator_stmt(ctx *Create_operator_stmtContext) {}

// EnterCreate_operator_class_opt is called when production create_operator_class_opt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_operator_class_opt(ctx *Create_operator_class_optContext) {
}

// ExitCreate_operator_class_opt is called when production create_operator_class_opt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_operator_class_opt(ctx *Create_operator_class_optContext) {
}

// EnterCreate_operator_class_stmt is called when production create_operator_class_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_operator_class_stmt(ctx *Create_operator_class_stmtContext) {
}

// ExitCreate_operator_class_stmt is called when production create_operator_class_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_operator_class_stmt(ctx *Create_operator_class_stmtContext) {
}

// EnterCreate_operator_family_stmt is called when production create_operator_family_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_operator_family_stmt(ctx *Create_operator_family_stmtContext) {
}

// ExitCreate_operator_family_stmt is called when production create_operator_family_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_operator_family_stmt(ctx *Create_operator_family_stmtContext) {
}

// EnterCreate_policy_stmt is called when production create_policy_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_policy_stmt(ctx *Create_policy_stmtContext) {}

// ExitCreate_policy_stmt is called when production create_policy_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_policy_stmt(ctx *Create_policy_stmtContext) {}

// EnterCreate_role_stmt is called when production create_role_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_role_stmt(ctx *Create_role_stmtContext) {}

// ExitCreate_role_stmt is called when production create_role_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_role_stmt(ctx *Create_role_stmtContext) {}

// EnterCreate_rule_event is called when production create_rule_event is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_rule_event(ctx *Create_rule_eventContext) {}

// ExitCreate_rule_event is called when production create_rule_event is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_rule_event(ctx *Create_rule_eventContext) {}

// EnterCreate_rule_stmt is called when production create_rule_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_rule_stmt(ctx *Create_rule_stmtContext) {}

// ExitCreate_rule_stmt is called when production create_rule_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_rule_stmt(ctx *Create_rule_stmtContext) {}

// EnterCreate_schema_stmt is called when production create_schema_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_schema_stmt(ctx *Create_schema_stmtContext) {}

// ExitCreate_schema_stmt is called when production create_schema_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_schema_stmt(ctx *Create_schema_stmtContext) {}

// EnterCreate_sequence_stmt is called when production create_sequence_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_sequence_stmt(ctx *Create_sequence_stmtContext) {}

// ExitCreate_sequence_stmt is called when production create_sequence_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_sequence_stmt(ctx *Create_sequence_stmtContext) {}

// EnterCreate_server_stmt is called when production create_server_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_server_stmt(ctx *Create_server_stmtContext) {}

// ExitCreate_server_stmt is called when production create_server_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_server_stmt(ctx *Create_server_stmtContext) {}

// EnterCreate_statistics_stmt is called when production create_statistics_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_statistics_stmt(ctx *Create_statistics_stmtContext) {
}

// ExitCreate_statistics_stmt is called when production create_statistics_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_statistics_stmt(ctx *Create_statistics_stmtContext) {
}

// EnterCreate_subscription_stmt is called when production create_subscription_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_subscription_stmt(ctx *Create_subscription_stmtContext) {
}

// ExitCreate_subscription_stmt is called when production create_subscription_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_subscription_stmt(ctx *Create_subscription_stmtContext) {
}

// EnterCreate_table_stmt is called when production create_table_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_table_stmt(ctx *Create_table_stmtContext) {}

// ExitCreate_table_stmt is called when production create_table_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_table_stmt(ctx *Create_table_stmtContext) {}

// EnterCreate_definitions is called when production create_definitions is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_definitions(ctx *Create_definitionsContext) {}

// ExitCreate_definitions is called when production create_definitions is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_definitions(ctx *Create_definitionsContext) {}

// EnterCreate_definition is called when production create_definition is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_definition(ctx *Create_definitionContext) {}

// ExitCreate_definition is called when production create_definition is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_definition(ctx *Create_definitionContext) {}

// EnterCreate_table_as_stmt is called when production create_table_as_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_table_as_stmt(ctx *Create_table_as_stmtContext) {}

// ExitCreate_table_as_stmt is called when production create_table_as_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_table_as_stmt(ctx *Create_table_as_stmtContext) {}

// EnterCreate_tablespace_stmt is called when production create_tablespace_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_tablespace_stmt(ctx *Create_tablespace_stmtContext) {
}

// ExitCreate_tablespace_stmt is called when production create_tablespace_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_tablespace_stmt(ctx *Create_tablespace_stmtContext) {
}

// EnterCreate_text_search_config_stmt is called when production create_text_search_config_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_text_search_config_stmt(ctx *Create_text_search_config_stmtContext) {
}

// ExitCreate_text_search_config_stmt is called when production create_text_search_config_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_text_search_config_stmt(ctx *Create_text_search_config_stmtContext) {
}

// EnterCreate_text_search_dict_stmt is called when production create_text_search_dict_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_text_search_dict_stmt(ctx *Create_text_search_dict_stmtContext) {
}

// ExitCreate_text_search_dict_stmt is called when production create_text_search_dict_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_text_search_dict_stmt(ctx *Create_text_search_dict_stmtContext) {
}

// EnterCreate_text_search_parser_stmt is called when production create_text_search_parser_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_text_search_parser_stmt(ctx *Create_text_search_parser_stmtContext) {
}

// ExitCreate_text_search_parser_stmt is called when production create_text_search_parser_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_text_search_parser_stmt(ctx *Create_text_search_parser_stmtContext) {
}

// EnterCreate_text_search_template_stmt is called when production create_text_search_template_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_text_search_template_stmt(ctx *Create_text_search_template_stmtContext) {
}

// ExitCreate_text_search_template_stmt is called when production create_text_search_template_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_text_search_template_stmt(ctx *Create_text_search_template_stmtContext) {
}

// EnterCreate_transform_stmt is called when production create_transform_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_transform_stmt(ctx *Create_transform_stmtContext) {
}

// ExitCreate_transform_stmt is called when production create_transform_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_transform_stmt(ctx *Create_transform_stmtContext) {}

// EnterCreate_trigger_stmt is called when production create_trigger_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_trigger_stmt(ctx *Create_trigger_stmtContext) {}

// ExitCreate_trigger_stmt is called when production create_trigger_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_trigger_stmt(ctx *Create_trigger_stmtContext) {}

// EnterCreate_type_stmt is called when production create_type_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_type_stmt(ctx *Create_type_stmtContext) {}

// ExitCreate_type_stmt is called when production create_type_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_type_stmt(ctx *Create_type_stmtContext) {}

// EnterCreate_user_stmt is called when production create_user_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_user_stmt(ctx *Create_user_stmtContext) {}

// ExitCreate_user_stmt is called when production create_user_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_user_stmt(ctx *Create_user_stmtContext) {}

// EnterCreate_user_mapping_stmt is called when production create_user_mapping_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_user_mapping_stmt(ctx *Create_user_mapping_stmtContext) {
}

// ExitCreate_user_mapping_stmt is called when production create_user_mapping_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_user_mapping_stmt(ctx *Create_user_mapping_stmtContext) {
}

// EnterCreate_view_stmt is called when production create_view_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterCreate_view_stmt(ctx *Create_view_stmtContext) {}

// ExitCreate_view_stmt is called when production create_view_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitCreate_view_stmt(ctx *Create_view_stmtContext) {}

// EnterDeallocate_stmt is called when production deallocate_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDeallocate_stmt(ctx *Deallocate_stmtContext) {}

// ExitDeallocate_stmt is called when production deallocate_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDeallocate_stmt(ctx *Deallocate_stmtContext) {}

// EnterDeclare_stmt is called when production declare_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDeclare_stmt(ctx *Declare_stmtContext) {}

// ExitDeclare_stmt is called when production declare_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDeclare_stmt(ctx *Declare_stmtContext) {}

// EnterDelete_stmt is called when production delete_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDelete_stmt(ctx *Delete_stmtContext) {}

// ExitDelete_stmt is called when production delete_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDelete_stmt(ctx *Delete_stmtContext) {}

// EnterDiscard_stmt is called when production discard_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDiscard_stmt(ctx *Discard_stmtContext) {}

// ExitDiscard_stmt is called when production discard_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDiscard_stmt(ctx *Discard_stmtContext) {}

// EnterDrop_stmt is called when production drop_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_stmt(ctx *Drop_stmtContext) {}

// ExitDrop_stmt is called when production drop_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_stmt(ctx *Drop_stmtContext) {}

// EnterDrop_access_method_stmt is called when production drop_access_method_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_access_method_stmt(ctx *Drop_access_method_stmtContext) {
}

// ExitDrop_access_method_stmt is called when production drop_access_method_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_access_method_stmt(ctx *Drop_access_method_stmtContext) {
}

// EnterDrop_aggregate_stmt is called when production drop_aggregate_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_aggregate_stmt(ctx *Drop_aggregate_stmtContext) {}

// ExitDrop_aggregate_stmt is called when production drop_aggregate_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_aggregate_stmt(ctx *Drop_aggregate_stmtContext) {}

// EnterDrop_cast_stmt is called when production drop_cast_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_cast_stmt(ctx *Drop_cast_stmtContext) {}

// ExitDrop_cast_stmt is called when production drop_cast_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_cast_stmt(ctx *Drop_cast_stmtContext) {}

// EnterDrop_collation_stmt is called when production drop_collation_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_collation_stmt(ctx *Drop_collation_stmtContext) {}

// ExitDrop_collation_stmt is called when production drop_collation_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_collation_stmt(ctx *Drop_collation_stmtContext) {}

// EnterDrop_conversion_stmt is called when production drop_conversion_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_conversion_stmt(ctx *Drop_conversion_stmtContext) {}

// ExitDrop_conversion_stmt is called when production drop_conversion_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_conversion_stmt(ctx *Drop_conversion_stmtContext) {}

// EnterDrop_database_stmt is called when production drop_database_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_database_stmt(ctx *Drop_database_stmtContext) {}

// ExitDrop_database_stmt is called when production drop_database_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_database_stmt(ctx *Drop_database_stmtContext) {}

// EnterDrop_domain_stmt is called when production drop_domain_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_domain_stmt(ctx *Drop_domain_stmtContext) {}

// ExitDrop_domain_stmt is called when production drop_domain_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_domain_stmt(ctx *Drop_domain_stmtContext) {}

// EnterDrop_event_trigger_stmt is called when production drop_event_trigger_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_event_trigger_stmt(ctx *Drop_event_trigger_stmtContext) {
}

// ExitDrop_event_trigger_stmt is called when production drop_event_trigger_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_event_trigger_stmt(ctx *Drop_event_trigger_stmtContext) {
}

// EnterDrop_extension_stmt is called when production drop_extension_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_extension_stmt(ctx *Drop_extension_stmtContext) {}

// ExitDrop_extension_stmt is called when production drop_extension_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_extension_stmt(ctx *Drop_extension_stmtContext) {}

// EnterDrop_foreign_data_wrapper_stmt is called when production drop_foreign_data_wrapper_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_foreign_data_wrapper_stmt(ctx *Drop_foreign_data_wrapper_stmtContext) {
}

// ExitDrop_foreign_data_wrapper_stmt is called when production drop_foreign_data_wrapper_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_foreign_data_wrapper_stmt(ctx *Drop_foreign_data_wrapper_stmtContext) {
}

// EnterDrop_foreign_table_stmt is called when production drop_foreign_table_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_foreign_table_stmt(ctx *Drop_foreign_table_stmtContext) {
}

// ExitDrop_foreign_table_stmt is called when production drop_foreign_table_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_foreign_table_stmt(ctx *Drop_foreign_table_stmtContext) {
}

// EnterDrop_function_stmt is called when production drop_function_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_function_stmt(ctx *Drop_function_stmtContext) {}

// ExitDrop_function_stmt is called when production drop_function_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_function_stmt(ctx *Drop_function_stmtContext) {}

// EnterDrop_group_stmt is called when production drop_group_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_group_stmt(ctx *Drop_group_stmtContext) {}

// ExitDrop_group_stmt is called when production drop_group_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_group_stmt(ctx *Drop_group_stmtContext) {}

// EnterDrop_index_stmt is called when production drop_index_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_index_stmt(ctx *Drop_index_stmtContext) {}

// ExitDrop_index_stmt is called when production drop_index_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_index_stmt(ctx *Drop_index_stmtContext) {}

// EnterDrop_language_stmt is called when production drop_language_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_language_stmt(ctx *Drop_language_stmtContext) {}

// ExitDrop_language_stmt is called when production drop_language_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_language_stmt(ctx *Drop_language_stmtContext) {}

// EnterDrop_materialized_view_stmt is called when production drop_materialized_view_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_materialized_view_stmt(ctx *Drop_materialized_view_stmtContext) {
}

// ExitDrop_materialized_view_stmt is called when production drop_materialized_view_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_materialized_view_stmt(ctx *Drop_materialized_view_stmtContext) {
}

// EnterDrop_operator_stmt is called when production drop_operator_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_operator_stmt(ctx *Drop_operator_stmtContext) {}

// ExitDrop_operator_stmt is called when production drop_operator_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_operator_stmt(ctx *Drop_operator_stmtContext) {}

// EnterDrop_operator_class_stmt is called when production drop_operator_class_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_operator_class_stmt(ctx *Drop_operator_class_stmtContext) {
}

// ExitDrop_operator_class_stmt is called when production drop_operator_class_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_operator_class_stmt(ctx *Drop_operator_class_stmtContext) {
}

// EnterDrop_operator_family_stmt is called when production drop_operator_family_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_operator_family_stmt(ctx *Drop_operator_family_stmtContext) {
}

// ExitDrop_operator_family_stmt is called when production drop_operator_family_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_operator_family_stmt(ctx *Drop_operator_family_stmtContext) {
}

// EnterDrop_owned_stmt is called when production drop_owned_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_owned_stmt(ctx *Drop_owned_stmtContext) {}

// ExitDrop_owned_stmt is called when production drop_owned_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_owned_stmt(ctx *Drop_owned_stmtContext) {}

// EnterDrop_policy_stmt is called when production drop_policy_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_policy_stmt(ctx *Drop_policy_stmtContext) {}

// ExitDrop_policy_stmt is called when production drop_policy_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_policy_stmt(ctx *Drop_policy_stmtContext) {}

// EnterDrop_publication_stmt is called when production drop_publication_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_publication_stmt(ctx *Drop_publication_stmtContext) {
}

// ExitDrop_publication_stmt is called when production drop_publication_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_publication_stmt(ctx *Drop_publication_stmtContext) {}

// EnterDrop_role_stmt is called when production drop_role_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_role_stmt(ctx *Drop_role_stmtContext) {}

// ExitDrop_role_stmt is called when production drop_role_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_role_stmt(ctx *Drop_role_stmtContext) {}

// EnterDrop_rule_stmt is called when production drop_rule_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_rule_stmt(ctx *Drop_rule_stmtContext) {}

// ExitDrop_rule_stmt is called when production drop_rule_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_rule_stmt(ctx *Drop_rule_stmtContext) {}

// EnterDrop_schema_stmt is called when production drop_schema_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_schema_stmt(ctx *Drop_schema_stmtContext) {}

// ExitDrop_schema_stmt is called when production drop_schema_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_schema_stmt(ctx *Drop_schema_stmtContext) {}

// EnterDrop_sequence_stmt is called when production drop_sequence_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_sequence_stmt(ctx *Drop_sequence_stmtContext) {}

// ExitDrop_sequence_stmt is called when production drop_sequence_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_sequence_stmt(ctx *Drop_sequence_stmtContext) {}

// EnterDrop_server_stmt is called when production drop_server_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_server_stmt(ctx *Drop_server_stmtContext) {}

// ExitDrop_server_stmt is called when production drop_server_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_server_stmt(ctx *Drop_server_stmtContext) {}

// EnterDrop_statistics_stmt is called when production drop_statistics_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_statistics_stmt(ctx *Drop_statistics_stmtContext) {}

// ExitDrop_statistics_stmt is called when production drop_statistics_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_statistics_stmt(ctx *Drop_statistics_stmtContext) {}

// EnterDrop_subscription_stmt is called when production drop_subscription_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_subscription_stmt(ctx *Drop_subscription_stmtContext) {
}

// ExitDrop_subscription_stmt is called when production drop_subscription_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_subscription_stmt(ctx *Drop_subscription_stmtContext) {
}

// EnterDrop_table_stmt is called when production drop_table_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_table_stmt(ctx *Drop_table_stmtContext) {}

// ExitDrop_table_stmt is called when production drop_table_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_table_stmt(ctx *Drop_table_stmtContext) {}

// EnterDrop_tablespace_stmt is called when production drop_tablespace_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_tablespace_stmt(ctx *Drop_tablespace_stmtContext) {}

// ExitDrop_tablespace_stmt is called when production drop_tablespace_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_tablespace_stmt(ctx *Drop_tablespace_stmtContext) {}

// EnterDrop_text_search_config_stmt is called when production drop_text_search_config_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_text_search_config_stmt(ctx *Drop_text_search_config_stmtContext) {
}

// ExitDrop_text_search_config_stmt is called when production drop_text_search_config_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_text_search_config_stmt(ctx *Drop_text_search_config_stmtContext) {
}

// EnterDrop_text_search_dict_stmt is called when production drop_text_search_dict_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_text_search_dict_stmt(ctx *Drop_text_search_dict_stmtContext) {
}

// ExitDrop_text_search_dict_stmt is called when production drop_text_search_dict_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_text_search_dict_stmt(ctx *Drop_text_search_dict_stmtContext) {
}

// EnterDrop_text_search_parser_stmt is called when production drop_text_search_parser_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_text_search_parser_stmt(ctx *Drop_text_search_parser_stmtContext) {
}

// ExitDrop_text_search_parser_stmt is called when production drop_text_search_parser_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_text_search_parser_stmt(ctx *Drop_text_search_parser_stmtContext) {
}

// EnterDrop_text_search_template_stmt is called when production drop_text_search_template_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_text_search_template_stmt(ctx *Drop_text_search_template_stmtContext) {
}

// ExitDrop_text_search_template_stmt is called when production drop_text_search_template_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_text_search_template_stmt(ctx *Drop_text_search_template_stmtContext) {
}

// EnterDrop_transform_stmt is called when production drop_transform_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_transform_stmt(ctx *Drop_transform_stmtContext) {}

// ExitDrop_transform_stmt is called when production drop_transform_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_transform_stmt(ctx *Drop_transform_stmtContext) {}

// EnterDrop_trigger_stmt is called when production drop_trigger_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_trigger_stmt(ctx *Drop_trigger_stmtContext) {}

// ExitDrop_trigger_stmt is called when production drop_trigger_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_trigger_stmt(ctx *Drop_trigger_stmtContext) {}

// EnterDrop_type_stmt is called when production drop_type_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_type_stmt(ctx *Drop_type_stmtContext) {}

// ExitDrop_type_stmt is called when production drop_type_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_type_stmt(ctx *Drop_type_stmtContext) {}

// EnterDrop_user_stmt is called when production drop_user_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_user_stmt(ctx *Drop_user_stmtContext) {}

// ExitDrop_user_stmt is called when production drop_user_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_user_stmt(ctx *Drop_user_stmtContext) {}

// EnterDrop_user_mapping_stmt is called when production drop_user_mapping_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_user_mapping_stmt(ctx *Drop_user_mapping_stmtContext) {
}

// ExitDrop_user_mapping_stmt is called when production drop_user_mapping_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_user_mapping_stmt(ctx *Drop_user_mapping_stmtContext) {
}

// EnterDrop_view_stmt is called when production drop_view_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterDrop_view_stmt(ctx *Drop_view_stmtContext) {}

// ExitDrop_view_stmt is called when production drop_view_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitDrop_view_stmt(ctx *Drop_view_stmtContext) {}

// EnterExecute_stmt is called when production execute_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterExecute_stmt(ctx *Execute_stmtContext) {}

// ExitExecute_stmt is called when production execute_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitExecute_stmt(ctx *Execute_stmtContext) {}

// EnterExplain_stmt is called when production explain_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterExplain_stmt(ctx *Explain_stmtContext) {}

// ExitExplain_stmt is called when production explain_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitExplain_stmt(ctx *Explain_stmtContext) {}

// EnterFetch_stmt is called when production fetch_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterFetch_stmt(ctx *Fetch_stmtContext) {}

// ExitFetch_stmt is called when production fetch_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitFetch_stmt(ctx *Fetch_stmtContext) {}

// EnterGrant_stmt is called when production grant_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterGrant_stmt(ctx *Grant_stmtContext) {}

// ExitGrant_stmt is called when production grant_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitGrant_stmt(ctx *Grant_stmtContext) {}

// EnterImport_foreign_schema_stmt is called when production import_foreign_schema_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterImport_foreign_schema_stmt(ctx *Import_foreign_schema_stmtContext) {
}

// ExitImport_foreign_schema_stmt is called when production import_foreign_schema_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitImport_foreign_schema_stmt(ctx *Import_foreign_schema_stmtContext) {
}

// EnterInsert_stmt is called when production insert_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterInsert_stmt(ctx *Insert_stmtContext) {}

// ExitInsert_stmt is called when production insert_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitInsert_stmt(ctx *Insert_stmtContext) {}

// EnterListen_stmt is called when production listen_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterListen_stmt(ctx *Listen_stmtContext) {}

// ExitListen_stmt is called when production listen_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitListen_stmt(ctx *Listen_stmtContext) {}

// EnterLoad_stmt is called when production load_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterLoad_stmt(ctx *Load_stmtContext) {}

// ExitLoad_stmt is called when production load_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitLoad_stmt(ctx *Load_stmtContext) {}

// EnterLock_stmt is called when production lock_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterLock_stmt(ctx *Lock_stmtContext) {}

// ExitLock_stmt is called when production lock_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitLock_stmt(ctx *Lock_stmtContext) {}

// EnterMove_stmt is called when production move_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterMove_stmt(ctx *Move_stmtContext) {}

// ExitMove_stmt is called when production move_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitMove_stmt(ctx *Move_stmtContext) {}

// EnterNotify_stmt is called when production notify_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterNotify_stmt(ctx *Notify_stmtContext) {}

// ExitNotify_stmt is called when production notify_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitNotify_stmt(ctx *Notify_stmtContext) {}

// EnterPrepare_stmt is called when production prepare_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterPrepare_stmt(ctx *Prepare_stmtContext) {}

// ExitPrepare_stmt is called when production prepare_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitPrepare_stmt(ctx *Prepare_stmtContext) {}

// EnterPrepare_transaction_stmt is called when production prepare_transaction_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterPrepare_transaction_stmt(ctx *Prepare_transaction_stmtContext) {
}

// ExitPrepare_transaction_stmt is called when production prepare_transaction_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitPrepare_transaction_stmt(ctx *Prepare_transaction_stmtContext) {
}

// EnterReassign_owned_stmt is called when production reassign_owned_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterReassign_owned_stmt(ctx *Reassign_owned_stmtContext) {}

// ExitReassign_owned_stmt is called when production reassign_owned_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitReassign_owned_stmt(ctx *Reassign_owned_stmtContext) {}

// EnterRefresh_materialized_view_stmt is called when production refresh_materialized_view_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterRefresh_materialized_view_stmt(ctx *Refresh_materialized_view_stmtContext) {
}

// ExitRefresh_materialized_view_stmt is called when production refresh_materialized_view_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitRefresh_materialized_view_stmt(ctx *Refresh_materialized_view_stmtContext) {
}

// EnterReindex_stmt is called when production reindex_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterReindex_stmt(ctx *Reindex_stmtContext) {}

// ExitReindex_stmt is called when production reindex_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitReindex_stmt(ctx *Reindex_stmtContext) {}

// EnterRelease_savepoint_stmt is called when production release_savepoint_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterRelease_savepoint_stmt(ctx *Release_savepoint_stmtContext) {
}

// ExitRelease_savepoint_stmt is called when production release_savepoint_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitRelease_savepoint_stmt(ctx *Release_savepoint_stmtContext) {
}

// EnterReset_stmt is called when production reset_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterReset_stmt(ctx *Reset_stmtContext) {}

// ExitReset_stmt is called when production reset_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitReset_stmt(ctx *Reset_stmtContext) {}

// EnterRevoke_stmt is called when production revoke_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterRevoke_stmt(ctx *Revoke_stmtContext) {}

// ExitRevoke_stmt is called when production revoke_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitRevoke_stmt(ctx *Revoke_stmtContext) {}

// EnterRollback_stmt is called when production rollback_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterRollback_stmt(ctx *Rollback_stmtContext) {}

// ExitRollback_stmt is called when production rollback_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitRollback_stmt(ctx *Rollback_stmtContext) {}

// EnterRollback_prepared_stmt is called when production rollback_prepared_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterRollback_prepared_stmt(ctx *Rollback_prepared_stmtContext) {
}

// ExitRollback_prepared_stmt is called when production rollback_prepared_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitRollback_prepared_stmt(ctx *Rollback_prepared_stmtContext) {
}

// EnterRollback_to_savepoint_stmt is called when production rollback_to_savepoint_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterRollback_to_savepoint_stmt(ctx *Rollback_to_savepoint_stmtContext) {
}

// ExitRollback_to_savepoint_stmt is called when production rollback_to_savepoint_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitRollback_to_savepoint_stmt(ctx *Rollback_to_savepoint_stmtContext) {
}

// EnterSavepoint_stmt is called when production savepoint_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSavepoint_stmt(ctx *Savepoint_stmtContext) {}

// ExitSavepoint_stmt is called when production savepoint_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSavepoint_stmt(ctx *Savepoint_stmtContext) {}

// EnterSecurity_label_stmt is called when production security_label_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSecurity_label_stmt(ctx *Security_label_stmtContext) {}

// ExitSecurity_label_stmt is called when production security_label_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSecurity_label_stmt(ctx *Security_label_stmtContext) {}

// EnterSelect_stmt is called when production select_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSelect_stmt(ctx *Select_stmtContext) {}

// ExitSelect_stmt is called when production select_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSelect_stmt(ctx *Select_stmtContext) {}

// EnterSelect_into_stmt is called when production select_into_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSelect_into_stmt(ctx *Select_into_stmtContext) {}

// ExitSelect_into_stmt is called when production select_into_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSelect_into_stmt(ctx *Select_into_stmtContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BasePostgreSQLParserListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BasePostgreSQLParserListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterWith_expr is called when production with_expr is entered.
func (s *BasePostgreSQLParserListener) EnterWith_expr(ctx *With_exprContext) {}

// ExitWith_expr is called when production with_expr is exited.
func (s *BasePostgreSQLParserListener) ExitWith_expr(ctx *With_exprContext) {}

// EnterSet_stmt is called when production set_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSet_stmt(ctx *Set_stmtContext) {}

// ExitSet_stmt is called when production set_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSet_stmt(ctx *Set_stmtContext) {}

// EnterSet_constraints_stmt is called when production set_constraints_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSet_constraints_stmt(ctx *Set_constraints_stmtContext) {}

// ExitSet_constraints_stmt is called when production set_constraints_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSet_constraints_stmt(ctx *Set_constraints_stmtContext) {}

// EnterSet_role_stmt is called when production set_role_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSet_role_stmt(ctx *Set_role_stmtContext) {}

// ExitSet_role_stmt is called when production set_role_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSet_role_stmt(ctx *Set_role_stmtContext) {}

// EnterSet_session_authorization_stmt is called when production set_session_authorization_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSet_session_authorization_stmt(ctx *Set_session_authorization_stmtContext) {
}

// ExitSet_session_authorization_stmt is called when production set_session_authorization_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSet_session_authorization_stmt(ctx *Set_session_authorization_stmtContext) {
}

// EnterTransaction_mode is called when production transaction_mode is entered.
func (s *BasePostgreSQLParserListener) EnterTransaction_mode(ctx *Transaction_modeContext) {}

// ExitTransaction_mode is called when production transaction_mode is exited.
func (s *BasePostgreSQLParserListener) ExitTransaction_mode(ctx *Transaction_modeContext) {}

// EnterTransaction_mode_list is called when production transaction_mode_list is entered.
func (s *BasePostgreSQLParserListener) EnterTransaction_mode_list(ctx *Transaction_mode_listContext) {
}

// ExitTransaction_mode_list is called when production transaction_mode_list is exited.
func (s *BasePostgreSQLParserListener) ExitTransaction_mode_list(ctx *Transaction_mode_listContext) {}

// EnterSet_transaction_stmt is called when production set_transaction_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterSet_transaction_stmt(ctx *Set_transaction_stmtContext) {}

// ExitSet_transaction_stmt is called when production set_transaction_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitSet_transaction_stmt(ctx *Set_transaction_stmtContext) {}

// EnterShow_stmt is called when production show_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterShow_stmt(ctx *Show_stmtContext) {}

// ExitShow_stmt is called when production show_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitShow_stmt(ctx *Show_stmtContext) {}

// EnterTruncate_stmt is called when production truncate_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterTruncate_stmt(ctx *Truncate_stmtContext) {}

// ExitTruncate_stmt is called when production truncate_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitTruncate_stmt(ctx *Truncate_stmtContext) {}

// EnterUnlisten_stmt is called when production unlisten_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterUnlisten_stmt(ctx *Unlisten_stmtContext) {}

// ExitUnlisten_stmt is called when production unlisten_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitUnlisten_stmt(ctx *Unlisten_stmtContext) {}

// EnterUpdate_stmt is called when production update_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterUpdate_stmt(ctx *Update_stmtContext) {}

// ExitUpdate_stmt is called when production update_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitUpdate_stmt(ctx *Update_stmtContext) {}

// EnterVacuum_opt is called when production vacuum_opt is entered.
func (s *BasePostgreSQLParserListener) EnterVacuum_opt(ctx *Vacuum_optContext) {}

// ExitVacuum_opt is called when production vacuum_opt is exited.
func (s *BasePostgreSQLParserListener) ExitVacuum_opt(ctx *Vacuum_optContext) {}

// EnterVacuum_opt_list is called when production vacuum_opt_list is entered.
func (s *BasePostgreSQLParserListener) EnterVacuum_opt_list(ctx *Vacuum_opt_listContext) {}

// ExitVacuum_opt_list is called when production vacuum_opt_list is exited.
func (s *BasePostgreSQLParserListener) ExitVacuum_opt_list(ctx *Vacuum_opt_listContext) {}

// EnterVacuum_stmt is called when production vacuum_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterVacuum_stmt(ctx *Vacuum_stmtContext) {}

// ExitVacuum_stmt is called when production vacuum_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitVacuum_stmt(ctx *Vacuum_stmtContext) {}

// EnterValues_stmt is called when production values_stmt is entered.
func (s *BasePostgreSQLParserListener) EnterValues_stmt(ctx *Values_stmtContext) {}

// ExitValues_stmt is called when production values_stmt is exited.
func (s *BasePostgreSQLParserListener) ExitValues_stmt(ctx *Values_stmtContext) {}

// EnterSelector_clause is called when production selector_clause is entered.
func (s *BasePostgreSQLParserListener) EnterSelector_clause(ctx *Selector_clauseContext) {}

// ExitSelector_clause is called when production selector_clause is exited.
func (s *BasePostgreSQLParserListener) ExitSelector_clause(ctx *Selector_clauseContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BasePostgreSQLParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BasePostgreSQLParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BasePostgreSQLParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BasePostgreSQLParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterGroup_by_clause is called when production group_by_clause is entered.
func (s *BasePostgreSQLParserListener) EnterGroup_by_clause(ctx *Group_by_clauseContext) {}

// ExitGroup_by_clause is called when production group_by_clause is exited.
func (s *BasePostgreSQLParserListener) ExitGroup_by_clause(ctx *Group_by_clauseContext) {}

// EnterGrouping_elem is called when production grouping_elem is entered.
func (s *BasePostgreSQLParserListener) EnterGrouping_elem(ctx *Grouping_elemContext) {}

// ExitGrouping_elem is called when production grouping_elem is exited.
func (s *BasePostgreSQLParserListener) ExitGrouping_elem(ctx *Grouping_elemContext) {}

// EnterGrouping_elem_list is called when production grouping_elem_list is entered.
func (s *BasePostgreSQLParserListener) EnterGrouping_elem_list(ctx *Grouping_elem_listContext) {}

// ExitGrouping_elem_list is called when production grouping_elem_list is exited.
func (s *BasePostgreSQLParserListener) ExitGrouping_elem_list(ctx *Grouping_elem_listContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BasePostgreSQLParserListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BasePostgreSQLParserListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterColumn_list is called when production column_list is entered.
func (s *BasePostgreSQLParserListener) EnterColumn_list(ctx *Column_listContext) {}

// ExitColumn_list is called when production column_list is exited.
func (s *BasePostgreSQLParserListener) ExitColumn_list(ctx *Column_listContext) {}

// EnterExplain_parameter is called when production explain_parameter is entered.
func (s *BasePostgreSQLParserListener) EnterExplain_parameter(ctx *Explain_parameterContext) {}

// ExitExplain_parameter is called when production explain_parameter is exited.
func (s *BasePostgreSQLParserListener) ExitExplain_parameter(ctx *Explain_parameterContext) {}

// EnterFrame is called when production frame is entered.
func (s *BasePostgreSQLParserListener) EnterFrame(ctx *FrameContext) {}

// ExitFrame is called when production frame is exited.
func (s *BasePostgreSQLParserListener) ExitFrame(ctx *FrameContext) {}

// EnterFrame_start is called when production frame_start is entered.
func (s *BasePostgreSQLParserListener) EnterFrame_start(ctx *Frame_startContext) {}

// ExitFrame_start is called when production frame_start is exited.
func (s *BasePostgreSQLParserListener) ExitFrame_start(ctx *Frame_startContext) {}

// EnterFrame_end is called when production frame_end is entered.
func (s *BasePostgreSQLParserListener) EnterFrame_end(ctx *Frame_endContext) {}

// ExitFrame_end is called when production frame_end is exited.
func (s *BasePostgreSQLParserListener) ExitFrame_end(ctx *Frame_endContext) {}

// EnterFrame_clause is called when production frame_clause is entered.
func (s *BasePostgreSQLParserListener) EnterFrame_clause(ctx *Frame_clauseContext) {}

// ExitFrame_clause is called when production frame_clause is exited.
func (s *BasePostgreSQLParserListener) ExitFrame_clause(ctx *Frame_clauseContext) {}

// EnterWindow_definition is called when production window_definition is entered.
func (s *BasePostgreSQLParserListener) EnterWindow_definition(ctx *Window_definitionContext) {}

// ExitWindow_definition is called when production window_definition is exited.
func (s *BasePostgreSQLParserListener) ExitWindow_definition(ctx *Window_definitionContext) {}

// EnterWindow_clause is called when production window_clause is entered.
func (s *BasePostgreSQLParserListener) EnterWindow_clause(ctx *Window_clauseContext) {}

// ExitWindow_clause is called when production window_clause is exited.
func (s *BasePostgreSQLParserListener) ExitWindow_clause(ctx *Window_clauseContext) {}

// EnterCombine_clause is called when production combine_clause is entered.
func (s *BasePostgreSQLParserListener) EnterCombine_clause(ctx *Combine_clauseContext) {}

// ExitCombine_clause is called when production combine_clause is exited.
func (s *BasePostgreSQLParserListener) ExitCombine_clause(ctx *Combine_clauseContext) {}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BasePostgreSQLParserListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BasePostgreSQLParserListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterOrder_by_item is called when production order_by_item is entered.
func (s *BasePostgreSQLParserListener) EnterOrder_by_item(ctx *Order_by_itemContext) {}

// ExitOrder_by_item is called when production order_by_item is exited.
func (s *BasePostgreSQLParserListener) ExitOrder_by_item(ctx *Order_by_itemContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BasePostgreSQLParserListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BasePostgreSQLParserListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterOffset_clause is called when production offset_clause is entered.
func (s *BasePostgreSQLParserListener) EnterOffset_clause(ctx *Offset_clauseContext) {}

// ExitOffset_clause is called when production offset_clause is exited.
func (s *BasePostgreSQLParserListener) ExitOffset_clause(ctx *Offset_clauseContext) {}

// EnterFetch_clause is called when production fetch_clause is entered.
func (s *BasePostgreSQLParserListener) EnterFetch_clause(ctx *Fetch_clauseContext) {}

// ExitFetch_clause is called when production fetch_clause is exited.
func (s *BasePostgreSQLParserListener) ExitFetch_clause(ctx *Fetch_clauseContext) {}

// EnterFor_clause is called when production for_clause is entered.
func (s *BasePostgreSQLParserListener) EnterFor_clause(ctx *For_clauseContext) {}

// ExitFor_clause is called when production for_clause is exited.
func (s *BasePostgreSQLParserListener) ExitFor_clause(ctx *For_clauseContext) {}

// EnterUpdater_clause is called when production updater_clause is entered.
func (s *BasePostgreSQLParserListener) EnterUpdater_clause(ctx *Updater_clauseContext) {}

// ExitUpdater_clause is called when production updater_clause is exited.
func (s *BasePostgreSQLParserListener) ExitUpdater_clause(ctx *Updater_clauseContext) {}

// EnterUpdater_expr is called when production updater_expr is entered.
func (s *BasePostgreSQLParserListener) EnterUpdater_expr(ctx *Updater_exprContext) {}

// ExitUpdater_expr is called when production updater_expr is exited.
func (s *BasePostgreSQLParserListener) ExitUpdater_expr(ctx *Updater_exprContext) {}

// EnterReturning_clause is called when production returning_clause is entered.
func (s *BasePostgreSQLParserListener) EnterReturning_clause(ctx *Returning_clauseContext) {}

// ExitReturning_clause is called when production returning_clause is exited.
func (s *BasePostgreSQLParserListener) ExitReturning_clause(ctx *Returning_clauseContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePostgreSQLParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePostgreSQLParserListener) ExitExpr(ctx *ExprContext) {}

// EnterBool_expr is called when production bool_expr is entered.
func (s *BasePostgreSQLParserListener) EnterBool_expr(ctx *Bool_exprContext) {}

// ExitBool_expr is called when production bool_expr is exited.
func (s *BasePostgreSQLParserListener) ExitBool_expr(ctx *Bool_exprContext) {}

// EnterCase_expr is called when production case_expr is entered.
func (s *BasePostgreSQLParserListener) EnterCase_expr(ctx *Case_exprContext) {}

// ExitCase_expr is called when production case_expr is exited.
func (s *BasePostgreSQLParserListener) ExitCase_expr(ctx *Case_exprContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BasePostgreSQLParserListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BasePostgreSQLParserListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterExpr_list_list is called when production expr_list_list is entered.
func (s *BasePostgreSQLParserListener) EnterExpr_list_list(ctx *Expr_list_listContext) {}

// ExitExpr_list_list is called when production expr_list_list is exited.
func (s *BasePostgreSQLParserListener) ExitExpr_list_list(ctx *Expr_list_listContext) {}

// EnterFunc_sig_arg is called when production func_sig_arg is entered.
func (s *BasePostgreSQLParserListener) EnterFunc_sig_arg(ctx *Func_sig_argContext) {}

// ExitFunc_sig_arg is called when production func_sig_arg is exited.
func (s *BasePostgreSQLParserListener) ExitFunc_sig_arg(ctx *Func_sig_argContext) {}

// EnterFunc_sig_arg_list is called when production func_sig_arg_list is entered.
func (s *BasePostgreSQLParserListener) EnterFunc_sig_arg_list(ctx *Func_sig_arg_listContext) {}

// ExitFunc_sig_arg_list is called when production func_sig_arg_list is exited.
func (s *BasePostgreSQLParserListener) ExitFunc_sig_arg_list(ctx *Func_sig_arg_listContext) {}

// EnterFunc_sig is called when production func_sig is entered.
func (s *BasePostgreSQLParserListener) EnterFunc_sig(ctx *Func_sigContext) {}

// ExitFunc_sig is called when production func_sig is exited.
func (s *BasePostgreSQLParserListener) ExitFunc_sig(ctx *Func_sigContext) {}

// EnterFunc_sig_list is called when production func_sig_list is entered.
func (s *BasePostgreSQLParserListener) EnterFunc_sig_list(ctx *Func_sig_listContext) {}

// ExitFunc_sig_list is called when production func_sig_list is exited.
func (s *BasePostgreSQLParserListener) ExitFunc_sig_list(ctx *Func_sig_listContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BasePostgreSQLParserListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BasePostgreSQLParserListener) ExitType_name(ctx *Type_nameContext) {}

// EnterTimezone is called when production timezone is entered.
func (s *BasePostgreSQLParserListener) EnterTimezone(ctx *TimezoneContext) {}

// ExitTimezone is called when production timezone is exited.
func (s *BasePostgreSQLParserListener) ExitTimezone(ctx *TimezoneContext) {}

// EnterOper is called when production oper is entered.
func (s *BasePostgreSQLParserListener) EnterOper(ctx *OperContext) {}

// ExitOper is called when production oper is exited.
func (s *BasePostgreSQLParserListener) ExitOper(ctx *OperContext) {}

// EnterAggregate is called when production aggregate is entered.
func (s *BasePostgreSQLParserListener) EnterAggregate(ctx *AggregateContext) {}

// ExitAggregate is called when production aggregate is exited.
func (s *BasePostgreSQLParserListener) ExitAggregate(ctx *AggregateContext) {}

// EnterName_ is called when production name_ is entered.
func (s *BasePostgreSQLParserListener) EnterName_(ctx *Name_Context) {}

// ExitName_ is called when production name_ is exited.
func (s *BasePostgreSQLParserListener) ExitName_(ctx *Name_Context) {}

// EnterName_list is called when production name_list is entered.
func (s *BasePostgreSQLParserListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BasePostgreSQLParserListener) ExitName_list(ctx *Name_listContext) {}

// EnterIdentifier_list is called when production identifier_list is entered.
func (s *BasePostgreSQLParserListener) EnterIdentifier_list(ctx *Identifier_listContext) {}

// ExitIdentifier_list is called when production identifier_list is exited.
func (s *BasePostgreSQLParserListener) ExitIdentifier_list(ctx *Identifier_listContext) {}

// EnterOption_expr is called when production option_expr is entered.
func (s *BasePostgreSQLParserListener) EnterOption_expr(ctx *Option_exprContext) {}

// ExitOption_expr is called when production option_expr is exited.
func (s *BasePostgreSQLParserListener) ExitOption_expr(ctx *Option_exprContext) {}

// EnterOption_list is called when production option_list is entered.
func (s *BasePostgreSQLParserListener) EnterOption_list(ctx *Option_listContext) {}

// ExitOption_list is called when production option_list is exited.
func (s *BasePostgreSQLParserListener) ExitOption_list(ctx *Option_listContext) {}

// EnterTable_name_ is called when production table_name_ is entered.
func (s *BasePostgreSQLParserListener) EnterTable_name_(ctx *Table_name_Context) {}

// ExitTable_name_ is called when production table_name_ is exited.
func (s *BasePostgreSQLParserListener) ExitTable_name_(ctx *Table_name_Context) {}

// EnterData_type is called when production data_type is entered.
func (s *BasePostgreSQLParserListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BasePostgreSQLParserListener) ExitData_type(ctx *Data_typeContext) {}

// EnterData_type_list is called when production data_type_list is entered.
func (s *BasePostgreSQLParserListener) EnterData_type_list(ctx *Data_type_listContext) {}

// ExitData_type_list is called when production data_type_list is exited.
func (s *BasePostgreSQLParserListener) ExitData_type_list(ctx *Data_type_listContext) {}

// EnterIndex_method is called when production index_method is entered.
func (s *BasePostgreSQLParserListener) EnterIndex_method(ctx *Index_methodContext) {}

// ExitIndex_method is called when production index_method is exited.
func (s *BasePostgreSQLParserListener) ExitIndex_method(ctx *Index_methodContext) {}

// EnterFunc_name is called when production func_name is entered.
func (s *BasePostgreSQLParserListener) EnterFunc_name(ctx *Func_nameContext) {}

// ExitFunc_name is called when production func_name is exited.
func (s *BasePostgreSQLParserListener) ExitFunc_name(ctx *Func_nameContext) {}

// EnterFunc_call is called when production func_call is entered.
func (s *BasePostgreSQLParserListener) EnterFunc_call(ctx *Func_callContext) {}

// ExitFunc_call is called when production func_call is exited.
func (s *BasePostgreSQLParserListener) ExitFunc_call(ctx *Func_callContext) {}

// EnterArray_cons_expr is called when production array_cons_expr is entered.
func (s *BasePostgreSQLParserListener) EnterArray_cons_expr(ctx *Array_cons_exprContext) {}

// ExitArray_cons_expr is called when production array_cons_expr is exited.
func (s *BasePostgreSQLParserListener) ExitArray_cons_expr(ctx *Array_cons_exprContext) {}

// EnterFrom_item is called when production from_item is entered.
func (s *BasePostgreSQLParserListener) EnterFrom_item(ctx *From_itemContext) {}

// ExitFrom_item is called when production from_item is exited.
func (s *BasePostgreSQLParserListener) ExitFrom_item(ctx *From_itemContext) {}

// EnterWith_column_alias is called when production with_column_alias is entered.
func (s *BasePostgreSQLParserListener) EnterWith_column_alias(ctx *With_column_aliasContext) {}

// ExitWith_column_alias is called when production with_column_alias is exited.
func (s *BasePostgreSQLParserListener) ExitWith_column_alias(ctx *With_column_aliasContext) {}

// EnterJoin_type is called when production join_type is entered.
func (s *BasePostgreSQLParserListener) EnterJoin_type(ctx *Join_typeContext) {}

// ExitJoin_type is called when production join_type is exited.
func (s *BasePostgreSQLParserListener) ExitJoin_type(ctx *Join_typeContext) {}

// EnterJoin_clause is called when production join_clause is entered.
func (s *BasePostgreSQLParserListener) EnterJoin_clause(ctx *Join_clauseContext) {}

// ExitJoin_clause is called when production join_clause is exited.
func (s *BasePostgreSQLParserListener) ExitJoin_clause(ctx *Join_clauseContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BasePostgreSQLParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BasePostgreSQLParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterAggregate_signature is called when production aggregate_signature is entered.
func (s *BasePostgreSQLParserListener) EnterAggregate_signature(ctx *Aggregate_signatureContext) {}

// ExitAggregate_signature is called when production aggregate_signature is exited.
func (s *BasePostgreSQLParserListener) ExitAggregate_signature(ctx *Aggregate_signatureContext) {}

// EnterColumn_constraint is called when production column_constraint is entered.
func (s *BasePostgreSQLParserListener) EnterColumn_constraint(ctx *Column_constraintContext) {}

// ExitColumn_constraint is called when production column_constraint is exited.
func (s *BasePostgreSQLParserListener) ExitColumn_constraint(ctx *Column_constraintContext) {}

// EnterColumn_constraints is called when production column_constraints is entered.
func (s *BasePostgreSQLParserListener) EnterColumn_constraints(ctx *Column_constraintsContext) {}

// ExitColumn_constraints is called when production column_constraints is exited.
func (s *BasePostgreSQLParserListener) ExitColumn_constraints(ctx *Column_constraintsContext) {}

// EnterIndex_parameters is called when production index_parameters is entered.
func (s *BasePostgreSQLParserListener) EnterIndex_parameters(ctx *Index_parametersContext) {}

// ExitIndex_parameters is called when production index_parameters is exited.
func (s *BasePostgreSQLParserListener) ExitIndex_parameters(ctx *Index_parametersContext) {}

// EnterExclude_element is called when production exclude_element is entered.
func (s *BasePostgreSQLParserListener) EnterExclude_element(ctx *Exclude_elementContext) {}

// ExitExclude_element is called when production exclude_element is exited.
func (s *BasePostgreSQLParserListener) ExitExclude_element(ctx *Exclude_elementContext) {}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *BasePostgreSQLParserListener) EnterTable_constraint(ctx *Table_constraintContext) {}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *BasePostgreSQLParserListener) ExitTable_constraint(ctx *Table_constraintContext) {}

// EnterRole_name is called when production role_name is entered.
func (s *BasePostgreSQLParserListener) EnterRole_name(ctx *Role_nameContext) {}

// ExitRole_name is called when production role_name is exited.
func (s *BasePostgreSQLParserListener) ExitRole_name(ctx *Role_nameContext) {}

// EnterRole_name_list is called when production role_name_list is entered.
func (s *BasePostgreSQLParserListener) EnterRole_name_list(ctx *Role_name_listContext) {}

// ExitRole_name_list is called when production role_name_list is exited.
func (s *BasePostgreSQLParserListener) ExitRole_name_list(ctx *Role_name_listContext) {}

// EnterParam_value is called when production param_value is entered.
func (s *BasePostgreSQLParserListener) EnterParam_value(ctx *Param_valueContext) {}

// ExitParam_value is called when production param_value is exited.
func (s *BasePostgreSQLParserListener) ExitParam_value(ctx *Param_valueContext) {}

// EnterNon_reserved_keyword is called when production non_reserved_keyword is entered.
func (s *BasePostgreSQLParserListener) EnterNon_reserved_keyword(ctx *Non_reserved_keywordContext) {}

// ExitNon_reserved_keyword is called when production non_reserved_keyword is exited.
func (s *BasePostgreSQLParserListener) ExitNon_reserved_keyword(ctx *Non_reserved_keywordContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BasePostgreSQLParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BasePostgreSQLParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterTodo_fill_in is called when production todo_fill_in is entered.
func (s *BasePostgreSQLParserListener) EnterTodo_fill_in(ctx *Todo_fill_inContext) {}

// ExitTodo_fill_in is called when production todo_fill_in is exited.
func (s *BasePostgreSQLParserListener) ExitTodo_fill_in(ctx *Todo_fill_inContext) {}

// EnterTodo_implement is called when production todo_implement is entered.
func (s *BasePostgreSQLParserListener) EnterTodo_implement(ctx *Todo_implementContext) {}

// ExitTodo_implement is called when production todo_implement is exited.
func (s *BasePostgreSQLParserListener) ExitTodo_implement(ctx *Todo_implementContext) {}

// EnterCorrelation_name is called when production correlation_name is entered.
func (s *BasePostgreSQLParserListener) EnterCorrelation_name(ctx *Correlation_nameContext) {}

// ExitCorrelation_name is called when production correlation_name is exited.
func (s *BasePostgreSQLParserListener) ExitCorrelation_name(ctx *Correlation_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BasePostgreSQLParserListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BasePostgreSQLParserListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BasePostgreSQLParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BasePostgreSQLParserListener) ExitAlias(ctx *AliasContext) {}

// EnterColumn_alias is called when production column_alias is entered.
func (s *BasePostgreSQLParserListener) EnterColumn_alias(ctx *Column_aliasContext) {}

// ExitColumn_alias is called when production column_alias is exited.
func (s *BasePostgreSQLParserListener) ExitColumn_alias(ctx *Column_aliasContext) {}

// EnterColumn_definition is called when production column_definition is entered.
func (s *BasePostgreSQLParserListener) EnterColumn_definition(ctx *Column_definitionContext) {}

// ExitColumn_definition is called when production column_definition is exited.
func (s *BasePostgreSQLParserListener) ExitColumn_definition(ctx *Column_definitionContext) {}

// EnterWindow_name is called when production window_name is entered.
func (s *BasePostgreSQLParserListener) EnterWindow_name(ctx *Window_nameContext) {}

// ExitWindow_name is called when production window_name is exited.
func (s *BasePostgreSQLParserListener) ExitWindow_name(ctx *Window_nameContext) {}
