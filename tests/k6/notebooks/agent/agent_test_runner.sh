#!/bin/bash

NUM_MODELS=1
NUM_ITERS=5
NUM_VUS=5
MODEL_NAME="tfsimple_string" # check model.js
EXTRA="--out influxdb=http://localhost:8086/k6db ../../scenarios/predict_agent.js"
DIR="results"

mkdir -p $DIR


export SCHEDULER_ENDPOINT=0.0.0.0:8100 
export INFER_GRPC_ENDPOINT=0.0.0.0:9998
export INFER_HTTP_ENDPOINT=http://0.0.0.0:9999
export INFER_HTTP_ITERATIONS=10 
export INFER_GRPC_ITERATIONS=10 
export MODEL_TYPE=$MODEL_NAME
export MAX_NUM_MODELS=$NUM_MODELS
export SCHEDULER_PROXY="true"
export ENVOY="false"
export INFER_BATCH_SIZE=8
#export SKIP_UNLOAD_MODEL=1

k6 run -u $NUM_VUS -i $NUM_ITERS --summary-export $DIR/base.json --out csv=$DIR/base.gz $EXTRA


export INFER_GRPC_ENDPOINT=0.0.0.0:9998 
export INFER_HTTP_ENDPOINT=http://0.0.0.0:9999 


#k6 run -u $NUM_VUS -i $NUM_ITERS --summary-export $DIR/oc_$MAX_NUM_MODELS.json --out csv=$DIR/$MAX_NUM_MODELS.gz $EXTRA

export MAX_NUM_MODELS=$(( $NUM_MODELS + 5 ))
#k6 run -u $NUM_VUS -i $NUM_ITERS --summary-export $DIR/oc_$MAX_NUM_MODELS.json --out csv=$DIR/$MAX_NUM_MODELS.gz $EXTRA

export MAX_NUM_MODELS=$(( $NUM_MODELS + 10 ))
#k6 run -u $NUM_VUS -i $NUM_ITERS --summary-export $DIR/oc_$MAX_NUM_MODELS.json --out csv=$DIR/$MAX_NUM_MODELS.gz $EXTRA

export MAX_NUM_MODELS=$(( $NUM_MODELS + 25 ))
#k6 run -u $NUM_VUS -i $NUM_ITERS --summary-export $DIR/oc_$MAX_NUM_MODELS.json --out csv=$DIR/$MAX_NUM_MODELS.gz $EXTRA