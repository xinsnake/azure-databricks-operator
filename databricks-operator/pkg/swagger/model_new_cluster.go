/*
 * Databricks Controller
 *
 * Manage Databricks
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NewCluster struct {
	SparkVersion      string       `json:"spark_version,omitempty"`
	SparkConf         *interface{} `json:"spark_conf,omitempty"`
	NodeTypeId        string       `json:"node_type_id,omitempty"`
	SparkEnvVars      *interface{} `json:"spark_env_vars,omitempty"`
	EnableElasticDisk bool         `json:"enable_elastic_disk,omitempty"`
	NumWorkers        int32        `json:"num_workers,omitempty"`
}
