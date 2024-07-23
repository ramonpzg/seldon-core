/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: chainer.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package io.seldon.mlops.chainer;

@kotlin.jvm.JvmName("-initializepipelineSubscriptionRequest")
public inline fun pipelineSubscriptionRequest(block: io.seldon.mlops.chainer.PipelineSubscriptionRequestKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineSubscriptionRequest =
  io.seldon.mlops.chainer.PipelineSubscriptionRequestKt.Dsl._create(io.seldon.mlops.chainer.ChainerOuterClass.PipelineSubscriptionRequest.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `seldon.mlops.chainer.PipelineSubscriptionRequest`
 */
public object PipelineSubscriptionRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineSubscriptionRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: io.seldon.mlops.chainer.ChainerOuterClass.PipelineSubscriptionRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): io.seldon.mlops.chainer.ChainerOuterClass.PipelineSubscriptionRequest = _builder.build()

    /**
     * `string name = 1;`
     */
    public var name: kotlin.String
      @JvmName("getName")
      get() = _builder.getName()
      @JvmName("setName")
      set(value) {
        _builder.setName(value)
      }
    /**
     * `string name = 1;`
     */
    public fun clearName() {
      _builder.clearName()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun io.seldon.mlops.chainer.ChainerOuterClass.PipelineSubscriptionRequest.copy(block: `io.seldon.mlops.chainer`.PipelineSubscriptionRequestKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.chainer.ChainerOuterClass.PipelineSubscriptionRequest =
  `io.seldon.mlops.chainer`.PipelineSubscriptionRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()

