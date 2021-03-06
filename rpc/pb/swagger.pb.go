package rpcpb

const (
	swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "rpc.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/block": {
      "get": {
        "operationId": "GetBlock",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/rpcpbBlockResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "hash",
            "description": "Block hash. Or the string \"genesis\", \"confirmed\", \"tail\".",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ApiService"
        ]
      }
    },
    "/v1/node/medstate": {
      "get": {
        "operationId": "GetMedState",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/rpcpbGetMedStateResponse"
            }
          }
        },
        "tags": [
          "ApiService"
        ]
      }
    },
    "/v1/transaction": {
      "get": {
        "operationId": "GetTransaction",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/rpcpbTransactionResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "hash",
            "description": "Transaction hash.",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ApiService"
        ]
      },
      "post": {
        "operationId": "SendTransaction",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/rpcpbSendTransactionResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/rpcpbSendTransactionRequest"
            }
          }
        ],
        "tags": [
          "ApiService"
        ]
      }
    },
    "/v1/user/accountstate": {
      "get": {
        "operationId": "GetAccountState",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/rpcpbGetAccountStateResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "address",
            "description": "Hex string of the account addresss.",
            "in": "query",
            "required": false,
            "type": "string"
          },
          {
            "name": "height",
            "description": "block account state with height. Or the string \"genesis\", \"confirmed\", \"tail\".",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ApiService"
        ]
      }
    }
  },
  "definitions": {
    "rpcpbBlockResponse": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "string",
          "title": "Block hash"
        },
        "parent_hash": {
          "type": "string",
          "title": "Block parent hash"
        },
        "coinbase": {
          "type": "string",
          "title": "Block coinbase address"
        },
        "timestamp": {
          "type": "string",
          "format": "int64",
          "title": "Block timestamp"
        },
        "chain_id": {
          "type": "integer",
          "format": "int64",
          "title": "Block chain id"
        },
        "alg": {
          "type": "integer",
          "format": "int64",
          "title": "Block signature algorithm"
        },
        "sign": {
          "type": "string",
          "title": "Block signature"
        },
        "accs_root": {
          "type": "string",
          "title": "Root hash of accounts trie"
        },
        "txs_root": {
          "type": "string",
          "title": "Root hash of transactions trie"
        },
        "usage_root": {
          "type": "string",
          "title": "Root hash of usage trie"
        },
        "records_root": {
          "type": "string",
          "title": "Root hash of records trie"
        },
        "consensus_root": {
          "type": "string",
          "title": "Root hash of consensus trie"
        },
        "transactions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/rpcpbTransactionResponse"
          },
          "title": "Transactions in block"
        },
        "height": {
          "type": "string",
          "format": "uint64",
          "title": "Block height"
        }
      }
    },
    "rpcpbGetAccountStateResponse": {
      "type": "object",
      "properties": {
        "balance": {
          "type": "string",
          "description": "Current balance in unit of 1/(10^18) nas."
        },
        "nonce": {
          "type": "string",
          "format": "uint64",
          "description": "Current transaction count."
        },
        "type": {
          "type": "integer",
          "format": "int64",
          "title": "Account type"
        }
      }
    },
    "rpcpbGetMedStateResponse": {
      "type": "object",
      "properties": {
        "chain_id": {
          "type": "integer",
          "format": "int64",
          "title": "Block chain id"
        },
        "tail": {
          "type": "string",
          "title": "Current tail block hash"
        },
        "height": {
          "type": "string",
          "format": "uint64",
          "title": "Current tail block height"
        },
        "protocol_version": {
          "type": "string",
          "description": "The current med protocol version."
        },
        "version": {
          "type": "string",
          "title": "Med version"
        }
      }
    },
    "rpcpbSendTransactionRequest": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "string",
          "title": "Transaction hash"
        },
        "from": {
          "type": "string",
          "description": "Hex string of the sender account addresss."
        },
        "to": {
          "type": "string",
          "description": "Hex string of the receiver account addresss."
        },
        "value": {
          "type": "string",
          "description": "Amount of value sending with this transaction."
        },
        "timestamp": {
          "type": "string",
          "format": "int64",
          "description": "Transaction timestamp."
        },
        "data": {
          "$ref": "#/definitions/rpcpbTransactionData",
          "description": "Transaction Data type."
        },
        "nonce": {
          "type": "string",
          "format": "uint64",
          "description": "Transaction nonce."
        },
        "chain_id": {
          "type": "integer",
          "format": "int64",
          "description": "Transaction chain ID."
        },
        "alg": {
          "type": "integer",
          "format": "int64",
          "description": "Transaction algorithm."
        },
        "sign": {
          "type": "string",
          "description": "Transaction sign."
        },
        "payer_sign": {
          "type": "string",
          "description": "Transaction payer's sign."
        }
      }
    },
    "rpcpbSendTransactionResponse": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "string",
          "description": "Hex string of transaction hash."
        }
      }
    },
    "rpcpbTransactionData": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "description": "Transaction data type."
        },
        "payload": {
          "type": "string",
          "description": "Transaction data payload."
        }
      }
    },
    "rpcpbTransactionResponse": {
      "type": "object",
      "properties": {
        "hash": {
          "type": "string",
          "title": "Transaction hash"
        },
        "from": {
          "type": "string",
          "description": "Hex string of the sender account addresss."
        },
        "to": {
          "type": "string",
          "description": "Hex string of the receiver account addresss."
        },
        "value": {
          "type": "string",
          "description": "Amount of value sending with this transaction."
        },
        "timestamp": {
          "type": "string",
          "format": "int64",
          "description": "Transaction timestamp."
        },
        "data": {
          "$ref": "#/definitions/rpcpbTransactionData",
          "description": "Transaction Data type."
        },
        "nonce": {
          "type": "string",
          "format": "uint64",
          "description": "Transaction nonce."
        },
        "chain_id": {
          "type": "integer",
          "format": "int64",
          "description": "Transaction chain ID."
        },
        "alg": {
          "type": "integer",
          "format": "int64",
          "description": "Transaction algorithm."
        },
        "sign": {
          "type": "string",
          "description": "Transaction sign."
        },
        "payer_sign": {
          "type": "string",
          "description": "Transaction payer's sign."
        }
      }
    }
  }
}
`
)
