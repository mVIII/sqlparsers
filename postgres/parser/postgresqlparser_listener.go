// Code generated from /home/kovacs/Projects/github.com/mVIII/sqlparsers/postgres/PostgreSQLParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PostgreSQLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PostgreSQLParserListener is a complete listener for a parse tree produced by PostgreSQLParser.
type PostgreSQLParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAbort_stmt is called when entering the abort_stmt production.
	EnterAbort_stmt(c *Abort_stmtContext)

	// EnterAlter_stmt is called when entering the alter_stmt production.
	EnterAlter_stmt(c *Alter_stmtContext)

	// EnterAlter_aggregate_stmt is called when entering the alter_aggregate_stmt production.
	EnterAlter_aggregate_stmt(c *Alter_aggregate_stmtContext)

	// EnterAlter_collation_stmt is called when entering the alter_collation_stmt production.
	EnterAlter_collation_stmt(c *Alter_collation_stmtContext)

	// EnterAlter_conversion_stmt is called when entering the alter_conversion_stmt production.
	EnterAlter_conversion_stmt(c *Alter_conversion_stmtContext)

	// EnterAlter_database_stmt is called when entering the alter_database_stmt production.
	EnterAlter_database_stmt(c *Alter_database_stmtContext)

	// EnterAlter_default_privileges_stmt is called when entering the alter_default_privileges_stmt production.
	EnterAlter_default_privileges_stmt(c *Alter_default_privileges_stmtContext)

	// EnterAlter_domain_stmt is called when entering the alter_domain_stmt production.
	EnterAlter_domain_stmt(c *Alter_domain_stmtContext)

	// EnterAlter_event_trigger_stmt is called when entering the alter_event_trigger_stmt production.
	EnterAlter_event_trigger_stmt(c *Alter_event_trigger_stmtContext)

	// EnterAlter_extension_stmt is called when entering the alter_extension_stmt production.
	EnterAlter_extension_stmt(c *Alter_extension_stmtContext)

	// EnterAlter_foreign_data_wrapper_stmt is called when entering the alter_foreign_data_wrapper_stmt production.
	EnterAlter_foreign_data_wrapper_stmt(c *Alter_foreign_data_wrapper_stmtContext)

	// EnterAlter_foreign_table_action is called when entering the alter_foreign_table_action production.
	EnterAlter_foreign_table_action(c *Alter_foreign_table_actionContext)

	// EnterAlter_foreign_table_action_list is called when entering the alter_foreign_table_action_list production.
	EnterAlter_foreign_table_action_list(c *Alter_foreign_table_action_listContext)

	// EnterAlter_foreign_table_stmt is called when entering the alter_foreign_table_stmt production.
	EnterAlter_foreign_table_stmt(c *Alter_foreign_table_stmtContext)

	// EnterAlter_function_stmt is called when entering the alter_function_stmt production.
	EnterAlter_function_stmt(c *Alter_function_stmtContext)

	// EnterAlter_group_stmt is called when entering the alter_group_stmt production.
	EnterAlter_group_stmt(c *Alter_group_stmtContext)

	// EnterAlter_index_stmt is called when entering the alter_index_stmt production.
	EnterAlter_index_stmt(c *Alter_index_stmtContext)

	// EnterAlter_language_stmt is called when entering the alter_language_stmt production.
	EnterAlter_language_stmt(c *Alter_language_stmtContext)

	// EnterAlter_large_object_stmt is called when entering the alter_large_object_stmt production.
	EnterAlter_large_object_stmt(c *Alter_large_object_stmtContext)

	// EnterAlter_materialize_view_stmt is called when entering the alter_materialize_view_stmt production.
	EnterAlter_materialize_view_stmt(c *Alter_materialize_view_stmtContext)

	// EnterAlter_operator_stmt is called when entering the alter_operator_stmt production.
	EnterAlter_operator_stmt(c *Alter_operator_stmtContext)

	// EnterAlter_operator_class_stmt is called when entering the alter_operator_class_stmt production.
	EnterAlter_operator_class_stmt(c *Alter_operator_class_stmtContext)

	// EnterAlter_operator_family_stmt is called when entering the alter_operator_family_stmt production.
	EnterAlter_operator_family_stmt(c *Alter_operator_family_stmtContext)

	// EnterAlter_policy_stmt is called when entering the alter_policy_stmt production.
	EnterAlter_policy_stmt(c *Alter_policy_stmtContext)

	// EnterAlter_publication_stmt is called when entering the alter_publication_stmt production.
	EnterAlter_publication_stmt(c *Alter_publication_stmtContext)

	// EnterAlter_role_options is called when entering the alter_role_options production.
	EnterAlter_role_options(c *Alter_role_optionsContext)

	// EnterAlter_role_stmt is called when entering the alter_role_stmt production.
	EnterAlter_role_stmt(c *Alter_role_stmtContext)

	// EnterAlter_rule_stmt is called when entering the alter_rule_stmt production.
	EnterAlter_rule_stmt(c *Alter_rule_stmtContext)

	// EnterAlter_schema_stmt is called when entering the alter_schema_stmt production.
	EnterAlter_schema_stmt(c *Alter_schema_stmtContext)

	// EnterAlter_sequence_stmt is called when entering the alter_sequence_stmt production.
	EnterAlter_sequence_stmt(c *Alter_sequence_stmtContext)

	// EnterAlter_server_options_list is called when entering the alter_server_options_list production.
	EnterAlter_server_options_list(c *Alter_server_options_listContext)

	// EnterAlter_server_stmt is called when entering the alter_server_stmt production.
	EnterAlter_server_stmt(c *Alter_server_stmtContext)

	// EnterAlter_statistics_stmt is called when entering the alter_statistics_stmt production.
	EnterAlter_statistics_stmt(c *Alter_statistics_stmtContext)

	// EnterAlter_subscription_stmt is called when entering the alter_subscription_stmt production.
	EnterAlter_subscription_stmt(c *Alter_subscription_stmtContext)

	// EnterAlter_system_stmt is called when entering the alter_system_stmt production.
	EnterAlter_system_stmt(c *Alter_system_stmtContext)

	// EnterAlter_table_stmt is called when entering the alter_table_stmt production.
	EnterAlter_table_stmt(c *Alter_table_stmtContext)

	// EnterAlter_tablespace_stmt is called when entering the alter_tablespace_stmt production.
	EnterAlter_tablespace_stmt(c *Alter_tablespace_stmtContext)

	// EnterAlter_text_search_config_stmt is called when entering the alter_text_search_config_stmt production.
	EnterAlter_text_search_config_stmt(c *Alter_text_search_config_stmtContext)

	// EnterAlter_text_search_dict_stmt is called when entering the alter_text_search_dict_stmt production.
	EnterAlter_text_search_dict_stmt(c *Alter_text_search_dict_stmtContext)

	// EnterAlter_text_search_parser_stmt is called when entering the alter_text_search_parser_stmt production.
	EnterAlter_text_search_parser_stmt(c *Alter_text_search_parser_stmtContext)

	// EnterAlter_text_search_template_stmt is called when entering the alter_text_search_template_stmt production.
	EnterAlter_text_search_template_stmt(c *Alter_text_search_template_stmtContext)

	// EnterAlter_trigger_stmt is called when entering the alter_trigger_stmt production.
	EnterAlter_trigger_stmt(c *Alter_trigger_stmtContext)

	// EnterAlter_type_stmt is called when entering the alter_type_stmt production.
	EnterAlter_type_stmt(c *Alter_type_stmtContext)

	// EnterAlter_user_stmt is called when entering the alter_user_stmt production.
	EnterAlter_user_stmt(c *Alter_user_stmtContext)

	// EnterAlter_user_mapping_stmt is called when entering the alter_user_mapping_stmt production.
	EnterAlter_user_mapping_stmt(c *Alter_user_mapping_stmtContext)

	// EnterAlter_view_stmt is called when entering the alter_view_stmt production.
	EnterAlter_view_stmt(c *Alter_view_stmtContext)

	// EnterAnalyze_stmt is called when entering the analyze_stmt production.
	EnterAnalyze_stmt(c *Analyze_stmtContext)

	// EnterClose_stmt is called when entering the close_stmt production.
	EnterClose_stmt(c *Close_stmtContext)

	// EnterCluster_stmt is called when entering the cluster_stmt production.
	EnterCluster_stmt(c *Cluster_stmtContext)

	// EnterComment_stmt is called when entering the comment_stmt production.
	EnterComment_stmt(c *Comment_stmtContext)

	// EnterCommit_stmt is called when entering the commit_stmt production.
	EnterCommit_stmt(c *Commit_stmtContext)

	// EnterCommit_prepared_stmt is called when entering the commit_prepared_stmt production.
	EnterCommit_prepared_stmt(c *Commit_prepared_stmtContext)

	// EnterCopy_stmt is called when entering the copy_stmt production.
	EnterCopy_stmt(c *Copy_stmtContext)

	// EnterCreate_stmt is called when entering the create_stmt production.
	EnterCreate_stmt(c *Create_stmtContext)

	// EnterCreate_access_method_stmt is called when entering the create_access_method_stmt production.
	EnterCreate_access_method_stmt(c *Create_access_method_stmtContext)

	// EnterCreate_aggregate_stmt is called when entering the create_aggregate_stmt production.
	EnterCreate_aggregate_stmt(c *Create_aggregate_stmtContext)

	// EnterCreate_cast_stmt is called when entering the create_cast_stmt production.
	EnterCreate_cast_stmt(c *Create_cast_stmtContext)

	// EnterCreate_collation_opt is called when entering the create_collation_opt production.
	EnterCreate_collation_opt(c *Create_collation_optContext)

	// EnterCreate_collation_opt_list is called when entering the create_collation_opt_list production.
	EnterCreate_collation_opt_list(c *Create_collation_opt_listContext)

	// EnterCreate_collation_stmt is called when entering the create_collation_stmt production.
	EnterCreate_collation_stmt(c *Create_collation_stmtContext)

	// EnterCreate_conversion_stmt is called when entering the create_conversion_stmt production.
	EnterCreate_conversion_stmt(c *Create_conversion_stmtContext)

	// EnterCreate_database_stmt is called when entering the create_database_stmt production.
	EnterCreate_database_stmt(c *Create_database_stmtContext)

	// EnterDomain_constraint is called when entering the domain_constraint production.
	EnterDomain_constraint(c *Domain_constraintContext)

	// EnterCreate_domain_stmt is called when entering the create_domain_stmt production.
	EnterCreate_domain_stmt(c *Create_domain_stmtContext)

	// EnterCreate_event_trigger_cond is called when entering the create_event_trigger_cond production.
	EnterCreate_event_trigger_cond(c *Create_event_trigger_condContext)

	// EnterCreate_event_trigger_stmt is called when entering the create_event_trigger_stmt production.
	EnterCreate_event_trigger_stmt(c *Create_event_trigger_stmtContext)

	// EnterCreate_foreign_data_options is called when entering the create_foreign_data_options production.
	EnterCreate_foreign_data_options(c *Create_foreign_data_optionsContext)

	// EnterCreate_foreign_data_stmt is called when entering the create_foreign_data_stmt production.
	EnterCreate_foreign_data_stmt(c *Create_foreign_data_stmtContext)

	// EnterCreate_foreign_table_stmt is called when entering the create_foreign_table_stmt production.
	EnterCreate_foreign_table_stmt(c *Create_foreign_table_stmtContext)

	// EnterCreate_function_stmt is called when entering the create_function_stmt production.
	EnterCreate_function_stmt(c *Create_function_stmtContext)

	// EnterCreate_group_stmt is called when entering the create_group_stmt production.
	EnterCreate_group_stmt(c *Create_group_stmtContext)

	// EnterCreate_index_stmt is called when entering the create_index_stmt production.
	EnterCreate_index_stmt(c *Create_index_stmtContext)

	// EnterCreate_language_stmt is called when entering the create_language_stmt production.
	EnterCreate_language_stmt(c *Create_language_stmtContext)

	// EnterCreate_materialized_view_stmt is called when entering the create_materialized_view_stmt production.
	EnterCreate_materialized_view_stmt(c *Create_materialized_view_stmtContext)

	// EnterCreate_operator_stmt is called when entering the create_operator_stmt production.
	EnterCreate_operator_stmt(c *Create_operator_stmtContext)

	// EnterCreate_operator_class_opt is called when entering the create_operator_class_opt production.
	EnterCreate_operator_class_opt(c *Create_operator_class_optContext)

	// EnterCreate_operator_class_stmt is called when entering the create_operator_class_stmt production.
	EnterCreate_operator_class_stmt(c *Create_operator_class_stmtContext)

	// EnterCreate_operator_family_stmt is called when entering the create_operator_family_stmt production.
	EnterCreate_operator_family_stmt(c *Create_operator_family_stmtContext)

	// EnterCreate_policy_stmt is called when entering the create_policy_stmt production.
	EnterCreate_policy_stmt(c *Create_policy_stmtContext)

	// EnterCreate_role_stmt is called when entering the create_role_stmt production.
	EnterCreate_role_stmt(c *Create_role_stmtContext)

	// EnterCreate_rule_event is called when entering the create_rule_event production.
	EnterCreate_rule_event(c *Create_rule_eventContext)

	// EnterCreate_rule_stmt is called when entering the create_rule_stmt production.
	EnterCreate_rule_stmt(c *Create_rule_stmtContext)

	// EnterCreate_schema_stmt is called when entering the create_schema_stmt production.
	EnterCreate_schema_stmt(c *Create_schema_stmtContext)

	// EnterCreate_sequence_stmt is called when entering the create_sequence_stmt production.
	EnterCreate_sequence_stmt(c *Create_sequence_stmtContext)

	// EnterCreate_server_stmt is called when entering the create_server_stmt production.
	EnterCreate_server_stmt(c *Create_server_stmtContext)

	// EnterCreate_statistics_stmt is called when entering the create_statistics_stmt production.
	EnterCreate_statistics_stmt(c *Create_statistics_stmtContext)

	// EnterCreate_subscription_stmt is called when entering the create_subscription_stmt production.
	EnterCreate_subscription_stmt(c *Create_subscription_stmtContext)

	// EnterCreate_table_stmt is called when entering the create_table_stmt production.
	EnterCreate_table_stmt(c *Create_table_stmtContext)

	// EnterCreate_definitions is called when entering the create_definitions production.
	EnterCreate_definitions(c *Create_definitionsContext)

	// EnterCreate_definition is called when entering the create_definition production.
	EnterCreate_definition(c *Create_definitionContext)

	// EnterCreate_table_as_stmt is called when entering the create_table_as_stmt production.
	EnterCreate_table_as_stmt(c *Create_table_as_stmtContext)

	// EnterCreate_tablespace_stmt is called when entering the create_tablespace_stmt production.
	EnterCreate_tablespace_stmt(c *Create_tablespace_stmtContext)

	// EnterCreate_text_search_config_stmt is called when entering the create_text_search_config_stmt production.
	EnterCreate_text_search_config_stmt(c *Create_text_search_config_stmtContext)

	// EnterCreate_text_search_dict_stmt is called when entering the create_text_search_dict_stmt production.
	EnterCreate_text_search_dict_stmt(c *Create_text_search_dict_stmtContext)

	// EnterCreate_text_search_parser_stmt is called when entering the create_text_search_parser_stmt production.
	EnterCreate_text_search_parser_stmt(c *Create_text_search_parser_stmtContext)

	// EnterCreate_text_search_template_stmt is called when entering the create_text_search_template_stmt production.
	EnterCreate_text_search_template_stmt(c *Create_text_search_template_stmtContext)

	// EnterCreate_transform_stmt is called when entering the create_transform_stmt production.
	EnterCreate_transform_stmt(c *Create_transform_stmtContext)

	// EnterCreate_trigger_stmt is called when entering the create_trigger_stmt production.
	EnterCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// EnterCreate_type_stmt is called when entering the create_type_stmt production.
	EnterCreate_type_stmt(c *Create_type_stmtContext)

	// EnterCreate_user_stmt is called when entering the create_user_stmt production.
	EnterCreate_user_stmt(c *Create_user_stmtContext)

	// EnterCreate_user_mapping_stmt is called when entering the create_user_mapping_stmt production.
	EnterCreate_user_mapping_stmt(c *Create_user_mapping_stmtContext)

	// EnterCreate_view_stmt is called when entering the create_view_stmt production.
	EnterCreate_view_stmt(c *Create_view_stmtContext)

	// EnterDeallocate_stmt is called when entering the deallocate_stmt production.
	EnterDeallocate_stmt(c *Deallocate_stmtContext)

	// EnterDeclare_stmt is called when entering the declare_stmt production.
	EnterDeclare_stmt(c *Declare_stmtContext)

	// EnterDelete_stmt is called when entering the delete_stmt production.
	EnterDelete_stmt(c *Delete_stmtContext)

	// EnterDiscard_stmt is called when entering the discard_stmt production.
	EnterDiscard_stmt(c *Discard_stmtContext)

	// EnterDrop_stmt is called when entering the drop_stmt production.
	EnterDrop_stmt(c *Drop_stmtContext)

	// EnterDrop_access_method_stmt is called when entering the drop_access_method_stmt production.
	EnterDrop_access_method_stmt(c *Drop_access_method_stmtContext)

	// EnterDrop_aggregate_stmt is called when entering the drop_aggregate_stmt production.
	EnterDrop_aggregate_stmt(c *Drop_aggregate_stmtContext)

	// EnterDrop_cast_stmt is called when entering the drop_cast_stmt production.
	EnterDrop_cast_stmt(c *Drop_cast_stmtContext)

	// EnterDrop_collation_stmt is called when entering the drop_collation_stmt production.
	EnterDrop_collation_stmt(c *Drop_collation_stmtContext)

	// EnterDrop_conversion_stmt is called when entering the drop_conversion_stmt production.
	EnterDrop_conversion_stmt(c *Drop_conversion_stmtContext)

	// EnterDrop_database_stmt is called when entering the drop_database_stmt production.
	EnterDrop_database_stmt(c *Drop_database_stmtContext)

	// EnterDrop_domain_stmt is called when entering the drop_domain_stmt production.
	EnterDrop_domain_stmt(c *Drop_domain_stmtContext)

	// EnterDrop_event_trigger_stmt is called when entering the drop_event_trigger_stmt production.
	EnterDrop_event_trigger_stmt(c *Drop_event_trigger_stmtContext)

	// EnterDrop_extension_stmt is called when entering the drop_extension_stmt production.
	EnterDrop_extension_stmt(c *Drop_extension_stmtContext)

	// EnterDrop_foreign_data_wrapper_stmt is called when entering the drop_foreign_data_wrapper_stmt production.
	EnterDrop_foreign_data_wrapper_stmt(c *Drop_foreign_data_wrapper_stmtContext)

	// EnterDrop_foreign_table_stmt is called when entering the drop_foreign_table_stmt production.
	EnterDrop_foreign_table_stmt(c *Drop_foreign_table_stmtContext)

	// EnterDrop_function_stmt is called when entering the drop_function_stmt production.
	EnterDrop_function_stmt(c *Drop_function_stmtContext)

	// EnterDrop_group_stmt is called when entering the drop_group_stmt production.
	EnterDrop_group_stmt(c *Drop_group_stmtContext)

	// EnterDrop_index_stmt is called when entering the drop_index_stmt production.
	EnterDrop_index_stmt(c *Drop_index_stmtContext)

	// EnterDrop_language_stmt is called when entering the drop_language_stmt production.
	EnterDrop_language_stmt(c *Drop_language_stmtContext)

	// EnterDrop_materialized_view_stmt is called when entering the drop_materialized_view_stmt production.
	EnterDrop_materialized_view_stmt(c *Drop_materialized_view_stmtContext)

	// EnterDrop_operator_stmt is called when entering the drop_operator_stmt production.
	EnterDrop_operator_stmt(c *Drop_operator_stmtContext)

	// EnterDrop_operator_class_stmt is called when entering the drop_operator_class_stmt production.
	EnterDrop_operator_class_stmt(c *Drop_operator_class_stmtContext)

	// EnterDrop_operator_family_stmt is called when entering the drop_operator_family_stmt production.
	EnterDrop_operator_family_stmt(c *Drop_operator_family_stmtContext)

	// EnterDrop_owned_stmt is called when entering the drop_owned_stmt production.
	EnterDrop_owned_stmt(c *Drop_owned_stmtContext)

	// EnterDrop_policy_stmt is called when entering the drop_policy_stmt production.
	EnterDrop_policy_stmt(c *Drop_policy_stmtContext)

	// EnterDrop_publication_stmt is called when entering the drop_publication_stmt production.
	EnterDrop_publication_stmt(c *Drop_publication_stmtContext)

	// EnterDrop_role_stmt is called when entering the drop_role_stmt production.
	EnterDrop_role_stmt(c *Drop_role_stmtContext)

	// EnterDrop_rule_stmt is called when entering the drop_rule_stmt production.
	EnterDrop_rule_stmt(c *Drop_rule_stmtContext)

	// EnterDrop_schema_stmt is called when entering the drop_schema_stmt production.
	EnterDrop_schema_stmt(c *Drop_schema_stmtContext)

	// EnterDrop_sequence_stmt is called when entering the drop_sequence_stmt production.
	EnterDrop_sequence_stmt(c *Drop_sequence_stmtContext)

	// EnterDrop_server_stmt is called when entering the drop_server_stmt production.
	EnterDrop_server_stmt(c *Drop_server_stmtContext)

	// EnterDrop_statistics_stmt is called when entering the drop_statistics_stmt production.
	EnterDrop_statistics_stmt(c *Drop_statistics_stmtContext)

	// EnterDrop_subscription_stmt is called when entering the drop_subscription_stmt production.
	EnterDrop_subscription_stmt(c *Drop_subscription_stmtContext)

	// EnterDrop_table_stmt is called when entering the drop_table_stmt production.
	EnterDrop_table_stmt(c *Drop_table_stmtContext)

	// EnterDrop_tablespace_stmt is called when entering the drop_tablespace_stmt production.
	EnterDrop_tablespace_stmt(c *Drop_tablespace_stmtContext)

	// EnterDrop_text_search_config_stmt is called when entering the drop_text_search_config_stmt production.
	EnterDrop_text_search_config_stmt(c *Drop_text_search_config_stmtContext)

	// EnterDrop_text_search_dict_stmt is called when entering the drop_text_search_dict_stmt production.
	EnterDrop_text_search_dict_stmt(c *Drop_text_search_dict_stmtContext)

	// EnterDrop_text_search_parser_stmt is called when entering the drop_text_search_parser_stmt production.
	EnterDrop_text_search_parser_stmt(c *Drop_text_search_parser_stmtContext)

	// EnterDrop_text_search_template_stmt is called when entering the drop_text_search_template_stmt production.
	EnterDrop_text_search_template_stmt(c *Drop_text_search_template_stmtContext)

	// EnterDrop_transform_stmt is called when entering the drop_transform_stmt production.
	EnterDrop_transform_stmt(c *Drop_transform_stmtContext)

	// EnterDrop_trigger_stmt is called when entering the drop_trigger_stmt production.
	EnterDrop_trigger_stmt(c *Drop_trigger_stmtContext)

	// EnterDrop_type_stmt is called when entering the drop_type_stmt production.
	EnterDrop_type_stmt(c *Drop_type_stmtContext)

	// EnterDrop_user_stmt is called when entering the drop_user_stmt production.
	EnterDrop_user_stmt(c *Drop_user_stmtContext)

	// EnterDrop_user_mapping_stmt is called when entering the drop_user_mapping_stmt production.
	EnterDrop_user_mapping_stmt(c *Drop_user_mapping_stmtContext)

	// EnterDrop_view_stmt is called when entering the drop_view_stmt production.
	EnterDrop_view_stmt(c *Drop_view_stmtContext)

	// EnterExecute_stmt is called when entering the execute_stmt production.
	EnterExecute_stmt(c *Execute_stmtContext)

	// EnterExplain_stmt is called when entering the explain_stmt production.
	EnterExplain_stmt(c *Explain_stmtContext)

	// EnterFetch_stmt is called when entering the fetch_stmt production.
	EnterFetch_stmt(c *Fetch_stmtContext)

	// EnterGrant_stmt is called when entering the grant_stmt production.
	EnterGrant_stmt(c *Grant_stmtContext)

	// EnterImport_foreign_schema_stmt is called when entering the import_foreign_schema_stmt production.
	EnterImport_foreign_schema_stmt(c *Import_foreign_schema_stmtContext)

	// EnterInsert_stmt is called when entering the insert_stmt production.
	EnterInsert_stmt(c *Insert_stmtContext)

	// EnterListen_stmt is called when entering the listen_stmt production.
	EnterListen_stmt(c *Listen_stmtContext)

	// EnterLoad_stmt is called when entering the load_stmt production.
	EnterLoad_stmt(c *Load_stmtContext)

	// EnterLock_stmt is called when entering the lock_stmt production.
	EnterLock_stmt(c *Lock_stmtContext)

	// EnterMove_stmt is called when entering the move_stmt production.
	EnterMove_stmt(c *Move_stmtContext)

	// EnterNotify_stmt is called when entering the notify_stmt production.
	EnterNotify_stmt(c *Notify_stmtContext)

	// EnterPrepare_stmt is called when entering the prepare_stmt production.
	EnterPrepare_stmt(c *Prepare_stmtContext)

	// EnterPrepare_transaction_stmt is called when entering the prepare_transaction_stmt production.
	EnterPrepare_transaction_stmt(c *Prepare_transaction_stmtContext)

	// EnterReassign_owned_stmt is called when entering the reassign_owned_stmt production.
	EnterReassign_owned_stmt(c *Reassign_owned_stmtContext)

	// EnterRefresh_materialized_view_stmt is called when entering the refresh_materialized_view_stmt production.
	EnterRefresh_materialized_view_stmt(c *Refresh_materialized_view_stmtContext)

	// EnterReindex_stmt is called when entering the reindex_stmt production.
	EnterReindex_stmt(c *Reindex_stmtContext)

	// EnterRelease_savepoint_stmt is called when entering the release_savepoint_stmt production.
	EnterRelease_savepoint_stmt(c *Release_savepoint_stmtContext)

	// EnterReset_stmt is called when entering the reset_stmt production.
	EnterReset_stmt(c *Reset_stmtContext)

	// EnterRevoke_stmt is called when entering the revoke_stmt production.
	EnterRevoke_stmt(c *Revoke_stmtContext)

	// EnterRollback_stmt is called when entering the rollback_stmt production.
	EnterRollback_stmt(c *Rollback_stmtContext)

	// EnterRollback_prepared_stmt is called when entering the rollback_prepared_stmt production.
	EnterRollback_prepared_stmt(c *Rollback_prepared_stmtContext)

	// EnterRollback_to_savepoint_stmt is called when entering the rollback_to_savepoint_stmt production.
	EnterRollback_to_savepoint_stmt(c *Rollback_to_savepoint_stmtContext)

	// EnterSavepoint_stmt is called when entering the savepoint_stmt production.
	EnterSavepoint_stmt(c *Savepoint_stmtContext)

	// EnterSecurity_label_stmt is called when entering the security_label_stmt production.
	EnterSecurity_label_stmt(c *Security_label_stmtContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterSelect_into_stmt is called when entering the select_into_stmt production.
	EnterSelect_into_stmt(c *Select_into_stmtContext)

	// EnterWith_clause is called when entering the with_clause production.
	EnterWith_clause(c *With_clauseContext)

	// EnterWith_expr is called when entering the with_expr production.
	EnterWith_expr(c *With_exprContext)

	// EnterSet_stmt is called when entering the set_stmt production.
	EnterSet_stmt(c *Set_stmtContext)

	// EnterSet_constraints_stmt is called when entering the set_constraints_stmt production.
	EnterSet_constraints_stmt(c *Set_constraints_stmtContext)

	// EnterSet_role_stmt is called when entering the set_role_stmt production.
	EnterSet_role_stmt(c *Set_role_stmtContext)

	// EnterSet_session_authorization_stmt is called when entering the set_session_authorization_stmt production.
	EnterSet_session_authorization_stmt(c *Set_session_authorization_stmtContext)

	// EnterTransaction_mode is called when entering the transaction_mode production.
	EnterTransaction_mode(c *Transaction_modeContext)

	// EnterTransaction_mode_list is called when entering the transaction_mode_list production.
	EnterTransaction_mode_list(c *Transaction_mode_listContext)

	// EnterSet_transaction_stmt is called when entering the set_transaction_stmt production.
	EnterSet_transaction_stmt(c *Set_transaction_stmtContext)

	// EnterShow_stmt is called when entering the show_stmt production.
	EnterShow_stmt(c *Show_stmtContext)

	// EnterTruncate_stmt is called when entering the truncate_stmt production.
	EnterTruncate_stmt(c *Truncate_stmtContext)

	// EnterUnlisten_stmt is called when entering the unlisten_stmt production.
	EnterUnlisten_stmt(c *Unlisten_stmtContext)

	// EnterUpdate_stmt is called when entering the update_stmt production.
	EnterUpdate_stmt(c *Update_stmtContext)

	// EnterVacuum_opt is called when entering the vacuum_opt production.
	EnterVacuum_opt(c *Vacuum_optContext)

	// EnterVacuum_opt_list is called when entering the vacuum_opt_list production.
	EnterVacuum_opt_list(c *Vacuum_opt_listContext)

	// EnterVacuum_stmt is called when entering the vacuum_stmt production.
	EnterVacuum_stmt(c *Vacuum_stmtContext)

	// EnterValues_stmt is called when entering the values_stmt production.
	EnterValues_stmt(c *Values_stmtContext)

	// EnterSelector_clause is called when entering the selector_clause production.
	EnterSelector_clause(c *Selector_clauseContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterGroup_by_clause is called when entering the group_by_clause production.
	EnterGroup_by_clause(c *Group_by_clauseContext)

	// EnterGrouping_elem is called when entering the grouping_elem production.
	EnterGrouping_elem(c *Grouping_elemContext)

	// EnterGrouping_elem_list is called when entering the grouping_elem_list production.
	EnterGrouping_elem_list(c *Grouping_elem_listContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterColumn_list is called when entering the column_list production.
	EnterColumn_list(c *Column_listContext)

	// EnterExplain_parameter is called when entering the explain_parameter production.
	EnterExplain_parameter(c *Explain_parameterContext)

	// EnterFrame is called when entering the frame production.
	EnterFrame(c *FrameContext)

	// EnterFrame_start is called when entering the frame_start production.
	EnterFrame_start(c *Frame_startContext)

	// EnterFrame_end is called when entering the frame_end production.
	EnterFrame_end(c *Frame_endContext)

	// EnterFrame_clause is called when entering the frame_clause production.
	EnterFrame_clause(c *Frame_clauseContext)

	// EnterWindow_definition is called when entering the window_definition production.
	EnterWindow_definition(c *Window_definitionContext)

	// EnterWindow_clause is called when entering the window_clause production.
	EnterWindow_clause(c *Window_clauseContext)

	// EnterCombine_clause is called when entering the combine_clause production.
	EnterCombine_clause(c *Combine_clauseContext)

	// EnterOrder_by_clause is called when entering the order_by_clause production.
	EnterOrder_by_clause(c *Order_by_clauseContext)

	// EnterOrder_by_item is called when entering the order_by_item production.
	EnterOrder_by_item(c *Order_by_itemContext)

	// EnterLimit_clause is called when entering the limit_clause production.
	EnterLimit_clause(c *Limit_clauseContext)

	// EnterOffset_clause is called when entering the offset_clause production.
	EnterOffset_clause(c *Offset_clauseContext)

	// EnterFetch_clause is called when entering the fetch_clause production.
	EnterFetch_clause(c *Fetch_clauseContext)

	// EnterFor_clause is called when entering the for_clause production.
	EnterFor_clause(c *For_clauseContext)

	// EnterUpdater_clause is called when entering the updater_clause production.
	EnterUpdater_clause(c *Updater_clauseContext)

	// EnterUpdater_expr is called when entering the updater_expr production.
	EnterUpdater_expr(c *Updater_exprContext)

	// EnterReturning_clause is called when entering the returning_clause production.
	EnterReturning_clause(c *Returning_clauseContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterBool_expr is called when entering the bool_expr production.
	EnterBool_expr(c *Bool_exprContext)

	// EnterCase_expr is called when entering the case_expr production.
	EnterCase_expr(c *Case_exprContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterExpr_list_list is called when entering the expr_list_list production.
	EnterExpr_list_list(c *Expr_list_listContext)

	// EnterFunc_sig_arg is called when entering the func_sig_arg production.
	EnterFunc_sig_arg(c *Func_sig_argContext)

	// EnterFunc_sig_arg_list is called when entering the func_sig_arg_list production.
	EnterFunc_sig_arg_list(c *Func_sig_arg_listContext)

	// EnterFunc_sig is called when entering the func_sig production.
	EnterFunc_sig(c *Func_sigContext)

	// EnterFunc_sig_list is called when entering the func_sig_list production.
	EnterFunc_sig_list(c *Func_sig_listContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterTimezone is called when entering the timezone production.
	EnterTimezone(c *TimezoneContext)

	// EnterOper is called when entering the oper production.
	EnterOper(c *OperContext)

	// EnterAggregate is called when entering the aggregate production.
	EnterAggregate(c *AggregateContext)

	// EnterName_ is called when entering the name_ production.
	EnterName_(c *Name_Context)

	// EnterName_list is called when entering the name_list production.
	EnterName_list(c *Name_listContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterOption_expr is called when entering the option_expr production.
	EnterOption_expr(c *Option_exprContext)

	// EnterOption_list is called when entering the option_list production.
	EnterOption_list(c *Option_listContext)

	// EnterTable_name_ is called when entering the table_name_ production.
	EnterTable_name_(c *Table_name_Context)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterData_type_list is called when entering the data_type_list production.
	EnterData_type_list(c *Data_type_listContext)

	// EnterIndex_method is called when entering the index_method production.
	EnterIndex_method(c *Index_methodContext)

	// EnterFunc_name is called when entering the func_name production.
	EnterFunc_name(c *Func_nameContext)

	// EnterFunc_call is called when entering the func_call production.
	EnterFunc_call(c *Func_callContext)

	// EnterArray_cons_expr is called when entering the array_cons_expr production.
	EnterArray_cons_expr(c *Array_cons_exprContext)

	// EnterFrom_item is called when entering the from_item production.
	EnterFrom_item(c *From_itemContext)

	// EnterWith_column_alias is called when entering the with_column_alias production.
	EnterWith_column_alias(c *With_column_aliasContext)

	// EnterJoin_type is called when entering the join_type production.
	EnterJoin_type(c *Join_typeContext)

	// EnterJoin_clause is called when entering the join_clause production.
	EnterJoin_clause(c *Join_clauseContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterAggregate_signature is called when entering the aggregate_signature production.
	EnterAggregate_signature(c *Aggregate_signatureContext)

	// EnterColumn_constraint is called when entering the column_constraint production.
	EnterColumn_constraint(c *Column_constraintContext)

	// EnterColumn_constraints is called when entering the column_constraints production.
	EnterColumn_constraints(c *Column_constraintsContext)

	// EnterIndex_parameters is called when entering the index_parameters production.
	EnterIndex_parameters(c *Index_parametersContext)

	// EnterExclude_element is called when entering the exclude_element production.
	EnterExclude_element(c *Exclude_elementContext)

	// EnterTable_constraint is called when entering the table_constraint production.
	EnterTable_constraint(c *Table_constraintContext)

	// EnterRole_name is called when entering the role_name production.
	EnterRole_name(c *Role_nameContext)

	// EnterRole_name_list is called when entering the role_name_list production.
	EnterRole_name_list(c *Role_name_listContext)

	// EnterParam_value is called when entering the param_value production.
	EnterParam_value(c *Param_valueContext)

	// EnterNon_reserved_keyword is called when entering the non_reserved_keyword production.
	EnterNon_reserved_keyword(c *Non_reserved_keywordContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterTodo_fill_in is called when entering the todo_fill_in production.
	EnterTodo_fill_in(c *Todo_fill_inContext)

	// EnterTodo_implement is called when entering the todo_implement production.
	EnterTodo_implement(c *Todo_implementContext)

	// EnterCorrelation_name is called when entering the correlation_name production.
	EnterCorrelation_name(c *Correlation_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterColumn_alias is called when entering the column_alias production.
	EnterColumn_alias(c *Column_aliasContext)

	// EnterColumn_definition is called when entering the column_definition production.
	EnterColumn_definition(c *Column_definitionContext)

	// EnterWindow_name is called when entering the window_name production.
	EnterWindow_name(c *Window_nameContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAbort_stmt is called when exiting the abort_stmt production.
	ExitAbort_stmt(c *Abort_stmtContext)

	// ExitAlter_stmt is called when exiting the alter_stmt production.
	ExitAlter_stmt(c *Alter_stmtContext)

	// ExitAlter_aggregate_stmt is called when exiting the alter_aggregate_stmt production.
	ExitAlter_aggregate_stmt(c *Alter_aggregate_stmtContext)

	// ExitAlter_collation_stmt is called when exiting the alter_collation_stmt production.
	ExitAlter_collation_stmt(c *Alter_collation_stmtContext)

	// ExitAlter_conversion_stmt is called when exiting the alter_conversion_stmt production.
	ExitAlter_conversion_stmt(c *Alter_conversion_stmtContext)

	// ExitAlter_database_stmt is called when exiting the alter_database_stmt production.
	ExitAlter_database_stmt(c *Alter_database_stmtContext)

	// ExitAlter_default_privileges_stmt is called when exiting the alter_default_privileges_stmt production.
	ExitAlter_default_privileges_stmt(c *Alter_default_privileges_stmtContext)

	// ExitAlter_domain_stmt is called when exiting the alter_domain_stmt production.
	ExitAlter_domain_stmt(c *Alter_domain_stmtContext)

	// ExitAlter_event_trigger_stmt is called when exiting the alter_event_trigger_stmt production.
	ExitAlter_event_trigger_stmt(c *Alter_event_trigger_stmtContext)

	// ExitAlter_extension_stmt is called when exiting the alter_extension_stmt production.
	ExitAlter_extension_stmt(c *Alter_extension_stmtContext)

	// ExitAlter_foreign_data_wrapper_stmt is called when exiting the alter_foreign_data_wrapper_stmt production.
	ExitAlter_foreign_data_wrapper_stmt(c *Alter_foreign_data_wrapper_stmtContext)

	// ExitAlter_foreign_table_action is called when exiting the alter_foreign_table_action production.
	ExitAlter_foreign_table_action(c *Alter_foreign_table_actionContext)

	// ExitAlter_foreign_table_action_list is called when exiting the alter_foreign_table_action_list production.
	ExitAlter_foreign_table_action_list(c *Alter_foreign_table_action_listContext)

	// ExitAlter_foreign_table_stmt is called when exiting the alter_foreign_table_stmt production.
	ExitAlter_foreign_table_stmt(c *Alter_foreign_table_stmtContext)

	// ExitAlter_function_stmt is called when exiting the alter_function_stmt production.
	ExitAlter_function_stmt(c *Alter_function_stmtContext)

	// ExitAlter_group_stmt is called when exiting the alter_group_stmt production.
	ExitAlter_group_stmt(c *Alter_group_stmtContext)

	// ExitAlter_index_stmt is called when exiting the alter_index_stmt production.
	ExitAlter_index_stmt(c *Alter_index_stmtContext)

	// ExitAlter_language_stmt is called when exiting the alter_language_stmt production.
	ExitAlter_language_stmt(c *Alter_language_stmtContext)

	// ExitAlter_large_object_stmt is called when exiting the alter_large_object_stmt production.
	ExitAlter_large_object_stmt(c *Alter_large_object_stmtContext)

	// ExitAlter_materialize_view_stmt is called when exiting the alter_materialize_view_stmt production.
	ExitAlter_materialize_view_stmt(c *Alter_materialize_view_stmtContext)

	// ExitAlter_operator_stmt is called when exiting the alter_operator_stmt production.
	ExitAlter_operator_stmt(c *Alter_operator_stmtContext)

	// ExitAlter_operator_class_stmt is called when exiting the alter_operator_class_stmt production.
	ExitAlter_operator_class_stmt(c *Alter_operator_class_stmtContext)

	// ExitAlter_operator_family_stmt is called when exiting the alter_operator_family_stmt production.
	ExitAlter_operator_family_stmt(c *Alter_operator_family_stmtContext)

	// ExitAlter_policy_stmt is called when exiting the alter_policy_stmt production.
	ExitAlter_policy_stmt(c *Alter_policy_stmtContext)

	// ExitAlter_publication_stmt is called when exiting the alter_publication_stmt production.
	ExitAlter_publication_stmt(c *Alter_publication_stmtContext)

	// ExitAlter_role_options is called when exiting the alter_role_options production.
	ExitAlter_role_options(c *Alter_role_optionsContext)

	// ExitAlter_role_stmt is called when exiting the alter_role_stmt production.
	ExitAlter_role_stmt(c *Alter_role_stmtContext)

	// ExitAlter_rule_stmt is called when exiting the alter_rule_stmt production.
	ExitAlter_rule_stmt(c *Alter_rule_stmtContext)

	// ExitAlter_schema_stmt is called when exiting the alter_schema_stmt production.
	ExitAlter_schema_stmt(c *Alter_schema_stmtContext)

	// ExitAlter_sequence_stmt is called when exiting the alter_sequence_stmt production.
	ExitAlter_sequence_stmt(c *Alter_sequence_stmtContext)

	// ExitAlter_server_options_list is called when exiting the alter_server_options_list production.
	ExitAlter_server_options_list(c *Alter_server_options_listContext)

	// ExitAlter_server_stmt is called when exiting the alter_server_stmt production.
	ExitAlter_server_stmt(c *Alter_server_stmtContext)

	// ExitAlter_statistics_stmt is called when exiting the alter_statistics_stmt production.
	ExitAlter_statistics_stmt(c *Alter_statistics_stmtContext)

	// ExitAlter_subscription_stmt is called when exiting the alter_subscription_stmt production.
	ExitAlter_subscription_stmt(c *Alter_subscription_stmtContext)

	// ExitAlter_system_stmt is called when exiting the alter_system_stmt production.
	ExitAlter_system_stmt(c *Alter_system_stmtContext)

	// ExitAlter_table_stmt is called when exiting the alter_table_stmt production.
	ExitAlter_table_stmt(c *Alter_table_stmtContext)

	// ExitAlter_tablespace_stmt is called when exiting the alter_tablespace_stmt production.
	ExitAlter_tablespace_stmt(c *Alter_tablespace_stmtContext)

	// ExitAlter_text_search_config_stmt is called when exiting the alter_text_search_config_stmt production.
	ExitAlter_text_search_config_stmt(c *Alter_text_search_config_stmtContext)

	// ExitAlter_text_search_dict_stmt is called when exiting the alter_text_search_dict_stmt production.
	ExitAlter_text_search_dict_stmt(c *Alter_text_search_dict_stmtContext)

	// ExitAlter_text_search_parser_stmt is called when exiting the alter_text_search_parser_stmt production.
	ExitAlter_text_search_parser_stmt(c *Alter_text_search_parser_stmtContext)

	// ExitAlter_text_search_template_stmt is called when exiting the alter_text_search_template_stmt production.
	ExitAlter_text_search_template_stmt(c *Alter_text_search_template_stmtContext)

	// ExitAlter_trigger_stmt is called when exiting the alter_trigger_stmt production.
	ExitAlter_trigger_stmt(c *Alter_trigger_stmtContext)

	// ExitAlter_type_stmt is called when exiting the alter_type_stmt production.
	ExitAlter_type_stmt(c *Alter_type_stmtContext)

	// ExitAlter_user_stmt is called when exiting the alter_user_stmt production.
	ExitAlter_user_stmt(c *Alter_user_stmtContext)

	// ExitAlter_user_mapping_stmt is called when exiting the alter_user_mapping_stmt production.
	ExitAlter_user_mapping_stmt(c *Alter_user_mapping_stmtContext)

	// ExitAlter_view_stmt is called when exiting the alter_view_stmt production.
	ExitAlter_view_stmt(c *Alter_view_stmtContext)

	// ExitAnalyze_stmt is called when exiting the analyze_stmt production.
	ExitAnalyze_stmt(c *Analyze_stmtContext)

	// ExitClose_stmt is called when exiting the close_stmt production.
	ExitClose_stmt(c *Close_stmtContext)

	// ExitCluster_stmt is called when exiting the cluster_stmt production.
	ExitCluster_stmt(c *Cluster_stmtContext)

	// ExitComment_stmt is called when exiting the comment_stmt production.
	ExitComment_stmt(c *Comment_stmtContext)

	// ExitCommit_stmt is called when exiting the commit_stmt production.
	ExitCommit_stmt(c *Commit_stmtContext)

	// ExitCommit_prepared_stmt is called when exiting the commit_prepared_stmt production.
	ExitCommit_prepared_stmt(c *Commit_prepared_stmtContext)

	// ExitCopy_stmt is called when exiting the copy_stmt production.
	ExitCopy_stmt(c *Copy_stmtContext)

	// ExitCreate_stmt is called when exiting the create_stmt production.
	ExitCreate_stmt(c *Create_stmtContext)

	// ExitCreate_access_method_stmt is called when exiting the create_access_method_stmt production.
	ExitCreate_access_method_stmt(c *Create_access_method_stmtContext)

	// ExitCreate_aggregate_stmt is called when exiting the create_aggregate_stmt production.
	ExitCreate_aggregate_stmt(c *Create_aggregate_stmtContext)

	// ExitCreate_cast_stmt is called when exiting the create_cast_stmt production.
	ExitCreate_cast_stmt(c *Create_cast_stmtContext)

	// ExitCreate_collation_opt is called when exiting the create_collation_opt production.
	ExitCreate_collation_opt(c *Create_collation_optContext)

	// ExitCreate_collation_opt_list is called when exiting the create_collation_opt_list production.
	ExitCreate_collation_opt_list(c *Create_collation_opt_listContext)

	// ExitCreate_collation_stmt is called when exiting the create_collation_stmt production.
	ExitCreate_collation_stmt(c *Create_collation_stmtContext)

	// ExitCreate_conversion_stmt is called when exiting the create_conversion_stmt production.
	ExitCreate_conversion_stmt(c *Create_conversion_stmtContext)

	// ExitCreate_database_stmt is called when exiting the create_database_stmt production.
	ExitCreate_database_stmt(c *Create_database_stmtContext)

	// ExitDomain_constraint is called when exiting the domain_constraint production.
	ExitDomain_constraint(c *Domain_constraintContext)

	// ExitCreate_domain_stmt is called when exiting the create_domain_stmt production.
	ExitCreate_domain_stmt(c *Create_domain_stmtContext)

	// ExitCreate_event_trigger_cond is called when exiting the create_event_trigger_cond production.
	ExitCreate_event_trigger_cond(c *Create_event_trigger_condContext)

	// ExitCreate_event_trigger_stmt is called when exiting the create_event_trigger_stmt production.
	ExitCreate_event_trigger_stmt(c *Create_event_trigger_stmtContext)

	// ExitCreate_foreign_data_options is called when exiting the create_foreign_data_options production.
	ExitCreate_foreign_data_options(c *Create_foreign_data_optionsContext)

	// ExitCreate_foreign_data_stmt is called when exiting the create_foreign_data_stmt production.
	ExitCreate_foreign_data_stmt(c *Create_foreign_data_stmtContext)

	// ExitCreate_foreign_table_stmt is called when exiting the create_foreign_table_stmt production.
	ExitCreate_foreign_table_stmt(c *Create_foreign_table_stmtContext)

	// ExitCreate_function_stmt is called when exiting the create_function_stmt production.
	ExitCreate_function_stmt(c *Create_function_stmtContext)

	// ExitCreate_group_stmt is called when exiting the create_group_stmt production.
	ExitCreate_group_stmt(c *Create_group_stmtContext)

	// ExitCreate_index_stmt is called when exiting the create_index_stmt production.
	ExitCreate_index_stmt(c *Create_index_stmtContext)

	// ExitCreate_language_stmt is called when exiting the create_language_stmt production.
	ExitCreate_language_stmt(c *Create_language_stmtContext)

	// ExitCreate_materialized_view_stmt is called when exiting the create_materialized_view_stmt production.
	ExitCreate_materialized_view_stmt(c *Create_materialized_view_stmtContext)

	// ExitCreate_operator_stmt is called when exiting the create_operator_stmt production.
	ExitCreate_operator_stmt(c *Create_operator_stmtContext)

	// ExitCreate_operator_class_opt is called when exiting the create_operator_class_opt production.
	ExitCreate_operator_class_opt(c *Create_operator_class_optContext)

	// ExitCreate_operator_class_stmt is called when exiting the create_operator_class_stmt production.
	ExitCreate_operator_class_stmt(c *Create_operator_class_stmtContext)

	// ExitCreate_operator_family_stmt is called when exiting the create_operator_family_stmt production.
	ExitCreate_operator_family_stmt(c *Create_operator_family_stmtContext)

	// ExitCreate_policy_stmt is called when exiting the create_policy_stmt production.
	ExitCreate_policy_stmt(c *Create_policy_stmtContext)

	// ExitCreate_role_stmt is called when exiting the create_role_stmt production.
	ExitCreate_role_stmt(c *Create_role_stmtContext)

	// ExitCreate_rule_event is called when exiting the create_rule_event production.
	ExitCreate_rule_event(c *Create_rule_eventContext)

	// ExitCreate_rule_stmt is called when exiting the create_rule_stmt production.
	ExitCreate_rule_stmt(c *Create_rule_stmtContext)

	// ExitCreate_schema_stmt is called when exiting the create_schema_stmt production.
	ExitCreate_schema_stmt(c *Create_schema_stmtContext)

	// ExitCreate_sequence_stmt is called when exiting the create_sequence_stmt production.
	ExitCreate_sequence_stmt(c *Create_sequence_stmtContext)

	// ExitCreate_server_stmt is called when exiting the create_server_stmt production.
	ExitCreate_server_stmt(c *Create_server_stmtContext)

	// ExitCreate_statistics_stmt is called when exiting the create_statistics_stmt production.
	ExitCreate_statistics_stmt(c *Create_statistics_stmtContext)

	// ExitCreate_subscription_stmt is called when exiting the create_subscription_stmt production.
	ExitCreate_subscription_stmt(c *Create_subscription_stmtContext)

	// ExitCreate_table_stmt is called when exiting the create_table_stmt production.
	ExitCreate_table_stmt(c *Create_table_stmtContext)

	// ExitCreate_definitions is called when exiting the create_definitions production.
	ExitCreate_definitions(c *Create_definitionsContext)

	// ExitCreate_definition is called when exiting the create_definition production.
	ExitCreate_definition(c *Create_definitionContext)

	// ExitCreate_table_as_stmt is called when exiting the create_table_as_stmt production.
	ExitCreate_table_as_stmt(c *Create_table_as_stmtContext)

	// ExitCreate_tablespace_stmt is called when exiting the create_tablespace_stmt production.
	ExitCreate_tablespace_stmt(c *Create_tablespace_stmtContext)

	// ExitCreate_text_search_config_stmt is called when exiting the create_text_search_config_stmt production.
	ExitCreate_text_search_config_stmt(c *Create_text_search_config_stmtContext)

	// ExitCreate_text_search_dict_stmt is called when exiting the create_text_search_dict_stmt production.
	ExitCreate_text_search_dict_stmt(c *Create_text_search_dict_stmtContext)

	// ExitCreate_text_search_parser_stmt is called when exiting the create_text_search_parser_stmt production.
	ExitCreate_text_search_parser_stmt(c *Create_text_search_parser_stmtContext)

	// ExitCreate_text_search_template_stmt is called when exiting the create_text_search_template_stmt production.
	ExitCreate_text_search_template_stmt(c *Create_text_search_template_stmtContext)

	// ExitCreate_transform_stmt is called when exiting the create_transform_stmt production.
	ExitCreate_transform_stmt(c *Create_transform_stmtContext)

	// ExitCreate_trigger_stmt is called when exiting the create_trigger_stmt production.
	ExitCreate_trigger_stmt(c *Create_trigger_stmtContext)

	// ExitCreate_type_stmt is called when exiting the create_type_stmt production.
	ExitCreate_type_stmt(c *Create_type_stmtContext)

	// ExitCreate_user_stmt is called when exiting the create_user_stmt production.
	ExitCreate_user_stmt(c *Create_user_stmtContext)

	// ExitCreate_user_mapping_stmt is called when exiting the create_user_mapping_stmt production.
	ExitCreate_user_mapping_stmt(c *Create_user_mapping_stmtContext)

	// ExitCreate_view_stmt is called when exiting the create_view_stmt production.
	ExitCreate_view_stmt(c *Create_view_stmtContext)

	// ExitDeallocate_stmt is called when exiting the deallocate_stmt production.
	ExitDeallocate_stmt(c *Deallocate_stmtContext)

	// ExitDeclare_stmt is called when exiting the declare_stmt production.
	ExitDeclare_stmt(c *Declare_stmtContext)

	// ExitDelete_stmt is called when exiting the delete_stmt production.
	ExitDelete_stmt(c *Delete_stmtContext)

	// ExitDiscard_stmt is called when exiting the discard_stmt production.
	ExitDiscard_stmt(c *Discard_stmtContext)

	// ExitDrop_stmt is called when exiting the drop_stmt production.
	ExitDrop_stmt(c *Drop_stmtContext)

	// ExitDrop_access_method_stmt is called when exiting the drop_access_method_stmt production.
	ExitDrop_access_method_stmt(c *Drop_access_method_stmtContext)

	// ExitDrop_aggregate_stmt is called when exiting the drop_aggregate_stmt production.
	ExitDrop_aggregate_stmt(c *Drop_aggregate_stmtContext)

	// ExitDrop_cast_stmt is called when exiting the drop_cast_stmt production.
	ExitDrop_cast_stmt(c *Drop_cast_stmtContext)

	// ExitDrop_collation_stmt is called when exiting the drop_collation_stmt production.
	ExitDrop_collation_stmt(c *Drop_collation_stmtContext)

	// ExitDrop_conversion_stmt is called when exiting the drop_conversion_stmt production.
	ExitDrop_conversion_stmt(c *Drop_conversion_stmtContext)

	// ExitDrop_database_stmt is called when exiting the drop_database_stmt production.
	ExitDrop_database_stmt(c *Drop_database_stmtContext)

	// ExitDrop_domain_stmt is called when exiting the drop_domain_stmt production.
	ExitDrop_domain_stmt(c *Drop_domain_stmtContext)

	// ExitDrop_event_trigger_stmt is called when exiting the drop_event_trigger_stmt production.
	ExitDrop_event_trigger_stmt(c *Drop_event_trigger_stmtContext)

	// ExitDrop_extension_stmt is called when exiting the drop_extension_stmt production.
	ExitDrop_extension_stmt(c *Drop_extension_stmtContext)

	// ExitDrop_foreign_data_wrapper_stmt is called when exiting the drop_foreign_data_wrapper_stmt production.
	ExitDrop_foreign_data_wrapper_stmt(c *Drop_foreign_data_wrapper_stmtContext)

	// ExitDrop_foreign_table_stmt is called when exiting the drop_foreign_table_stmt production.
	ExitDrop_foreign_table_stmt(c *Drop_foreign_table_stmtContext)

	// ExitDrop_function_stmt is called when exiting the drop_function_stmt production.
	ExitDrop_function_stmt(c *Drop_function_stmtContext)

	// ExitDrop_group_stmt is called when exiting the drop_group_stmt production.
	ExitDrop_group_stmt(c *Drop_group_stmtContext)

	// ExitDrop_index_stmt is called when exiting the drop_index_stmt production.
	ExitDrop_index_stmt(c *Drop_index_stmtContext)

	// ExitDrop_language_stmt is called when exiting the drop_language_stmt production.
	ExitDrop_language_stmt(c *Drop_language_stmtContext)

	// ExitDrop_materialized_view_stmt is called when exiting the drop_materialized_view_stmt production.
	ExitDrop_materialized_view_stmt(c *Drop_materialized_view_stmtContext)

	// ExitDrop_operator_stmt is called when exiting the drop_operator_stmt production.
	ExitDrop_operator_stmt(c *Drop_operator_stmtContext)

	// ExitDrop_operator_class_stmt is called when exiting the drop_operator_class_stmt production.
	ExitDrop_operator_class_stmt(c *Drop_operator_class_stmtContext)

	// ExitDrop_operator_family_stmt is called when exiting the drop_operator_family_stmt production.
	ExitDrop_operator_family_stmt(c *Drop_operator_family_stmtContext)

	// ExitDrop_owned_stmt is called when exiting the drop_owned_stmt production.
	ExitDrop_owned_stmt(c *Drop_owned_stmtContext)

	// ExitDrop_policy_stmt is called when exiting the drop_policy_stmt production.
	ExitDrop_policy_stmt(c *Drop_policy_stmtContext)

	// ExitDrop_publication_stmt is called when exiting the drop_publication_stmt production.
	ExitDrop_publication_stmt(c *Drop_publication_stmtContext)

	// ExitDrop_role_stmt is called when exiting the drop_role_stmt production.
	ExitDrop_role_stmt(c *Drop_role_stmtContext)

	// ExitDrop_rule_stmt is called when exiting the drop_rule_stmt production.
	ExitDrop_rule_stmt(c *Drop_rule_stmtContext)

	// ExitDrop_schema_stmt is called when exiting the drop_schema_stmt production.
	ExitDrop_schema_stmt(c *Drop_schema_stmtContext)

	// ExitDrop_sequence_stmt is called when exiting the drop_sequence_stmt production.
	ExitDrop_sequence_stmt(c *Drop_sequence_stmtContext)

	// ExitDrop_server_stmt is called when exiting the drop_server_stmt production.
	ExitDrop_server_stmt(c *Drop_server_stmtContext)

	// ExitDrop_statistics_stmt is called when exiting the drop_statistics_stmt production.
	ExitDrop_statistics_stmt(c *Drop_statistics_stmtContext)

	// ExitDrop_subscription_stmt is called when exiting the drop_subscription_stmt production.
	ExitDrop_subscription_stmt(c *Drop_subscription_stmtContext)

	// ExitDrop_table_stmt is called when exiting the drop_table_stmt production.
	ExitDrop_table_stmt(c *Drop_table_stmtContext)

	// ExitDrop_tablespace_stmt is called when exiting the drop_tablespace_stmt production.
	ExitDrop_tablespace_stmt(c *Drop_tablespace_stmtContext)

	// ExitDrop_text_search_config_stmt is called when exiting the drop_text_search_config_stmt production.
	ExitDrop_text_search_config_stmt(c *Drop_text_search_config_stmtContext)

	// ExitDrop_text_search_dict_stmt is called when exiting the drop_text_search_dict_stmt production.
	ExitDrop_text_search_dict_stmt(c *Drop_text_search_dict_stmtContext)

	// ExitDrop_text_search_parser_stmt is called when exiting the drop_text_search_parser_stmt production.
	ExitDrop_text_search_parser_stmt(c *Drop_text_search_parser_stmtContext)

	// ExitDrop_text_search_template_stmt is called when exiting the drop_text_search_template_stmt production.
	ExitDrop_text_search_template_stmt(c *Drop_text_search_template_stmtContext)

	// ExitDrop_transform_stmt is called when exiting the drop_transform_stmt production.
	ExitDrop_transform_stmt(c *Drop_transform_stmtContext)

	// ExitDrop_trigger_stmt is called when exiting the drop_trigger_stmt production.
	ExitDrop_trigger_stmt(c *Drop_trigger_stmtContext)

	// ExitDrop_type_stmt is called when exiting the drop_type_stmt production.
	ExitDrop_type_stmt(c *Drop_type_stmtContext)

	// ExitDrop_user_stmt is called when exiting the drop_user_stmt production.
	ExitDrop_user_stmt(c *Drop_user_stmtContext)

	// ExitDrop_user_mapping_stmt is called when exiting the drop_user_mapping_stmt production.
	ExitDrop_user_mapping_stmt(c *Drop_user_mapping_stmtContext)

	// ExitDrop_view_stmt is called when exiting the drop_view_stmt production.
	ExitDrop_view_stmt(c *Drop_view_stmtContext)

	// ExitExecute_stmt is called when exiting the execute_stmt production.
	ExitExecute_stmt(c *Execute_stmtContext)

	// ExitExplain_stmt is called when exiting the explain_stmt production.
	ExitExplain_stmt(c *Explain_stmtContext)

	// ExitFetch_stmt is called when exiting the fetch_stmt production.
	ExitFetch_stmt(c *Fetch_stmtContext)

	// ExitGrant_stmt is called when exiting the grant_stmt production.
	ExitGrant_stmt(c *Grant_stmtContext)

	// ExitImport_foreign_schema_stmt is called when exiting the import_foreign_schema_stmt production.
	ExitImport_foreign_schema_stmt(c *Import_foreign_schema_stmtContext)

	// ExitInsert_stmt is called when exiting the insert_stmt production.
	ExitInsert_stmt(c *Insert_stmtContext)

	// ExitListen_stmt is called when exiting the listen_stmt production.
	ExitListen_stmt(c *Listen_stmtContext)

	// ExitLoad_stmt is called when exiting the load_stmt production.
	ExitLoad_stmt(c *Load_stmtContext)

	// ExitLock_stmt is called when exiting the lock_stmt production.
	ExitLock_stmt(c *Lock_stmtContext)

	// ExitMove_stmt is called when exiting the move_stmt production.
	ExitMove_stmt(c *Move_stmtContext)

	// ExitNotify_stmt is called when exiting the notify_stmt production.
	ExitNotify_stmt(c *Notify_stmtContext)

	// ExitPrepare_stmt is called when exiting the prepare_stmt production.
	ExitPrepare_stmt(c *Prepare_stmtContext)

	// ExitPrepare_transaction_stmt is called when exiting the prepare_transaction_stmt production.
	ExitPrepare_transaction_stmt(c *Prepare_transaction_stmtContext)

	// ExitReassign_owned_stmt is called when exiting the reassign_owned_stmt production.
	ExitReassign_owned_stmt(c *Reassign_owned_stmtContext)

	// ExitRefresh_materialized_view_stmt is called when exiting the refresh_materialized_view_stmt production.
	ExitRefresh_materialized_view_stmt(c *Refresh_materialized_view_stmtContext)

	// ExitReindex_stmt is called when exiting the reindex_stmt production.
	ExitReindex_stmt(c *Reindex_stmtContext)

	// ExitRelease_savepoint_stmt is called when exiting the release_savepoint_stmt production.
	ExitRelease_savepoint_stmt(c *Release_savepoint_stmtContext)

	// ExitReset_stmt is called when exiting the reset_stmt production.
	ExitReset_stmt(c *Reset_stmtContext)

	// ExitRevoke_stmt is called when exiting the revoke_stmt production.
	ExitRevoke_stmt(c *Revoke_stmtContext)

	// ExitRollback_stmt is called when exiting the rollback_stmt production.
	ExitRollback_stmt(c *Rollback_stmtContext)

	// ExitRollback_prepared_stmt is called when exiting the rollback_prepared_stmt production.
	ExitRollback_prepared_stmt(c *Rollback_prepared_stmtContext)

	// ExitRollback_to_savepoint_stmt is called when exiting the rollback_to_savepoint_stmt production.
	ExitRollback_to_savepoint_stmt(c *Rollback_to_savepoint_stmtContext)

	// ExitSavepoint_stmt is called when exiting the savepoint_stmt production.
	ExitSavepoint_stmt(c *Savepoint_stmtContext)

	// ExitSecurity_label_stmt is called when exiting the security_label_stmt production.
	ExitSecurity_label_stmt(c *Security_label_stmtContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitSelect_into_stmt is called when exiting the select_into_stmt production.
	ExitSelect_into_stmt(c *Select_into_stmtContext)

	// ExitWith_clause is called when exiting the with_clause production.
	ExitWith_clause(c *With_clauseContext)

	// ExitWith_expr is called when exiting the with_expr production.
	ExitWith_expr(c *With_exprContext)

	// ExitSet_stmt is called when exiting the set_stmt production.
	ExitSet_stmt(c *Set_stmtContext)

	// ExitSet_constraints_stmt is called when exiting the set_constraints_stmt production.
	ExitSet_constraints_stmt(c *Set_constraints_stmtContext)

	// ExitSet_role_stmt is called when exiting the set_role_stmt production.
	ExitSet_role_stmt(c *Set_role_stmtContext)

	// ExitSet_session_authorization_stmt is called when exiting the set_session_authorization_stmt production.
	ExitSet_session_authorization_stmt(c *Set_session_authorization_stmtContext)

	// ExitTransaction_mode is called when exiting the transaction_mode production.
	ExitTransaction_mode(c *Transaction_modeContext)

	// ExitTransaction_mode_list is called when exiting the transaction_mode_list production.
	ExitTransaction_mode_list(c *Transaction_mode_listContext)

	// ExitSet_transaction_stmt is called when exiting the set_transaction_stmt production.
	ExitSet_transaction_stmt(c *Set_transaction_stmtContext)

	// ExitShow_stmt is called when exiting the show_stmt production.
	ExitShow_stmt(c *Show_stmtContext)

	// ExitTruncate_stmt is called when exiting the truncate_stmt production.
	ExitTruncate_stmt(c *Truncate_stmtContext)

	// ExitUnlisten_stmt is called when exiting the unlisten_stmt production.
	ExitUnlisten_stmt(c *Unlisten_stmtContext)

	// ExitUpdate_stmt is called when exiting the update_stmt production.
	ExitUpdate_stmt(c *Update_stmtContext)

	// ExitVacuum_opt is called when exiting the vacuum_opt production.
	ExitVacuum_opt(c *Vacuum_optContext)

	// ExitVacuum_opt_list is called when exiting the vacuum_opt_list production.
	ExitVacuum_opt_list(c *Vacuum_opt_listContext)

	// ExitVacuum_stmt is called when exiting the vacuum_stmt production.
	ExitVacuum_stmt(c *Vacuum_stmtContext)

	// ExitValues_stmt is called when exiting the values_stmt production.
	ExitValues_stmt(c *Values_stmtContext)

	// ExitSelector_clause is called when exiting the selector_clause production.
	ExitSelector_clause(c *Selector_clauseContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitGroup_by_clause is called when exiting the group_by_clause production.
	ExitGroup_by_clause(c *Group_by_clauseContext)

	// ExitGrouping_elem is called when exiting the grouping_elem production.
	ExitGrouping_elem(c *Grouping_elemContext)

	// ExitGrouping_elem_list is called when exiting the grouping_elem_list production.
	ExitGrouping_elem_list(c *Grouping_elem_listContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitColumn_list is called when exiting the column_list production.
	ExitColumn_list(c *Column_listContext)

	// ExitExplain_parameter is called when exiting the explain_parameter production.
	ExitExplain_parameter(c *Explain_parameterContext)

	// ExitFrame is called when exiting the frame production.
	ExitFrame(c *FrameContext)

	// ExitFrame_start is called when exiting the frame_start production.
	ExitFrame_start(c *Frame_startContext)

	// ExitFrame_end is called when exiting the frame_end production.
	ExitFrame_end(c *Frame_endContext)

	// ExitFrame_clause is called when exiting the frame_clause production.
	ExitFrame_clause(c *Frame_clauseContext)

	// ExitWindow_definition is called when exiting the window_definition production.
	ExitWindow_definition(c *Window_definitionContext)

	// ExitWindow_clause is called when exiting the window_clause production.
	ExitWindow_clause(c *Window_clauseContext)

	// ExitCombine_clause is called when exiting the combine_clause production.
	ExitCombine_clause(c *Combine_clauseContext)

	// ExitOrder_by_clause is called when exiting the order_by_clause production.
	ExitOrder_by_clause(c *Order_by_clauseContext)

	// ExitOrder_by_item is called when exiting the order_by_item production.
	ExitOrder_by_item(c *Order_by_itemContext)

	// ExitLimit_clause is called when exiting the limit_clause production.
	ExitLimit_clause(c *Limit_clauseContext)

	// ExitOffset_clause is called when exiting the offset_clause production.
	ExitOffset_clause(c *Offset_clauseContext)

	// ExitFetch_clause is called when exiting the fetch_clause production.
	ExitFetch_clause(c *Fetch_clauseContext)

	// ExitFor_clause is called when exiting the for_clause production.
	ExitFor_clause(c *For_clauseContext)

	// ExitUpdater_clause is called when exiting the updater_clause production.
	ExitUpdater_clause(c *Updater_clauseContext)

	// ExitUpdater_expr is called when exiting the updater_expr production.
	ExitUpdater_expr(c *Updater_exprContext)

	// ExitReturning_clause is called when exiting the returning_clause production.
	ExitReturning_clause(c *Returning_clauseContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitBool_expr is called when exiting the bool_expr production.
	ExitBool_expr(c *Bool_exprContext)

	// ExitCase_expr is called when exiting the case_expr production.
	ExitCase_expr(c *Case_exprContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitExpr_list_list is called when exiting the expr_list_list production.
	ExitExpr_list_list(c *Expr_list_listContext)

	// ExitFunc_sig_arg is called when exiting the func_sig_arg production.
	ExitFunc_sig_arg(c *Func_sig_argContext)

	// ExitFunc_sig_arg_list is called when exiting the func_sig_arg_list production.
	ExitFunc_sig_arg_list(c *Func_sig_arg_listContext)

	// ExitFunc_sig is called when exiting the func_sig production.
	ExitFunc_sig(c *Func_sigContext)

	// ExitFunc_sig_list is called when exiting the func_sig_list production.
	ExitFunc_sig_list(c *Func_sig_listContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitTimezone is called when exiting the timezone production.
	ExitTimezone(c *TimezoneContext)

	// ExitOper is called when exiting the oper production.
	ExitOper(c *OperContext)

	// ExitAggregate is called when exiting the aggregate production.
	ExitAggregate(c *AggregateContext)

	// ExitName_ is called when exiting the name_ production.
	ExitName_(c *Name_Context)

	// ExitName_list is called when exiting the name_list production.
	ExitName_list(c *Name_listContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitOption_expr is called when exiting the option_expr production.
	ExitOption_expr(c *Option_exprContext)

	// ExitOption_list is called when exiting the option_list production.
	ExitOption_list(c *Option_listContext)

	// ExitTable_name_ is called when exiting the table_name_ production.
	ExitTable_name_(c *Table_name_Context)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitData_type_list is called when exiting the data_type_list production.
	ExitData_type_list(c *Data_type_listContext)

	// ExitIndex_method is called when exiting the index_method production.
	ExitIndex_method(c *Index_methodContext)

	// ExitFunc_name is called when exiting the func_name production.
	ExitFunc_name(c *Func_nameContext)

	// ExitFunc_call is called when exiting the func_call production.
	ExitFunc_call(c *Func_callContext)

	// ExitArray_cons_expr is called when exiting the array_cons_expr production.
	ExitArray_cons_expr(c *Array_cons_exprContext)

	// ExitFrom_item is called when exiting the from_item production.
	ExitFrom_item(c *From_itemContext)

	// ExitWith_column_alias is called when exiting the with_column_alias production.
	ExitWith_column_alias(c *With_column_aliasContext)

	// ExitJoin_type is called when exiting the join_type production.
	ExitJoin_type(c *Join_typeContext)

	// ExitJoin_clause is called when exiting the join_clause production.
	ExitJoin_clause(c *Join_clauseContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitAggregate_signature is called when exiting the aggregate_signature production.
	ExitAggregate_signature(c *Aggregate_signatureContext)

	// ExitColumn_constraint is called when exiting the column_constraint production.
	ExitColumn_constraint(c *Column_constraintContext)

	// ExitColumn_constraints is called when exiting the column_constraints production.
	ExitColumn_constraints(c *Column_constraintsContext)

	// ExitIndex_parameters is called when exiting the index_parameters production.
	ExitIndex_parameters(c *Index_parametersContext)

	// ExitExclude_element is called when exiting the exclude_element production.
	ExitExclude_element(c *Exclude_elementContext)

	// ExitTable_constraint is called when exiting the table_constraint production.
	ExitTable_constraint(c *Table_constraintContext)

	// ExitRole_name is called when exiting the role_name production.
	ExitRole_name(c *Role_nameContext)

	// ExitRole_name_list is called when exiting the role_name_list production.
	ExitRole_name_list(c *Role_name_listContext)

	// ExitParam_value is called when exiting the param_value production.
	ExitParam_value(c *Param_valueContext)

	// ExitNon_reserved_keyword is called when exiting the non_reserved_keyword production.
	ExitNon_reserved_keyword(c *Non_reserved_keywordContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitTodo_fill_in is called when exiting the todo_fill_in production.
	ExitTodo_fill_in(c *Todo_fill_inContext)

	// ExitTodo_implement is called when exiting the todo_implement production.
	ExitTodo_implement(c *Todo_implementContext)

	// ExitCorrelation_name is called when exiting the correlation_name production.
	ExitCorrelation_name(c *Correlation_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitColumn_alias is called when exiting the column_alias production.
	ExitColumn_alias(c *Column_aliasContext)

	// ExitColumn_definition is called when exiting the column_definition production.
	ExitColumn_definition(c *Column_definitionContext)

	// ExitWindow_name is called when exiting the window_name production.
	ExitWindow_name(c *Window_nameContext)
}
