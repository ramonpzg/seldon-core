/*
Copyright 2023 Seldon Technologies Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: chainer.proto

package io.seldon.mlops.chainer;

@kotlin.jvm.JvmName("-initializepipelineUpdateStatusResponse")
public inline fun pipelineUpdateStatusResponse(block: io.seldon.mlops.chainer.PipelineUpdateStatusResponseKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusResponse =
  io.seldon.mlops.chainer.PipelineUpdateStatusResponseKt.Dsl._create(io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusResponse.newBuilder()).apply { block() }._build()
public object PipelineUpdateStatusResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusResponse = _builder.build()
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusResponse.copy(block: io.seldon.mlops.chainer.PipelineUpdateStatusResponseKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineUpdateStatusResponse =
  io.seldon.mlops.chainer.PipelineUpdateStatusResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()

