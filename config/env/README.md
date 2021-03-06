Environment Vars Config
=======================

The default environment variables based config is a simple bridge config with
env variables for all important fields.

Values for `BENTHOS_INPUT` and `BENTHOS_OUTPUT` should be chosen from [here][0]
and [here][1] respectively.

The full list of variables and their default values:

``` sh
BENTHOS_HTTP_ADDRESS = 0.0.0.0:4195
BENTHOS_INPUT        =
BENTHOS_OUTPUT       =
BENTHOS_LOG_LEVEL    = INFO

METRICS_TYPE   = http_server
STATSD_ADDRESS = localhost:9125
STATSD_NETWORK = udp

BENTHOS_MAX_PARTS        = 100
BENTHOS_MIN_PARTS        = 1
BENTHOS_MAX_PART_SIZE    = 1073741824
BENTHOS_INPUT_PROCESSOR  = noop
BENTHOS_OUTPUT_PROCESSOR = noop

AMAZON_INPUT_REGION              = eu-west-1
AMAZON_INPUT_BUCKET              =
AMAZON_INPUT_PREFIX              =
AMAZON_INPUT_DELETE_OBJECTS      = false
AMAZON_INPUT_SQS_URL             =
AMAZON_INPUT_SQS_BODY_PATH       = Records.s3.object.key
AMAZON_INPUT_CREDENTIALS_ID      =
AMAZON_INPUT_CREDENTIALS_SECRET  =
AMAZON_INPUT_CREDENTIALS_TOKEN   =
AMAZON_INPUT_TIMEOUT_S           = 5
AMAZON_OUTPUT_REGION             = eu-west-1
AMAZON_OUTPUT_BUCKET             =
AMAZON_OUTPUT_SQS_URL            =
AMAZON_OUTPUT_CREDENTIALS_ID     =
AMAZON_OUTPUT_CREDENTIALS_SECRET =
AMAZON_OUTPUT_CREDENTIALS_TOKEN  =
AMAZON_OUTPUT_TIMEOUT_S          = 5

AMQP_INPUT_URL            =
AMQP_INPUT_EXCHANGE       = benthos-exchange
AMQP_INPUT_EXCHANGE_TYPE  = direct
AMQP_INPUT_QUEUE          = benthos-stream
AMQP_INPUT_KEY            = benthos-key
AMQP_INPUT_CONSUMER_TAG   = benthos-consumer
AMQP_OUTPUT_URL           =
AMQP_OUTPUT_EXCHANGE      = benthos-exchange
AMQP_OUTPUT_EXCHANGE_TYPE = direct
AMQP_OUTPUT_QUEUE         = benthos-stream
AMQP_OUTPUT_KEY           = benthos-key
AMQP_OUTPUT_CONSUMER_TAG  = benthos-consumer

DYNAMIC_INPUT_PREFIX      =
DYNAMIC_INPUT_TIMEOUT_MS  = 5000
DYNAMIC_OUTPUT_PREFIX     =
DYNAMIC_OUTPUT_TIMEOUT_MS = 5000

FILE_INPUT_PATH        =
FILE_INPUT_MULTIPART   = false
FILE_INPUT_MAX_BUFFER  = 65536
FILE_OUTPUT_PATH       =
FILE_OUTPUT_MULTIPART  = false
FILE_OUTPUT_MAX_BUFFER = 65536

HTTP_CLIENT_INPUT_URL                  =
HTTP_CLIENT_INPUT_VERB                 = GET
HTTP_CLIENT_INPUT_PAYLOAD              =
HTTP_CLIENT_INPUT_CONTENT_TYPE         = application/octet-stream
HTTP_CLIENT_INPUT_STREAM               = false
HTTP_CLIENT_INPUT_STREAM_MULTIPART     = false
HTTP_CLIENT_INPUT_STREAM_MAX_BUFFER    = 65536
HTTP_CLIENT_INPUT_STREAM_DELIMITER     =
HTTP_CLIENT_INPUT_OAUTH_ENABLED        = false
HTTP_CLIENT_INPUT_OAUTH_KEY            =
HTTP_CLIENT_INPUT_OAUTH_SECRET         =
HTTP_CLIENT_INPUT_OAUTH_TOKEN          =
HTTP_CLIENT_INPUT_OAUTH_TOKEN_SECRET   =
HTTP_CLIENT_INPUT_OAUTH_URL            =
HTTP_CLIENT_INPUT_BASIC_AUTH_ENABLED   = false
HTTP_CLIENT_INPUT_BASIC_AUTH_USERNAME  =
HTTP_CLIENT_INPUT_BASIC_AUTH_PASSWORD  =
HTTP_CLIENT_INPUT_TIMEOUT_MS           = 5000
HTTP_CLIENT_INPUT_SKIP_CERT_VERIFY     = false
HTTP_CLIENT_OUTPUT_URL                 =
HTTP_CLIENT_OUTPUT_VERB                = POST
HTTP_CLIENT_OUTPUT_CONTENT_TYPE        = application/octet-stream
HTTP_CLIENT_OUTPUT_OAUTH_ENABLED       = false
HTTP_CLIENT_OUTPUT_OAUTH_KEY           =
HTTP_CLIENT_OUTPUT_OAUTH_SECRET        =
HTTP_CLIENT_OUTPUT_OAUTH_TOKEN         =
HTTP_CLIENT_OUTPUT_OAUTH_TOKEN_SECRET  =
HTTP_CLIENT_OUTPUT_OAUTH_URL           =
HTTP_CLIENT_OUTPUT_BASIC_AUTH_ENABLED  = false
HTTP_CLIENT_OUTPUT_BASIC_AUTH_USERNAME =
HTTP_CLIENT_OUTPUT_BASIC_AUTH_PASSWORD =
HTTP_CLIENT_OUTPUT_TIMEOUT_MS          = 5000
HTTP_CLIENT_OUTPUT_SKIP_CERT_VERIFY    = false

HTTP_SERVER_INPUT_ADDRESS      =
HTTP_SERVER_INPUT_PATH         = /post
HTTP_SERVER_OUTPUT_ADDRESS     =
HTTP_SERVER_OUTPUT_PATH        = /get
HTTP_SERVER_OUTPUT_STREAM_PATH = /get/stream

KAFKA_INPUT_BROKER_ADDRESSES  =
KAFKA_INPUT_CLIENT_ID         = benthos-client
KAFKA_INPUT_CONSUMER_GROUP    = benthos-group
KAFKA_INPUT_TOPIC             = benthos-stream
KAFKA_INPUT_PARTITION         = 0
KAFKA_INPUT_START_OLDEST      = true
KAFKA_OUTPUT_BROKER_ADDRESSES =
KAFKA_OUTPUT_CLIENT_ID        = benthos-client
KAFKA_OUTPUT_KEY              =
KAFKA_OUTPUT_ROUND_ROBIN      = false
KAFKA_OUTPUT_CONSUMER_GROUP   = benthos-group
KAFKA_OUTPUT_TOPIC            = benthos-stream
KAFKA_OUTPUT_PARTITION        = 0
KAFKA_OUTPUT_MAX_MSG_BYTES    = 1000000
KAFKA_OUTPUT_START_OLDEST     = true
KAFKA_OUTPUT_ACK_REP          = true

NATS_INPUT_URLS         =
NATS_INPUT_SUBJECT      = benthos-stream
NATS_INPUT_CLUSTER_ID   = benthos-cluster  # Used only for nats_stream
NATS_INPUT_CLIENT_ID    = benthos-consumer # ^
NATS_INPUT_QUEUE        = benthos-queue    # ^
NATS_INPUT_DURABLE_NAME = benthos-offset   # ^
NATS_OUTPUT_URLS        =
NATS_OUTPUT_SUBJECT     = benthos-stream
NATS_OUTPUT_CLUSTER_ID  = benthos-cluster  # Used only for nats_stream
NATS_OUTPUT_CLIENT_ID   = benthos-consumer # ^

NSQD_INPUT_TCP_ADDRESSES    =
NSQD_INPUT_LOOKUP_ADDRESSES =
NSQ_INPUT_TOPIC             = benthos-messages
NSQ_INPUT_CHANNEL           = benthos-stream
NSQ_INPUT_USER_AGENT        = benthos-consumer
NSQ_OUTPUT_TCP_ADDRESS      =
NSQ_OUTPUT_TOPIC            = benthos-messages
NSQ_OUTPUT_CHANNEL          = benthos-stream
NSQ_OUTPUT_USER_AGENT       = benthos-consumer

REDIS_INPUT_URL        =
REDIS_INPUT_CHANNEL    = benthos-stream
REDIS_INPUT_LIST       = benthos_list
REDIS_INPUT_TIMEOUT_MS = 5000
REDIS_OUTPUT_URL       =
REDIS_OUTPUT_CHANNEL   = benthos-stream
REDIS_OUTPUT_LIST      = benthos_list

SCALE_PROTO_INPUT_URLS    =
SCALE_PROTO_INPUT_BIND    = true
SCALE_PROTO_INPUT_SOCKET  = PULL
SCALE_PROTO_OUTPUT_URLS   =
SCALE_PROTO_OUTPUT_BIND   = false
SCALE_PROTO_OUTPUT_SOCKET = PUSH

ZMQ_INPUT_URLS    =
ZMQ_INPUT_BIND    = true
ZMQ_INPUT_SOCKET  = PULL
ZMQ_OUTPUT_URLS   =
ZMQ_OUTPUT_BIND   = true
ZMQ_OUTPUT_SOCKET = PULL
```

[0]: ../../resources/docs/inputs/README.md
[1]: ../../resources/docs/outputs/README.md
