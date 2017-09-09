package io.autonomy.devise;

import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;

/**
 * <pre>
 * The devise service definition.
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.3.0)",
    comments = "Source: api.proto")
public final class DeviseGrpc {

  private DeviseGrpc() {}

  public static final String SERVICE_NAME = "api.Devise";

  // Static method descriptors that strictly reflect the proto.
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<io.autonomy.devise.DeviseProto.OpenTemplateRequest,
      io.autonomy.devise.DeviseProto.OpenTemplateReply> METHOD_OPEN_TEMPLATE =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "api.Devise", "OpenTemplate"),
          io.grpc.protobuf.ProtoUtils.marshaller(io.autonomy.devise.DeviseProto.OpenTemplateRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(io.autonomy.devise.DeviseProto.OpenTemplateReply.getDefaultInstance()));
  @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/1901")
  public static final io.grpc.MethodDescriptor<io.autonomy.devise.DeviseProto.RenderTemplateRequest,
      io.autonomy.devise.DeviseProto.RenderTemplateReply> METHOD_RENDER_TEMPLATE =
      io.grpc.MethodDescriptor.create(
          io.grpc.MethodDescriptor.MethodType.UNARY,
          generateFullMethodName(
              "api.Devise", "RenderTemplate"),
          io.grpc.protobuf.ProtoUtils.marshaller(io.autonomy.devise.DeviseProto.RenderTemplateRequest.getDefaultInstance()),
          io.grpc.protobuf.ProtoUtils.marshaller(io.autonomy.devise.DeviseProto.RenderTemplateReply.getDefaultInstance()));

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static DeviseStub newStub(io.grpc.Channel channel) {
    return new DeviseStub(channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static DeviseBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    return new DeviseBlockingStub(channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary and streaming output calls on the service
   */
  public static DeviseFutureStub newFutureStub(
      io.grpc.Channel channel) {
    return new DeviseFutureStub(channel);
  }

  /**
   * <pre>
   * The devise service definition.
   * </pre>
   */
  public static abstract class DeviseImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     * Opens the template.
     * </pre>
     */
    public void openTemplate(io.autonomy.devise.DeviseProto.OpenTemplateRequest request,
        io.grpc.stub.StreamObserver<io.autonomy.devise.DeviseProto.OpenTemplateReply> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_OPEN_TEMPLATE, responseObserver);
    }

    /**
     * <pre>
     * Renders a template.
     * </pre>
     */
    public void renderTemplate(io.autonomy.devise.DeviseProto.RenderTemplateRequest request,
        io.grpc.stub.StreamObserver<io.autonomy.devise.DeviseProto.RenderTemplateReply> responseObserver) {
      asyncUnimplementedUnaryCall(METHOD_RENDER_TEMPLATE, responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            METHOD_OPEN_TEMPLATE,
            asyncUnaryCall(
              new MethodHandlers<
                io.autonomy.devise.DeviseProto.OpenTemplateRequest,
                io.autonomy.devise.DeviseProto.OpenTemplateReply>(
                  this, METHODID_OPEN_TEMPLATE)))
          .addMethod(
            METHOD_RENDER_TEMPLATE,
            asyncUnaryCall(
              new MethodHandlers<
                io.autonomy.devise.DeviseProto.RenderTemplateRequest,
                io.autonomy.devise.DeviseProto.RenderTemplateReply>(
                  this, METHODID_RENDER_TEMPLATE)))
          .build();
    }
  }

  /**
   * <pre>
   * The devise service definition.
   * </pre>
   */
  public static final class DeviseStub extends io.grpc.stub.AbstractStub<DeviseStub> {
    private DeviseStub(io.grpc.Channel channel) {
      super(channel);
    }

    private DeviseStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DeviseStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new DeviseStub(channel, callOptions);
    }

    /**
     * <pre>
     * Opens the template.
     * </pre>
     */
    public void openTemplate(io.autonomy.devise.DeviseProto.OpenTemplateRequest request,
        io.grpc.stub.StreamObserver<io.autonomy.devise.DeviseProto.OpenTemplateReply> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_OPEN_TEMPLATE, getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Renders a template.
     * </pre>
     */
    public void renderTemplate(io.autonomy.devise.DeviseProto.RenderTemplateRequest request,
        io.grpc.stub.StreamObserver<io.autonomy.devise.DeviseProto.RenderTemplateReply> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(METHOD_RENDER_TEMPLATE, getCallOptions()), request, responseObserver);
    }
  }

  /**
   * <pre>
   * The devise service definition.
   * </pre>
   */
  public static final class DeviseBlockingStub extends io.grpc.stub.AbstractStub<DeviseBlockingStub> {
    private DeviseBlockingStub(io.grpc.Channel channel) {
      super(channel);
    }

    private DeviseBlockingStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DeviseBlockingStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new DeviseBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Opens the template.
     * </pre>
     */
    public io.autonomy.devise.DeviseProto.OpenTemplateReply openTemplate(io.autonomy.devise.DeviseProto.OpenTemplateRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_OPEN_TEMPLATE, getCallOptions(), request);
    }

    /**
     * <pre>
     * Renders a template.
     * </pre>
     */
    public io.autonomy.devise.DeviseProto.RenderTemplateReply renderTemplate(io.autonomy.devise.DeviseProto.RenderTemplateRequest request) {
      return blockingUnaryCall(
          getChannel(), METHOD_RENDER_TEMPLATE, getCallOptions(), request);
    }
  }

  /**
   * <pre>
   * The devise service definition.
   * </pre>
   */
  public static final class DeviseFutureStub extends io.grpc.stub.AbstractStub<DeviseFutureStub> {
    private DeviseFutureStub(io.grpc.Channel channel) {
      super(channel);
    }

    private DeviseFutureStub(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected DeviseFutureStub build(io.grpc.Channel channel,
        io.grpc.CallOptions callOptions) {
      return new DeviseFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Opens the template.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.autonomy.devise.DeviseProto.OpenTemplateReply> openTemplate(
        io.autonomy.devise.DeviseProto.OpenTemplateRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_OPEN_TEMPLATE, getCallOptions()), request);
    }

    /**
     * <pre>
     * Renders a template.
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<io.autonomy.devise.DeviseProto.RenderTemplateReply> renderTemplate(
        io.autonomy.devise.DeviseProto.RenderTemplateRequest request) {
      return futureUnaryCall(
          getChannel().newCall(METHOD_RENDER_TEMPLATE, getCallOptions()), request);
    }
  }

  private static final int METHODID_OPEN_TEMPLATE = 0;
  private static final int METHODID_RENDER_TEMPLATE = 1;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final DeviseImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(DeviseImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_OPEN_TEMPLATE:
          serviceImpl.openTemplate((io.autonomy.devise.DeviseProto.OpenTemplateRequest) request,
              (io.grpc.stub.StreamObserver<io.autonomy.devise.DeviseProto.OpenTemplateReply>) responseObserver);
          break;
        case METHODID_RENDER_TEMPLATE:
          serviceImpl.renderTemplate((io.autonomy.devise.DeviseProto.RenderTemplateRequest) request,
              (io.grpc.stub.StreamObserver<io.autonomy.devise.DeviseProto.RenderTemplateReply>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static final class DeviseDescriptorSupplier implements io.grpc.protobuf.ProtoFileDescriptorSupplier {
    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return io.autonomy.devise.DeviseProto.getDescriptor();
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (DeviseGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new DeviseDescriptorSupplier())
              .addMethod(METHOD_OPEN_TEMPLATE)
              .addMethod(METHOD_RENDER_TEMPLATE)
              .build();
        }
      }
    }
    return result;
  }
}
