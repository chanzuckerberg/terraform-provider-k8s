
# resource "k8s_core_v1_replication_controller"

ReplicationController represents the configuration of a replication controller.

  
<details>
<summary>metadata</summary><blockquote>

    
- [annotations](#annotations)
- [creation_timestamp](#creation_timestamp)
- [deletion_grace_period_seconds](#deletion_grace_period_seconds)
- [deletion_timestamp](#deletion_timestamp)
- [labels](#labels)
- [name](#name)
- [namespace](#namespace)
- [self_link](#self_link)
- [uid](#uid)

    
</details>

<details>
<summary>spec</summary><blockquote>

    
- [min_ready_seconds](#min_ready_seconds)
- [replicas](#replicas)
- [selector](#selector)

    
<details>
<summary>template</summary><blockquote>

    

    
<details>
<summary>metadata</summary><blockquote>

    
- [annotations](#annotations)
- [creation_timestamp](#creation_timestamp)
- [deletion_grace_period_seconds](#deletion_grace_period_seconds)
- [deletion_timestamp](#deletion_timestamp)
- [labels](#labels)
- [name](#name)
- [namespace](#namespace)
- [self_link](#self_link)
- [uid](#uid)

    
</details>

<details>
<summary>spec</summary><blockquote>

    
- [active_deadline_seconds](#active_deadline_seconds)
- [automount_service_account_token](#automount_service_account_token)
- [dns_policy](#dns_policy)
- [enable_service_links](#enable_service_links)
- [host_ipc](#host_ipc)
- [host_network](#host_network)
- [host_pid](#host_pid)
- [hostname](#hostname)
- [node_name](#node_name)
- [node_selector](#node_selector)
- [priority](#priority)
- [priority_class_name](#priority_class_name)
- [restart_policy](#restart_policy)
- [runtime_class_name](#runtime_class_name)
- [scheduler_name](#scheduler_name)
- [service_account](#service_account)
- [service_account_name](#service_account_name)
- [share_process_namespace](#share_process_namespace)
- [subdomain](#subdomain)
- [termination_grace_period_seconds](#termination_grace_period_seconds)

    
<details>
<summary>affinity</summary><blockquote>

    

    
<details>
<summary>node_affinity</summary><blockquote>

    

    
<details>
<summary>preferred_during_scheduling_ignored_during_execution</summary><blockquote>

    
- [weight](#weight)*

    
<details>
<summary>preference</summary><blockquote>

    

    
<details>
<summary>match_expressions</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

<details>
<summary>match_fields</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

</details>

</details>

<details>
<summary>required_during_scheduling_ignored_during_execution</summary><blockquote>

    

    
<details>
<summary>node_selector_terms</summary><blockquote>

    

    
<details>
<summary>match_expressions</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

<details>
<summary>match_fields</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

</details>

</details>

</details>

<details>
<summary>pod_affinity</summary><blockquote>

    

    
<details>
<summary>preferred_during_scheduling_ignored_during_execution</summary><blockquote>

    
- [weight](#weight)*

    
<details>
<summary>pod_affinity_term</summary><blockquote>

    
- [namespaces](#namespaces)
- [topology_key](#topology_key)*

    
<details>
<summary>label_selector</summary><blockquote>

    
- [match_labels](#match_labels)

    
<details>
<summary>match_expressions</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

</details>

</details>

</details>

<details>
<summary>required_during_scheduling_ignored_during_execution</summary><blockquote>

    
- [namespaces](#namespaces)
- [topology_key](#topology_key)*

    
<details>
<summary>label_selector</summary><blockquote>

    
- [match_labels](#match_labels)

    
<details>
<summary>match_expressions</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

</details>

</details>

</details>

<details>
<summary>pod_anti_affinity</summary><blockquote>

    

    
<details>
<summary>preferred_during_scheduling_ignored_during_execution</summary><blockquote>

    
- [weight](#weight)*

    
<details>
<summary>pod_affinity_term</summary><blockquote>

    
- [namespaces](#namespaces)
- [topology_key](#topology_key)*

    
<details>
<summary>label_selector</summary><blockquote>

    
- [match_labels](#match_labels)

    
<details>
<summary>match_expressions</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

</details>

</details>

</details>

<details>
<summary>required_during_scheduling_ignored_during_execution</summary><blockquote>

    
- [namespaces](#namespaces)
- [topology_key](#topology_key)*

    
<details>
<summary>label_selector</summary><blockquote>

    
- [match_labels](#match_labels)

    
<details>
<summary>match_expressions</summary><blockquote>

    
- [key](#key)*
- [operator](#operator)*
- [values](#values)

    
</details>

</details>

</details>

</details>

</details>

<details>
<summary>containers</summary><blockquote>

    
- [args](#args)
- [command](#command)
- [image](#image)
- [image_pull_policy](#image_pull_policy)
- [name](#name)*
- [stdin](#stdin)
- [stdin_once](#stdin_once)
- [termination_message_path](#termination_message_path)
- [termination_message_policy](#termination_message_policy)
- [tty](#tty)
- [working_dir](#working_dir)

    
<details>
<summary>env</summary><blockquote>

    
- [name](#name)*
- [value](#value)

    
<details>
<summary>value_from</summary><blockquote>

    

    
<details>
<summary>config_map_keyref</summary><blockquote>

    
- [key](#key)*
- [name](#name)
- [optional](#optional)

    
</details>

<details>
<summary>field_ref</summary><blockquote>

    
- [api_version](#api_version)
- [field_path](#field_path)*

    
</details>

<details>
<summary>resource_field_ref</summary><blockquote>

    
- [container_name](#container_name)
- [divisor](#divisor)
- [resource](#resource)*

    
</details>

<details>
<summary>secret_key_ref</summary><blockquote>

    
- [key](#key)*
- [name](#name)
- [optional](#optional)

    
</details>

</details>

</details>

<details>
<summary>env_from</summary><blockquote>

    
- [prefix](#prefix)

    
<details>
<summary>config_map_ref</summary><blockquote>

    
- [name](#name)
- [optional](#optional)

    
</details>

<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)
- [optional](#optional)

    
</details>

</details>

<details>
<summary>lifecycle</summary><blockquote>

    

    
<details>
<summary>post_start</summary><blockquote>

    

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

<details>
<summary>pre_stop</summary><blockquote>

    

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

</details>

<details>
<summary>liveness_probe</summary><blockquote>

    
- [failure_threshold](#failure_threshold)
- [initial_delay_seconds](#initial_delay_seconds)
- [period_seconds](#period_seconds)
- [success_threshold](#success_threshold)
- [timeout_seconds](#timeout_seconds)

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

<details>
<summary>ports</summary><blockquote>

    
- [container_port](#container_port)*
- [host_ip](#host_ip)
- [host_port](#host_port)
- [name](#name)
- [protocol](#protocol)

    
</details>

<details>
<summary>readiness_probe</summary><blockquote>

    
- [failure_threshold](#failure_threshold)
- [initial_delay_seconds](#initial_delay_seconds)
- [period_seconds](#period_seconds)
- [success_threshold](#success_threshold)
- [timeout_seconds](#timeout_seconds)

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

<details>
<summary>resources</summary><blockquote>

    
- [limits](#limits)
- [requests](#requests)

    
</details>

<details>
<summary>security_context</summary><blockquote>

    
- [allow_privilege_escalation](#allow_privilege_escalation)
- [privileged](#privileged)
- [proc_mount](#proc_mount)
- [read_only_root_filesystem](#read_only_root_filesystem)
- [run_asgroup](#run_asgroup)
- [run_asnon_root](#run_asnon_root)
- [run_asuser](#run_asuser)

    
<details>
<summary>capabilities</summary><blockquote>

    
- [add](#add)
- [drop](#drop)

    
</details>

<details>
<summary>selinux_options</summary><blockquote>

    
- [level](#level)
- [role](#role)
- [type](#type)
- [user](#user)

    
</details>

</details>

<details>
<summary>volume_devices</summary><blockquote>

    
- [device_path](#device_path)*
- [name](#name)*

    
</details>

<details>
<summary>volume_mounts</summary><blockquote>

    
- [mount_path](#mount_path)*
- [mount_propagation](#mount_propagation)
- [name](#name)*
- [read_only](#read_only)
- [sub_path](#sub_path)
- [sub_path_expr](#sub_path_expr)

    
</details>

</details>

<details>
<summary>dns_config</summary><blockquote>

    
- [nameservers](#nameservers)
- [searches](#searches)

    
<details>
<summary>options</summary><blockquote>

    
- [name](#name)
- [value](#value)

    
</details>

</details>

<details>
<summary>host_aliases</summary><blockquote>

    
- [hostnames](#hostnames)
- [ip](#ip)

    
</details>

<details>
<summary>image_pull_secrets</summary><blockquote>

    
- [name](#name)

    
</details>

<details>
<summary>init_containers</summary><blockquote>

    
- [args](#args)
- [command](#command)
- [image](#image)
- [image_pull_policy](#image_pull_policy)
- [name](#name)*
- [stdin](#stdin)
- [stdin_once](#stdin_once)
- [termination_message_path](#termination_message_path)
- [termination_message_policy](#termination_message_policy)
- [tty](#tty)
- [working_dir](#working_dir)

    
<details>
<summary>env</summary><blockquote>

    
- [name](#name)*
- [value](#value)

    
<details>
<summary>value_from</summary><blockquote>

    

    
<details>
<summary>config_map_keyref</summary><blockquote>

    
- [key](#key)*
- [name](#name)
- [optional](#optional)

    
</details>

<details>
<summary>field_ref</summary><blockquote>

    
- [api_version](#api_version)
- [field_path](#field_path)*

    
</details>

<details>
<summary>resource_field_ref</summary><blockquote>

    
- [container_name](#container_name)
- [divisor](#divisor)
- [resource](#resource)*

    
</details>

<details>
<summary>secret_key_ref</summary><blockquote>

    
- [key](#key)*
- [name](#name)
- [optional](#optional)

    
</details>

</details>

</details>

<details>
<summary>env_from</summary><blockquote>

    
- [prefix](#prefix)

    
<details>
<summary>config_map_ref</summary><blockquote>

    
- [name](#name)
- [optional](#optional)

    
</details>

<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)
- [optional](#optional)

    
</details>

</details>

<details>
<summary>lifecycle</summary><blockquote>

    

    
<details>
<summary>post_start</summary><blockquote>

    

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

<details>
<summary>pre_stop</summary><blockquote>

    

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

</details>

<details>
<summary>liveness_probe</summary><blockquote>

    
- [failure_threshold](#failure_threshold)
- [initial_delay_seconds](#initial_delay_seconds)
- [period_seconds](#period_seconds)
- [success_threshold](#success_threshold)
- [timeout_seconds](#timeout_seconds)

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

<details>
<summary>ports</summary><blockquote>

    
- [container_port](#container_port)*
- [host_ip](#host_ip)
- [host_port](#host_port)
- [name](#name)
- [protocol](#protocol)

    
</details>

<details>
<summary>readiness_probe</summary><blockquote>

    
- [failure_threshold](#failure_threshold)
- [initial_delay_seconds](#initial_delay_seconds)
- [period_seconds](#period_seconds)
- [success_threshold](#success_threshold)
- [timeout_seconds](#timeout_seconds)

    
<details>
<summary>exec</summary><blockquote>

    
- [command](#command)

    
</details>

<details>
<summary>http_get</summary><blockquote>

    
- [host](#host)
- [path](#path)
- [port](#port)*
- [scheme](#scheme)

    
<details>
<summary>http_headers</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tcp_socket</summary><blockquote>

    
- [host](#host)
- [port](#port)*

    
</details>

</details>

<details>
<summary>resources</summary><blockquote>

    
- [limits](#limits)
- [requests](#requests)

    
</details>

<details>
<summary>security_context</summary><blockquote>

    
- [allow_privilege_escalation](#allow_privilege_escalation)
- [privileged](#privileged)
- [proc_mount](#proc_mount)
- [read_only_root_filesystem](#read_only_root_filesystem)
- [run_asgroup](#run_asgroup)
- [run_asnon_root](#run_asnon_root)
- [run_asuser](#run_asuser)

    
<details>
<summary>capabilities</summary><blockquote>

    
- [add](#add)
- [drop](#drop)

    
</details>

<details>
<summary>selinux_options</summary><blockquote>

    
- [level](#level)
- [role](#role)
- [type](#type)
- [user](#user)

    
</details>

</details>

<details>
<summary>volume_devices</summary><blockquote>

    
- [device_path](#device_path)*
- [name](#name)*

    
</details>

<details>
<summary>volume_mounts</summary><blockquote>

    
- [mount_path](#mount_path)*
- [mount_propagation](#mount_propagation)
- [name](#name)*
- [read_only](#read_only)
- [sub_path](#sub_path)
- [sub_path_expr](#sub_path_expr)

    
</details>

</details>

<details>
<summary>readiness_gates</summary><blockquote>

    
- [condition_type](#condition_type)*

    
</details>

<details>
<summary>security_context</summary><blockquote>

    
- [fsgroup](#fsgroup)
- [run_asgroup](#run_asgroup)
- [run_asnon_root](#run_asnon_root)
- [run_asuser](#run_asuser)
- [supplemental_groups](#supplemental_groups)

    
<details>
<summary>selinux_options</summary><blockquote>

    
- [level](#level)
- [role](#role)
- [type](#type)
- [user](#user)

    
</details>

<details>
<summary>sysctls</summary><blockquote>

    
- [name](#name)*
- [value](#value)*

    
</details>

</details>

<details>
<summary>tolerations</summary><blockquote>

    
- [effect](#effect)
- [key](#key)
- [operator](#operator)
- [toleration_seconds](#toleration_seconds)
- [value](#value)

    
</details>

<details>
<summary>volumes</summary><blockquote>

    
- [name](#name)*

    
<details>
<summary>aws_elastic_block_store</summary><blockquote>

    
- [fstype](#fstype)
- [partition](#partition)
- [read_only](#read_only)
- [volume_id](#volume_id)*

    
</details>

<details>
<summary>azure_disk</summary><blockquote>

    
- [caching_mode](#caching_mode)
- [disk_name](#disk_name)*
- [disk_uri](#disk_uri)*
- [fstype](#fstype)
- [kind](#kind)
- [read_only](#read_only)

    
</details>

<details>
<summary>azure_file</summary><blockquote>

    
- [read_only](#read_only)
- [secret_name](#secret_name)*
- [share_name](#share_name)*

    
</details>

<details>
<summary>cephfs</summary><blockquote>

    
- [monitors](#monitors)*
- [path](#path)
- [read_only](#read_only)
- [secret_file](#secret_file)
- [user](#user)

    
<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>cinder</summary><blockquote>

    
- [fstype](#fstype)
- [read_only](#read_only)
- [volume_id](#volume_id)*

    
<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>config_map</summary><blockquote>

    
- [default_mode](#default_mode)
- [name](#name)
- [optional](#optional)

    
<details>
<summary>items</summary><blockquote>

    
- [key](#key)*
- [mode](#mode)
- [path](#path)*

    
</details>

</details>

<details>
<summary>csi</summary><blockquote>

    
- [driver](#driver)*
- [fstype](#fstype)
- [read_only](#read_only)
- [volume_attributes](#volume_attributes)

    
<details>
<summary>node_publish_secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>downward_api</summary><blockquote>

    
- [default_mode](#default_mode)

    
<details>
<summary>items</summary><blockquote>

    
- [mode](#mode)
- [path](#path)*

    
<details>
<summary>field_ref</summary><blockquote>

    
- [api_version](#api_version)
- [field_path](#field_path)*

    
</details>

<details>
<summary>resource_field_ref</summary><blockquote>

    
- [container_name](#container_name)
- [divisor](#divisor)
- [resource](#resource)*

    
</details>

</details>

</details>

<details>
<summary>empty_dir</summary><blockquote>

    
- [medium](#medium)
- [size_limit](#size_limit)

    
</details>

<details>
<summary>fc</summary><blockquote>

    
- [fstype](#fstype)
- [lun](#lun)
- [read_only](#read_only)
- [target_wwns](#target_wwns)
- [wwids](#wwids)

    
</details>

<details>
<summary>flex_volume</summary><blockquote>

    
- [driver](#driver)*
- [fstype](#fstype)
- [options](#options)
- [read_only](#read_only)

    
<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>flocker</summary><blockquote>

    
- [dataset_name](#dataset_name)
- [dataset_uuid](#dataset_uuid)

    
</details>

<details>
<summary>gce_persistent_disk</summary><blockquote>

    
- [fstype](#fstype)
- [partition](#partition)
- [pdname](#pdname)*
- [read_only](#read_only)

    
</details>

<details>
<summary>git_repo</summary><blockquote>

    
- [directory](#directory)
- [repository](#repository)*
- [revision](#revision)

    
</details>

<details>
<summary>glusterfs</summary><blockquote>

    
- [endpoints](#endpoints)*
- [path](#path)*
- [read_only](#read_only)

    
</details>

<details>
<summary>host_path</summary><blockquote>

    
- [path](#path)*
- [type](#type)

    
</details>

<details>
<summary>iscsi</summary><blockquote>

    
- [chap_auth_discovery](#chap_auth_discovery)
- [chap_auth_session](#chap_auth_session)
- [fstype](#fstype)
- [initiator_name](#initiator_name)
- [iqn](#iqn)*
- [iscsi_interface](#iscsi_interface)
- [lun](#lun)*
- [portals](#portals)
- [read_only](#read_only)
- [target_portal](#target_portal)*

    
<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>nfs</summary><blockquote>

    
- [path](#path)*
- [read_only](#read_only)
- [server](#server)*

    
</details>

<details>
<summary>persistent_volume_claim</summary><blockquote>

    
- [claim_name](#claim_name)*
- [read_only](#read_only)

    
</details>

<details>
<summary>photon_persistent_disk</summary><blockquote>

    
- [fstype](#fstype)
- [pdid](#pdid)*

    
</details>

<details>
<summary>portworx_volume</summary><blockquote>

    
- [fstype](#fstype)
- [read_only](#read_only)
- [volume_id](#volume_id)*

    
</details>

<details>
<summary>projected</summary><blockquote>

    
- [default_mode](#default_mode)

    
<details>
<summary>sources</summary><blockquote>

    

    
<details>
<summary>config_map</summary><blockquote>

    
- [name](#name)
- [optional](#optional)

    
<details>
<summary>items</summary><blockquote>

    
- [key](#key)*
- [mode](#mode)
- [path](#path)*

    
</details>

</details>

<details>
<summary>downward_api</summary><blockquote>

    

    
<details>
<summary>items</summary><blockquote>

    
- [mode](#mode)
- [path](#path)*

    
<details>
<summary>field_ref</summary><blockquote>

    
- [api_version](#api_version)
- [field_path](#field_path)*

    
</details>

<details>
<summary>resource_field_ref</summary><blockquote>

    
- [container_name](#container_name)
- [divisor](#divisor)
- [resource](#resource)*

    
</details>

</details>

</details>

<details>
<summary>secret</summary><blockquote>

    
- [name](#name)
- [optional](#optional)

    
<details>
<summary>items</summary><blockquote>

    
- [key](#key)*
- [mode](#mode)
- [path](#path)*

    
</details>

</details>

<details>
<summary>service_account_token</summary><blockquote>

    
- [audience](#audience)
- [expiration_seconds](#expiration_seconds)
- [path](#path)*

    
</details>

</details>

</details>

<details>
<summary>quobyte</summary><blockquote>

    
- [group](#group)
- [read_only](#read_only)
- [registry](#registry)*
- [tenant](#tenant)
- [user](#user)
- [volume](#volume)*

    
</details>

<details>
<summary>rbd</summary><blockquote>

    
- [fstype](#fstype)
- [image](#image)*
- [keyring](#keyring)
- [monitors](#monitors)*
- [pool](#pool)
- [read_only](#read_only)
- [user](#user)

    
<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>scale_io</summary><blockquote>

    
- [fstype](#fstype)
- [gateway](#gateway)*
- [protection_domain](#protection_domain)
- [read_only](#read_only)
- [ssl_enabled](#ssl_enabled)
- [storage_mode](#storage_mode)
- [storage_pool](#storage_pool)
- [system](#system)*
- [volume_name](#volume_name)

    
<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>secret</summary><blockquote>

    
- [default_mode](#default_mode)
- [optional](#optional)
- [secret_name](#secret_name)

    
<details>
<summary>items</summary><blockquote>

    
- [key](#key)*
- [mode](#mode)
- [path](#path)*

    
</details>

</details>

<details>
<summary>storageos</summary><blockquote>

    
- [fstype](#fstype)
- [read_only](#read_only)
- [volume_name](#volume_name)
- [volume_namespace](#volume_namespace)

    
<details>
<summary>secret_ref</summary><blockquote>

    
- [name](#name)

    
</details>

</details>

<details>
<summary>vsphere_volume</summary><blockquote>

    
- [fstype](#fstype)
- [storage_policy_id](#storage_policy_id)
- [storage_policy_name](#storage_policy_name)
- [volume_path](#volume_path)*

    
</details>

</details>

</details>

</details>

</details>


<details>
<summary>example</summary><blockquote>

```hcl
resource "k8s_core_v1_replication_controller" "this" {

  metadata {
    annotations = { "key" = "TypeString" }
    labels      = { "key" = "TypeString" }
    name        = "TypeString"
    namespace   = "TypeString"
  }

  spec {
    min_ready_seconds = "TypeInt"
    replicas          = "TypeInt"
    selector          = { "key" = "TypeString" }

    template {

      metadata {
        annotations = { "key" = "TypeString" }
        labels      = { "key" = "TypeString" }
        name        = "TypeString"
        namespace   = "TypeString"
      }

      spec {
        active_deadline_seconds = "TypeInt"

        affinity {

          node_affinity {

            preferred_during_scheduling_ignored_during_execution {

              preference {

                match_expressions {
                  key      = "TypeString*"
                  operator = "TypeString*"
                  values   = ["TypeString"]
                }

                match_fields {
                  key      = "TypeString*"
                  operator = "TypeString*"
                  values   = ["TypeString"]
                }
              }
              weight = "TypeInt*"
            }

            required_during_scheduling_ignored_during_execution {

              node_selector_terms {

                match_expressions {
                  key      = "TypeString*"
                  operator = "TypeString*"
                  values   = ["TypeString"]
                }

                match_fields {
                  key      = "TypeString*"
                  operator = "TypeString*"
                  values   = ["TypeString"]
                }
              }
            }
          }

          pod_affinity {

            preferred_during_scheduling_ignored_during_execution {

              pod_affinity_term {

                label_selector {

                  match_expressions {
                    key      = "TypeString*"
                    operator = "TypeString*"
                    values   = ["TypeString"]
                  }
                  match_labels = { "key" = "TypeString" }
                }
                namespaces   = ["TypeString"]
                topology_key = "TypeString*"
              }
              weight = "TypeInt*"
            }

            required_during_scheduling_ignored_during_execution {

              label_selector {

                match_expressions {
                  key      = "TypeString*"
                  operator = "TypeString*"
                  values   = ["TypeString"]
                }
                match_labels = { "key" = "TypeString" }
              }
              namespaces   = ["TypeString"]
              topology_key = "TypeString*"
            }
          }

          pod_anti_affinity {

            preferred_during_scheduling_ignored_during_execution {

              pod_affinity_term {

                label_selector {

                  match_expressions {
                    key      = "TypeString*"
                    operator = "TypeString*"
                    values   = ["TypeString"]
                  }
                  match_labels = { "key" = "TypeString" }
                }
                namespaces   = ["TypeString"]
                topology_key = "TypeString*"
              }
              weight = "TypeInt*"
            }

            required_during_scheduling_ignored_during_execution {

              label_selector {

                match_expressions {
                  key      = "TypeString*"
                  operator = "TypeString*"
                  values   = ["TypeString"]
                }
                match_labels = { "key" = "TypeString" }
              }
              namespaces   = ["TypeString"]
              topology_key = "TypeString*"
            }
          }
        }
        automount_service_account_token = "TypeBool"

        containers {
          args    = ["TypeString"]
          command = ["TypeString"]

          env {
            name  = "TypeString*"
            value = "TypeString"

            value_from {

              config_map_keyref {
                key      = "TypeString*"
                name     = "TypeString"
                optional = "TypeBool"
              }

              field_ref {
                api_version = "TypeString"
                field_path  = "TypeString*"
              }

              resource_field_ref {
                container_name = "TypeString"
                divisor        = "TypeString"
                resource       = "TypeString*"
              }

              secret_key_ref {
                key      = "TypeString*"
                name     = "TypeString"
                optional = "TypeBool"
              }
            }
          }

          env_from {

            config_map_ref {
              name     = "TypeString"
              optional = "TypeBool"
            }
            prefix = "TypeString"

            secret_ref {
              name     = "TypeString"
              optional = "TypeBool"
            }
          }
          image             = "TypeString"
          image_pull_policy = "TypeString"

          lifecycle {

            post_start {

              exec {
                command = ["TypeString"]
              }

              http_get {
                host = "TypeString"

                http_headers {
                  name  = "TypeString*"
                  value = "TypeString*"
                }
                path   = "TypeString"
                port   = "TypeString*"
                scheme = "TypeString"
              }

              tcp_socket {
                host = "TypeString"
                port = "TypeString*"
              }
            }

            pre_stop {

              exec {
                command = ["TypeString"]
              }

              http_get {
                host = "TypeString"

                http_headers {
                  name  = "TypeString*"
                  value = "TypeString*"
                }
                path   = "TypeString"
                port   = "TypeString*"
                scheme = "TypeString"
              }

              tcp_socket {
                host = "TypeString"
                port = "TypeString*"
              }
            }
          }

          liveness_probe {

            exec {
              command = ["TypeString"]
            }
            failure_threshold = "TypeInt"

            http_get {
              host = "TypeString"

              http_headers {
                name  = "TypeString*"
                value = "TypeString*"
              }
              path   = "TypeString"
              port   = "TypeString*"
              scheme = "TypeString"
            }
            initial_delay_seconds = "TypeInt"
            period_seconds        = "TypeInt"
            success_threshold     = "TypeInt"

            tcp_socket {
              host = "TypeString"
              port = "TypeString*"
            }
            timeout_seconds = "TypeInt"
          }
          name = "TypeString*"

          ports {
            container_port = "TypeInt*"
            host_ip        = "TypeString"
            host_port      = "TypeInt"
            name           = "TypeString"
            protocol       = "TypeString"
          }

          readiness_probe {

            exec {
              command = ["TypeString"]
            }
            failure_threshold = "TypeInt"

            http_get {
              host = "TypeString"

              http_headers {
                name  = "TypeString*"
                value = "TypeString*"
              }
              path   = "TypeString"
              port   = "TypeString*"
              scheme = "TypeString"
            }
            initial_delay_seconds = "TypeInt"
            period_seconds        = "TypeInt"
            success_threshold     = "TypeInt"

            tcp_socket {
              host = "TypeString"
              port = "TypeString*"
            }
            timeout_seconds = "TypeInt"
          }

          resources {
            limits   = { "key" = "TypeString" }
            requests = { "key" = "TypeString" }
          }

          security_context {
            allow_privilege_escalation = "TypeBool"

            capabilities {
              add  = ["TypeString"]
              drop = ["TypeString"]
            }
            privileged                = "TypeBool"
            proc_mount                = "TypeString"
            read_only_root_filesystem = "TypeBool"
            run_asgroup               = "TypeInt"
            run_asnon_root            = "TypeBool"
            run_asuser                = "TypeInt"

            selinux_options {
              level = "TypeString"
              role  = "TypeString"
              type  = "TypeString"
              user  = "TypeString"
            }
          }
          stdin                      = "TypeBool"
          stdin_once                 = "TypeBool"
          termination_message_path   = "TypeString"
          termination_message_policy = "TypeString"
          tty                        = "TypeBool"

          volume_devices {
            device_path = "TypeString*"
            name        = "TypeString*"
          }

          volume_mounts {
            mount_path        = "TypeString*"
            mount_propagation = "TypeString"
            name              = "TypeString*"
            read_only         = "TypeBool"
            sub_path          = "TypeString"
            sub_path_expr     = "TypeString"
          }
          working_dir = "TypeString"
        }

        dns_config {
          nameservers = ["TypeString"]

          options {
            name  = "TypeString"
            value = "TypeString"
          }
          searches = ["TypeString"]
        }
        dns_policy           = "TypeString"
        enable_service_links = "TypeBool"

        host_aliases {
          hostnames = ["TypeString"]
          ip        = "TypeString"
        }
        host_ipc     = "TypeBool"
        host_network = "TypeBool"
        host_pid     = "TypeBool"
        hostname     = "TypeString"

        image_pull_secrets {
          name = "TypeString"
        }

        init_containers {
          args    = ["TypeString"]
          command = ["TypeString"]

          env {
            name  = "TypeString*"
            value = "TypeString"

            value_from {

              config_map_keyref {
                key      = "TypeString*"
                name     = "TypeString"
                optional = "TypeBool"
              }

              field_ref {
                api_version = "TypeString"
                field_path  = "TypeString*"
              }

              resource_field_ref {
                container_name = "TypeString"
                divisor        = "TypeString"
                resource       = "TypeString*"
              }

              secret_key_ref {
                key      = "TypeString*"
                name     = "TypeString"
                optional = "TypeBool"
              }
            }
          }

          env_from {

            config_map_ref {
              name     = "TypeString"
              optional = "TypeBool"
            }
            prefix = "TypeString"

            secret_ref {
              name     = "TypeString"
              optional = "TypeBool"
            }
          }
          image             = "TypeString"
          image_pull_policy = "TypeString"

          lifecycle {

            post_start {

              exec {
                command = ["TypeString"]
              }

              http_get {
                host = "TypeString"

                http_headers {
                  name  = "TypeString*"
                  value = "TypeString*"
                }
                path   = "TypeString"
                port   = "TypeString*"
                scheme = "TypeString"
              }

              tcp_socket {
                host = "TypeString"
                port = "TypeString*"
              }
            }

            pre_stop {

              exec {
                command = ["TypeString"]
              }

              http_get {
                host = "TypeString"

                http_headers {
                  name  = "TypeString*"
                  value = "TypeString*"
                }
                path   = "TypeString"
                port   = "TypeString*"
                scheme = "TypeString"
              }

              tcp_socket {
                host = "TypeString"
                port = "TypeString*"
              }
            }
          }

          liveness_probe {

            exec {
              command = ["TypeString"]
            }
            failure_threshold = "TypeInt"

            http_get {
              host = "TypeString"

              http_headers {
                name  = "TypeString*"
                value = "TypeString*"
              }
              path   = "TypeString"
              port   = "TypeString*"
              scheme = "TypeString"
            }
            initial_delay_seconds = "TypeInt"
            period_seconds        = "TypeInt"
            success_threshold     = "TypeInt"

            tcp_socket {
              host = "TypeString"
              port = "TypeString*"
            }
            timeout_seconds = "TypeInt"
          }
          name = "TypeString*"

          ports {
            container_port = "TypeInt*"
            host_ip        = "TypeString"
            host_port      = "TypeInt"
            name           = "TypeString"
            protocol       = "TypeString"
          }

          readiness_probe {

            exec {
              command = ["TypeString"]
            }
            failure_threshold = "TypeInt"

            http_get {
              host = "TypeString"

              http_headers {
                name  = "TypeString*"
                value = "TypeString*"
              }
              path   = "TypeString"
              port   = "TypeString*"
              scheme = "TypeString"
            }
            initial_delay_seconds = "TypeInt"
            period_seconds        = "TypeInt"
            success_threshold     = "TypeInt"

            tcp_socket {
              host = "TypeString"
              port = "TypeString*"
            }
            timeout_seconds = "TypeInt"
          }

          resources {
            limits   = { "key" = "TypeString" }
            requests = { "key" = "TypeString" }
          }

          security_context {
            allow_privilege_escalation = "TypeBool"

            capabilities {
              add  = ["TypeString"]
              drop = ["TypeString"]
            }
            privileged                = "TypeBool"
            proc_mount                = "TypeString"
            read_only_root_filesystem = "TypeBool"
            run_asgroup               = "TypeInt"
            run_asnon_root            = "TypeBool"
            run_asuser                = "TypeInt"

            selinux_options {
              level = "TypeString"
              role  = "TypeString"
              type  = "TypeString"
              user  = "TypeString"
            }
          }
          stdin                      = "TypeBool"
          stdin_once                 = "TypeBool"
          termination_message_path   = "TypeString"
          termination_message_policy = "TypeString"
          tty                        = "TypeBool"

          volume_devices {
            device_path = "TypeString*"
            name        = "TypeString*"
          }

          volume_mounts {
            mount_path        = "TypeString*"
            mount_propagation = "TypeString"
            name              = "TypeString*"
            read_only         = "TypeBool"
            sub_path          = "TypeString"
            sub_path_expr     = "TypeString"
          }
          working_dir = "TypeString"
        }
        node_name           = "TypeString"
        node_selector       = { "key" = "TypeString" }
        priority            = "TypeInt"
        priority_class_name = "TypeString"

        readiness_gates {
          condition_type = "TypeString*"
        }
        restart_policy     = "TypeString"
        runtime_class_name = "TypeString"
        scheduler_name     = "TypeString"

        security_context {
          fsgroup        = "TypeInt"
          run_asgroup    = "TypeInt"
          run_asnon_root = "TypeBool"
          run_asuser     = "TypeInt"

          selinux_options {
            level = "TypeString"
            role  = "TypeString"
            type  = "TypeString"
            user  = "TypeString"
          }
          supplemental_groups = ["TypeInt"]

          sysctls {
            name  = "TypeString*"
            value = "TypeString*"
          }
        }
        service_account                  = "TypeString"
        service_account_name             = "TypeString"
        share_process_namespace          = "TypeBool"
        subdomain                        = "TypeString"
        termination_grace_period_seconds = "TypeInt"

        tolerations {
          effect             = "TypeString"
          key                = "TypeString"
          operator           = "TypeString"
          toleration_seconds = "TypeInt"
          value              = "TypeString"
        }

        volumes {

          aws_elastic_block_store {
            fstype    = "TypeString"
            partition = "TypeInt"
            read_only = "TypeBool"
            volume_id = "TypeString*"
          }

          azure_disk {
            caching_mode = "TypeString"
            disk_name    = "TypeString*"
            disk_uri     = "TypeString*"
            fstype       = "TypeString"
            kind         = "TypeString"
            read_only    = "TypeBool"
          }

          azure_file {
            read_only   = "TypeBool"
            secret_name = "TypeString*"
            share_name  = "TypeString*"
          }

          cephfs {
            monitors    = ["TypeString*"]
            path        = "TypeString"
            read_only   = "TypeBool"
            secret_file = "TypeString"

            secret_ref {
              name = "TypeString"
            }
            user = "TypeString"
          }

          cinder {
            fstype    = "TypeString"
            read_only = "TypeBool"

            secret_ref {
              name = "TypeString"
            }
            volume_id = "TypeString*"
          }

          config_map {
            default_mode = "TypeInt"

            items {
              key  = "TypeString*"
              mode = "TypeInt"
              path = "TypeString*"
            }
            name     = "TypeString"
            optional = "TypeBool"
          }

          csi {
            driver = "TypeString*"
            fstype = "TypeString"

            node_publish_secret_ref {
              name = "TypeString"
            }
            read_only         = "TypeBool"
            volume_attributes = { "key" = "TypeString" }
          }

          downward_api {
            default_mode = "TypeInt"

            items {

              field_ref {
                api_version = "TypeString"
                field_path  = "TypeString*"
              }
              mode = "TypeInt"
              path = "TypeString*"

              resource_field_ref {
                container_name = "TypeString"
                divisor        = "TypeString"
                resource       = "TypeString*"
              }
            }
          }

          empty_dir {
            medium     = "TypeString"
            size_limit = "TypeString"
          }

          fc {
            fstype      = "TypeString"
            lun         = "TypeInt"
            read_only   = "TypeBool"
            target_wwns = ["TypeString"]
            wwids       = ["TypeString"]
          }

          flex_volume {
            driver    = "TypeString*"
            fstype    = "TypeString"
            options   = { "key" = "TypeString" }
            read_only = "TypeBool"

            secret_ref {
              name = "TypeString"
            }
          }

          flocker {
            dataset_name = "TypeString"
            dataset_uuid = "TypeString"
          }

          gce_persistent_disk {
            fstype    = "TypeString"
            partition = "TypeInt"
            pdname    = "TypeString*"
            read_only = "TypeBool"
          }

          git_repo {
            directory  = "TypeString"
            repository = "TypeString*"
            revision   = "TypeString"
          }

          glusterfs {
            endpoints = "TypeString*"
            path      = "TypeString*"
            read_only = "TypeBool"
          }

          host_path {
            path = "TypeString*"
            type = "TypeString"
          }

          iscsi {
            chap_auth_discovery = "TypeBool"
            chap_auth_session   = "TypeBool"
            fstype              = "TypeString"
            initiator_name      = "TypeString"
            iqn                 = "TypeString*"
            iscsi_interface     = "TypeString"
            lun                 = "TypeInt*"
            portals             = ["TypeString"]
            read_only           = "TypeBool"

            secret_ref {
              name = "TypeString"
            }
            target_portal = "TypeString*"
          }
          name = "TypeString*"

          nfs {
            path      = "TypeString*"
            read_only = "TypeBool"
            server    = "TypeString*"
          }

          persistent_volume_claim {
            claim_name = "TypeString*"
            read_only  = "TypeBool"
          }

          photon_persistent_disk {
            fstype = "TypeString"
            pdid   = "TypeString*"
          }

          portworx_volume {
            fstype    = "TypeString"
            read_only = "TypeBool"
            volume_id = "TypeString*"
          }

          projected {
            default_mode = "TypeInt"

            sources {

              config_map {

                items {
                  key  = "TypeString*"
                  mode = "TypeInt"
                  path = "TypeString*"
                }
                name     = "TypeString"
                optional = "TypeBool"
              }

              downward_api {

                items {

                  field_ref {
                    api_version = "TypeString"
                    field_path  = "TypeString*"
                  }
                  mode = "TypeInt"
                  path = "TypeString*"

                  resource_field_ref {
                    container_name = "TypeString"
                    divisor        = "TypeString"
                    resource       = "TypeString*"
                  }
                }
              }

              secret {

                items {
                  key  = "TypeString*"
                  mode = "TypeInt"
                  path = "TypeString*"
                }
                name     = "TypeString"
                optional = "TypeBool"
              }

              service_account_token {
                audience           = "TypeString"
                expiration_seconds = "TypeInt"
                path               = "TypeString*"
              }
            }
          }

          quobyte {
            group     = "TypeString"
            read_only = "TypeBool"
            registry  = "TypeString*"
            tenant    = "TypeString"
            user      = "TypeString"
            volume    = "TypeString*"
          }

          rbd {
            fstype    = "TypeString"
            image     = "TypeString*"
            keyring   = "TypeString"
            monitors  = ["TypeString*"]
            pool      = "TypeString"
            read_only = "TypeBool"

            secret_ref {
              name = "TypeString"
            }
            user = "TypeString"
          }

          scale_io {
            fstype            = "TypeString"
            gateway           = "TypeString*"
            protection_domain = "TypeString"
            read_only         = "TypeBool"

            secret_ref {
              name = "TypeString"
            }
            ssl_enabled  = "TypeBool"
            storage_mode = "TypeString"
            storage_pool = "TypeString"
            system       = "TypeString*"
            volume_name  = "TypeString"
          }

          secret {
            default_mode = "TypeInt"

            items {
              key  = "TypeString*"
              mode = "TypeInt"
              path = "TypeString*"
            }
            optional    = "TypeBool"
            secret_name = "TypeString"
          }

          storageos {
            fstype    = "TypeString"
            read_only = "TypeBool"

            secret_ref {
              name = "TypeString"
            }
            volume_name      = "TypeString"
            volume_namespace = "TypeString"
          }

          vsphere_volume {
            fstype              = "TypeString"
            storage_policy_id   = "TypeString"
            storage_policy_name = "TypeString"
            volume_path         = "TypeString*"
          }
        }
      }
    }
  }
}


```

</details>

  
## metadata

If the Labels of a ReplicationController are empty, they are defaulted to be the same as the Pod(s) that the replication controller manages. Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

    
#### annotations

######  TypeMap

Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
#### creation_timestamp

######  ReadOnly • TypeString

CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.

Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
#### deletion_grace_period_seconds

######  ReadOnly • TypeInt

Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.
#### deletion_timestamp

######  ReadOnly • TypeString

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.

Populated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
#### labels

######  TypeMap

Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
#### name

######  TypeString

Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
#### namespace

######  TypeString

Namespace defines the space within each name must be unique. An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.

Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces
#### self_link

######  ReadOnly • TypeString

SelfLink is a URL representing this object. Populated by the system. Read-only.
#### uid

######  ReadOnly • TypeString

UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.

Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
## spec

Spec defines the specification of the desired behavior of the replication controller. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

    
#### min_ready_seconds

######  TypeInt

Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
#### replicas

######  TypeInt

Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#what-is-a-replicationcontroller
#### selector

######  TypeMap

Selector is a label query over pods that should match the Replicas count. If Selector is empty, it is defaulted to the labels present on the Pod template. Label keys and values that must match in order to be controlled by this replication controller, if empty defaulted to labels on Pod template. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
## template

Template is the object that describes the pod that will be created if insufficient replicas are detected. This takes precedence over a TemplateRef. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicationcontroller#pod-template

    
## metadata

Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata

    
#### annotations

######  TypeMap

Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations
#### creation_timestamp

######  ReadOnly • TypeString

CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC.

Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
#### deletion_grace_period_seconds

######  ReadOnly • TypeInt

Number of seconds allowed for this object to gracefully terminate before it will be removed from the system. Only set when deletionTimestamp is also set. May only be shortened. Read-only.
#### deletion_timestamp

######  ReadOnly • TypeString

DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This field is set by the server when a graceful deletion is requested by the user, and is not directly settable by a client. The resource is expected to be deleted (no longer visible from resource lists, and not reachable by name) after the time in this field, once the finalizers list is empty. As long as the finalizers list contains items, deletion is blocked. Once the deletionTimestamp is set, this value may not be unset or be set further into the future, although it may be shortened or the resource may be deleted prior to this time. For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react by sending a graceful termination signal to the containers in the pod. After that 30 seconds, the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup, remove the pod from the API. In the presence of network partitions, this object may still exist after this timestamp, until an administrator or automated process can determine the resource is fully terminated. If not set, graceful deletion of the object has not been requested.

Populated by the system when a graceful deletion is requested. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
#### labels

######  TypeMap

Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels
#### name

######  TypeString

Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names
#### namespace

######  TypeString

Namespace defines the space within each name must be unique. An empty namespace is equivalent to the "default" namespace, but "default" is the canonical representation. Not all objects are required to be scoped to a namespace - the value of this field for those objects will be empty.

Must be a DNS_LABEL. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/namespaces
#### self_link

######  ReadOnly • TypeString

SelfLink is a URL representing this object. Populated by the system. Read-only.
#### uid

######  ReadOnly • TypeString

UID is the unique in time and space value for this object. It is typically generated by the server on successful creation of a resource and is not allowed to change on PUT operations.

Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/identifiers#uids
## spec

Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status

    
#### active_deadline_seconds

######  TypeInt

Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.
## affinity

If specified, the pod's scheduling constraints

    
## node_affinity

Describes node affinity scheduling rules for the pod.

    
## preferred_during_scheduling_ignored_during_execution

The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.

    
## preference

A node selector term, associated with the corresponding weight.

    
## match_expressions

A list of node selector requirements by node's labels.

    
#### key

###### Required •  TypeString

The label key that the selector applies to.
#### operator

###### Required •  TypeString

Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
#### values

######  TypeList

An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
## match_fields

A list of node selector requirements by node's fields.

    
#### key

###### Required •  TypeString

The label key that the selector applies to.
#### operator

###### Required •  TypeString

Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
#### values

######  TypeList

An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
#### weight

###### Required •  TypeInt

Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
## required_during_scheduling_ignored_during_execution

If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.

    
## node_selector_terms

Required. A list of node selector terms. The terms are ORed.

    
## match_expressions

A list of node selector requirements by node's labels.

    
#### key

###### Required •  TypeString

The label key that the selector applies to.
#### operator

###### Required •  TypeString

Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
#### values

######  TypeList

An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
## match_fields

A list of node selector requirements by node's fields.

    
#### key

###### Required •  TypeString

The label key that the selector applies to.
#### operator

###### Required •  TypeString

Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
#### values

######  TypeList

An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
## pod_affinity

Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).

    
## preferred_during_scheduling_ignored_during_execution

The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.

    
## pod_affinity_term

Required. A pod affinity term, associated with the corresponding weight.

    
## label_selector

A label query over a set of resources, in this case pods.

    
## match_expressions

matchExpressions is a list of label selector requirements. The requirements are ANDed.

    
#### key

###### Required •  TypeString

key is the label key that the selector applies to.
#### operator

###### Required •  TypeString

operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
#### values

######  TypeList

values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
#### match_labels

######  TypeMap

matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
#### namespaces

######  TypeList

namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means "this pod's namespace"
#### topology_key

###### Required •  TypeString

This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
#### weight

###### Required •  TypeInt

weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
## required_during_scheduling_ignored_during_execution

If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.

    
## label_selector

A label query over a set of resources, in this case pods.

    
## match_expressions

matchExpressions is a list of label selector requirements. The requirements are ANDed.

    
#### key

###### Required •  TypeString

key is the label key that the selector applies to.
#### operator

###### Required •  TypeString

operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
#### values

######  TypeList

values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
#### match_labels

######  TypeMap

matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
#### namespaces

######  TypeList

namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means "this pod's namespace"
#### topology_key

###### Required •  TypeString

This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
## pod_anti_affinity

Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).

    
## preferred_during_scheduling_ignored_during_execution

The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding "weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.

    
## pod_affinity_term

Required. A pod affinity term, associated with the corresponding weight.

    
## label_selector

A label query over a set of resources, in this case pods.

    
## match_expressions

matchExpressions is a list of label selector requirements. The requirements are ANDed.

    
#### key

###### Required •  TypeString

key is the label key that the selector applies to.
#### operator

###### Required •  TypeString

operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
#### values

######  TypeList

values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
#### match_labels

######  TypeMap

matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
#### namespaces

######  TypeList

namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means "this pod's namespace"
#### topology_key

###### Required •  TypeString

This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
#### weight

###### Required •  TypeInt

weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
## required_during_scheduling_ignored_during_execution

If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.

    
## label_selector

A label query over a set of resources, in this case pods.

    
## match_expressions

matchExpressions is a list of label selector requirements. The requirements are ANDed.

    
#### key

###### Required •  TypeString

key is the label key that the selector applies to.
#### operator

###### Required •  TypeString

operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
#### values

######  TypeList

values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
#### match_labels

######  TypeMap

matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
#### namespaces

######  TypeList

namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means "this pod's namespace"
#### topology_key

###### Required •  TypeString

This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.
#### automount_service_account_token

######  TypeBool

AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.
## containers

List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.

    
#### args

######  TypeList

Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
#### command

######  TypeList

Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
## env

List of environment variables to set in the container. Cannot be updated.

    
#### name

###### Required •  TypeString

Name of the environment variable. Must be a C_IDENTIFIER.
#### value

######  TypeString

Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
## value_from

Source for the environment variable's value. Cannot be used if value is not empty.

    
## config_map_keyref

Selects a key of a ConfigMap.

    
#### key

###### Required •  TypeString

The key to select.
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the ConfigMap or it's key must be defined
## field_ref

Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.

    
#### api_version

######  TypeString

Version of the schema the FieldPath is written in terms of, defaults to "v1".
#### field_path

###### Required •  TypeString

Path of the field to select in the specified API version.
## resource_field_ref

Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.

    
#### container_name

######  TypeString

Container name: required for volumes, optional for env vars
#### divisor

######  TypeString

Specifies the output format of the exposed resources, defaults to "1"
#### resource

###### Required •  TypeString

Required: resource to select
## secret_key_ref

Selects a key of a secret in the pod's namespace

    
#### key

###### Required •  TypeString

The key of the secret to select from.  Must be a valid secret key.
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the Secret or it's key must be defined
## env_from

List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.

    
## config_map_ref

The ConfigMap to select from

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the ConfigMap must be defined
#### prefix

######  TypeString

An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
## secret_ref

The Secret to select from

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the Secret must be defined
#### image

######  TypeString

Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
#### image_pull_policy

######  TypeString

Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
## lifecycle

Actions that the management system should take in response to container lifecycle events. Cannot be updated.

    
## post_start

PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
## pre_stop

PreStop is called immediately before a container is terminated due to an API request or management event such as liveness probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
## liveness_probe

Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
#### failure_threshold

######  TypeInt

Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
#### initial_delay_seconds

######  TypeInt

Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
#### period_seconds

######  TypeInt

How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
#### success_threshold

######  TypeInt

Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### timeout_seconds

######  TypeInt

Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
#### name

###### Required •  TypeString

Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
## ports

List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.

    
#### container_port

###### Required •  TypeInt

Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
#### host_ip

######  TypeString

What host IP to bind the external port to.
#### host_port

######  TypeInt

Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
#### name

######  TypeString

If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
#### protocol

######  TypeString

Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
## readiness_probe

Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
#### failure_threshold

######  TypeInt

Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
#### initial_delay_seconds

######  TypeInt

Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
#### period_seconds

######  TypeInt

How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
#### success_threshold

######  TypeInt

Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### timeout_seconds

######  TypeInt

Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
## resources

Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

    
#### limits

######  TypeMap

Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
#### requests

######  TypeMap

Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
## security_context

Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/

    
#### allow_privilege_escalation

######  TypeBool

AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
## capabilities

The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.

    
#### add

######  TypeList

Added capabilities
#### drop

######  TypeList

Removed capabilities
#### privileged

######  TypeBool

Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
#### proc_mount

######  TypeString

procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.
#### read_only_root_filesystem

######  TypeBool

Whether this container has a read-only root filesystem. Default is false.
#### run_asgroup

######  TypeInt

The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
#### run_asnon_root

######  TypeBool

Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
#### run_asuser

######  TypeInt

The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
## selinux_options

The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.

    
#### level

######  TypeString

Level is SELinux level label that applies to the container.
#### role

######  TypeString

Role is a SELinux role label that applies to the container.
#### type

######  TypeString

Type is a SELinux type label that applies to the container.
#### user

######  TypeString

User is a SELinux user label that applies to the container.
#### stdin

######  TypeBool

Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
#### stdin_once

######  TypeBool

Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
#### termination_message_path

######  TypeString

Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
#### termination_message_policy

######  TypeString

Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
#### tty

######  TypeBool

Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
## volume_devices

volumeDevices is the list of block devices to be used by the container. This is a beta feature.

    
#### device_path

###### Required •  TypeString

devicePath is the path inside of the container that the device will be mapped to.
#### name

###### Required •  TypeString

name must match the name of a persistentVolumeClaim in the pod
## volume_mounts

Pod volumes to mount into the container's filesystem. Cannot be updated.

    
#### mount_path

###### Required •  TypeString

Path within the container at which the volume should be mounted.  Must not contain ':'.
#### mount_propagation

######  TypeString

mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.
#### name

###### Required •  TypeString

This must match the Name of a Volume.
#### read_only

######  TypeBool

Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
#### sub_path

######  TypeString

Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
#### sub_path_expr

######  TypeString

Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to "" (volume's root). SubPathExpr and SubPath are mutually exclusive. This field is alpha in 1.14.
#### working_dir

######  TypeString

Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
## dns_config

Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.

    
#### nameservers

######  TypeList

A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
## options

A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.

    
#### name

######  TypeString

Required.
#### value

######  TypeString


#### searches

######  TypeList

A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
#### dns_policy

######  TypeString

Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
#### enable_service_links

######  TypeBool

EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.
## host_aliases

HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.

    
#### hostnames

######  TypeList

Hostnames for the above IP address.
#### ip

######  TypeString

IP address of the host file entry.
#### host_ipc

######  TypeBool

Use the host's ipc namespace. Optional: Default to false.
#### host_network

######  TypeBool

Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.
#### host_pid

######  TypeBool

Use the host's pid namespace. Optional: Default to false.
#### hostname

######  TypeString

Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.
## image_pull_secrets

ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
## init_containers

List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, or Liveness probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/

    
#### args

######  TypeList

Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
#### command

######  TypeList

Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
## env

List of environment variables to set in the container. Cannot be updated.

    
#### name

###### Required •  TypeString

Name of the environment variable. Must be a C_IDENTIFIER.
#### value

######  TypeString

Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "".
## value_from

Source for the environment variable's value. Cannot be used if value is not empty.

    
## config_map_keyref

Selects a key of a ConfigMap.

    
#### key

###### Required •  TypeString

The key to select.
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the ConfigMap or it's key must be defined
## field_ref

Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.

    
#### api_version

######  TypeString

Version of the schema the FieldPath is written in terms of, defaults to "v1".
#### field_path

###### Required •  TypeString

Path of the field to select in the specified API version.
## resource_field_ref

Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.

    
#### container_name

######  TypeString

Container name: required for volumes, optional for env vars
#### divisor

######  TypeString

Specifies the output format of the exposed resources, defaults to "1"
#### resource

###### Required •  TypeString

Required: resource to select
## secret_key_ref

Selects a key of a secret in the pod's namespace

    
#### key

###### Required •  TypeString

The key of the secret to select from.  Must be a valid secret key.
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the Secret or it's key must be defined
## env_from

List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.

    
## config_map_ref

The ConfigMap to select from

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the ConfigMap must be defined
#### prefix

######  TypeString

An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
## secret_ref

The Secret to select from

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the Secret must be defined
#### image

######  TypeString

Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.
#### image_pull_policy

######  TypeString

Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
## lifecycle

Actions that the management system should take in response to container lifecycle events. Cannot be updated.

    
## post_start

PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
## pre_stop

PreStop is called immediately before a container is terminated due to an API request or management event such as liveness probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
## liveness_probe

Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
#### failure_threshold

######  TypeInt

Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
#### initial_delay_seconds

######  TypeInt

Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
#### period_seconds

######  TypeInt

How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
#### success_threshold

######  TypeInt

Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### timeout_seconds

######  TypeInt

Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
#### name

###### Required •  TypeString

Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
## ports

List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default "0.0.0.0" address inside a container will be accessible from the network. Cannot be updated.

    
#### container_port

###### Required •  TypeInt

Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
#### host_ip

######  TypeString

What host IP to bind the external port to.
#### host_port

######  TypeInt

Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
#### name

######  TypeString

If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
#### protocol

######  TypeString

Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
## readiness_probe

Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes

    
## exec

One and only one of the following should be specified. Exec specifies the action to take.

    
#### command

######  TypeList

Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
#### failure_threshold

######  TypeInt

Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
## http_get

HTTPGet specifies the http request to perform.

    
#### host

######  TypeString

Host name to connect to, defaults to the pod IP. You probably want to set "Host" in httpHeaders instead.
## http_headers

Custom headers to set in the request. HTTP allows repeated headers.

    
#### name

###### Required •  TypeString

The header field name
#### value

###### Required •  TypeString

The header field value
#### path

######  TypeString

Path to access on the HTTP server.
#### port

###### Required •  TypeString

Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### scheme

######  TypeString

Scheme to use for connecting to the host. Defaults to HTTP.
#### initial_delay_seconds

######  TypeInt

Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
#### period_seconds

######  TypeInt

How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
#### success_threshold

######  TypeInt

Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
## tcp_socket

TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

    
#### host

######  TypeString

Optional: Host name to connect to, defaults to the pod IP.
#### port

###### Required •  TypeString

Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
#### timeout_seconds

######  TypeInt

Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
## resources

Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/

    
#### limits

######  TypeMap

Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
#### requests

######  TypeMap

Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
## security_context

Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/

    
#### allow_privilege_escalation

######  TypeBool

AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN
## capabilities

The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.

    
#### add

######  TypeList

Added capabilities
#### drop

######  TypeList

Removed capabilities
#### privileged

######  TypeBool

Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.
#### proc_mount

######  TypeString

procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.
#### read_only_root_filesystem

######  TypeBool

Whether this container has a read-only root filesystem. Default is false.
#### run_asgroup

######  TypeInt

The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
#### run_asnon_root

######  TypeBool

Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
#### run_asuser

######  TypeInt

The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
## selinux_options

The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.

    
#### level

######  TypeString

Level is SELinux level label that applies to the container.
#### role

######  TypeString

Role is a SELinux role label that applies to the container.
#### type

######  TypeString

Type is a SELinux type label that applies to the container.
#### user

######  TypeString

User is a SELinux user label that applies to the container.
#### stdin

######  TypeBool

Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
#### stdin_once

######  TypeBool

Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
#### termination_message_path

######  TypeString

Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
#### termination_message_policy

######  TypeString

Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
#### tty

######  TypeBool

Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
## volume_devices

volumeDevices is the list of block devices to be used by the container. This is a beta feature.

    
#### device_path

###### Required •  TypeString

devicePath is the path inside of the container that the device will be mapped to.
#### name

###### Required •  TypeString

name must match the name of a persistentVolumeClaim in the pod
## volume_mounts

Pod volumes to mount into the container's filesystem. Cannot be updated.

    
#### mount_path

###### Required •  TypeString

Path within the container at which the volume should be mounted.  Must not contain ':'.
#### mount_propagation

######  TypeString

mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.
#### name

###### Required •  TypeString

This must match the Name of a Volume.
#### read_only

######  TypeBool

Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.
#### sub_path

######  TypeString

Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
#### sub_path_expr

######  TypeString

Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to "" (volume's root). SubPathExpr and SubPath are mutually exclusive. This field is alpha in 1.14.
#### working_dir

######  TypeString

Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
#### node_name

######  TypeString

NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.
#### node_selector

######  TypeMap

NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
#### priority

######  TypeInt

The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.
#### priority_class_name

######  TypeString

If specified, indicates the pod's priority. "system-node-critical" and "system-cluster-critical" are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.
## readiness_gates

If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to "True" More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md

    
#### condition_type

###### Required •  TypeString

ConditionType refers to a condition in the pod's condition list with matching type.
#### restart_policy

######  TypeString

Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy
#### runtime_class_name

######  TypeString

RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is an alpha feature and may change in the future.
#### scheduler_name

######  TypeString

If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.
## security_context

SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.

    
#### fsgroup

######  TypeInt

A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:

1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----

If unset, the Kubelet will not modify the ownership and permissions of any volume.
#### run_asgroup

######  TypeInt

The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
#### run_asnon_root

######  TypeBool

Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
#### run_asuser

######  TypeInt

The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
## selinux_options

The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.

    
#### level

######  TypeString

Level is SELinux level label that applies to the container.
#### role

######  TypeString

Role is a SELinux role label that applies to the container.
#### type

######  TypeString

Type is a SELinux type label that applies to the container.
#### user

######  TypeString

User is a SELinux user label that applies to the container.
#### supplemental_groups

######  TypeList

A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container.
## sysctls

Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch.

    
#### name

###### Required •  TypeString

Name of a property to set
#### value

###### Required •  TypeString

Value of a property to set
#### service_account

######  TypeString

DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.
#### service_account_name

######  TypeString

ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
#### share_process_namespace

######  TypeBool

Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false. This field is beta-level and may be disabled with the PodShareProcessNamespace feature.
#### subdomain

######  TypeString

If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the pod will not have a domainname at all.
#### termination_grace_period_seconds

######  TypeInt

Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.
## tolerations

If specified, the pod's tolerations.

    
#### effect

######  TypeString

Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
#### key

######  TypeString

Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.
#### operator

######  TypeString

Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.
#### toleration_seconds

######  TypeInt

TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.
#### value

######  TypeString

Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.
## volumes

List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes

    
## aws_elastic_block_store

AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore

    
#### fstype

######  TypeString

Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
#### partition

######  TypeInt

The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
#### read_only

######  TypeBool

Specify "true" to force and set the ReadOnly property in VolumeMounts to "true". If omitted, the default is "false". More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
#### volume_id

###### Required •  TypeString

Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
## azure_disk

AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.

    
#### caching_mode

######  TypeString

Host Caching mode: None, Read Only, Read Write.
#### disk_name

###### Required •  TypeString

The Name of the data disk in the blob storage
#### disk_uri

###### Required •  TypeString

The URI the data disk in the blob storage
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
#### kind

######  TypeString

Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
#### read_only

######  TypeBool

Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
## azure_file

AzureFile represents an Azure File Service mount on the host and bind mount to the pod.

    
#### read_only

######  TypeBool

Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
#### secret_name

###### Required •  TypeString

the name of secret that contains Azure Storage Account Name and Key
#### share_name

###### Required •  TypeString

Share Name
## cephfs

CephFS represents a Ceph FS mount on the host that shares a pod's lifetime

    
#### monitors

###### Required •  TypeList

Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
#### path

######  TypeString

Optional: Used as the mounted root, rather than the full Ceph tree, default is /
#### read_only

######  TypeBool

Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
#### secret_file

######  TypeString

Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
## secret_ref

Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### user

######  TypeString

Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
## cinder

Cinder represents a cinder volume attached and mounted on kubelets host machine More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md

    
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
#### read_only

######  TypeBool

Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
## secret_ref

Optional: points to a secret object containing parameters used to connect to OpenStack.

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### volume_id

###### Required •  TypeString

volume id used to identify the volume in cinder More info: https://releases.k8s.io/HEAD/examples/mysql-cinder-pd/README.md
## config_map

ConfigMap represents a configMap that should populate this volume

    
#### default_mode

######  TypeInt

Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
## items

If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.

    
#### key

###### Required •  TypeString

The key to project.
#### mode

######  TypeInt

Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
#### path

###### Required •  TypeString

The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the ConfigMap or it's keys must be defined
## csi

CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).

    
#### driver

###### Required •  TypeString

Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.
#### fstype

######  TypeString

Filesystem type to mount. Ex. "ext4", "xfs", "ntfs". If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.
## node_publish_secret_ref

NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### read_only

######  TypeBool

Specifies a read-only configuration for the volume. Defaults to false (read/write).
#### volume_attributes

######  TypeMap

VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.
## downward_api

DownwardAPI represents downward API about the pod that should populate this volume

    
#### default_mode

######  TypeInt

Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
## items

Items is a list of downward API volume file

    
## field_ref

Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.

    
#### api_version

######  TypeString

Version of the schema the FieldPath is written in terms of, defaults to "v1".
#### field_path

###### Required •  TypeString

Path of the field to select in the specified API version.
#### mode

######  TypeInt

Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
#### path

###### Required •  TypeString

Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
## resource_field_ref

Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

    
#### container_name

######  TypeString

Container name: required for volumes, optional for env vars
#### divisor

######  TypeString

Specifies the output format of the exposed resources, defaults to "1"
#### resource

###### Required •  TypeString

Required: resource to select
## empty_dir

EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir

    
#### medium

######  TypeString

What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
#### size_limit

######  TypeString

Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
## fc

FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.

    
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
#### lun

######  TypeInt

Optional: FC target lun number
#### read_only

######  TypeBool

Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
#### target_wwns

######  TypeList

Optional: FC target worldwide names (WWNs)
#### wwids

######  TypeList

Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
## flex_volume

FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.

    
#### driver

###### Required •  TypeString

Driver is the name of the driver to use for this volume.
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
#### options

######  TypeMap

Optional: Extra command options if any.
#### read_only

######  TypeBool

Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
## secret_ref

Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
## flocker

Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running

    
#### dataset_name

######  TypeString

Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
#### dataset_uuid

######  TypeString

UUID of the dataset. This is unique identifier of a Flocker dataset
## gce_persistent_disk

GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk

    
#### fstype

######  TypeString

Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
#### partition

######  TypeInt

The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as "1". Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
#### pdname

###### Required •  TypeString

Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
#### read_only

######  TypeBool

ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
## git_repo

GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.

    
#### directory

######  TypeString

Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.
#### repository

###### Required •  TypeString

Repository URL
#### revision

######  TypeString

Commit hash for the specified revision.
## glusterfs

Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md

    
#### endpoints

###### Required •  TypeString

EndpointsName is the endpoint name that details Glusterfs topology. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
#### path

###### Required •  TypeString

Path is the Glusterfs volume path. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
#### read_only

######  TypeBool

ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/glusterfs/README.md#create-a-pod
## host_path

HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath

    
#### path

###### Required •  TypeString

Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
#### type

######  TypeString

Type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
## iscsi

ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://releases.k8s.io/HEAD/examples/volumes/iscsi/README.md

    
#### chap_auth_discovery

######  TypeBool

whether support iSCSI Discovery CHAP authentication
#### chap_auth_session

######  TypeBool

whether support iSCSI Session CHAP authentication
#### fstype

######  TypeString

Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
#### initiator_name

######  TypeString

Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
#### iqn

###### Required •  TypeString

Target iSCSI Qualified Name.
#### iscsi_interface

######  TypeString

iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
#### lun

###### Required •  TypeInt

iSCSI Target Lun number.
#### portals

######  TypeList

iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
#### read_only

######  TypeBool

ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
## secret_ref

CHAP Secret for iSCSI target and initiator authentication

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### target_portal

###### Required •  TypeString

iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
#### name

###### Required •  TypeString

Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
## nfs

NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs

    
#### path

###### Required •  TypeString

Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
#### read_only

######  TypeBool

ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
#### server

###### Required •  TypeString

Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
## persistent_volume_claim

PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims

    
#### claim_name

###### Required •  TypeString

ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
#### read_only

######  TypeBool

Will force the ReadOnly setting in VolumeMounts. Default false.
## photon_persistent_disk

PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine

    
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
#### pdid

###### Required •  TypeString

ID that identifies Photon Controller persistent disk
## portworx_volume

PortworxVolume represents a portworx volume attached and mounted on kubelets host machine

    
#### fstype

######  TypeString

FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
#### read_only

######  TypeBool

Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
#### volume_id

###### Required •  TypeString

VolumeID uniquely identifies a Portworx volume
## projected

Items for all in one resources secrets, configmaps, and downward API

    
#### default_mode

######  TypeInt

Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
## sources

list of volume projections

    
## config_map

information about the configMap data to project

    
## items

If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.

    
#### key

###### Required •  TypeString

The key to project.
#### mode

######  TypeInt

Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
#### path

###### Required •  TypeString

The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the ConfigMap or it's keys must be defined
## downward_api

information about the downwardAPI data to project

    
## items

Items is a list of DownwardAPIVolume file

    
## field_ref

Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.

    
#### api_version

######  TypeString

Version of the schema the FieldPath is written in terms of, defaults to "v1".
#### field_path

###### Required •  TypeString

Path of the field to select in the specified API version.
#### mode

######  TypeInt

Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
#### path

###### Required •  TypeString

Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
## resource_field_ref

Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.

    
#### container_name

######  TypeString

Container name: required for volumes, optional for env vars
#### divisor

######  TypeString

Specifies the output format of the exposed resources, defaults to "1"
#### resource

###### Required •  TypeString

Required: resource to select
## secret

information about the secret data to project

    
## items

If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.

    
#### key

###### Required •  TypeString

The key to project.
#### mode

######  TypeInt

Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
#### path

###### Required •  TypeString

The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### optional

######  TypeBool

Specify whether the Secret or its key must be defined
## service_account_token

information about the serviceAccountToken data to project

    
#### audience

######  TypeString

Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.
#### expiration_seconds

######  TypeInt

ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.
#### path

###### Required •  TypeString

Path is the path relative to the mount point of the file to project the token into.
## quobyte

Quobyte represents a Quobyte mount on the host that shares a pod's lifetime

    
#### group

######  TypeString

Group to map volume access to Default is no group
#### read_only

######  TypeBool

ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.
#### registry

###### Required •  TypeString

Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes
#### tenant

######  TypeString

Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin
#### user

######  TypeString

User to map volume access to Defaults to serivceaccount user
#### volume

###### Required •  TypeString

Volume is a string that references an already created Quobyte volume by name.
## rbd

RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md

    
#### fstype

######  TypeString

Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
#### image

###### Required •  TypeString

The rados image name. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
#### keyring

######  TypeString

Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
#### monitors

###### Required •  TypeList

A collection of Ceph monitors. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
#### pool

######  TypeString

The rados pool name. Default is rbd. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
#### read_only

######  TypeBool

ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
## secret_ref

SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### user

######  TypeString

The rados user name. Default is admin. More info: https://releases.k8s.io/HEAD/examples/volumes/rbd/README.md#how-to-use-it
## scale_io

ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.

    
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".
#### gateway

###### Required •  TypeString

The host address of the ScaleIO API Gateway.
#### protection_domain

######  TypeString

The name of the ScaleIO Protection Domain for the configured storage.
#### read_only

######  TypeBool

Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
## secret_ref

SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### ssl_enabled

######  TypeBool

Flag to enable/disable SSL communication with Gateway, default false
#### storage_mode

######  TypeString

Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
#### storage_pool

######  TypeString

The ScaleIO Storage Pool associated with the protection domain.
#### system

###### Required •  TypeString

The name of the storage system as configured in ScaleIO.
#### volume_name

######  TypeString

The name of a volume already created in the ScaleIO system that is associated with this volume source.
## secret

Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret

    
#### default_mode

######  TypeInt

Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
## items

If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.

    
#### key

###### Required •  TypeString

The key to project.
#### mode

######  TypeInt

Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
#### path

###### Required •  TypeString

The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.
#### optional

######  TypeBool

Specify whether the Secret or it's keys must be defined
#### secret_name

######  TypeString

Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
## storageos

StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.

    
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
#### read_only

######  TypeBool

Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
## secret_ref

SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.

    
#### name

######  TypeString

Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
#### volume_name

######  TypeString

VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.
#### volume_namespace

######  TypeString

VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to "default" if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.
## vsphere_volume

VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine

    
#### fstype

######  TypeString

Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
#### storage_policy_id

######  TypeString

Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
#### storage_policy_name

######  TypeString

Storage Policy Based Management (SPBM) profile name.
#### volume_path

###### Required •  TypeString

Path that identifies vSphere volume vmdk