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

@kotlin.jvm.JvmName("-initializerepositoryModelLoadResponse")
public inline fun repositoryModelLoadResponse(block: io.seldon.mlops.inference.v2.RepositoryModelLoadResponseKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.inference.v2.V2Dataplane.RepositoryModelLoadResponse =
  io.seldon.mlops.inference.v2.RepositoryModelLoadResponseKt.Dsl._create(io.seldon.mlops.inference.v2.V2Dataplane.RepositoryModelLoadResponse.newBuilder()).apply { block() }._build()
/**
 * ```
 * @@
 * @@.. cpp:var:: message RepositoryModelLoadResponse
 * @@
 * @@   Response message for RepositoryModelLoad.
 * @@
 * ```
 *
 * Protobuf type `inference.RepositoryModelLoadResponse`
 */
public object RepositoryModelLoadResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: io.seldon.mlops.inference.v2.V2Dataplane.RepositoryModelLoadResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: io.seldon.mlops.inference.v2.V2Dataplane.RepositoryModelLoadResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): io.seldon.mlops.inference.v2.V2Dataplane.RepositoryModelLoadResponse = _builder.build()
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun io.seldon.mlops.inference.v2.V2Dataplane.RepositoryModelLoadResponse.copy(block: `io.seldon.mlops.inference.v2`.RepositoryModelLoadResponseKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.inference.v2.V2Dataplane.RepositoryModelLoadResponse =
  `io.seldon.mlops.inference.v2`.RepositoryModelLoadResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()

