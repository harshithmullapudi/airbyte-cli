package airbyte

var SOURCES = "/api/v1/sources"
var SOURCE_DEFINITIONS = "/api/v1/source_definitions"
var WORKSPACES = "/api/v1/workspaces"
var CONNECTIONS = "/api/v1/web_backend/connections"
var JOBS = "/api/v1/jobs"
var DESTINATIONS = "/api/v1/destinations"

var GET_WORKSPACES = WORKSPACES + "/get"

var GET_SOURCES = SOURCES + "/list"
var GET_SOURCE = SOURCES + "/get"
var SOURCE_CHECK_CONNECTION = SOURCES + "/check_connection"

var GET_CONNECTIONS = CONNECTIONS + "/list"
var GET_CONNECTION = CONNECTIONS + "/get"

var GET_JOBS = JOBS + "/list"
var GET_JOB = JOBS + "/get"

var GET_DESTINATIONS = DESTINATIONS + "/list"
var GET_DESTINATION = DESTINATIONS + "/get"

var GET_SOURCE_DEFINITIONS = SOURCE_DEFINITIONS + "/list"
