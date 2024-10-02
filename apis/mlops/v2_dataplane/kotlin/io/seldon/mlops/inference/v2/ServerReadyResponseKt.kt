/*
Copyright (c) 2024 Seldon Technologies Ltd.

Use of this software is governed BY
(1) the license included in the LICENSE file or
(2) if the license included in the LICENSE file is the Business Source License 1.1,
the Change License after the Change Date as each is defined in accordance with the LICENSE file.
*/

// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: v2_dataplane.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package io.seldon.mlops.inference.v2;

@kotlin.jvm.JvmName("-initializeserverReadyResponse")
public inline fun serverReadyResponse(block: io.seldon.mlops.inference.v2.ServerReadyResponseKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyResponse =
  io.seldon.mlops.inference.v2.ServerReadyResponseKt.Dsl._create(io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyResponse.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `inference.ServerReadyResponse`
 */
public object ServerReadyResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyResponse = _builder.build()

    /**
     * ```
     * True if the inference server is ready, false if not ready.
     * ```
     *
     * `bool ready = 1;`
     */
    public var ready: kotlin.Boolean
      @JvmName("getReady")
      get() = _builder.getReady()
      @JvmName("setReady")
      set(value) {
        _builder.setReady(value)
      }
    /**
     * ```
     * True if the inference server is ready, false if not ready.
     * ```
     *
     * `bool ready = 1;`
     */
    public fun clearReady() {
      _builder.clearReady()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyResponse.copy(block: `io.seldon.mlops.inference.v2`.ServerReadyResponseKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyResponse =
  `io.seldon.mlops.inference.v2`.ServerReadyResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()

