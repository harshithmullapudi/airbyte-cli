package api

var SOURCES = "/api/v1/sources"
var SOURCE_DEFINITIONS = "/api/v1/source_definitions"
var WORKSPACES = "/api/v1/workspaces"
var CONNECTIONS = "/api/v1/web_backend/connections"
var JOBS = "/api/v1/jobs"
var DESTINATIONS = "/api/v1/destinations"

var GET_WORKSPACES = WORKSPACES + "/list"
var GET_WORKSPACE = WORKSPACES + "/get"

var CREATE_SOURCE = SOURCES + "/create"
var GET_SOURCES = SOURCES + "/list"
var GET_SOURCE = SOURCES + "/get"
var SOURCE_CHECK_CONNECTION = SOURCES + "/check_connection"

var CREATE_CONNECTION = CONNECTIONS + "/create"
var GET_CONNECTIONS = CONNECTIONS + "/list"
var GET_CONNECTION = CONNECTIONS + "/get"

var GET_JOBS = JOBS + "/list"
var GET_JOB = JOBS + "/get"

var CREATE_DESTINATION = DESTINATIONS + "/create"
var GET_DESTINATIONS = DESTINATIONS + "/list"
var GET_DESTINATION = DESTINATIONS + "/get"

var GET_SOURCE_DEFINITIONS = SOURCE_DEFINITIONS + "/list"

var DOWNLOAD_CONFIG = "/api/v1/deployment/export"

var SOURCE_CONNECTION_CHECK = "/api/v1/scheduler/sources/check_connection"
var DESTINATION_CONNECTION_CHECK = "/api/v1/scheduler/destinations/check_connection"
