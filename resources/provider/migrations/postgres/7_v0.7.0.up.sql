-- Autogenerated by migration tool on 2022-02-15 03:23:58
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: network.interfaces
CREATE TABLE IF NOT EXISTS "azure_network_interfaces" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"subscription_id" text,
	"etag" text,
	"extended_location_name" text,
	"extended_location_type" text,
	"id" text,
	"location" text,
	"name" text,
	"dns_settings_applied_dns_servers" text[],
	"dns_settings_dns_servers" text[],
	"dns_settings_internal_dns_name_label" text,
	"dns_settings_internal_domain_name_suffix" text,
	"dns_settings_internal_fqdn" text,
	"dscp_configuration_id" text,
	"enable_accelerated_networking" boolean,
	"enable_ip_forwarding" boolean,
	"hosted_workloads" text[],
	"mac_address" text,
	"migration_phase" text,
	"network_security_group" jsonb,
	"nic_type" text,
	"primary" boolean,
	"private_endpoint" jsonb,
	"private_link_service" jsonb,
	"provisioning_state" text,
	"resource_guid" text,
	"tap_configurations" jsonb,
	"virtual_machine_id" text,
	"tags" jsonb,
	"type" text,
	CONSTRAINT azure_network_interfaces_pk PRIMARY KEY(subscription_id,id),
	UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "azure_network_interface_ip_configurations" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"interface_cq_id" uuid,
	"id" text,
	"name" text,
	"etag" text,
	"type" text,
	"primary" boolean,
	"application_gateway_backend_address_pools" jsonb,
	"application_security_groups" jsonb,
	"load_balancer_backend_address_pools" jsonb,
	"load_balancer_inbound_nat_rules" jsonb,
	"private_ip_address" text,
	"private_ip_address_version" text,
	"private_ip_allocation_method" text,
	"private_link_connection_properties" text,
	"provisioning_state" text,
	"public_ip_address" text,
	"subnet_id" text,
	"virtual_network_taps" jsonb,
	CONSTRAINT azure_network_interface_ip_configurations_pk PRIMARY KEY(interface_cq_id,id),
	UNIQUE(cq_id),
	FOREIGN KEY (interface_cq_id) REFERENCES azure_network_interfaces(cq_id) ON DELETE CASCADE
);
